FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache make

COPY go.* .

RUN go mod download

COPY . .

CMD ["sh"]
