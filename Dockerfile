FROM golang:1.19 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api

RUN go install github.com/go-delve/delve/cmd/dlv
RUN go build -gcflags="all=-N -l" -o main
CMD ./main
