FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY backend/ ./
RUN go build -o babybox-mock

FROM alpine:latest
COPY --from=builder /app/babybox-mock /babybox-mock
CMD ["/babybox-mock"]
