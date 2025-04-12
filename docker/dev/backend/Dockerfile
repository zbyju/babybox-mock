FROM golang:1.22-alpine

WORKDIR /app
RUN apk add --no-cache reflex
COPY backend/go.mod backend/go.sum ./
RUN go mod download
CMD ["reflex", "-r", "\\.go$", "--", "go", "run", "main.go"]
