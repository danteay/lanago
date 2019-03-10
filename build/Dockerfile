FROM golang:1.12

ENV CONFIG_FILE /app/properties.env
ENV GIN_MODE release
ENV PORT 8080

ADD out/api /app/api
ADD ./prod.env ${CONFIG_FILE}

RUN chmod +x /app/api

WORKDIR /app

EXPOSE ${PORT}

CMD [ "/app/api" ]