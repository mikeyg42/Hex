FROM golang:1.19-alpine

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy

CMD ["air", "-c", ".air.toml"]