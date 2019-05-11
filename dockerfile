FROM golang:latest
COPY . /usr/home/go/src/
RUN  mkdir /app
WORKDIR /app
ENV AMCO_HOME=/app
ENV GOPATH=/usr/home/go
RUN wget http://archive.ubuntu.com/ubuntu/pool/universe/libr/librdkafka/librdkafka1_0.11.5-1_amd64.deb && \
    wget http://archive.ubuntu.com/ubuntu/pool/universe/libr/librdkafka/librdkafka++1_0.11.5-1_amd64.deb && \
    wget http://archive.ubuntu.com/ubuntu/pool/universe/libr/librdkafka/librdkafka-dev_0.11.5-1_amd64.deb && \
    apt install -y  ./librdkafka1_0.11.5-1_amd64.deb && \
    apt install  -y ./librdkafka++1_0.11.5-1_amd64.deb && \
    apt install -y  ./librdkafka-dev_0.11.5-1_amd64.deb 
# RUN go build -tags prod /usr/home/go/src/AMCO
# RUN cp -R /usr/home/go/src/AMCO/serverConfig /app/serverConfig
# RUN ls -a /app/serverConfig
# RUN rm -rf /usr/home/go/src/AMCO

RUN /usr/home/go/src/KafkaMessageQ-API/script-build.sh && \
  cp /usr/home/go/src/KafkaMessageQ-API/build/KafkaMessageQ-API.tar.gz /app && \
  tar -xf /app/KafkaMessageQ-API.tar.gz
EXPOSE 7890
CMD ./AMCO
