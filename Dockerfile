FROM busybox
EXPOSE 8081
COPY Server /Server
COPY Server/cert.pem /etc/ssl/cert.pem
WORKDIR /Server
CMD ["./DOORSamd64", "&"]
