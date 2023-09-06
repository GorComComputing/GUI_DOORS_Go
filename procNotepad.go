package main

import (
	//"strings"
)

var memNotepad *Node
var menuNotepad *Node
var menuFileNotepad *Node
var menuEditNotepad *Node


func startNotepad(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 300
    frmMain.obj.(*tForm).y = 200
    frmMain.obj.(*tForm).sizeX = 900
    frmMain.obj.(*tForm).sizeY = 800
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
    
	memNotepad = CreateMemo(frmMain, "memNotepad", 2, 18+21, 900-4, 800-17-4-21, 0xF8FCF8, 0x000000, nil)
	memNotepad.obj.(*tMemo).list = []string{"String 1", "Test string", "Two string"}
	
	listNotepad := []tMenuList{{"File", nil}, {"Edit", nil}}
	menuNotepad = CreateMenu(frmMain, "menuNotepad", 2, 18, 900-1-4, 20, 0xd8dcc0, 0x0, LINE, listNotepad, menuNotepadClick, nil)
	
	listFileNotepad := []tMenuList{{"New", bmpNew_file}, {"Open", bmpOpen_file}, {"Save", bmpSave_file}}
	menuFileNotepad = CreateMenu(frmMain, "menuFileNotepad", 2, 18+20, 100, len(listFileNotepad)*20, 0xd8dcc0, 0x0, NONE, listFileNotepad, menuFileNotepadClick, nil)
	menuFileNotepad.obj.(*tMenu).visible = false
	
	listEditNotepad := []tMenuList{{"Test", nil}}
	menuEditNotepad = CreateMenu(frmMain, "menuEditNotepad", 2+60, 18+20, 100, len(listEditNotepad)*20, 0xd8dcc0, 0x0, NONE, listEditNotepad, menuEditNotepadClick, nil)
	menuEditNotepad.obj.(*tMenu).visible = false
}


func menuNotepadClick(node *Node, x int, y int){
	if node.obj.(*tMenu).selected == 0 {
		menuFileNotepad.obj.(*tMenu).visible = true
		menuEditNotepad.obj.(*tMenu).visible = false
	} else if node.obj.(*tMenu).selected == 1 {
		menuFileNotepad.obj.(*tMenu).visible = false
		menuEditNotepad.obj.(*tMenu).visible = true
	} else {
		menuFileNotepad.obj.(*tMenu).visible = false
		menuEditNotepad.obj.(*tMenu).visible = false
	}
}

//var str string
func menuFileNotepadClick(node *Node, x int, y int){
	node.obj.(*tMenu).visible = false
	if node.obj.(*tMenu).selected == 0 {
		memNotepad.obj.(*tMemo).list = nil
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
	} else if node.obj.(*tMenu).selected == 1 {
		OpenDialog(RootDir, &(memNotepad.obj.(*tMemo).list))
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
	} else if node.obj.(*tMenu).selected == 2 {
		//SaveDialog(RootDir, &(memNotepad.obj.(*tMemo).text))
	}
}


func menuEditNotepadClick(node *Node, x int, y int){
	node.obj.(*tMenu).visible = false
	if node.obj.(*tMenu).selected == 0 {
		
	} 
}







