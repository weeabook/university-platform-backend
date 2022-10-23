FROM golang:latest


RUN go version
ENV GOPATH=/

COPY . ./

RUN go build -o main ./src/cmd/main.go

CMD ["./main"]