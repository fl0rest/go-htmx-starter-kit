FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

RUN addgroup -g 1001 -S appgroup && \
    adduser -S -u 1001 -G appgroup appuser

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/assets ./assets
COPY /.env .

RUN mkdir -p /app/log
RUN touch /app/log/error.log /app/log/app.log
RUN chown -R appuser:appgroup /app
RUN chmod 666 /app/log/*.log

USER appuser
EXPOSE 8000

HEALTHCHECK --interval=30s --timeout=3s --start-period=15s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8000/ || exit 1

ENTRYPOINT ["./main"]
