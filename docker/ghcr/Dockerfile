FROM alpine

LABEL org.opencontainers.image.description Simulator for UHPPOTE access controllers

EXPOSE 8000
EXPOSE 60000/udp

RUN mkdir -p /opt/uhppoted
RUN mkdir -p /var/uhppoted/simulator

COPY 405419896.json /var/uhppoted/simulator

WORKDIR /opt/uhppote
COPY uhppote-simulator .

ENTRYPOINT /opt/uhppote/uhppote-simulator --debug --devices /var/uhppoted/simulator