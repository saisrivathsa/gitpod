FROM caddy:2.3.0-builder AS builder

WORKDIR /plugin

COPY plugin/go.* /plugin/

RUN go mod download

COPY plugin /plugin

RUN xcaddy build \
  --output /usr/bin/caddy \
  --with github.com/gitpod-io/gitpod/proxy/plugin=/plugin

FROM caddy:2.3.0

COPY --from=builder /usr/bin/caddy /usr/bin/caddy
