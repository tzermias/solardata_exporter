FROM busybox:latest

COPY ./bin/solardata_exporter /bin/solardata_exporter

EXPOSE 9101
USER nobody
ENTRYPOINT [ "/bin/solardata_exporter" ]
