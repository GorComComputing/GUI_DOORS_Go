package main

import (
    "fmt"
    //"math/rand"
    //"math"
    "syscall/js"
    "strings"
    //"time"
    //"strconv"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    //"encoding/json"
)


var btnSendTrap *Node
var btnTrapServer *Node
var btnSendHelp *Node
var lblMessage *Node
var editMessage *Node
var lblFontTest *Node
var btnSendGet *Node
var btnSendGetHex *Node
var btnBrowser *Node


func startSNMP(frmMain *Node){
	//frmTrap = CreateForm(&layout, 400, 400, 200, 130, 0xD8DCC0, WIN, "Trap", false, nil)
	frmMain.obj.(*tForm).x = 100
	frmMain.obj.(*tForm).y = 50
	frmMain.obj.(*tForm).sizeX = 550
	frmMain.obj.(*tForm).sizeY = 300
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	
	btnSendTrap = CreateBtn(frmMain, 200, 30, 60, 24, 0xD8DCC0, 0x000000, "Trap", btnSendTrapClick)
	btnTrapServer = CreateBtn(frmMain, 270, 30, 60, 24, 0xD8DCC0, 0x000000, "Server", btnTrapServerClick)
		
	lblMessage = CreateLabel(frmMain, 12, 32, 120, 20, 0xD8DCC0, 0x000000, "Message", nil)
	editMessage = CreateEdit(frmMain, 80, 30, 100, 20, 0xF8FCF8, 0x000000, "", nil)

	btnSendGet = CreateBtn(frmMain, 12, 100, 60, 24, 0xD8DCC0, 0x000000, "Get", btnSendGetClick)
	btnSendGetHex = CreateBtn(frmMain, 12 + 70, 100, 70, 24, 0xD8DCC0, 0x000000, "Get Hex", btnSendGetHexClick)
	btnSendHelp = CreateBtn(frmMain, 12 + 70 + 80, 100, 60, 24, 0xD8DCC0, 0x000000, "Help", btnSendHelpClick)
	
	btnBrowser = CreateBtn(frmMain, 12, 130, 80, 24, 0xD8DCC0, 0x000000, "Browser", btnBrowserClick)

	lblFontTest = CreateLabel(frmMain, 12, 200, 500, 20, 0xD8DCC0, 0x000000, "abcdefghijklmnopqrstuvwxyz !\"#$%&'()*+,-./ :;<=>?@ [\\]^_`  {|}~", nil)
}


func btnSendTrapClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://localhost:8087/api?cmd=trap_v2 " + editMessage.obj.(*tEdit).text).Get("response").String()
	fmt.Println("Responsed: ", result)
	
	memTerminal.obj.(*tMemo).text = result
}


func btnTrapServerClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://localhost:8087/api?cmd=trap_srv").Get("response").String()
	fmt.Println("Responsed: ", result)
	
	memTerminal.obj.(*tMemo).text = result
}


func btnSendHelpClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://localhost:8087/api?cmd=.help").Get("response").String()
	fmt.Println("Responsed: ", result)
	
	result = strings.Replace(result, "\n", string(13), -1)
	
	memTerminal.obj.(*tMemo).text = result
}


func btnSendGetClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://localhost:8087/api?cmd=get").Get("response").String()
	fmt.Println("Responsed: ", result)
	
	result = strings.Replace(result, "\n", string(13), -1)
	
	memTerminal.obj.(*tMemo).text = result
}


func btnSendGetHexClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://localhost:8087/api?cmd=get_hex").Get("response").String()
	fmt.Println("Responsed: ", result)
	
	result = strings.Replace(result, "\n", string(13), -1)
	
	memTerminal.obj.(*tMemo).text = result
}


func btnBrowserClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://info.cern.ch/hypertext/WWW/TheProject.html").Get("response").String()
	fmt.Println("Responsed: ", result)
	
	result = strings.Replace(result, "\n", string(13), -1)
	
	memTerminal.obj.(*tMemo).text = result
}
