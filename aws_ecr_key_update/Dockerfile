FROM ubuntu:16.04

USER root
ENV PATH $PATH:/usr/local/go/bin

RUN apt update && apt upgrade -y && apt install python3-pip wget -y && \
  pip3 install -U pip && pip install -U awscli

WORKDIR /opt
RUN wget https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go1.10.1.linux-amd64.tar.gz && \
  rm -rf go1.10.1.linux-amd64.tar.gz

ADD aws_ecr_key_update/run.go /opt/aws_ecr_key_update/run.go

WORKDIR /opt/aws_ecr_key_update
RUN go build

WORKDIR /opt
CMD ["/opt/aws_ecr_key_update/aws_ecr_key_update"]
