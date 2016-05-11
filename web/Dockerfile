FROM golang:1.6.2-wheezy
MAINTAINER Joe Kratzat "joe.kratzat@gmail.com"

# Setup non root user
RUN useradd --user-group --create-home --shell /bin/false demo
ENV HOME=/home/demo

# Copy the local package files to the container's workspace.
ADD . /home/demo/go/src/web
RUN chown -R demo:demo $HOME/*

# Change user
USER demo

ENV GOPATH=/home/demo/go
ENV PATH=$PATH:/home/demo/go/bin

WORKDIR /home/demo/go/src/web

EXPOSE 8080

RUN go get github.com/jinzhu/gorm
RUN go get github.com/joho/godotenv
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get bitbucket.org/liamstask/goose/cmd/goose

RUN go build -o main .
CMD ["/home/demo/go/src/web/main"]
