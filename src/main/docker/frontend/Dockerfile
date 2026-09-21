# Provides the containerized Parcel development environment for the frontend.
# Authored by: OpenCode
FROM node:24.21.0-bookworm-slim@sha256:0e0ff40c39bc087845bfb27465a0df4ea419520094bc35842ff83dd8cbe6f9b6 AS dev
LABEL authors="OpenCode"

WORKDIR /frontend

COPY package.json package-lock.json ./
RUN npm ci

COPY .parcelrc .proxyrc.js root.html tsconfig.json ./
COPY websrc ./websrc

EXPOSE 8000

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--cache-dir", ".parcel-cache", "--dist-dir", "dist"]
