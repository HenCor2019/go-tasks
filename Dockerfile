FROM golang:1.19-buster AS build
ARG DEBUG_SERVER_PORT
ARG PORT

WORKDIR $GOPATH/app/api

# Delve for debugging and air for live reload
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY . .
RUN go build -gcflags="all=-N -l"  -o main
EXPOSE ${PORT}

CMD ["./main"]

# Entrypoints (It doesn't work)
# COPY --from=build /go/app/main ~/main
# COPY --from=build /go/bin/dlv ~/dlv
# CMD dlv --listen=:${DEBUG_SERVER_PORT} --headless=true --api-version=2 --accept-multiclient exec ./main

