FROM golang:1.8

WORKDIR /go/src/app

RUN go-wrapper download
RUN go-wrapper install
RUN govendor sync

EXPOSE 27010

CMD ["go-wrapper", "run"]
