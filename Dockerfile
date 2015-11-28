FROM golang:1.5.1
MAINTAINER sctlee "sctlee221@gmail.com"
COPY . /go/src/github.com/sctlee/hazel/
RUN go get gopkg.in/yaml.v2
RUN go get github.com/jackc/pgx
RUN go get github.com/garyburd/redigo/redis
RUN go get github.com/nu7hatch/gouuid
RUN go install github.com/sctlee/hazel
