FROM golang:1.11.1 as builder
WORKDIR /go/src/github.com/lbrindze/dyndns-util
COPY ./ /go/src/github.com/lbrindze/dyndns-util

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -a -ldflags '-w' -o ./bin/dyndns-util

FROM aarch64/busybox 
LABEL authors="Loren Brindze"

COPY --from=builder /go/src/github.com/lbrindze/dyndns-util /bin/dyndns-util

ENTRYPOINT ["/bin/dyndns-util"]
