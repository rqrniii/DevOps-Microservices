FROM oven/bun:latest AS builder
WORKDIR /app

COPY package.json bun.lockb* ./

RUN bun install --no-save
COPY . .

# Build the app
RUN bun run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html

RUN echo 'server { \
    listen 80; \
    location / { \
        root /usr/share/nginx/html; \
        index index.html; \
        try_files $uri $uri/ /index.html; \
    } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]