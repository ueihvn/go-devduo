FROM golang:1.16.5-alpine3.13

RUN mkdir /go-devduo

ADD . /go-devduo

WORKDIR /go-devduo

RUN go build -o go-devduo .

CMD ["./go-devduo"]