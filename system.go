package main

import (
	//"fmt"
    "syscall/js"
    "strings"
)


func WriteFile(name string, str string) string {
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":8081/save?name=" + name, str).Get("response").String() 
	//fmt.Println("Responsed: ", result)
	return result
}



func ReadFile(name string) string {
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":8081/api?cmd=read " + name, "").Get("response").String() 
	//fmt.Println("Responsed: ", result)
	return result
}


type Catalog struct {
	name string 
    typ string
}

func GetCatalogList(name string) []Catalog {
	output := make([]Catalog, 0)
	var typCat string

	//name = strings.Replace(name, ".", "%2e", -1)
	
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":8081/api?cmd=ls -l " + name, "").Get("response").String() 
	words := strings.Split(result, "\n")

	for i := 1; i < len(words)-1; i++ {
		row := strings.Fields(words[i])
		if row[0][0] == byte('d') {
			typCat = "D"
		} else if row[0][0] == byte('-') && row[0][9] == byte('x') {
			typCat = "X"
		} else if strings.HasSuffix(row[8], ".dor") && row[0][0] != byte('d') {
			typCat = "B"
		} else if row[0][0] == byte('-') && row[0][9] == byte('-') {
			typCat = "F"
		} 
		cat := Catalog{row[8], typCat}
		output = append(output, cat)
	}
	//fmt.Println("Responsed: ", output)
	return output
}


func Get(url string, s string, body string) string {
	if s != "" {
		url += "?" + s
	}
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":8081/api?cmd=curl_get " + url, body).Get("response").String()
	//fmt.Println("Responsed: ", result)
	return result
}
