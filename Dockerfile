FROM golang:alpine AS builder

ARG ARCH
ARG OS

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN GOOS=$OS GOARCH=$ARCH go build -o app .

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]