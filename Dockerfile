FROM alpine:3.6

WORKDIR /usr/app

ADD .env.template .env
ADD backend-cert2addr backend-cert2addr
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE 9990
CMD /usr/app/backend-cert2addr
