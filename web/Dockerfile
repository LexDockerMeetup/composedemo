FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/web

WORKDIR /go/src/web

# Document that the service listens on port 8080.
EXPOSE 8080

RUN go get github.com/jinzhu/gorm
RUN go get github.com/joho/godotenv
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql

RUN go build -o main .
CMD ["/go/src/web/main"]
