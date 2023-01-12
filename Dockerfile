FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o ./dist/se214 ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /app/dist/se214 /bin/se214
EXPOSE 8080
CMD ["se214", "server", "start"]