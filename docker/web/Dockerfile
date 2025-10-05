# --- Build the Vue app ---
FROM node:20-alpine AS build
WORKDIR /app

ARG PKG_MGR=npm

# Install deps first for better caching
COPY frontend/package*.json ./
RUN if [ "$PKG_MGR" = "npm" ]; then \
      if [ -f package-lock.json ] || [ -f npm-shrinkwrap.json ]; then npm ci; else npm install; fi; \
    elif [ "$PKG_MGR" = "pnpm" ]; then \
      corepack enable && corepack prepare pnpm@latest --activate && pnpm install --frozen-lockfile; \
    else \
      corepack enable && corepack prepare yarn@stable --activate && yarn install --frozen-lockfile; \
    fi

# Copy source and build
COPY frontend/ ./
RUN if [ "$PKG_MGR" = "npm" ]; then npm run build; \
    elif [ "$PKG_MGR" = "pnpm" ]; then pnpm build; \
    else yarn build; fi

# --- Nginx to serve the built app and proxy /api ---
FROM nginx:1.27-alpine
RUN rm -f /etc/nginx/conf.d/default.conf
COPY docker/web/nginx.conf /etc/nginx/conf.d/default.conf

# Static files
COPY --from=build /app/dist /usr/share/nginx/html

# Healthcheck (optional)
HEALTHCHECK --interval=10s --timeout=3s --retries=10 CMD wget -qO- http://localhost/ >/dev/null 2>&1 || exit 1