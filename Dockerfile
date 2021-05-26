FROM golang:latest as development

WORKDIR /app
VOLUME /app
ADD . /app

ENV GO_ENV=development

RUN apt-get update && apt-get -y install default-mysql-client
RUN go get github.com/cosmtrek/air && go get bitbucket.org/liamstask/goose/cmd/goose

CMD ["./entry.sh"]
