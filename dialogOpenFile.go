package main

import (
	//"fmt"
	"strings"
)


var frmOpenDialog *Node
var btnOpenDialogOpen *Node
var edtOpenDialogPath *Node
var edtOpenDialogFile *Node
var btnOpenDialogCancel *Node
var btnOpenDialogUp *Node
var lsfOpenDialog *Node

var TO *[]string

func OpenDialog(path string, to *[]string){ 
    frmOpenDialog := CreateForm(&layout, "frmOpenDialog", nil, BITMAP_WIDTH/2-465/2, BITMAP_HEIGHT/2-310/2, 465, 310, 0xD8DCC0, DIALOG, "Open File", true, nil)
    btnOpenDialogOpen = CreateBtn(frmOpenDialog, "btnOpenDialogOpen", 389, 251, 70, 24, 0xd8dcc0, 0x0, "Open", btnOpenDialogOpenClick)
    edtOpenDialogPath = CreateEdit(frmOpenDialog, "edtOpenDialogPath", 60, 22, 400, 20, 0xf8fcf8, 0x0, path, nil, nil)
    edtOpenDialogFile = CreateEdit(frmOpenDialog, "edtOpenDialogFile", 5, 252, 375, 20, 0xf8fcf8, 0x0, "", nil, nil)
    btnOpenDialogCancel = CreateBtn(frmOpenDialog, "btnOpenDialogCancel", 389, 281, 70, 24, 0xd8dcc0, 0x0, "Cancel", btnOpenDialogCancelClick)
    btnOpenDialogUp = CreateBitBtn(frmOpenDialog, "btnOpenDialogUp", bmpUp, 9, 20, 40, 24, 0xd8dcc0, 0x0, "", FLAT, btnOpenDialogUpClick)
    listOpenDialog := GetCatalogList(path)
    lsfOpenDialog = CreateListFileBox(frmOpenDialog, "lsfOpenDialog", 5, 48, 455, 200, 0xF8FCF8, 0x0, listOpenDialog, LISTICON, lsfOpenDialogClick, nil)
    TO = to
}


func btnOpenDialogUpClick(node *Node){
	words := strings.Split(edtOpenDialogPath.obj.(*tEdit).text, "/")
	edtOpenDialogPath.obj.(*tEdit).text = ""
	edtOpenDialogFile.obj.(*tEdit).text = ""
	for i := 0; i < len(words)-2; i++ {
		edtOpenDialogPath.obj.(*tEdit).text += words[i] + "/"
	}
	lsfOpenDialog.obj.(*tListFileBox).list = GetCatalogList(edtOpenDialogPath.obj.(*tEdit).text)
}


func lsfOpenDialogClick(node *Node, x int, y int){
	if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "F" {
		edtOpenDialogFile.obj.(*tEdit).text = node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
	} else if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "D" {
		edtOpenDialogFile.obj.(*tEdit).text = ""
		edtOpenDialogPath.obj.(*tEdit).text += node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name + "/"
		lsfOpenDialog.obj.(*tListFileBox).list = GetCatalogList(edtOpenDialogPath.obj.(*tEdit).text)
		
	}
}


func btnOpenDialogOpenClick(node *Node){
	if edtOpenDialogFile.obj.(*tEdit).text != "" {
		result := ReadFile(edtOpenDialogPath.obj.(*tEdit).text + edtOpenDialogFile.obj.(*tEdit).text)		
		*TO = strings.Split(result, string(10))
		
		// Удаляет форму
		for i := 0; i < len(layout.children); i++ {
			if node.parent == layout.children[i] {
				copy(layout.children[i:], layout.children[i+1:])
				layout.children[len(layout.children)-1] = nil
				layout.children = layout.children[:len(layout.children)-1]
				break
			}
		}
	}
}


func btnOpenDialogCancelClick(node *Node){
	// Удаляет форму
		for i := 0; i < len(layout.children); i++ {
			if node.parent == layout.children[i] {
				copy(layout.children[i:], layout.children[i+1:])
				layout.children[len(layout.children)-1] = nil
				layout.children = layout.children[:len(layout.children)-1]
				break
			}
		}
}


