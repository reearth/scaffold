ARG NODE_VERSION=22.11.0
ARG NGINX_VERSION=1.27
ARG GITHUB_SHA

FROM node:${NODE_VERSION}-slim AS builder

# https://pnpm.io/ja/next/docker
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN npm install -g corepack && corepack enable

WORKDIR /app

# web/appsの必要なファイルだけをコピー
COPY web/apps/package.json ./web/apps/
COPY web/pnpm-lock.yaml* ./web/
COPY web/apps ./web/apps/

# web/appsディレクトリで依存関係をインストール
WORKDIR /app/web/apps
RUN pnpm install

ENV GITHUB_SHA=$GITHUB_SHA

# ビルド
RUN pnpm build

FROM nginx:${NGINX_VERSION}-alpine
WORKDIR /usr/share/nginx/html

# Viteのビルド結果をコピー
COPY --from=builder --chown=nginx:nginx /app/web/apps/dist /usr/share/nginx/html

# nginx設定をコピー
COPY --chown=nginx:nginx web/apps/docker/nginx.conf.template /etc/nginx/templates/nginx.conf.template

# Nginxの設定
ENV NGINX_ENTRYPOINT_QUIET_LOGS=true
ENV PORT=8080
ENV REAL_IP_HEADER=X-Forwarded-For

CMD ["nginx", "-g", "daemon off;"]