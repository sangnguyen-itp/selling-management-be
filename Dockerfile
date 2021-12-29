FROM golang:alpine

WORKDIR /

COPY ./.env .
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go .
RUN go build -o /app

EXPOSE 8080

CMD ["/app"]



