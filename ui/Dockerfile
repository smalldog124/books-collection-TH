FROM nginx:1.15.7-alpine

ENV NGINX_ROOT /server
RUN mkdir -p ${NGINX_ROOT}

# Set working directory
WORKDIR ${NGINX_ROOT}

# Override Caddy settings
COPY ./nginx.conf /etc/nginx/conf.d/books.conf
COPY ./public/ /usr/share/nginx/html/

RUN rm /etc/nginx/conf.d/default.conf

EXPOSE 80