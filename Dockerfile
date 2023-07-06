FROM alpine:3.16

COPY gin-demo /usr/bin/

EXPOSE 8080

CMD ["gin-demo"]