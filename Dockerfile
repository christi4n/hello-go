FROM golang:1.19.4-alpine as build

RUN mkdir /build
ADD go.mod go.sum hello.go /build/

WORKDIR /build
RUN go build

FROM alpine:3.17.0

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=build /build/hello-go /app/
COPY views/ /app/views

WORKDIR /app

EXPOSE 8181
ENTRYPOINT ["./hello-go"]
