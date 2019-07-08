FROM golang:1.9

ENV LANG en_US.UTF-8

RUN go get github.com/codegangsta/gin \
    golang.org/x/tools/cmd/cover \
    github.com/mattn/goveralls \
    github.com/gorilla/mux \
    github.com/go-sql-driver/mysql \
    googlemaps.github.io/maps \
    github.com/kr/pretty

ENV GOPATH /go
WORKDIR /go/src/github.com/sharypovandrey/GoOrders

CMD ["gin", "-i", "run", "GoOrders"]