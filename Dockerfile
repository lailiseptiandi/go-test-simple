FROM golang:1.26.8-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/app ./cmd/main.go

FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata

USER nobody:nobody

COPY --from=builder /app/bin/app /usr/local/bin/app

CMD ["app"]
