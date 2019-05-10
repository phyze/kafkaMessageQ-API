FROM centos:7
COPY . /app
WORKDIR /app
ENV AMCO_HOME=/app
RUN  yum install -y tar curl lsof which
RUN yum install -y http://download-ib01.fedoraproject.org/pub/epel/7/x86_64/Packages/l/librdkafka-0.11.5-1.el7.x86_64.rpm
RUN chmod a+x /app/amco
EXPOSE 7890
CMD ./amco