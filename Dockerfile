FROM golang:1.26.8-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG APP=api

RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" \
    -o /app/bin/app ./cmd/${APP}

FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /usr/src/app

COPY --from=builder --chown=nobody:nobody \
    /app/bin/app /usr/local/bin/app

USER nobody:nobody

CMD ["/usr/local/bin/app"]
