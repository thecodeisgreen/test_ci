FROM golang:1.13-alpine

#RUN apk add --update git

WORKDIR /go/src/test_ci

COPY functions_test.go \
  test_ci.go ./

#RUN go get -u github.com/golang/dep/...
#RUN dep ensure --vendor-only

CMD ["go", "run", "test_ci.go"]
