# Stage 1: Build the binary
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Stage 2: Create the final tiny image
FROM alpine:latest
WORKDIR /root/
# Copy only the compiled program from the builder stage
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]