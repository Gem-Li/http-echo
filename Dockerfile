FROM alpine:3.22.1

USER root
RUN echo 'work:x:1000:1000::/home/work:/bin/sh' >> /etc/passwd
RUN echo 'work:x:1000:' >> /etc/group
RUN mkdir -p /home/work

COPY http-echo-linux-amd64 /home/work/bin/
RUN chown -R work:work /home/work

ENTRYPOINT [ "/home/work/bin/http-echo-linux-amd64"]