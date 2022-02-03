# Basic GoLang server

## Plain HTML response (with already defined route)

    http://localhost:8181/hello

## Plain HTML with template

    http://localhost:8181/

## JSON response

    curl --location --request GET 'http://localhost:8181/api/ping'

## Build a binary

    go build