FROM alpine

EXPOSE 8000
EXPOSE 60000/udp
EXPOSE 60000/tcp

RUN mkdir -p /opt/uhppote
RUN mkdir -p /usr/local/etc/uhppoted/simulator

COPY 201020304.json /usr/local/etc/uhppoted/simulator
COPY 303986753.json /usr/local/etc/uhppoted/simulator
COPY 405419896.json /usr/local/etc/uhppoted/simulator

WORKDIR /opt/uhppote
COPY uhppote-simulator .

ENTRYPOINT /opt/uhppote/uhppote-simulator --debug --devices /usr/local/etc/uhppoted/simulator
