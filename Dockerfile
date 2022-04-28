FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .

ARG opts
RUN env ${opts} go mod download

COPY . .

RUN env ${opts} go build ./*.go

WORKDIR /app
ENTRYPOINT ["./main"]