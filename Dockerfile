FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /api ./cmd/main.go

FROM builder

WORKDIR /

COPY --from=builder /api /api

EXPOSE 8080

ENTRYPOINT ["./api"]
