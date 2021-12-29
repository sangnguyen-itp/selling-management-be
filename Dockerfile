FROM golang:alpine

WORKDIR /

COPY ./.env /.env
COPY ./go.mod /go.mod
COPY ./go.sum /go.sum
RUN go mod download

COPY *.go .
RUN go build -o /app

EXPOSE 8080

CMD ["/app"]



