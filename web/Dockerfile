ARG NODE_VERSION=22.11.0
ARG NGINX_VERSION=1.27
ARG NODE_OPTIONS="--max-old-space-size=4096"
ARG GITHUB_SHA

FROM node:${NODE_VERSION}-slim AS builder

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV NODE_OPTIONS=$NODE_OPTIONS
ENV GITHUB_SHA=$GITHUB_SHA

RUN npm install -g corepack && corepack enable

WORKDIR /app

RUN --mount=type=bind,source=package.json,target=package.json \
  --mount=type=bind,source=pnpm-lock.yaml,target=pnpm-lock.yaml \
  --mount=type=cache,id=pnpm,target=/pnpm/store \
  pnpm install --frozen-lockfile

COPY . /app

RUN --mount=type=cache,id=pnpm,target=/pnpm/store \
  pnpm run build

FROM nginx:${NGINX_VERSION}-alpine
WORKDIR /usr/share/nginx/html

RUN --mount=type=cache,target=/var/cache/apk \
  apk add --no-cache jq

COPY --from=builder --chown=nginx:nginx /app/dist /usr/share/nginx/html
COPY --chown=nginx:nginx docker/nginx.conf.template /etc/nginx/templates/nginx.conf.template
COPY --chown=nginx:nginx docker/40-envsubst-on-reearth-config.sh /docker-entrypoint.d
COPY --chown=nginx:nginx docker/reearth_config.json.template /opt/reearth/reearth_config.json.template

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
