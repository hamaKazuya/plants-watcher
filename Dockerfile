FROM golang:1.12

RUN mkdir /go
WORKDIR /go
ADD . /go

COPY go.mod .
COPY go.sum .
CMD ["go", "run", "main.go"]
