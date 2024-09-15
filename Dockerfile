FROM debian:stable-slim

ARG BIN_DIR=bin
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    apt-get clean all && \
    ln -sf /dev/null /dev/log
COPY $BIN_DIR/solardata_exporter /bin

EXPOSE 9101
USER nobody
ENTRYPOINT [ "/bin/solardata_exporter" ]
