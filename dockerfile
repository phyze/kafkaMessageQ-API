FROM golang:latest
COPY . /usr/home/go/src/KafkaMessageQ-API
WORKDIR /app
ENV GOPATH=/usr/home/go
ENV AMCO_HOME=/app
RUN wget http://archive.ubuntu.com/ubuntu/pool/universe/libr/librdkafka/librdkafka1_0.11.5-1_amd64.deb && \
    wget http://archive.ubuntu.com/ubuntu/pool/universe/libr/librdkafka/librdkafka++1_0.11.5-1_amd64.deb && \
    wget http://archive.ubuntu.com/ubuntu/pool/universe/libr/librdkafka/librdkafka-dev_0.11.5-1_amd64.deb && \
    apt install -y  ./librdkafka1_0.11.5-1_amd64.deb && \
    apt install  -y ./librdkafka++1_0.11.5-1_amd64.deb && \
    apt install -y  ./librdkafka-dev_0.11.5-1_amd64.deb 
RUN cd /usr/home/go/src/KafkaMessageQ-API && \
  ./script-build.sh -env prod && \
  cp build/KafkaMessageQ-API.tar.gz /app && \
  cd /app && \
  tar -xf /app/KafkaMessageQ-API.tar.gz && \
  rm -rf /usr/home/go/src/*
EXPOSE 7890
CMD ./KafkaMessageQ-API
