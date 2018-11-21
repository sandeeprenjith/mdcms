FROM golang
MAINTAINER sandeep.renjith@gmail.com
RUN go get github.com/sandeeprenjith/mdcms
RUN go build -o /go/src/github.com/sandeeprenjith/mdcms/runserver github.com/sandeeprenjith/mdcms/ 
CMD bash -c "cd /go/src/github.com/sandeeprenjith/mdcms/;./runserver"

