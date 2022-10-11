# Build stage
FROM golang:1.19

ENV GOPATH=/
COPY ./ ./

RUN go build -o main .
CMD ["/main"]



