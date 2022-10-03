FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go get -d -v

EXPOSE 3000

CMD ["go", "run", "main.go"]
