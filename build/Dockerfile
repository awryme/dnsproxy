FROM golang:1.22 AS build
COPY ./ /go/src/dnsproxy
WORKDIR /go/src/dnsproxy
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN go build -v -o /bin/dnsproxy ./cmd/dnsproxy

FROM alpine
COPY --from=build /bin/dnsproxy /bin/dnsproxy

ENTRYPOINT [ "/bin/dnsproxy" ]