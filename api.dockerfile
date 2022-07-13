# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o eApp ./api/cmd

RUN chmod +x /app/eApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/eApp /app

CMD [ "/app/eApp" ]