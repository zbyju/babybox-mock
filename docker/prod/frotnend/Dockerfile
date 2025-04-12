FROM node:18-alpine AS builder
WORKDIR /app
COPY frontend/ ./
RUN pnpm install && pnpm run build

FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
