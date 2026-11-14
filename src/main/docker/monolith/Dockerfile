# Builds the production monolith that serves the immutable Parcel distribution
# and Go API from one Gin process.
# Authored by: OpenCode

# Builds the immutable production frontend from lockfile-managed dependencies.
FROM node:26.13.1-bookworm-slim@sha256:5df79989ed685e09f0a98019db4cebfb22fb9df1a53c0161664a011f7ff6b0f1 AS frontend-build

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
FROM gcr.io/distroless/static-debian13:nonroot@sha256:25ac8f0c3b2334a9eb05452266bd4697864a8c8ef0ef71e5e6480d7f92868329 AS runtime
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
