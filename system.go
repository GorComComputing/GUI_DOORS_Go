package main

import (
	//"fmt"
    "syscall/js"
    "strings"
    "strconv"
)


type tFileType int
const (
	fBIN tFileType = iota
    fUTF8 	
)

func WriteFile(name string, str string, fileType tFileType) string {
	var result string = ""
	//fmt.Println("Before: ", []byte(str))
	switch fileType {
	case fBIN:
		var tmp string = ""
		for i := 0; i < len(str); i++ {
			tmp += strconv.Itoa(int(str[i])) + " "
		}
		result = js.Global().Call("HttpRequest", "http://"+ServerIP+":8081/savebyte?name=" + name, tmp).Get("response").String()
	case fUTF8:
		result = js.Global().Call("HttpRequest", "http://"+ServerIP+":8081/saveutf8?name=" + name, str).Get("response").String()
	} 
	//fmt.Println("Responsed: ", result)
	//fmt.Println("Responsed: ", []byte(result))
	return result
}



func ReadFileByte(name string) []byte {
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":8081/api?cmd=read_byte " + name, "").Get("response").String() 
		
	//fmt.Println(result)
	
	strArr := strings.Split(result, " ")
	
	tmp := make([]byte, 0)
	for i := 0; i < len(strArr); i++ {
		asci,_ := strconv.Atoi(strArr[i]) 
		tmp = append(tmp, byte(asci)) 
	}

	//fmt.Println(tmp)

	//fmt.Println("Responsed: ", result)
	return tmp
}


func ReadFileUTF8(name string) string {
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":8081/api?cmd=read_utf8 " + name, "").Get("response").String()
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
			typCat = ".dor"
		} else if strings.HasSuffix(row[8], ".go") && row[0][0] != byte('d') {
			typCat = ".go"
		} else if strings.HasSuffix(row[8], ".c") && row[0][0] != byte('d') {
			typCat = ".c"
		} else if (strings.HasSuffix(row[8], ".html") || strings.HasSuffix(row[8], ".htm")) && row[0][0] != byte('d') {
			typCat = ".html"
		} else if strings.HasSuffix(row[8], ".asm") && row[0][0] != byte('d') {
			typCat = ".asm"
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


func drawDo() {
	js.Global().Call("drawDo")
}
