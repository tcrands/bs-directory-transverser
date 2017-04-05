FROM golang:1.7

RUN mkdir /go-directory-transverser
RUN mkdir /source

WORKDIR /go-directory-transverser

ADD . /go-directory-transverser

RUN go build

RUN cp ./go-directory-transverser /usr/bin

WORKDIR /source

CMD ["go-directory-transverser"]
