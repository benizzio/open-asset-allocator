/**
 * Validates that the project Playwright packages and runner image use one version.
 *
 * Run `npm run validate:playwright` from the E2E package, or pass
 * `--dockerfile=/path/to/Dockerfile` to validate a source Dockerfile explicitly.
 * Authored by: OpenCode
 */
import { readFile } from "node:fs/promises";
import { existsSync } from "node:fs";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const scriptDirectory = dirname(fileURLToPath(import.meta.url));
const packageDirectory = resolve(scriptDirectory, "..");
const repositoryDirectory = resolve(scriptDirectory, "../../../..");
const defaultDockerfile = resolve(
  repositoryDirectory,
  "src/main/docker/e2e/Dockerfile",
);
const argumentsByName = new Map(
  process.argv.slice(2).map((argument) => {
    const [name, value] = argument.split("=", 2);
    return [name, value];
  }),
);

const packageJson = await readJson(resolve(packageDirectory, "package.json"));
const packageLock = await readJson(
  resolve(packageDirectory, "package-lock.json"),
);
const versions = new Set([
  packageJson.devDependencies?.["@playwright/test"],
  packageLock.packages?.[""]?.devDependencies?.["@playwright/test"],
  packageLock.packages?.["node_modules/@playwright/test"]?.version,
  packageLock.packages?.["node_modules/playwright"]?.version,
  packageLock.packages?.["node_modules/playwright-core"]?.version,
]);

if (versions.has(undefined) || versions.size !== 1) {
  fail(
    `Playwright package versions must match exactly; found: ${[...versions].join(", ")}`,
  );
}

const [projectVersion] = versions;
if (!/^\d+\.\d+\.\d+$/.test(projectVersion)) {
  fail(
    `Playwright package version must be an exact stable version; found: ${projectVersion}`,
  );
}

const requestedDockerfile = argumentsByName.get("--dockerfile");
const dockerfile =
  requestedDockerfile ??
  (existsSync(defaultDockerfile) ? defaultDockerfile : undefined);
if (dockerfile !== undefined) {
  const imageVersion = parseImageVersion(
    await readFile(dockerfile, "utf8"),
    dockerfile,
  );
  assertEqual(
    projectVersion,
    imageVersion,
    `Docker image version in ${dockerfile}`,
  );
}

const expectedImageVersion = argumentsByName.get("--image-version");
if (expectedImageVersion !== undefined) {
  assertEqual(
    projectVersion,
    expectedImageVersion,
    "expected Docker image version",
  );
}

const dockerInfoPath = "/ms-playwright/.docker-info";
if (existsSync(dockerInfoPath)) {
  const dockerInfo = await readJson(dockerInfoPath);
  assertEqual(
    projectVersion,
    dockerInfo.driverVersion,
    "installed Docker image version",
  );
}

console.log(`Playwright package and image versions match: ${projectVersion}`);

/** Reads and parses a JSON file. */
async function readJson(path) {
  return JSON.parse(await readFile(path, "utf8"));
}

/** Extracts the exact Playwright version from a literal, pinned image instruction. */
function parseImageVersion(content, path) {
  const matchingLines = content
    .split("\n")
    .filter((line) =>
      /^FROM mcr\.microsoft\.com\/playwright:v\d+\.\d+\.\d+-noble(?: AS [A-Za-z0-9_-]+)?\s*$/.test(
        line,
      ),
    );

  if (matchingLines.length !== 1) {
    fail(
      `${path} must contain exactly one literal Playwright Noble image FROM instruction.`,
    );
  }

  return matchingLines[0].match(/:v(\d+\.\d+\.\d+)-noble/)?.[1];
}

/** Throws a consistent validation error when two version sources differ. */
function assertEqual(expected, actual, source) {
  if (expected !== actual) {
    fail(
      `Playwright package version ${expected} does not match ${source} ${actual ?? "<missing>"}.`,
    );
  }
}

/** Terminates validation with a machine-readable error message. */
function fail(message) {
  throw new Error(`Playwright version validation failed: ${message}`);
}
