FROM alpine:latest
MAINTAINER ys.cho <ssalbatore@gmail.com>

ENV PATH /usr/local/bin:$PATH
RUN ulimit -n 65535
RUN ulimit -u 65535
ENV TZ=Asia/Seoul
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/locatime && echo $TZ > /etc/timezone
COPY ./heartbeat /usr/local/bin/

CMD ["heartbeat"]

EXPOSE 5000


