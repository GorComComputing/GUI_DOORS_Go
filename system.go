package main

import (
	//"fmt"
    "syscall/js"
)



func WriteFile(name string, str string) string {
	result := js.Global().Call("HttpRequest", "http://localhost:8081/save?name=" + name, str).Get("response").String() 
	//fmt.Println("Responsed: ", result)
	return result
}



func ReadFile(name string) string {
	result := js.Global().Call("HttpRequest", "http://localhost:8081/api?cmd=read " + name, "").Get("response").String() 
	//fmt.Println("Responsed: ", result)
	return result
}


func GetCatalog(name string) string {
	result := js.Global().Call("HttpRequest", "http://localhost:8081/api?cmd=ls " + name, "").Get("response").String() 
	//fmt.Println("Responsed: ", result)
	return result
}


func Get(url string, s string, body string) string {
	if s != "" {
		url += "?" + s
	}
	result := js.Global().Call("HttpRequest", "http://localhost:8081/api?cmd=curl_get " + url, body).Get("response").String()
	//fmt.Println("Responsed: ", result)
	return result
}
