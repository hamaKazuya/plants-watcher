FROM golang:1.13
LABEL maintainer "[hamanKazuya]"

WORKDIR /go/src
COPY go.mod .
COPY go.sum .
ENV GO111MODULE=on
RUN go mod download
EXPOSE 8088

CMD ["go", "run", "/go/src/main.go"]
