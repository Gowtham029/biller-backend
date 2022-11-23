FROM golang:1.19-alpine as builder
RUN apk add build-base
WORKDIR /app
ADD . /app
RUN go build -o /identity

FROM alpine:3.14
COPY --from=builder /identity /
EXPOSE 9090
CMD [ "/identity" ]
