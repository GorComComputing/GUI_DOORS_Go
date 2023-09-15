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
    setSize(frmMain, 572, 552)
    frmMain.obj.(*tForm).x = BITMAP_WIDTH/2 - frmMain.obj.(*tForm).sizeX/2
	frmMain.obj.(*tForm).y = BITMAP_HEIGHT/2 - frmMain.obj.(*tForm).sizeY/2
    
    listTabBrowser := []string{"Google", "Web"} 
    pnlBrowser1 = CreatePanel(frmMain, "pnlBrowser1", 2, 41, 568, 509, 0xd8dcc0, NONE, nil)
    pnlBrowser1.obj.(*tPanel).align = CLIENT
    pnlBrowser2 = CreatePanel(frmMain, "pnlBrowser2", 2, 41, 568, 509, 0xd8dcc0, NONE, nil)
    pnlBrowser2.obj.(*tPanel).align = CLIENT
    pnlBrowser2.obj.(*tPanel).visible = false
	tabBrowser = CreateTab(frmMain, "tabBrowser", 2, 20, 90, 20, 0xd8dcc0, 0x0, listTabBrowser, tabBrowserClick, nil)
	

    btnGo = CreateBitBtn(pnlBrowser1, "btnGo", bmpRefresh, 2, 2, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnGoClick)
    //btnRefreshDevices = CreateBitBtn(pnlDevicesDisp, "btnRefreshDevices", bmpRefresh, 60, 7, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnRefreshDevicesClick)
    edtUrl = CreateEdit(pnlBrowser1, "edtUrl", 43, 2, 522, 24, 0xf8fcf8, 0x0, "http://google.com", nil, nil)
    memWebPage = CreateMemo(pnlBrowser1, "memWebPage", 2, 27, 563, 480, 0xf8fcf8, 0x0, nil)
    memWebPage.obj.(*tMemo).align = CLIENT
    
    btnGoClick(btnGo)
    
    btnGo2 = CreateBitBtn(pnlBrowser2, "btnGo2", bmpRefresh, 2, 2, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnGoClick)
    //btnGo2 = CreateBtn(pnlBrowser2, "btnGo2", 2, 2, 40, 24, 0xd8dcc0, 0x0, "GO", btnGoClick)
    edtUrl2 = CreateEdit(pnlBrowser2, "edtUrl2", 43, 2, 522, 24, 0xf8fcf8, 0x0, "http://info.cern.ch/hypertext/WWW/TheProject.html", nil, nil)
    memWebPage2 = CreateMemo(pnlBrowser2, "memWebPage2", 2, 27, 563, 480, 0xf8fcf8, 0x0, nil)
    memWebPage2.obj.(*tMemo).align = CLIENT
    
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
	result = strings.Replace(result, "\n", string(10), -1)
	memWebPage.obj.(*tMemo).list = strings.Split(result, string(10))
}


func btnGo2Click(node *Node){
	result := Get(edtUrl2.obj.(*tEdit).text, "", "")	
	result = strings.Replace(result, "\n", string(10), -1)
	memWebPage2.obj.(*tMemo).list = strings.Split(result, string(10))
}


func newPageBrowser(name string, url string){
	pnlBrowser1.obj.(*tPanel).visible = true
	pnlBrowser2.obj.(*tPanel).visible = false
	edtUrl.obj.(*tEdit).text = url
	tabBrowser.obj.(*tTab).list[0] = name
	result := Get(edtUrl.obj.(*tEdit).text, "", "")	
	result = strings.Replace(result, "\n", string(10), -1)
	memWebPage.obj.(*tMemo).list = strings.Split(result, string(10))
}
