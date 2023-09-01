package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	//"log"
	//"net/http"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
/*
func main() {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile("./ArchiumWASM_Demo_Tux.png")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)
	
	
	
	
	
	//encoded := "b25lIPCfkJggYW5kIHRocmVlIPCfkIs="
    //fmt.Println(encoded)

    data := make([]byte, base64.StdEncoding.DecodedLen(len(base64Encoding)))
    _, err = base64.StdEncoding.Decode(data, []byte(base64Encoding))

    //fmt.Println(data)

    if err != nil {
        log.Fatal(err)
    }

    //fmt.Println(string(data[:n]))
	
}
*/


func main() {
  //Читаем png файл
  pngData, err := ioutil.ReadFile("./Tree.bmp")
  if err != nil {
    //ВСЕ ПРОПАЛО!!1!!ОДИНОДИН
    panic(err)
  }	

  //Кодируем в base64
  b64String := base64.StdEncoding.EncodeToString(pngData)	
  
  fmt.Println(b64String)

  //Записываем результат кодирования
  err = ioutil.WriteFile("sample.b64", []byte(b64String), 0644)

  if err != nil {
    panic(err)
  }

  //Читаем base64
  b64Data, err := ioutil.ReadFile("sample.b64")
  if err !=nil {
    panic(err)
  }

  //Декодируем base64
  outPngData, err := base64.StdEncoding.DecodeString(string(b64Data))

  if err != nil {
    panic(err)
  }

  //Записываем результат декордирования
  err = ioutil.WriteFile("sample.out.png", outPngData, 0644)

  if err != nil {
    panic(err)
  }	
}



