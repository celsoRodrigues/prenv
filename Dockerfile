FROM golang:1.19 AS binarybuilder
WORKDIR /go/src/github.com/repowatchdog
COPY . .
RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o repowatchdog-linux main.go

FROM alpine:3.10
RUN apk add --update --no-cache bash ca-certificates
WORKDIR /app
COPY --from=binarybuilder /go/src/github.com/repowatchdog/repowatchdog-linux /usr/local/bin/
EXPOSE 80
CMD ["repowatchdog-linux"]

