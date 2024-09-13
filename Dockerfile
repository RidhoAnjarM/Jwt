FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy
RUN go build -o app

ENTRYPOINT ["/app/app"]