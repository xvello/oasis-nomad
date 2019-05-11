FROM golang:1.12-alpine
ENV CGO_ENABLED=0

RUN apk add --no-cache python3 git openssh-client bash \
 && pip3 install invoke \
 && go get -u github.com/golang/dep/cmd/dep \
 && go get -u golang.org/x/lint/golint \
 && go get -u github.com/gordonklaus/ineffassign \
 && go get -u github.com/client9/misspell/cmd/misspell
