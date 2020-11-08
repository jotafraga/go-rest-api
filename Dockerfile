FROM golang:1.14

WORKDIR /go/github.com/jotafraga/go-rest-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]