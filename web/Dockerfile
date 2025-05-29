ARG NODE_VERSION=22.11.0
ARG NGINX_VERSION=1.27
ARG GITHUB_SHA

FROM node:${NODE_VERSION}-slim AS builder

# https://pnpm.io/ja/next/docker
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN npm install -g pnpm

WORKDIR /app

RUN echo "manage-package-manager-versions=true" > .npmrc
COPY pnpm-lock.yaml ./
RUN pnpm fetch

ENV GITHUB_SHA=$GITHUB_SHA

COPY . .

WORKDIR /app/apps

RUN pnpm i --frozen-lockfile
RUN pnpm run build

FROM nginx:${NGINX_VERSION}-alpine
WORKDIR /usr/share/nginx/html

COPY --from=builder --chown=nginx:nginx /app/apps/dist /usr/share/nginx/html

COPY --chown=nginx:nginx docker/nginx.conf.template /etc/nginx/templates/nginx.conf.template

ENV NGINX_ENTRYPOINT_QUIET_LOGS=true
ENV PORT=8080
ENV REAL_IP_HEADER=X-Forwarded-For

CMD ["nginx", "-g", "daemon off;"]
