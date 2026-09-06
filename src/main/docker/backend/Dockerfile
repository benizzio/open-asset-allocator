# Builds the standalone API-only backend from the Go source context.
# Co-authored by: OpenCode and Igor Benicio de Mesquita
FROM cosmtrek/air:v1.67.4 AS dev

RUN go install github.com/go-delve/delve/cmd/dlv@v1.27.0

FROM golang:1.27.1
LABEL authors="benizzio, OpenCode"

WORKDIR /backend/src
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ../bin/open-asset-allocator

ENV API_ONLY=true

CMD ["/backend/bin/open-asset-allocator"]
