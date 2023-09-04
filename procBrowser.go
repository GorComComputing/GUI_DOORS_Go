package main

import (
	"strings"
)

var btnGo *Node
var edtUrl *Node
var memWebPage *Node

var btnGo2 *Node
var edtUrl2 *Node
var memWebPage2 *Node

var pnlBrowser1 *Node
var pnlBrowser2 *Node
var tabBrowser *Node


func startBrowser(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 661
    frmMain.obj.(*tForm).y = 218
    frmMain.obj.(*tForm).sizeX = 572
    frmMain.obj.(*tForm).sizeY = 552
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
    
    listTabBrowser := []string{"Google", "Web"} 
    pnlBrowser1 = CreatePanel(frmMain, "pnlBrowser1", 2, 41, 568, 509, 0xd8dcc0, NONE, nil)
    pnlBrowser2 = CreatePanel(frmMain, "pnlBrowser2", 2, 41, 568, 509, 0xd8dcc0, NONE, nil)
    pnlBrowser2.obj.(*tPanel).visible = false
	tabBrowser = CreateTab(frmMain, "tabBrowser", 2, 20, 90, 20, 0xd8dcc0, 0x0, listTabBrowser, tabBrowserClick, nil)
	

    btnGo = CreateBtn(pnlBrowser1, "btnGo", 2, 2, 40, 24, 0xd8dcc0, 0x0, "GO", btnGoClick)
    edtUrl = CreateEdit(pnlBrowser1, "edtUrl", 43, 2, 522, 24, 0xf8fcf8, 0x0, "http://google.com", nil, nil)
    memWebPage = CreateMemo(pnlBrowser1, "memWebPage", 2, 27, 563, 480, 0xf8fcf8, 0x0, "", nil)
    
    btnGoClick(btnGo)
    
    btnGo2 = CreateBtn(pnlBrowser2, "btnGo2", 2, 2, 40, 24, 0xd8dcc0, 0x0, "GO", btnGoClick)
    edtUrl2 = CreateEdit(pnlBrowser2, "edtUrl2", 43, 2, 522, 24, 0xf8fcf8, 0x0, "http://info.cern.ch/hypertext/WWW/TheProject.html", nil, nil)
    memWebPage2 = CreateMemo(pnlBrowser2, "memWebPage2", 2, 27, 563, 480, 0xf8fcf8, 0x0, "", nil)
    
    //btnGo2Click(btnGo2)
}


func tabBrowserClick(node *Node, x int, y int) {
	if node.obj.(*tTab).selected == 0 {
		pnlBrowser1.obj.(*tPanel).visible = true
		pnlBrowser2.obj.(*tPanel).visible = false
	} else {
		pnlBrowser2.obj.(*tPanel).visible = true
		pnlBrowser1.obj.(*tPanel).visible = false
	}
}


func btnGoClick(node *Node){
	result := Get(edtUrl.obj.(*tEdit).text, "", "")	
	result = strings.Replace(result, "\n", string(13), -1)
	memWebPage.obj.(*tMemo).text = result
}


func btnGo2Click(node *Node){
	result := Get(edtUrl2.obj.(*tEdit).text, "", "")	
	result = strings.Replace(result, "\n", string(13), -1)
	memWebPage2.obj.(*tMemo).text = result
}


func newPageBrowser(name string, url string){
	pnlBrowser1.obj.(*tPanel).visible = true
	pnlBrowser2.obj.(*tPanel).visible = false
	edtUrl.obj.(*tEdit).text = url
	tabBrowser.obj.(*tTab).list[0] = name
	result := Get(edtUrl.obj.(*tEdit).text, "", "")	
	result = strings.Replace(result, "\n", string(13), -1)
	memWebPage.obj.(*tMemo).text = result
}
