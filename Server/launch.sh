#!/bin/sh

# Этот скрипт написан на языке оболочки (shell script) и предназначен 
# для проверки и перезапуска сервера "DOORS", если он не работает.

#echo "DOORS Server upper is START:"
#echo "-----------------------------------------------------"

while true
do
ret=$(ps | grep [D]OORS | wc -l)				# Проверка количества процессов "DOORS"
	if [ "$ret" -eq 0 ]							# Если процессов "DOORS" нет
	then {
		#echo "DOWN: Server is DOWN:       "
		killall DOORS							# Завершение всех процессов "DOORS"
		chmod +x /DOORS/DOORS 					# Установка права на выполнение для файла
		chmod +x /DOORS/startDOORS.sh			# Установка права на выполнение для файла
        sleep 1  								# Ожидание 1 секунду 
        echo "UP1:   Start  SERVER:        "
        /DOORS/startDOORS.sh > /dev/null 2>&1	# Запуск скрипта
        echo "UP2:   Start  SERVER:        "
        #/Server/DOORS
        #echo "UP:   Start  SERVER:        "	
		#exit 1
	}
	#else 
	#{
		#echo "DOORS already running!"
		#exit 1
	#}
	fi;
	
	sleep 1										# Ожидание 1 секунду перед следующей проверкой

done




