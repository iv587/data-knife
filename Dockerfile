FROM busybox
MAINTAINER lijunfei developerljf2020@163.com
WORKDIR /usr/app
COPY ./start /usr/app
ADD ./web /usr/app/web
VOLUME ["/usr/app/data" , "/usr/app/logs", "/usr/app/conf"]
EXPOSE 8888
ENTRYPOINT ["./start"]