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


var lblIPaddr *Node
var editIPaddr *Node
var btnSendTrap *Node
var lblPortGet *Node
var editPortGet *Node
var lblPortTrap *Node
var editPortTrap *Node
var lblOID *Node
var editOID *Node
var btnSendHelp *Node
var lblValue *Node
var editValue *Node
var lblFontTest *Node
var lblFontTest2 *Node
var btnSendGet *Node
var btnSet *Node
var btnBrowser *Node
var btnTrapServer *Node
var cbxVersion1 *Node
var cbxVersion2 *Node
var cbxVersion3 *Node


func startSNMP(frmMain *Node){
	
	frmMain.obj.(*tForm).x = 190
	frmMain.obj.(*tForm).y = 70
	frmMain.obj.(*tForm).sizeX = 550
	frmMain.obj.(*tForm).sizeY = 300
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	
	lblIPaddr = CreateLabel(frmMain, "lblIPaddr", 12, 32, 120, 20, 0xD8DCC0, 0x000000, "IP address", nil)
	editIPaddr = CreateEdit(frmMain, "editIPaddr", 100, 30, 100, 20, 0xF8FCF8, 0x000000, "127.0.0.1", nil, nil)
	
	lblPortGet = CreateLabel(frmMain, "lblPortGet", 12, 70, 120, 20, 0xD8DCC0, 0x000000, "Port Get", nil)
	editPortGet = CreateEdit(frmMain, "editPortGet", 100, 68, 100, 20, 0xF8FCF8, 0x000000, "161", nil, nil)
	
	lblPortTrap = CreateLabel(frmMain, "lblPortTrap", 220, 70, 120, 20, 0xD8DCC0, 0x000000, "Port Trap", nil)
	editPortTrap = CreateEdit(frmMain, "editPortTrap", 300, 68, 100, 20, 0xF8FCF8, 0x000000, "9161", nil, nil)
	
	btnTrapServer = CreateBtn(frmMain, "btnTrapServer", 300, 30, 100, 24, 0xD8DCC0, 0x000000, "Run Server", btnTrapServerClick)
	
	
	btnSendGet = CreateBtn(frmMain, "btnSendGet", 50, 100, 70, 24, 0xD8DCC0, 0x000000, "Get", btnSendGetClick)
	btnSet = CreateBtn(frmMain, "btnSet", 50 + 80, 100, 70, 24, 0xD8DCC0, 0x000000, "Set", btnSetClick)
	btnSendTrap = CreateBtn(frmMain, "btnSendTrap", 300, 100, 70, 24, 0xD8DCC0, 0x000000, "Trap", btnSendTrapClick)
	
	
	lblOID = CreateLabel(frmMain, "lblOID", 165, 170, 120, 20, 0xD8DCC0, 0x000000, "OID", nil)
	editOID = CreateEdit(frmMain, "editOID", 200, 170, 200, 20, 0xF8FCF8, 0x000000, "1.3.6.1.2.1.1.4.0", nil, nil)
		
	lblValue = CreateLabel(frmMain, "lblValue", 150, 200, 120, 20, 0xD8DCC0, 0x000000, "Value", nil)
	editValue = CreateEdit(frmMain, "editValue", 200, 200, 200, 20, 0xF8FCF8, 0x000000, "", nil, nil)

	
	btnSendHelp = CreateBtn(frmMain, "btnSendHelp", 470, 270, 70, 24, 0xD8DCC0, 0x000000, "Help", btnSendHelpClick)
	btnBrowser = CreateBtn(frmMain, "btnBrowser", 12, 270, 90, 24, 0xD8DCC0, 0x000000, "Browser", btnBrowserClick)
	
	cbxVersion1 = CreateCheckBox(frmMain, "cbxVersion1", 430, 30, 140, 16, 0xD8DCC0, 0x000000, "Version 1", false, cbxVersion1Click)
	cbxVersion2 = CreateCheckBox(frmMain, "cbxVersion2", 430, 60, 140, 16, 0xD8DCC0, 0x000000, "Version 2", true, cbxVersion2Click)
	cbxVersion3 = CreateCheckBox(frmMain, "cbxVersion3", 430, 90, 140, 16, 0xD8DCC0, 0x000000, "Version 3", false, cbxVersion3Click)

	//lblFontTest = CreateLabel(frmMain, "lblFontTest", 12, 200, 500, 20, 0xD8DCC0, 0x000000, "abcdefghijklmnopqrstuvwxyz !\"#$%&'()*+,-./ :;<=>?@ [\\]^_`  {|}~", nil)
	//lblFontTest2 = CreateLabel(frmMain, "lblFontTest2", 12, 230, 500, 20, 0xD8DCC0, 0x000000, "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", nil)
}


func cbxVersion1Click(node *Node){
	node.obj.(*tCheckBox).checked = true
	cbxVersion2.obj.(*tCheckBox).checked = false
	cbxVersion3.obj.(*tCheckBox).checked = false
}


func cbxVersion2Click(node *Node){
	node.obj.(*tCheckBox).checked = true
	cbxVersion1.obj.(*tCheckBox).checked = false
	cbxVersion3.obj.(*tCheckBox).checked = false
}


func cbxVersion3Click(node *Node){
	node.obj.(*tCheckBox).checked = true
	cbxVersion1.obj.(*tCheckBox).checked = false
	cbxVersion2.obj.(*tCheckBox).checked = false
}


func btnSendTrapClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://localhost:8087/api?cmd=trap_v2 "  + editIPaddr.obj.(*tEdit).text + " " + editPortGet.obj.(*tEdit).text + " " + editOID.obj.(*tEdit).text + " " + editValue.obj.(*tEdit).text).Get("response").String()
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
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://localhost:8087/api?cmd=get_param " + editIPaddr.obj.(*tEdit).text + " " + editPortGet.obj.(*tEdit).text + " " + editOID.obj.(*tEdit).text).Get("response").String()
	fmt.Println("Responsed: ", result)
	
	result = strings.Replace(result, "\n", string(13), -1)
	
	memTerminal.obj.(*tMemo).text = result
}


func btnSetClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=curl_get " + "http://localhost:8087/api?cmd=set " + editIPaddr.obj.(*tEdit).text + " " + editPortGet.obj.(*tEdit).text + " " + editOID.obj.(*tEdit).text + " " + editValue.obj.(*tEdit).text).Get("response").String()
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
