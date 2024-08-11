# Build stage
FROM golang:1.21.5-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache build-base
RUN go build -o ./main ./cmd/server/main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .

# Install necessary packages
RUN apk add --no-cache pngquant git tzdata

# Copy other necessary files
COPY icons /app/icons
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
COPY cert.pem ./cert.pem

# Set timezone data
ENV TZ=Etc/UTC

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]