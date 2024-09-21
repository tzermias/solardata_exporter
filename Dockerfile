FROM busybox:latest

ARG BIN_DIR=bin
COPY $BIN_DIR/solardata_exporter /bin/

EXPOSE 9101
USER nobody
ENTRYPOINT [ "/bin/solardata_exporter" ]
