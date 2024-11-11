# First stage: Build
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api/main.go

# Second state: Optimize
FROM hatamiarash7/upx:latest AS packer

COPY --from=builder /app/app /app

RUN upx --best --lzma /app

# Third stage: Final image
FROM alpine:latest

COPY --from=packer /app /app

RUN chmod +x /app

EXPOSE 3000

ENTRYPOINT ["/app"]
