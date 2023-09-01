package main

import (
	"strings"
)


var frmSaveDialog *Node
var btnSaveDialogSave *Node
var edtSaveDialogPath *Node
var edtSaveDialogFile *Node
var btnSaveDialogCancel *Node
var btnSaveDialogUp *Node
var lsfSaveDialog *Node

var FROM *string

func SaveDialog(path string, from *string){ 
    frmSaveDialog := CreateForm(&layout, "frmSaveDialog", BITMAP_WIDTH/2-465/2, BITMAP_HEIGHT/2-310/2, 465, 310, 0xD8DCC0, DIALOG, "Save File", true, nil)
    btnSaveDialogSave = CreateBtn(frmSaveDialog, "btnSaveDialogSave", 389, 251, 70, 24, 0xd8dcc0, 0x0, "Save", btnSaveDialogSaveClick)
    edtSaveDialogPath = CreateEdit(frmSaveDialog, "edtSaveDialogPath", 60, 22, 400, 20, 0xf8fcf8, 0x0, path, nil, nil)
    edtSaveDialogFile = CreateEdit(frmSaveDialog, "edtSaveDialogFile", 5, 252, 375, 20, 0xf8fcf8, 0x0, "", nil, nil)
    btnSaveDialogCancel = CreateBtn(frmSaveDialog, "btnSaveDialogCancel", 389, 281, 70, 24, 0xd8dcc0, 0x0, "Cancel", btnSaveDialogCancelClick)
    btnSaveDialogUp = CreateBtn(frmSaveDialog, "btnSaveDialogUp", 9, 20, 40, 24, 0xd8dcc0, 0x0, "Up", btnSaveDialogUpClick)
    listSaveDialog := GetCatalogList(path)
    lsfSaveDialog = CreateListFileBox(frmSaveDialog, "lsfSaveDialog", 5, 48, 455, 200, 0xF8FCF8, 0x0, listSaveDialog, LISTICON, lsfSaveDialogClick, nil)
    FROM = from
}


func btnSaveDialogUpClick(node *Node){
	words := strings.Split(edtSaveDialogPath.obj.(*tEdit).text, "/")
	edtSaveDialogPath.obj.(*tEdit).text = ""
	edtSaveDialogFile.obj.(*tEdit).text = ""
	for i := 0; i < len(words)-2; i++ {
		edtSaveDialogPath.obj.(*tEdit).text += words[i] + "/"
	}
	lsfSaveDialog.obj.(*tListFileBox).list = GetCatalogList(edtSaveDialogPath.obj.(*tEdit).text)
}


func lsfSaveDialogClick(node *Node, x int, y int){
	if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "F" {
		edtSaveDialogFile.obj.(*tEdit).text = node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
	} else if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "D" {
		edtSaveDialogFile.obj.(*tEdit).text = ""
		edtSaveDialogPath.obj.(*tEdit).text += node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name + "/"
		lsfSaveDialog.obj.(*tListFileBox).list = GetCatalogList(edtSaveDialogPath.obj.(*tEdit).text)
	}
}


func btnSaveDialogSaveClick(node *Node){
	if edtSaveDialogFile.obj.(*tEdit).text != "" {
		tmp := strings.Replace(*FROM, string(13), "\r\n", -1)
		WriteFile(edtSaveDialogPath.obj.(*tEdit).text + edtSaveDialogFile.obj.(*tEdit).text, tmp)
		
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


func btnSaveDialogCancelClick(node *Node){
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


