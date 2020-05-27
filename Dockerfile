FROM golang:1.14.3 as builder

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -v -o golang-hello

FROM gcr.io/distroless/base
COPY --from=builder /app/golang-hello /golang-hello

CMD ["/golang-hello"]
