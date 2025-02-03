FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download

COPY . .
COPY .env .
RUN go build -o app cmd/main.go

FROM scratch

COPY --from=builder /app/app /app/app
COPY --from=builder /app/.env /app/.env

WORKDIR /app

ENTRYPOINT ["/app/app"]

