FROM golang:1.22rc2-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download

COPY *.go ./

EXPOSE 8080

CMD ["go", "run", "./main.go"]
