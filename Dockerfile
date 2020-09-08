FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /Users/lx/go-application/go-gin-example
COPY . /Users/lx/go-application/go-gin-example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]