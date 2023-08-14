.PHONY: main
main: *.go deps
	# GOOS=js GOARCH=wasm go build -o main.wasm . 
	tinygo build -o GUI.wasm -target wasm .
	cp GUI.wasm ../WebServer/www


.PHONY:deps
deps:
	# cp /usr/lib/go-1.18/misc/wasm/wasm_exec.js .   #$(go env GOROOT)
	# wget https://raw.githubusercontent.com/tinygo-org/tinygo/v0.19.0/targets/wasm_exec.js
	# go get marwan.io/wasm-fetch








