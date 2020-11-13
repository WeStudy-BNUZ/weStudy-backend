FROM golang:1.6

RUN go get github.com/kataras/iris

RUN apk update --no-cache

EXPOSE 8080

RUN mkdir /main
COPY main /main/

WORKDIR /main
ENTRYPOINT ["./main"]
