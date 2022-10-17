FROM golang:1.19

WORKDIR /gor/src
ENV PATH="/gor/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]