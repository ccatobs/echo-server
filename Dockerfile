FROM golang:1
WORKDIR /go/src/github.com/ccatp/echo-server
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install -a -v

FROM scratch
COPY --from=0 /go/bin/echo-server /
CMD ["/echo-server"]
