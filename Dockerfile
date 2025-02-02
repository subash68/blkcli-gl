FROM golang:1.23.5-alpine3.20 AS builder
# set environment variable
ARG CGO_ENABLE=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build

FROM scratch
COPY --from=builder /app/blkcli-gl /blkcli-gl
EXPOSE 8000
ENTRYPOINT ["/blkcli-gl"]
