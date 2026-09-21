FROM golang:1.26-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /out/sso ./cmd/sso \
    && CGO_ENABLED=0 go build -o /out/migrator ./cmd/migrator

FROM alpine:3.20
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=build /out/sso /usr/local/bin/sso
COPY --from=build /out/migrator /usr/local/bin/migrator
COPY migrations /app/migrations
COPY config/prod.yaml /app/config/prod.yaml

CMD ["/bin/sh", "-c", "migrator -database-url \"$DATABASE_URL\" -migrations-path /app/migrations && sso --config=/app/config/prod.yaml"]