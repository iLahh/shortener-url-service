FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/main
COPY --from=builder /build/.env /app/.env

WORKDIR /app

EXPOSE 3000

CMD ["./main"]