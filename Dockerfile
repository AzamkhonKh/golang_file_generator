# syntax=docker/dockerfile:1

FROM golang:alpine3.17
WORKDIR /build
COPY . .
RUN go install
RUN go build -o main main.go
CMD ["./main"]

EXPOSE 8080