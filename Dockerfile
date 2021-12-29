FROM golang:1.15.14 as build

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /go/bin/app


FROM gcr.io/distroless/base
WORKDIR /
COPY --from=build /go/bin/app /

COPY ./.env .

EXPOSE 8080

ENTRYPOINT [ "/app" ]




