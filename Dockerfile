FROM golang:1.22rc2-alpine3.18

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

CMD ["go", "run", "./main.go"]
