FROM cosmtrek/air:v1.67.4 AS dev

RUN go install github.com/go-delve/delve/cmd/dlv@v1.27.0

FROM golang:1.26.7
LABEL authors="benizzio, OpenCode"

WORKDIR /backend
COPY src-go ./src

WORKDIR /backend/src
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o ../bin/open-asset-allocator

WORKDIR /frontend
COPY dist-web-static ./dist

ENV WEB_STATIC_CONTENT_PATH=/frontend/dist

CMD ["/backend/bin/open-asset-allocator"]
