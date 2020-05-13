FROM golang:latest

ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPATH=/go \
    PATH=$PATH:$GOPATH/bin

RUN go get github.com/nigi4/fish-farm

WORKDIR /go/src/github.com/nigi4/fish-farm

COPY *.lock ./
COPY *.toml ./

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure 

COPY . .

RUN rm -rf bin
RUN go build -o bin/fish-farm .

CMD ["./bin/fish-farm"]