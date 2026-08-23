# Provides the containerized Parcel development environment for the frontend.
# Authored by: OpenCode
FROM node:24.12.0-bookworm-slim AS dev
LABEL authors="OpenCode"

WORKDIR /frontend

COPY package.json package-lock.json ./
RUN npm ci

COPY .parcelrc .proxyrc.js root.html tsconfig.json ./
COPY websrc ./websrc

EXPOSE 8000

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--cache-dir", ".parcel-cache", "--dist-dir", "dist"]
