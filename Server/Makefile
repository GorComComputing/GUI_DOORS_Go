.PHONY: main
main: *.go deps
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o DOORS .
	#GOOS=linux GOARCH=arm go build -o DOORS .
	GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o DOORS .


.PHONY:deps
deps:
#	go get github.com/gorilla/sessions
#	go get github.com/kabukky/httpscerts







