package main

import (
    "fmt"
    //"math/rand"
    //"math"
    "syscall/js"
    //"time"
    //"strconv"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    //"encoding/json"
)


//var frmTrap *Node
var btnSendTrap *Node
//var btnCancel *Node
//var btnOther *Node
var lblMessage *Node
//var lblPswd *Node
var editMessage *Node
//var editPswd *Node


func startTrap(frmMain *Node){
	//frmTrap = CreateForm(&layout, 400, 400, 200, 130, 0xD8DCC0, WIN, "Trap", false, nil)
	frmMain.obj.(*tForm).x = 400
	frmMain.obj.(*tForm).y = 400
	frmMain.obj.(*tForm).sizeX = 200
	frmMain.obj.(*tForm).sizeY = 130
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	
	btnSendTrap = CreateBtn(frmMain, 70, 20 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "Send", btnSendTrapClick)
	//btnCancel = CreateBtn(frmTrap, 40 + 70, 20 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "CANCEL", nil)
	//btnOther = CreateBtn(frmTrap, 80 + 60, 20 + 30 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "OTHER", nil)
	
	lblMessage = CreateLabel(frmMain, 12, 32, 120, 20, 0xD8DCC0, 0x000000, "Message", nil)
	//lblPswd = CreateLabel(frmTrap, 12, 22 + 30, 80, 20, 0xD8DCC0, 0x000000, "PASSWORD", nil)
	
	editMessage = CreateEdit(frmMain, 80, 30, 100, 20, 0xF8FCF8, 0x000000, "", nil)
	//editPswd = CreateEdit(frmTrap, 80, 20 + 30, 80, 20, 0xF8FCF8, 0x000000, "PSWD", nil)
}


func btnSendTrapClick(node *Node){
	result := js.Global().Call("HttpRequest", "http://localhost:8087/api?cmd=trap_v2 " + editMessage.obj.(*tEdit).text).Get("response").String()
	fmt.Println("Responsed: ", result)
}
