
FROM node:19-alpine as build

WORKDIR /app

COPY . /app

RUN npm ci

RUN npm run build

FROM nginx:1.21.0-alpine as deploy

COPY --from=build /app/build /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

# Install OpenSSL
RUN apk update && apk add openssl

# Create directory for NGINX SSL/TLS certificates
RUN mkdir -p /etc/nginx/ssl

# Generate SSL/TLS certificate
RUN openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 \
    -subj "/C=US/ST=California/L=Walnut Creek/O=CardHero/CN=cardhero.awate.in" \
    -keyout /etc/nginx/ssl/server.key -out /etc/nginx/ssl/server.crt

EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]