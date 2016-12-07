FROM golang:latest
ADD . /go/
WORKDIR /go
RUN go get gopkg.in/olivere/elastic.v5
RUN go get github.com/gorilla/mux
RUN go get github.com/fatih/structs
RUN go build -o search  main.go
EXPOSE 8000
CMD ["/go/search"]