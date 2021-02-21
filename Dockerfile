FROM ubuntu:18.04

RUN apt update && apt install -y --no-install-recommends wget git ca-certificates

RUN wget --no-check-certificate -O go.tar.gz https://dl.google.com/go/go1.15.7.linux-amd64.tar.gz \
&& tar -C /usr/local -xzf go.tar.gz \
&& chmod -R 755 /usr/local/go

COPY entrypoint.sh /usr/local/bin/entrypoint.sh

RUN chmod 777 /usr/local/bin/entrypoint.sh \
    && ln -s /usr/local/bin/entrypoint.sh /

ENTRYPOINT [ "bash", "entrypoint.sh" ]