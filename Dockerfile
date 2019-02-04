FROM golang:1.11
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
#RUN go install -v ./...
RUN go build -o app cmd/main.go
#RUN go build

CMD ["./app", "-dev-port", "10200",  "-human-port", "4200", "dir", "./db"]