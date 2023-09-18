package main

import (
	"strings"
	//"fmt"
)


var edtExplorerPath *Node
var btnExplorerUp *Node
var lsfExplorer *Node

func startExplorer(frmMain *Node){ 
    setSize(frmMain, 600, 400)
    frmMain.obj.(*tForm).x = BITMAP_WIDTH/2 - frmMain.obj.(*tForm).sizeX/2
	frmMain.obj.(*tForm).y = BITMAP_HEIGHT/2 - frmMain.obj.(*tForm).sizeY/2
    
	edtExplorerPath = CreateEdit(frmMain, "edtExplorerPath", 50, 22, 546, 20, 0xf8fcf8, 0x0, RootDir, nil, nil)
	btnExplorerUp = CreateBitBtn(frmMain, "btnExplorerUp", bmpUp, 4, 20, 40, 24, 0xd8dcc0, 0x0, "", FLAT, btnExplorerUpClick)
    listExplorer := GetCatalogList(edtExplorerPath.obj.(*tEdit).text)
    lsfExplorer = CreateListFileBox(frmMain, "lsfExplorer", 4, 48, 592, 348, 0xF8FCF8, 0x0, listExplorer, BIGICON, lsfExplorerClick, nil)
    lsfExplorer.obj.(*tListFileBox).align = CLIENT

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
	switch node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ {
	case "F":
		execProcess(1)  // Run Notepad
		result := ReadFile(edtExplorerPath.obj.(*tEdit).text + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		memNotepad.obj.(*tMemo).list = strings.Split(result, string(0x0D)+string(0x0A))
		curFileNameNotepad = edtExplorerPath.obj.(*tEdit).text + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
		if strings.HasSuffix(curFileNameNotepad, ".asm"){
			syntax(langASM)
		} else if strings.HasSuffix(curFileNameNotepad, ".c") || 
				  strings.HasSuffix(curFileNameNotepad, ".cpp") {
			syntax(langC)
		} else if strings.HasSuffix(curFileNameNotepad, ".go") {
			syntax(langGO)
		} else {
			memNotepad.obj.(*tMemo).color = nil
		}
	case "D":
		edtExplorerPath.obj.(*tEdit).text += node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name + "/"
		lsfExplorer.obj.(*tListFileBox).list = GetCatalogList(edtExplorerPath.obj.(*tEdit).text)
	case "B":
		result := ReadFile(edtExplorerPath.obj.(*tEdit).text + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		loadOVM(result)
		runVM()
	}
}

