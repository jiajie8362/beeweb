FROM golang

ADD . /go/src/beeweb
WORKDIR /go/src/beeweb
RUN go get github.com/tools/godep
RUN godep restore
RUN go install beeweb
ENTRYPOINT /go/bin/beeweb
EXPOSE 8080