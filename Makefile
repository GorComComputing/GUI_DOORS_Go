.PHONY: main
main: *.go deps
	#GOOS=js GOARCH=wasm go build -o GUI.wasm . 
	tinygo build -o DOORS.wasm -target wasm .
	cp DOORS.wasm ./Server/www
	#scp DOORS.wasm user@172.18.0.1:~/WORK/DOORS/www
	rm ./DOORS.wasm


.PHONY:deps
deps:
	#cp /usr/lib/go-1.18/misc/wasm/wasm_exec.js .   #$(go env GOROOT)
	#wget https://raw.githubusercontent.com/tinygo-org/tinygo/v0.19.0/targets/wasm_exec.js
	#go get marwan.io/wasm-fetch








