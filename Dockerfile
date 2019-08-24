FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o mygrpcserver /app/grpc-server/server.go
CMD ["/app/mygrpcserver"]
