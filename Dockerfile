FROM golang:1.17-alpine as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server


FROM alpine:3

WORKDIR /app

RUN apk add --no-cache ca-certificates && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Etc/UTC /etc/localtime && \
    echo "Etc/UTC" > /etc/timezone

COPY --from=builder /app/server /app/server
CMD ["/app/server"]
