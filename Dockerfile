# Modules caching
FROM golang:1.19-alpine3.17 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Builder
FROM golang:1.19-alpine3.17 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/app

# Final
FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
CMD ["/app"]
