FROM golang:1.7

COPY . /go/src/github.com/ahailu/go-helloworld
WORKDIR /go/src/github.com/ahailu/go-helloworld

RUN CGO_ENABLED=0 GOOS=$(uname -s | tr A-Z a-z) go install -ldflags "-X github.com/ahailu/go-helloworld/server.Revision=$(git rev-parse HEAD) -s" -a -installsuffix cgo github.com/ahailu/go-helloworld

EXPOSE 8080

CMD ["/go/bin/go-helloworld"]