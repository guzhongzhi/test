FROM centos
WORKDIR /app
ADD /app/test ./test
ENTRYPOINT ["./test"]

