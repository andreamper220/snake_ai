FROM golang:1.22.1-bookworm as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -v -o server ./cmd/server

FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/internal/templates /app/internal/templates
COPY --from=builder /app/server /app/server

CMD ["/app/server", \
     "-d", "host=snake_db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", \
     "-r", "redis://snake_redis:6379" \
     ]