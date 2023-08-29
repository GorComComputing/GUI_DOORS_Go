package main

import ()

var btnButton *Node
var edtEdit *Node
var memMemo *Node


func startProc(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 661
    frmMain.obj.(*tForm).y = 218
    frmMain.obj.(*tForm).sizeX = 568
    frmMain.obj.(*tForm).sizeY = 509
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17

    btnButton = CreateBtn(frmMain, "btnButton", 2, 18, 40, 24, 0xd8dcc0, 0x0, "GO", btngoclick)
    edtEdit = CreateEdit(frmMain, "edtEdit", 43, 18, 523, 24, 0xf8fcf8, 0x0, "Edit", nil, nil)
    memMemo = CreateMemo(frmMain, "memMemo", 2, 43, 564, 463, 0xf8fcf8, 0x0, "", nil)
}


func btngoclick(node *Node){

}


