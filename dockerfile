FROM golang:alpine as builder 

RUN go get -d -v
RUN go build -o service-discvoery main.go

