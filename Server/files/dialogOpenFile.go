package main

import ()

var btnButton *Node
var edtEdit *Node
var edtEdit *Node
var btnButton *Node
var btnButton *Node


func startProc(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 341
    frmMain.obj.(*tForm).y = 242
    frmMain.obj.(*tForm).sizeX = 320
    frmMain.obj.(*tForm).sizeY = 410
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17

    btnButton = CreateBtn(frmMain, "btnButton", 240, 352, 70, 24, 0xd8dcc0, 0x0, "OPEN", nil)
    edtEdit = CreateEdit(frmMain, "edtEdit", 10, 22, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
    edtEdit = CreateEdit(frmMain, "edtEdit", 11, 352, 220, 20, 0xf8fcf8, 0x0, "Edit", nil, nil)
    btnButton = CreateBtn(frmMain, "btnButton", 240, 380, 70, 24, 0xd8dcc0, 0x0, "CANCEL", nil)
    btnButton = CreateBtn(frmMain, "btnButton", 217, 20, 40, 24, 0xd8dcc0, 0x0, "UP", nil)
}


