FROM golang:1.11 as build

COPY ./ /go/src/github.com/innovate-technologies/yp-rover
WORKDIR /go/src/github.com/innovate-technologies/yp-rover

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o yp-rover ./cmd/rover

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/github.com/innovate-technologies/yp-rover/yp-rover /opt/yp-rover/yp-rover

RUN chmod +x /opt/yp-rover/yp-rover

WORKDIR /opt/yp-rover/

ENTRYPOINT ["/opt/yp-rover/yp-rover"]
