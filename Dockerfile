FROM golang:1.23.4-alpine3.20

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY .air.toml ./

COPY go.mod go.sum ./

ENV GOPROXY=https://goproxy.io,direct

RUN go mod download && go mod verify

COPY . .

CMD ["air", "-c", ".air.toml"]