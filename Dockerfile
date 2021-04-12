FROM  alpine:latest

WORKDIR /root/

COPY ./.bin .
COPY ./configs .

EXPOSE 80

CMD ["./bot"]