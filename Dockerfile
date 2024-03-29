FROM golang:1.19 AS builder
WORKDIR /code
COPY . .
RUN GOOS=linux go build -o gin-demo

FROM alpine:3.16

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache bash busybox-extras curl net-tools tcpdump nmap tzdata libc6-compat

COPY --from=builder /code/gin-demo /usr/bin/
ENV TZ=Asia/Shanghai

EXPOSE 8080

CMD ["gin-demo"]