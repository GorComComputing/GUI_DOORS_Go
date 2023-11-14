.PHONY: main
main: *.go deps
	#GOOS=js GOARCH=wasm go build -o GUI.wasm . 
	tinygo build -o DOORS.wasm -target wasm .
	cp DOORS.wasm ./Server/www
	rm ./DOORS.wasm


.PHONY:deps
deps:

	#go get marwan.io/wasm-fetch








