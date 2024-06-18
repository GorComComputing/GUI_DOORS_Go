#!/bin/sh

#echo "DOORS Server upper is START:"
#echo "-----------------------------------------------------"

while true
do

ret=$(ps | grep [D]OORS | wc -l)
	if [ "$ret" -eq 0 ]
	then {
		#echo "DOWN: Server is DOWN:       "
		killall DOORS
		chmod +x /Server/DOORS 
		chmod +x /root/startDOORS.sh
        sleep 1  	
        /root/startDOORS.sh > /dev/null 2>&1
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
	
	sleep 1

done




