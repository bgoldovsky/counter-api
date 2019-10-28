FROM golang

ADD . /go/src/github.com/bgoldovsky/counter-api

ENV PORT=8085
ENV EXPIRES=60
ENV STORE_PATH=./store.gob

RUN go install /go/src/github.com/bgoldovsky/counter-api/cmd/counter-api

ENTRYPOINT /go/bin/counter-api

EXPOSE 8085