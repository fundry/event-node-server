# small image to reduce resources
FROM golang:alpine as Image

ENV DIR = $GOPATH/src/github.com/vickywane/event-server/
WORKDIR $DIR

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . $DIR

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

WORKDIR /dist

COPY /build/server .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=Image /server.go .
COPY --from=Image /.env .

#My Playground Port
EXPOSE 4040

RUN go run server.go

ENTRYPOINT ["/server"]
#CMD ["/dist/server"]


