# Provides isolated Playwright runner and headed-debug images for E2E tests.
# Authored by: OpenCode
FROM mcr.microsoft.com/playwright:v1.62.1-noble AS dependencies
LABEL authors="OpenCode"

WORKDIR /workspace/src/test/e2e

ENV HOME=/workspace/.home \
    NPM_CONFIG_CACHE=/workspace/.npm \
    TMPDIR=/workspace/.tmp

COPY package.json package-lock.json ./
RUN npm ci

COPY . ./
RUN npm run validate:playwright -- --image-version=1.62.1 && \
    mkdir -p /workspace/.home /workspace/.npm /workspace/.tmp /workspace/target/e2e-results && \
    chmod -R 0777 /workspace/.home /workspace/.npm /workspace/.tmp /workspace/target

# Adds a virtual display and noVNC only for local headed browser inspection.
FROM dependencies AS debug
USER root

RUN apt-get update && \
    apt-get install -y --no-install-recommends novnc x11vnc xvfb && \
    rm -rf /var/lib/apt/lists/*

EXPOSE 7900
USER pwuser
ENTRYPOINT ["./scripts/start-debug.sh"]
CMD ["npm", "test", "--", "--headed"]

# Provides the minimal non-root image used by headless local and CI runs.
FROM dependencies AS runner
USER pwuser
CMD ["npm", "test"]
