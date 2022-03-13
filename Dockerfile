FROM centos
WORKDIR /app
ADD test
ENTRYPOINT ["./test"]

