# Build image
FROM golang:latest AS build-env

WORKDIR /app
ENV GO111MODULE=on
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
COPY . .
RUN go test ./...
RUN go build

# Run image
FROM alpine:latest
COPY --from=build-env /app/flyte-bamboo .

ENTRYPOINT ["./flyte-bamboo"]
