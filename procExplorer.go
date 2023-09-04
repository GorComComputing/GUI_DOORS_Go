package main

import (
	"strings"
)


var edtExplorerPath *Node
var btnExplorerUp *Node
var lsfExplorer *Node

func startExplorer(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 200
    frmMain.obj.(*tForm).y = 200
    frmMain.obj.(*tForm).sizeX = 600
    frmMain.obj.(*tForm).sizeY = 400
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
    
	edtExplorerPath = CreateEdit(frmMain, "edtExplorerPath", 50, 22, 546, 20, 0xf8fcf8, 0x0, RootDir, nil, nil)
	btnExplorerUp = CreateBitBtn(frmMain, "btnExplorerUp", bmpUp, 4, 20, 40, 24, 0xd8dcc0, 0x0, "", FLAT, btnExplorerUpClick)
    listExplorer := GetCatalogList(edtExplorerPath.obj.(*tEdit).text)
    lsfExplorer = CreateListFileBox(frmMain, "lsfExplorer", 4, 48, 592, 348, 0xF8FCF8, 0x0, listExplorer, BIGICON, lsfExplorerClick, nil)

}


func btnExplorerUpClick(node *Node){
	words := strings.Split(edtExplorerPath.obj.(*tEdit).text, "/")
	edtExplorerPath.obj.(*tEdit).text = ""
	for i := 0; i < len(words)-2; i++ {
		edtExplorerPath.obj.(*tEdit).text += words[i] + "/"
	}
	lsfExplorer.obj.(*tListFileBox).list = GetCatalogList(edtExplorerPath.obj.(*tEdit).text)
}


func lsfExplorerClick(node *Node, x int, y int){
	if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "F" {
		execProcess(1)  // Run Notepad
		result := ReadFile(edtExplorerPath.obj.(*tEdit).text + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		result = strings.Replace(result, "\r\n", string(13), -1)
		result = strings.Replace(result, "\t", string(0x20) + string(0x20) + string(0x20) + string(0x20), -1)
		memNotepad.obj.(*tMemo).text = result
	} else if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "D" {
		edtExplorerPath.obj.(*tEdit).text += node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name + "/"
		lsfExplorer.obj.(*tListFileBox).list = GetCatalogList(edtExplorerPath.obj.(*tEdit).text)
	}
}

