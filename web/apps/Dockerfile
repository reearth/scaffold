ARG NODE_VERSION=22.11.0
ARG NGINX_VERSION=1.27
ARG GITHUB_SHA

FROM node:${NODE_VERSION}-slim AS builder

# https://pnpm.io/ja/next/docker
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN npm install -g corepack && corepack enable

COPY . /app
WORKDIR /app

RUN --mount=type=cache,id=pnpm,target=/pnpm/store \
  pnpm install --frozen-lockfile

ENV GITHUB_SHA=$GITHUB_SHA

WORKDIR /app/web/apps
RUN pnpm --filter . build

FROM nginx:${NGINX_VERSION}-alpine
WORKDIR /usr/share/nginx/html

RUN --mount=type=cache,target=/var/cache/apk \
  apk add --no-cache jq

COPY --from=builder --chown=nginx:nginx /app/web/apps/.next/static /usr/share/nginx/html/_next/static
COPY --from=builder --chown=nginx:nginx /app/web/apps/out /usr/share/nginx/html
COPY --chown=nginx:nginx web/apps/docker/nginx.conf.template /etc/nginx/templates/nginx.conf.template
COPY --chown=nginx:nginx web/apps/docker/40-envsubst-on-reearth-config.sh /docker-entrypoint.d
COPY --chown=nginx:nginx web/apps/docker/reearth_config.json.template /opt/reearth/reearth_config.json.template

# Quite the Nginx startup logs.
ENV NGINX_ENTRYPOINT_QUIET_LOGS=true

# Default to Cloud Run port.
# Ref: https://cloud.google.com/run/docs/reference/container-contract#port
ENV PORT=8080

# Defaults Google Cloud Load Balancer header.
# Ref: https://cloud.google.com/load-balancing/docs/https#target-proxies
ENV REAL_IP_HEADER=X-Forwarded-For

ENV REEARTH_TITLE=
ENV REEARTH_FAVICON_URL=

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
