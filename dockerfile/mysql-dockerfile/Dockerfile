FROM mysql:8.0

COPY ./dockerfile/mysql-dockerfile/healthcheck.sh /usr/local/bin
RUN chmod +x /usr/local/bin/healthcheck.sh

HEALTHCHECK --interval=30s --timeout=10s CMD /usr/local/bin/healthcheck.sh