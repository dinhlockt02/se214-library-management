FROM nginx:alpine
COPY docker/nginx/nginx.conf /etc/nginx/
COPY ./docs /docs
EXPOSE 8080