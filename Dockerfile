FROM busybox:latest

ARG BIN_DIR=bin
COPY $BIN_DIR/solardata_exporter /usr/bin

EXPOSE 9101
USER nobody
ENTRYPOINT [ "/usr/bin/solardata_exporter" ]
