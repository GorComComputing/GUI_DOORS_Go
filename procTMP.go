package main

import ()

var btnButton *Node
var edtEdit *Node
var lblLabel *Node
var cbxChBox *Node
var cmbTest *Node
var lsbTest *Node
var tabPropEvents2 *Node


func startTMP(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 290
    frmMain.obj.(*tForm).y = 240
    frmMain.obj.(*tForm).sizeX = 500
    frmMain.obj.(*tForm).sizeY = 500
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17

    btnButton = CreateBtn(frmMain, "btnButton", 191, 263, 70, 24, 0xd8dcc0, 0x0, "OK", okclick)
    edtEdit = CreateEdit(frmMain, "edtEdit", 209, 208, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
    lblLabel = CreateLabel(frmMain, "lblLabel", 127, 206, 70, 20, 0xd8dcff, 0x0, "Label", nil)
    cbxChBox = CreateCheckBox(frmMain, "cbxChBox", 321, 260, 110, 20, 0xd8dcc0, 0x0, "CheckBox", false, check)
    
    list := []string{"true", "false"} 
    cmbTest = CreateComboBox(frmMain, "cmbTest", 191, 350, 100, 16, 0xf8fcf8, 0x0, "true", list, nil, nil)
    lsbTest = CreateListBox(frmMain, "lsbTest", 191, 80, 100, 60, 0xf8fcf8, 0x0, list, nil, nil)
    
    listTab := []string{"Properties", "Events"} 
    
	tabPropEvents2 = CreateTab(frmMain, "tabPropEvents", 2, 20, 90, 20, 0xd8dcc0, 0x0, listTab, tabtabPropEventsClick, nil)

}


func okclick(node *Node){
	lblLabel.obj.(*tLabel).caption = edtEdit.obj.(*tEdit).text
}


func check(node *Node){
	cbxChBox.obj.(*tCheckBox).checked = !(cbxChBox.obj.(*tCheckBox).checked)
	if !(cbxChBox.obj.(*tCheckBox).checked) {
		btnButton.obj.(*tBtn).enabled = false
	} else {
		btnButton.obj.(*tBtn).enabled = true
	}
}


