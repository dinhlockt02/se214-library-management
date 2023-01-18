FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./dist/se214 ./cmd

FROM alpine
WORKDIR /root/
COPY --from=0 /app/dist/se214 /bin/se214
CMD ["se214", "server", "start"]