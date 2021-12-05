#CGO_ENABLED=1 GOOS=linux go build -ldflags '-extldflags "-static"' -o rula
FROM alpine:dec2021
RUN apk --no-cache add curl
COPY ./rula /rula
COPY ./.env /.env
EXPOSE 3333/tcp
CMD ["/rula"]