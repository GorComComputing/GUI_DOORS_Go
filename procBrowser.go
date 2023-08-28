package main

import (
	    "strings"
)

var btnGo *Node
var edtUrl *Node
var memWebPage *Node


func startBrowser(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 661
    frmMain.obj.(*tForm).y = 218
    frmMain.obj.(*tForm).sizeX = 568
    frmMain.obj.(*tForm).sizeY = 509
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17

    btnGo = CreateBtn(frmMain, "btnGo", 2, 18, 40, 24, 0xd8dcc0, 0x0, "GO", btnGoClick)
    edtUrl = CreateEdit(frmMain, "edtUrl", 43, 18, 522, 24, 0xf8fcf8, 0x0, "http://info.cern.ch/hypertext/WWW/TheProject.html", nil, nil)
    memWebPage = CreateMemo(frmMain, "memWebPage", 2, 43, 563, 463, 0xf8fcf8, 0x0, "", nil)
    
    btnGoClick(btnGo)
}


func btnGoClick(node *Node){
	result := Get(edtUrl.obj.(*tEdit).text, "", "")	
	result = strings.Replace(result, "\n", string(13), -1)
	memWebPage.obj.(*tMemo).text = result
}


