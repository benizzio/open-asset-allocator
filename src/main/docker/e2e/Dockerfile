# Provides isolated Playwright runner and headed-debug images for E2E tests.
# Authored by: OpenCode
FROM mcr.microsoft.com/playwright:v1.63.0-noble AS dependencies
LABEL authors="OpenCode"

WORKDIR /workspace/src/test/e2e

ENV HOME=/workspace/.home \
    NPM_CONFIG_CACHE=/workspace/.npm \
    TMPDIR=/workspace/.tmp

COPY package.json package-lock.json ./
RUN npm ci

COPY scripts/validate-playwright-version.mjs ./scripts/
RUN npm run validate:playwright -- --require-image-metadata && \
    mkdir -p /workspace/.home /workspace/.npm /workspace/.tmp /workspace/target/e2e-results && \
    chmod -R 0777 /workspace/.home /workspace/.npm /workspace/.tmp /workspace/target

# Provides the immutable non-root image used by headless local and CI runs.
FROM dependencies AS runner

COPY playwright.config.ts tsconfig.json ./
COPY support ./support
COPY tests ./tests

USER pwuser
CMD ["npm", "test"]

# Adds a virtual display and noVNC only for local headed browser inspection.
FROM dependencies AS debug
USER root

RUN apt-get update && \
    apt-get install -y --no-install-recommends novnc x11vnc xvfb && \
    rm -rf /var/lib/apt/lists/*

COPY playwright.config.ts tsconfig.json ./
COPY support ./support
COPY tests ./tests
COPY scripts/start-debug.sh ./scripts/

EXPOSE 7900
USER pwuser
ENTRYPOINT ["./scripts/start-debug.sh"]
CMD ["npm", "test", "--", "--headed"]
