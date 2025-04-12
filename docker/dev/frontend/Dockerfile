FROM node:18-alpine

WORKDIR /app
RUN npm install -g pnpm
COPY frontend/pnpm-lock.yaml frontend/package.json ./
RUN pnpm install
CMD ["pnpm", "run", "dev", "--host"]
