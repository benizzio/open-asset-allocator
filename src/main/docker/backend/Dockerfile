FROM golang:1.23
LABEL authors="benizzio"

WORKDIR /backend
COPY src-go ./src

WORKDIR /backend/src
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ../bin/open-asset-allocator

WORKDIR /frontend
COPY dist-web-static ./dist

CMD ["/backend/bin/open-asset-allocator"]