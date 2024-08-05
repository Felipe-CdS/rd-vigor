FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN ["templ", "generate"]

RUN go build -o ./tmp/main ./cmd/

CMD ["./tmp/main"]
