# Gunakan base image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy semua file ke container
COPY . .

# Install dependensi
RUN go mod tidy

# Build binary
RUN go build -o main ./cmd/main.go

# Port yang digunakan
EXPOSE 8080

# Jalankan aplikasi
CMD ["/app/main"]
