

FROM golang:alpine as gobuild

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .



FROM alpine

RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=gobuild /build/main /app/
WORKDIR /app

CMD ["./main"]

