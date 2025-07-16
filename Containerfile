FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache make

COPY go.mod go.sum .

RUN go mod download

COPY . .

CMD ["sh"]