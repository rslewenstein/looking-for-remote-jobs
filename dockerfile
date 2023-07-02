FROM golang:1.20.5-alpine3.18

WORKDIR /go/src/looking-for-remote-jobs

COPY . .

RUN go build -o main .

# Expor a porta que o aplicativo irá utilizar (se necessário)
EXPOSE 8080

CMD ["./main"]