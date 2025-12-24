FROM oven/bun:latest AS builder
WORKDIR /app

COPY package.json bun.lockb* ./
RUN bun install --no-save

COPY . .
COPY .env.production ./

# Build the app
ENV NODE_ENV=production
RUN bun run build

RUN ls -lah /app/dist/assets/ && \
    echo "✓ Build complete. Checking for CSS files..." && \
    find /app/dist -name "*.css" -type f | head -5

# Stage 2: Production  
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost/health || exit 1

CMD ["nginx", "-g", "daemon off;"]