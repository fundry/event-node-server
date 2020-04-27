# small image to reduce resources
FROM golang:alpine

ENV DIR = $GOPATH/src/github.com/vickywane/event-server/
WORKDIR $DIR

COPY go.mod .
COPY go.sum .
COPY . $DIR 

RUN go build -o server . 

WORKDIR /dist

COPY /build/server .

FROM scratch

COPY --from=builder /dist/server /

#My Playground Port
EXPOSE 8080

RUN go run server.go

ENTRYPOINT ["/server"]
#CMD ["/dist/server"]


