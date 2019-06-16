FROM golang:1.12 as builder

RUN go get -u github.com/golang/dep/cmd/dep \
 && go get -u golang.org/x/lint/golint \
 && go get -u github.com/gordonklaus/ineffassign \
 && go get -u github.com/client9/misspell/cmd/misspell \
 && go get -u github.com/docker/distribution/cmd/registry



FROM golang:1.12

WORKDIR /go/src/github.com/xvello/oasis-nomad

RUN apt-get update \
 && apt-get install --no-install-recommends -y python3-pip \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN pip3 install invoke

COPY --from=builder /go/bin/* /go/bin/
