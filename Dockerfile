FROM busybox
EXPOSE 80

COPY Server/files /DOORS/files
COPY Server/www /DOORS/www
COPY Server/DOORSamd64 /DOORS/DOORS
COPY Server/launch.sh /DOORS/launch.sh
COPY Server/startDOORS.sh /DOORS/startDOORS.sh

COPY Server/cert.pem /etc/ssl/cert.pem
COPY Server/auto_doors /etc/init.d/auto_doors

RUN chmod +x /DOORS/DOORS
RUN chmod +x /DOORS/startDOORS.sh
RUN chmod +x /DOORS/launch.sh

WORKDIR /DOORS
CMD ["./launch.sh", "&"]


