FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o api .

FROM scratch

COPY --from=builder /app/api /api

EXPOSE 8000

ENTRYPOINT ["/api"]
