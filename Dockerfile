FROM golang:1.8.3

WORKDIR $GOPATH/src/github.com/guusvw/darmstadt
RUN mkdir -p $GOPATH/src/github.com/guusvw/darmstadt
COPY . .
RUN go install .

CMD [ "darmstadt" ]