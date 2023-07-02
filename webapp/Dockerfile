
FROM node:19-alpine as build

WORKDIR /app

COPY . /app

RUN npm ci

RUN npm run build

FROM caddy:2.6-alpine as deploy

COPY --from=build /app/build /usr/share/caddy/html

COPY Caddyfile /etc/Caddyfile

EXPOSE 443