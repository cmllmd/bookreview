FROM golang:1.16

WORKDIR /bookreview

COPY . .

RUN go get -u github.com/lib/pq

RUN go build -o main .

CMD ["./main"]