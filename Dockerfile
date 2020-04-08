FROM golang:alpine AS builder

ARG ARCH=amd64
ARG OS=linux

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN GOOS=$OS GOARCH=$ARCH go build -o app .

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]