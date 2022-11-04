FROM golang:1.19-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./dist/se214 ./cmd

RUN ln -s /app/dist/se214 /bin/se214

EXPOSE 8080

CMD ["go", "run", "./cmd", "server", "start"]