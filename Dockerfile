FROM alpine:latest

WORKDIR /app

ADD pixiv-api .

EXPOSE 9630

ENTRYPOINT ["/app/pixiv-api"]
