FROM golang:1.16.2-alpine3.13 AS builder

WORKDIR /project
COPY . .
RUN go mod download

RUN go build -o tls-exporter .


FROM alpine:3.13
WORKDIR /project
COPY --from=builder /project/tls-exporter /project/tls-exporter 

EXPOSE 80
ENTRYPOINT [ "/project/tls-exporter" ]