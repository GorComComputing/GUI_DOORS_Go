package main

import (
	"strings"
)


var frmOpenDialog *Node
var btnOpenDialogOpen *Node
var edtOpenDialogPath *Node
var edtOpenDialogFile *Node
var btnOpenDialogCancel *Node
var btnOpenDialogUp *Node
var lsfOpenDialog *Node

var TO *string

func OpenDialog(path string, to *string){ 
    frmOpenDialog := CreateForm(&layout, "frmOpenDialog", 200, 200, 320, 310, 0xD8DCC0, WIN, "Open File", true, nil)
    btnOpenDialogOpen = CreateBtn(frmOpenDialog, "btnOpenDialogOpen", 240, 252, 70, 24, 0xd8dcc0, 0x0, "Open", btnOpenDialogOpenClick)
    edtOpenDialogPath = CreateEdit(frmOpenDialog, "edtOpenDialogPath", 10, 22, 200, 20, 0xf8fcf8, 0x0, path, nil, nil)
    edtOpenDialogFile = CreateEdit(frmOpenDialog, "edtOpenDialogFile", 11, 252, 220, 20, 0xf8fcf8, 0x0, "", nil, nil)
    btnOpenDialogCancel = CreateBtn(frmOpenDialog, "btnOpenDialogCancel", 240, 280, 70, 24, 0xd8dcc0, 0x0, "Cancel", btnOpenDialogCancelClick)
    btnOpenDialogUp = CreateBtn(frmOpenDialog, "btnOpenDialogUp", 217, 20, 40, 24, 0xd8dcc0, 0x0, "Up", nil)
    listOpenDialog := GetCatalogList(path)
    lsfOpenDialog = CreateListFileBox(frmOpenDialog, "lsfOpenDialog", 10, 48, 300, 200, 0xF8FCF8, 0x0, listOpenDialog, lsfOpenDialogClick, nil)
    TO = to
}


func lsfOpenDialogClick(node *Node, x int, y int){
	if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "F" {
		edtOpenDialogFile.obj.(*tEdit).text = node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
	} else if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "D" {
		edtOpenDialogFile.obj.(*tEdit).text = ""
	}
}


func btnOpenDialogOpenClick(node *Node){
	if edtOpenDialogFile.obj.(*tEdit).text != "" {
		result := ReadFile(edtOpenDialogPath.obj.(*tEdit).text + edtOpenDialogFile.obj.(*tEdit).text)
		result = strings.Replace(result, "\r\n", string(13), -1)
		result = strings.Replace(result, "\t", string(0x20) + string(0x20) + string(0x20) + string(0x20), -1)
		*TO = result
		
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


