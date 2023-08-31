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
    frmSaveDialog := CreateForm(&layout, "frmSaveDialog", 200, 200, 320, 310, 0xD8DCC0, WIN, "Save File", true, nil)
    btnSaveDialogSave = CreateBtn(frmSaveDialog, "btnSaveDialogSave", 240, 252, 70, 24, 0xd8dcc0, 0x0, "Save", btnSaveDialogSaveClick)
    edtSaveDialogPath = CreateEdit(frmSaveDialog, "edtSaveDialogPath", 10, 22, 200, 20, 0xf8fcf8, 0x0, path, nil, nil)
    edtSaveDialogFile = CreateEdit(frmSaveDialog, "edtSaveDialogFile", 11, 252, 220, 20, 0xf8fcf8, 0x0, "", nil, nil)
    btnSaveDialogCancel = CreateBtn(frmSaveDialog, "btnSaveDialogCancel", 240, 280, 70, 24, 0xd8dcc0, 0x0, "Cancel", btnSaveDialogCancelClick)
    btnSaveDialogUp = CreateBtn(frmSaveDialog, "btnSaveDialogUp", 217, 20, 40, 24, 0xd8dcc0, 0x0, "Up", nil)
    listSaveDialog := GetCatalogList(path)
    lsfSaveDialog = CreateListFileBox(frmSaveDialog, "lsfSaveDialog", 10, 48, 300, 200, 0xF8FCF8, 0x0, listSaveDialog, lsfSaveDialogClick, nil)
    FROM = from
}


func lsfSaveDialogClick(node *Node, x int, y int){
	if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "F" {
		edtSaveDialogFile.obj.(*tEdit).text = node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
	} else if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "D" {
		edtSaveDialogFile.obj.(*tEdit).text = ""
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


