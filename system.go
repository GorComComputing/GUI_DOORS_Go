package main

import (
	"fmt"
    "syscall/js"
    "strings"
    "strconv"
    "marwan.io/wasm-fetch"
    //"golang.org/x/image/bmp"
	//"image/jpeg"
//	"image"
//	"bytes"
	//"io"
	//"bufio"
	//"io/ioutil"
	//asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v22"
)


type tFileType int
const (
	fBIN tFileType = iota
    fUTF8 	
)

func WriteFile(name string, str string, fileType tFileType) string {
	if mainUser != loginAdmin {
		fmt.Println("Access denied!")
		return "Access denied!"
	}
	var result string = ""
	//fmt.Println("Before: ", []byte(str))
	switch fileType {
	case fBIN:
		var tmp string = ""
		for i := 0; i < len(str); i++ {
			tmp += strconv.Itoa(int(str[i])) + " "
		}
		result = js.Global().Call("HttpRequest", "http://"+ServerIP+":"+ServerPort+"/savebyte?name=" + name, tmp).Get("response").String()
	case fUTF8:
		result = js.Global().Call("HttpRequest", "http://"+ServerIP+":"+ServerPort+"/saveutf8?name=" + name, str).Get("response").String()
	} 
	//fmt.Println("Responsed: ", result)
	//fmt.Println("Responsed: ", []byte(result))
	return result
}



func ReadFileByte(name string) []byte {
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":"+ServerPort+"/api?cmd=read_byte " + name, "").Get("response").String() 
		
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
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":"+ServerPort+"/api?cmd=read_utf8 " + name, "").Get("response").String()
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
	
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":"+ServerPort+"/api?cmd=ls -l " + name, "").Get("response").String() 
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
	result := js.Global().Call("HttpRequest", "http://"+ServerIP+":"+ServerPort+"/api?cmd=curl_get " + url, body).Get("response").String()
	//fmt.Println("Responsed: ", result)
	return result
}


func drawDo() {
	js.Global().Call("drawDo")
}


// Получает файл по HTTP
func fetchFile(path string) []byte {
	//n := map[string]string{"Content-Type": "image/bmp"}
	resp, err := fetch.Fetch(path, &fetch.Opts{
		//Headers:   n,
	})
	if err != nil {
		fmt.Println("NILL ERR")
	}
    return resp.Body
}


// Получает файл JPEG по HTTP
/*func fetchFileJPEG(path string) []byte {
	//n := map[string]string{"Content-Type": "image/jpeg"}
	resp, err := fetch.Fetch(path, &fetch.Opts{
		//Headers:   n,
	})
	if err != nil {
		fmt.Println("NILL JPEG ERR")
	}
	
	
	var src image.Image
	src, err = jpeg.Decode(bytes.NewReader(resp.Body))
	if err != nil {
		panic(err.Error())
	}
	
	var bmpFile []byte
	
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, src, nil)
	bmpFile = buf.Bytes()
*/	
/*	
	var w io.Writer
	//bmpFile := bufio.NewWriter(bmpFile)
	err = bmp.Encode(w, src)
	if err != nil {
		panic(err.Error())
	}
	
	bmpFile, err = ioutil.ReadAll(w)
	//bmpFile := bufio.NewWriter(bmpFile)
	
	*/
 /*   return bmpFile
}*/

/*
func testJPEG(path string) []byte {
	//n := map[string]string{"Content-Type": "image/jpeg"}
	resp, err := fetch.Fetch(path, &fetch.Opts{
		//Headers:   n,
	})
	if err != nil {
		fmt.Println("NILL JPEG ERR")
	}
	
	fmt.Println("WERE 1")
	fmt.Println(resp.Body)
	fmt.Println("WERE 2")
	imData, imType, err := image.Decode(bytes.NewReader(resp.Body))
    if err != nil {
        fmt.Println(err)
    }
 
    fmt.Println(imData)
    fmt.Println("WERE 3")
    fmt.Println(imType)
    
    return resp.Body
}*/
