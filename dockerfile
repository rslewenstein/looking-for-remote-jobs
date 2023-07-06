FROM golang:1.20.5-alpine3.18

WORKDIR /go/src/looking-for-remote-jobs

COPY . .

RUN go build -o main .

CMD ["./main"]