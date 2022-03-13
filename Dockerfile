FROM centos
WORKDIR /app
ADD ./test /app/test
ENTRYPOINT ["./test"]

