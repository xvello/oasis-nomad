FROM golang:1.10-alpine

RUN apk add --no-cache python3 git openssh-client bash \
 && pip3 install invoke \
 && go get -u github.com/golang/dep/cmd/dep \
 && go get -u github.com/golang/lint/golint \
 && go get -u github.com/gordonklaus/ineffassign \
 && go get -u github.com/client9/misspell/cmd/misspell
