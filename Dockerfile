FROM golang:1.13-alpine

#RUN apk add --update git

WORKDIR /go/src/test_ci

COPY functions_test.go \
  test_ci.go ./

COPY models .

#RUN go get -u github.com/golang/dep/...
#RUN dep ensure --vendor-only

RUN sed -i -- 's/__VERSION__/0.1.1/' test_ci.go

CMD ["go", "run", "test_ci.go"]
