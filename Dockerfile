FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/ignatov/boot-test

ENV USER ignatov
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET OgdUpblJosGklBeI

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://ignatov@localhost:5432/boot-test?sslmode=disable

WORKDIR /go/src/github.com/ignatov/boot-test

RUN godep go build

EXPOSE 8888
CMD ./boot-test