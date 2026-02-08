FROM cosmtrek/air:v1.64.4 AS dev

RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM golang:1.25.7
LABEL authors="benizzio"

WORKDIR /backend
COPY src-go ./src

WORKDIR /backend/src
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o ../bin/open-asset-allocator

WORKDIR /frontend
COPY dist-web-static ./dist

ENV WEB_STATIC_CONTENT_PATH=/frontend/dist

CMD ["/backend/bin/open-asset-allocator"]