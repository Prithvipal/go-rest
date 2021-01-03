FROM alpine

#WORKDIR /app

COPY todo /

CMD ["/todo"]