FROM ubuntu:18.04

ARG WORKDIR=/opt/test

WORKDIR $WORKDIR

COPY ./bin/test_linux ./test

RUN chmod +x ./test

CMD ["/opt/test/test"]