FROM golang:1.12

ENV GIN_MODE release
ENV PORT 8080

ADD out/api /app/api
ADD ./.env /app/properties.env

RUN chmod +x /app/api

WORKDIR /app

EXPOSE ${PORT}

CMD [ "/app/api" ]