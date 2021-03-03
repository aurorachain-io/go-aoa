# Build AOA in a stock Go builder container
FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-aoa
RUN cd /go-aoa && make aoa

# Pull AOA into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-aoa/build/bin/aoa /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["aoa"]