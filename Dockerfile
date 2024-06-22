# Используем минималистичный базовый образ busybox
FROM busybox

# Указываем, что контейнер будет прослушивать порт 80
EXPOSE 80

# Копируем необходимые файлы и директории в контейнер
COPY Server/files /DOORS/files
COPY Server/www /DOORS/www
COPY Server/DOORSamd64 /DOORS/DOORS
COPY Server/launch.sh /DOORS/launch.sh
COPY Server/startDOORS.sh /DOORS/startDOORS.sh
COPY Server/cert.pem /etc/ssl/cert.pem
COPY Server/auto_doors /etc/init.d/auto_doors

# Устанавливаем права на выполнение для определенных файлов
RUN chmod +x /DOORS/DOORS /DOORS/startDOORS.sh /DOORS/launch.sh

# Устанавливаем рабочую директорию
WORKDIR /DOORS

# Определяем команду для запуска скрипта launch.sh при старте контейнера
CMD ["./launch.sh"]
