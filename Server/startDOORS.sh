#!/bin/sh

# Убедимся, что директория /DOORS/ существует
if [ -d "/DOORS/" ]; then
    cd /DOORS/
else
    echo "Error: Directory /DOORS/ does not exist."
    exit 1
fi

# Убедимся, что файл DOORS существует и является исполняемым
if [ -f "DOORS" ]; then
    chmod +x DOORS
else
    echo "Error: File DOORS does not exist in /DOORS/ directory."
    exit 1
fi

# Запускаем файл DOORS в фоновом режиме
./DOORS &
if [ $? -eq 0 ]; then
    echo "DOORS started successfully."
else
    echo "Error: Failed to start DOORS."
    exit 1
fi
 

