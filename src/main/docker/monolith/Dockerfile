# Builds the production monolith that serves the immutable Parcel distribution
# and Go API from one Gin process.
# Authored by: OpenCode

# Builds the immutable production frontend from lockfile-managed dependencies.
FROM node:26.13.0-bookworm-slim@sha256:bc0a2afd3b74ce12222fc8bab99271e66257f021078f5f4a5d687a77d280c1d9 AS frontend-build

WORKDIR /frontend

COPY web-static/package.json web-static/package-lock.json ./
RUN npm ci

COPY web-static/.parcelrc web-static/root.html web-static/tsconfig.json ./
COPY web-static/websrc ./websrc

RUN npm run build -- --no-source-maps

# Builds the statically linked Linux backend from locked Go modules.
FROM golang:1.27.0 AS backend-build

WORKDIR /backend/src

COPY go/go.mod go/go.sum ./
RUN go mod download

COPY go/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend/bin/open-asset-allocator

# Packages only the runtime dependencies and immutable application artifacts.
FROM gcr.io/distroless/static-debian13:nonroot@sha256:e2e927ec666bae08560abb3c55d0659eceabb657f56b6782ab500a9fc7f555e3 AS runtime
LABEL authors="OpenCode"

COPY --from=backend-build /backend/bin/open-asset-allocator /backend/bin/open-asset-allocator
COPY --from=frontend-build /frontend/dist /frontend/dist

ENV API_ONLY=false \
    PORT=8080 \
    ROOT_HTML_FILENAME=root.html \
    WEB_STATIC_CONTENT_PATH=/frontend/dist \
    WEB_STATIC_SOURCE_REL_PATH=/websrc

EXPOSE 8080

USER 65532:65532
CMD ["/backend/bin/open-asset-allocator"]
