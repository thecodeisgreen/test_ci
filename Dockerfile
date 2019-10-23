FROM golang:1.13-alpine

#RUN apk add --update git

WORKDIR /go/src/test_ci

COPY \
  functions_test.go \
  test_ci.go ./

COPY models ./models
COPY tcig.io ./tcig.io

#RUN go get -u github.com/golang/dep/...
#RUN dep ensure --vendor-only

RUN sed -i -- "s/__VERSION__/${BUILD_VERSION}/" tcig.io/version/version.go 
RUN sed -i -- "s/__DATE__/$(date)/" tcig.io/version/version.go

CMD ["go", "run", "test_ci.go"]
