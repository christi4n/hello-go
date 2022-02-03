# Basic GoLang server

## Plain HTML response (with already defined route)

    http://localhost:8181/hello

## Plain HTML with template

    http://localhost:8181/

## JSON response

    curl --location --request GET 'http://localhost:8181/api/ping'

## Build a binary

    go build
    ./hello-go

## Build a Docker image

    nerdctl build -t hello-go .
    nerdctl images hello-go

REPOSITORY    TAG       IMAGE ID        CREATED          SIZE
hello-go      latest    e3fcd9cb6639    2 minutes ago    14.9 MiB

## Run the image with nerdctl container runtime

    nerdctl run --rm -p 8181:8181 hello-go
