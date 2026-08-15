FROM golang:1.25-alpine

WORKDIR /var/www/html

RUN apk add --no-cache make

CMD ["sh"]
