# Provides the containerized Parcel development environment for the frontend.
# Authored by: OpenCode
FROM node:26.13.0-bookworm-slim@sha256:bc0a2afd3b74ce12222fc8bab99271e66257f021078f5f4a5d687a77d280c1d9 AS dev
LABEL authors="OpenCode"

WORKDIR /frontend

COPY package.json package-lock.json ./
RUN npm ci

COPY .parcelrc .proxyrc.js root.html tsconfig.json ./
COPY websrc ./websrc

EXPOSE 8000

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--cache-dir", ".parcel-cache", "--dist-dir", "dist"]
