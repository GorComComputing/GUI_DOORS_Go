package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"
    "strings"
    //"reflect"
)


var RAD bool = false
var RADElement *Node
var RADFormElement *Node

var frmProperties *Node
var frmRAD *Node
var btnAddBtn *Node
var btnAddLabel *Node
var btnAddEdit *Node
var btnAddCheckBox *Node
var btnAddMemo *Node
var btnAddPanel *Node
var btnAddForm *Node 
var frmCode *Node 
var memCode *Node 

var menuWinRAD *Node
var menuFileRAD *Node
var menuBuildRAD *Node


var lblPropTop *Node
var editPropTop *Node
var lblPropLeft *Node
var editPropLeft *Node
var lblPropCaption *Node
var editPropCaption *Node
var lblPropBC *Node
var editPropBC *Node
var lblPropWidth *Node
var editPropWidth *Node
var lblPropHeight *Node
var editPropHeight *Node
var lblPropTC *Node
var editPropTC *Node
var lblPropText *Node
var editPropText *Node
var lblPropName *Node
var editPropName *Node
var lblPropMode *Node
var cmbPropMode *Node
var lblPropVisible *Node
var editPropVisible *Node
var cmbPropVisible *Node
var lblPropEnabled *Node
var cmbPropEnabled *Node
var lblPropChecked *Node
var cmbPropChecked *Node
var lblEvntClick *Node
var editEvntClick *Node
var lblEvntEnter *Node
var editEvntEnter *Node
var lblPropList *Node
var cmbPropList *Node

var lblPropSelected *Node
var editPropSelected *Node

var btnAddComboBox *Node
var btnAddListBox *Node

var tabPropEvents *Node
var pnlProperties *Node
var pnlEvents *Node

var listBool = []string{"true", "false"}
var listMode = []string{"NONE", "WIN", "DIALOG", "FLAT", "TASK"}


func startRAD(){
	frmRAD = CreateForm(&layout, "frmRAD", nil, 0, 0, BITMAP_WIDTH-1, 80, 0xD8DCC0, WIN, "RAD", false, nil)
	btnAddBtn = CreateBtn(frmRAD, "btnAddBtn", 2, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Btn", btnAddBtnClick)
	btnAddLabel = CreateBtn(frmRAD, "btnAddLabel", 2+41, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Lbl", btnAddLabelClick)
	btnAddEdit = CreateBtn(frmRAD, "btnAddEdit", 2+41+41, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Edt", btnAddEditClick)
	btnAddCheckBox = CreateBtn(frmRAD, "btnAddCheckBox", 2+41+41+41, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Cbx", btnAddCheckBoxClick)
	btnAddMemo = CreateBtn(frmRAD, "btnAddMemo", 2+41+41+41+41, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Mem", btnAddMemoClick)
	btnAddPanel = CreateBtn(frmRAD, "btnAddPanel", 2+41+41+41+41+41, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Pnl", btnAddPanelClick)
	btnAddForm = CreateBtn(frmRAD, "btnAddForm", 2+41+41+41+41+41+41, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Frm", btnAddFormClick)
	btnAddComboBox = CreateBtn(frmRAD, "btnAddComboBox", 2+41+41+41+41+41+41+41, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Cmb", btnAddComboBoxClick)
	btnAddListBox = CreateBtn(frmRAD, "btnAddListBox", 2+41+41+41+41+41+41+41+41, 18+21, 40, 40, 0xD8DCC0, 0x000000, "Lbx", btnAddListBoxClick)
	
	
	
	frmProperties = CreateForm(&layout, "frmProperties", nil, 0, 81, 185, 400, 0xD8DCC0, DIALOG, "Object Inspector", false, nil)
	listTab := []string{"Properties", "Events"} 
    pnlProperties = CreatePanel(frmProperties, "pnlPropertis", 2, 41, 181, 357, 0xd8dcc0, NONE, nil)
    pnlEvents = CreatePanel(frmProperties, "pnlEvents", 2, 41, 181, 357, 0xd8dcc0, NONE, nil)
    pnlEvents.obj.(*tPanel).visible = false
	tabPropEvents = CreateTab(frmProperties, "tabPropEvents", 2, 20, 90, 20, 0xd8dcc0, 0x0, listTab, tabPropEventsClick, nil)
	
	frmCode = CreateForm(&layout, "frmCode", nil, 185, 81, 900, 800, 0xD8DCC0, WIN, "Code", false, nil)
	memCode = CreateMemo(frmCode, "memCode", 2, 18, 900-4, 800-17-4, 0xF8FCF8, 0x000000, nil)
	
	listWinRAD := []tMenuList{{"File", nil}, {"Build", nil}}
	menuWinRAD = CreateMenu(frmRAD, "menuWinRAD", 2, 18, BITMAP_WIDTH-1-1-4, 20, 0xd8dcc0, 0x0, LINE, listWinRAD, menuWinRADClick, nil)
	
	listFileRAD := []tMenuList{{"New", bmpNew_file}, {"Open", bmpOpen_file}, {"Save", bmpSave_file}}
	menuFileRAD = CreateMenu(frmRAD, "menuFileRAD", 2, 18+20, 100, len(listFileRAD)*20, 0xd8dcc0, 0x0, NONE, listFileRAD, menuFileRADClick, nil)
	menuFileRAD.obj.(*tMenu).visible = false
	
	listBuildRAD := []tMenuList{{"Generate Go", nil}, {"Generate Asm", nil}, {"Assemble", nil}}
	menuBuildRAD = CreateMenu(frmRAD, "menuBuildRAD", 2+60, 18+20, 150, len(listBuildRAD)*20, 0xd8dcc0, 0x0, NONE, listBuildRAD, menuBuildRADClick, nil)
	menuBuildRAD.obj.(*tMenu).visible = false
}


func menuWinRADClick(node *Node, x int, y int){
	if node.obj.(*tMenu).selected == 0 {
		menuFileRAD.obj.(*tMenu).visible = true
		menuBuildRAD.obj.(*tMenu).visible = false
	} else if node.obj.(*tMenu).selected == 1 {
		menuFileRAD.obj.(*tMenu).visible = false
		menuBuildRAD.obj.(*tMenu).visible = true
	} else {
		menuFileRAD.obj.(*tMenu).visible = false
		menuBuildRAD.obj.(*tMenu).visible = false
	}
}


func menuFileRADClick(node *Node, x int, y int){
	node.obj.(*tMenu).visible = false
	if node.obj.(*tMenu).selected == 0 {
		//memCode.obj.(*tMemo).text = ""
	} else if node.obj.(*tMenu).selected == 1 {
		//OpenDialog(RootDir+"files/", &(memCode.obj.(*tMemo).text))
	} else if node.obj.(*tMenu).selected == 2 {
		//SaveDialog(RootDir+"files/", &(memCode.obj.(*tMemo).text))
	}
}


func menuBuildRADClick(node *Node, x int, y int){
	node.obj.(*tMenu).visible = false
	switch node.obj.(*tMenu).selected {
	case 0:
		//memCode.obj.(*tMemo).text = ""	
		//package main
		//import ()
		//memCode.obj.(*tMemo).text += "package main" + string(13) + string(13) + "import ()" + string(13) + string(13)
		
		PrintVarNode(RADFormElement)
		//memCode.obj.(*tMemo).text += string(13) + string(13)
		
		//func startSNMP(frmMain *Node){
		//memCode.obj.(*tMemo).text += "func startProc(frmMain *Node){ " + string(13)
	
		//frmMain.obj.(*tForm).x = 190
		//frmMain.obj.(*tForm).y = 70
		//frmMain.obj.(*tForm).sizeX = 550
		//frmMain.obj.(*tForm).sizeY = 300
		//frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
		
		//memCode.obj.(*tMemo).text += "    frmMain.obj.(*tForm).x = " + strconv.Itoa(RADFormElement.obj.(*tForm).x) + string(13)
		//memCode.obj.(*tMemo).text += "    frmMain.obj.(*tForm).y = " + strconv.Itoa(RADFormElement.obj.(*tForm).y) + string(13)
		//memCode.obj.(*tMemo).text += "    frmMain.obj.(*tForm).sizeX = " + strconv.Itoa(RADFormElement.obj.(*tForm).sizeX) + string(13)
		//memCode.obj.(*tMemo).text += "    frmMain.obj.(*tForm).sizeY = " + strconv.Itoa(RADFormElement.obj.(*tForm).sizeY) + string(13)
		//memCode.obj.(*tMemo).text += "    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17" + string(13) + string(13)
	
		PrintElementNode(RADFormElement, &layout)	
		// }		
		//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
	
		PrintFuncNode(RADFormElement)
	case 1:
	case 2:	
		/*var fileNameAsm string = "asm.asm"
		RAMasm = make([]int, 0) 
	
		textAsm = ReadFile(RootDir + fileNameAsm)
		textAsm = strings.Replace(textAsm, "\r\n", string(10), -1)
		InitNameTable()
		Assemble()
	
		fmt.Println(RAMasm)
	
		var tmp string = ""
		for i := 0; i < PC; i++ {
			tmp += string(RAMasm[i])
		}
		WriteFile(RootDir + fileNameAsm[:len(fileNameAsm) - 4] + ".dor", tmp)*/
	} 
}


func tabPropEventsClick(node *Node, x int, y int) {
	if node.obj.(*tTab).selected == 0 {
		pnlProperties.obj.(*tPanel).visible = true
		pnlEvents.obj.(*tPanel).visible = false
	} else {
		pnlEvents.obj.(*tPanel).visible = true
		pnlProperties.obj.(*tPanel).visible = false
	}
}


func editPropNameEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.name = node.obj.(*tEdit).text
	case *tBtn:
		obj.name  = node.obj.(*tEdit).text
	case *tEdit:
		obj.name  = node.obj.(*tEdit).text
	case *tLabel:
		obj.name  = node.obj.(*tEdit).text
	case *tPanel:
		obj.name  = node.obj.(*tEdit).text
	case *tMemo:
		obj.name  = node.obj.(*tEdit).text
	case *tBitBtn:
		obj.name  = node.obj.(*tEdit).text
	case *tCheckBox:
		obj.name  = node.obj.(*tEdit).text
	case *tCanvas:
		obj.name  = node.obj.(*tEdit).text
	case *tComboBox:
		obj.name  = node.obj.(*tEdit).text
	case *tListBox:
		obj.name  = node.obj.(*tEdit).text
	case *tTab:
		obj.name  = node.obj.(*tEdit).text
	case *tTable:
		obj.name  = node.obj.(*tEdit).text
	case *tMenu:
		obj.name  = node.obj.(*tEdit).text
	case *tListFileBox:
		obj.name  = node.obj.(*tEdit).text
	case *tImage:
		obj.name  = node.obj.(*tEdit).text
	}
}


func editPropLeftEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tComboBox:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListBox:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTab:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTable:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMenu:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListFileBox:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tImage:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropTopEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tComboBox:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListBox:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTab:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTable:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMenu:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListFileBox:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tImage:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropCaptionEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.caption = node.obj.(*tEdit).text
	case *tBtn:
		obj.caption = node.obj.(*tEdit).text
	case *tLabel:
		obj.caption = node.obj.(*tEdit).text
	case *tBitBtn:
		obj.caption = node.obj.(*tEdit).text
	case *tCheckBox:
		obj.caption = node.obj.(*tEdit).text
	}
}


func editPropTextEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tEdit:
		obj.text = node.obj.(*tEdit).text
	case *tMemo:
		//obj.text = node.obj.(*tEdit).text
	case *tComboBox:
		obj.text = node.obj.(*tEdit).text
	}
}


func editPropWidthEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tComboBox:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListBox:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTab:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTable:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMenu:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListFileBox:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tImage:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropHeightEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tComboBox:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListBox:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTab:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTable:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMenu:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListFileBox:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tImage:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropBCEnter(node *Node){
	val, _ := strconv.ParseInt(node.obj.(*tEdit).text, 16, 32)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.BC = uint32(val)
	case *tBtn:
		obj.BC = uint32(val)
	case *tEdit:
		obj.BC = uint32(val)
	case *tLabel:
		obj.BC = uint32(val)
	case *tPanel:
		obj.BC = uint32(val)
	case *tMemo:
		obj.BC = uint32(val)
	case *tBitBtn:
		obj.BC = uint32(val)
	case *tCheckBox:
		obj.BC = uint32(val)
	case *tComboBox:
		obj.BC = uint32(val)
	case *tListBox:
		obj.BC = uint32(val)
	case *tTab:
		obj.BC = uint32(val)
	case *tTable:
		obj.BC = uint32(val)
	case *tMenu:
		obj.BC = uint32(val)
	case *tListFileBox:
		obj.BC = uint32(val)
	}
}


func editPropTCEnter(node *Node){
	val, _ := strconv.ParseInt(node.obj.(*tEdit).text, 16, 32)
	switch obj := RADElement.obj.(type) {
	case *tBtn:
		obj.TC = uint32(val)
	case *tEdit:
		obj.TC = uint32(val)
	case *tLabel:
		obj.TC = uint32(val)
	case *tMemo:
		obj.TC = uint32(val)
	case *tBitBtn:
		obj.TC = uint32(val)
	case *tCheckBox:
		obj.TC = uint32(val)
	case *tComboBox:
		obj.TC = uint32(val)
	case *tListBox:
		obj.TC = uint32(val)
	case *tTab:
		obj.TC = uint32(val)
	case *tTable:
		obj.TC = uint32(val)
	case *tMenu:
		obj.TC = uint32(val)
	case *tListFileBox:
		obj.TC = uint32(val)
	}
}


func cmbPropModeEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		if node.obj.(*tComboBox).text == "NONE" {
			obj.mode = NONE	
		} else if node.obj.(*tComboBox).text == "WIN" {
			obj.mode = WIN
		} else if node.obj.(*tComboBox).text == "DIALOG" {
			obj.mode = DIALOG
		} else if node.obj.(*tComboBox).text == "FLAT" {
			obj.mode = FLAT
		} else if node.obj.(*tComboBox).text == "TASK" {
			obj.mode = TASK
		}
	case *tPanel:
		if node.obj.(*tComboBox).text == "NONE" {
			obj.mode = NONE	
		} else if node.obj.(*tComboBox).text == "WIN" {
			obj.mode = WIN
		} else if node.obj.(*tComboBox).text == "DIALOG" {
			obj.mode = DIALOG
		} else if node.obj.(*tComboBox).text == "FLAT" {
			obj.mode = FLAT
		} else if node.obj.(*tComboBox).text == "TASK" {
			obj.mode = TASK
		}
	case *tMenu:
		if node.obj.(*tComboBox).text == "NONE" {
			obj.mode = NONE	
		} else if node.obj.(*tComboBox).text == "WIN" {
			obj.mode = WIN
		} else if node.obj.(*tComboBox).text == "DIALOG" {
			obj.mode = DIALOG
		} else if node.obj.(*tComboBox).text == "FLAT" {
			obj.mode = FLAT
		} else if node.obj.(*tComboBox).text == "TASK" {
			obj.mode = TASK
		}
	}
}


func cmbPropVisibleEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tBtn:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tEdit:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tLabel:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tPanel:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tMemo:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tBitBtn:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tCheckBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tCanvas:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tComboBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tListBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tTab:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tTable:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tMenu:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tListFileBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	case *tImage:
		if node.obj.(*tComboBox).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.visible = false
		}
	}
}


func cmbPropEnabledEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tBtn:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tEdit:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tMemo:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tBitBtn:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tCheckBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tComboBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tListBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tTab:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tTable:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tMenu:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	case *tListFileBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.enabled = false
		}
	}
}


func cmbPropCheckedEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tCheckBox:
		if node.obj.(*tComboBox).text == "true" {
			obj.checked = true	
		} else if node.obj.(*tComboBox).text == "false" {
			obj.checked = false
		}
	}
}


func cmbPropListEnter(node *Node){
	node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tComboBox:
		obj.list = node.obj.(*tComboBox).list
	case *tListBox:
		obj.list = node.obj.(*tComboBox).list
	case *tTab:
		obj.list = node.obj.(*tComboBox).list
	case *tMenu:
		//obj.list = node.obj.(*tComboBox).list
	//case *tListFileBox:
	//	obj.list = node.obj.(*tComboBox).list
	}
}


func editPropSelectedEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tListBox:
		obj.selected, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tComboBox:
		obj.selected, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tTab:
		obj.selected, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMenu:
		obj.selected, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tListFileBox:
		obj.selected, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editEvntClickEnter(node *Node){
	node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tBtn:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tEdit:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tLabel:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tPanel:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tMemo:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tBitBtn:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tCheckBox:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tCanvas:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tComboBox:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tListBox:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tTab:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tTable:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tMenu:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tListFileBox:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tImage:
		obj.onClickStr = node.obj.(*tEdit).text
	}
}


func editEvntEnterEnter(node *Node){
	node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tEdit:
		obj.onEnterStr = node.obj.(*tEdit).text
	case *tComboBox:
		obj.onEnterStr = node.obj.(*tEdit).text
	case *tListBox:
		obj.onEnterStr = node.obj.(*tEdit).text
	case *tTab:
		obj.onEnterStr = node.obj.(*tEdit).text
	case *tTable:
		obj.onEnterStr = node.obj.(*tEdit).text
	case *tMenu:
		obj.onEnterStr = node.obj.(*tEdit).text
	case *tListFileBox:
		obj.onEnterStr = node.obj.(*tEdit).text
	}
}



func btnAddBtnClick(node *Node){
	CreateBtn(RADFormElement, "btnButton", 10, 24, 70, 24, 0xD8DCC0, 0x000000, "Button", nil)
}

func btnAddLabelClick(node *Node){
	CreateLabel(RADFormElement, "lblLabel", 10, 24, 70, 20, 0xD8DCC0, 0x000000, "Label", nil)
}

func btnAddEditClick(node *Node){
	CreateEdit(RADFormElement, "edtEdit", 10, 24, 70, 20, 0xF8FCF8, 0x000000, "Edit", nil, nil)
}

func btnAddCheckBoxClick(node *Node){
	CreateCheckBox(RADFormElement, "cbxChBox", 10, 24, 110, 20, 0xD8DCC0, 0x000000, "CheckBox", false, nil)
}

func btnAddMemoClick(node *Node){
	CreateMemo(RADFormElement, "memMemo", 10, 24, 100, 100, 0xF8FCF8, 0x000000, nil)
}

func btnAddPanelClick(node *Node){
	CreatePanel(RADFormElement, "pnlPanel", 10, 24, 100, 100, 0xD8DCC0, NONE, nil)
}

func btnAddFormClick(node *Node){
	obj := CreateForm(&layout, "frmForm", nil, 190, 70, 500, 300, 0xD8DCC0, WIN, "Form", true, nil)
	i := findNode(obj)
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
}

func btnAddComboBoxClick(node *Node){
	CreateComboBox(RADFormElement, "cbxComboBox", 10, 24, 100, 16, 0xf8fcf8, 0x0, "", nil, nil, nil)
    
    	
}

func btnAddListBoxClick(node *Node){
	CreateListBox(RADFormElement, "lbxListBox", 10, 24, 100, 100, 0xf8fcf8, 0x0, nil, nil, nil)
}


func PrintFuncNode(node *Node){
	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tForm:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tBtn:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tLabel:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tEdit:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tPanel:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tCheckBox:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tCanvas:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tBitBtn:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tMemo:
			if obj.onClickStr != ""{
				//memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				//memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		}
	}
	
	if node.children != nil {
			for i := 0; i < len(node.children); i++ { 
				PrintFuncNode(node.children[i])
			}
	}
	return
}


func PrintVarNode(node *Node){
/*	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tForm:

		case *tBtn:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tLabel:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tEdit:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tPanel:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tCheckBox:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tCanvas:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tBitBtn:
			
		case *tMemo:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		}
	}
	
	if node.children != nil {
			for i := 0; i < len(node.children); i++ { 
				PrintVarNode(node.children[i])
			}
	}
	return*/
}


func PrintElementNode(node *Node, parent *Node){
/*	if node.obj != nil {
		var parentName string = "frmMain"
		switch obj := parent.obj.(type) {
		case *tForm:
			if parent == &layout {
				parentName = "&layout"
			} else {
				parentName = "frmMain"
			} 
		case *tBtn:
			if parent.typ != FORM {
				parentName = obj.name
			} else {
				parentName = "frmMain"
			}
		case *tLabel:
			if parent.typ != FORM {
				parentName = obj.name
			} else {
				parentName = "frmMain"
			}
		case *tEdit:
			if parent.typ != FORM {
				parentName = obj.name
			} else {
				parentName = "frmMain"
			}
		case *tPanel:
			if parent.typ != FORM {
				parentName = obj.name
			} else {
				parentName = "frmMain"
			}
		case *tCheckBox:
			if parent.typ != FORM {
				parentName = obj.name
			} else {
				parentName = "frmMain"
			}
		case *tCanvas:
			if parent.typ != FORM {
				parentName = obj.name
			} else {
				parentName = "frmMain"
			}
		case *tBitBtn:
			if parent.typ != FORM {
				parentName = obj.name
			} else {
				parentName = "frmMain"
			}
		case *tMemo:
			if parent.typ != FORM {
				parentName = obj.name
			} else {
				parentName = "frmMain"
			}
		}
		
		switch obj := node.obj.(type) {
		case *tForm:
			
		case *tBtn:
			var onClickStr string = "nil"
			if obj.onClickStr != "" {
				onClickStr = obj.onClickStr
			}
			//btnTrapServer = CreateBtn(frmMain, 300, 30, 100, 24, 0xD8DCC0, 0x000000, "Run Server", btnTrapServerClick)
			memCode.obj.(*tMemo).text += "    " + obj.name + " = CreateBtn(" + parentName + ", " +
				"\"" + obj.name + "\", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				fmt.Sprintf("0x%x", obj.BC) + ", " +
				fmt.Sprintf("0x%x", obj.TC) + ", " +
				"\"" + obj.caption + "\"" + ", " +
				onClickStr + ")" + string(13)
		case *tLabel:
			var onClickStr string = "nil"
			if obj.onClickStr != "" {
				onClickStr = obj.onClickStr
			}
			//lblIPaddr = CreateLabel(frmMain, 12, 32, 120, 20, 0xD8DCC0, 0x000000, "IP address", nil)
			memCode.obj.(*tMemo).text += "    " + obj.name + " = CreateLabel(" + parentName + ", " +
				"\"" + obj.name + "\", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				fmt.Sprintf("0x%x", obj.BC) + ", " +
				fmt.Sprintf("0x%x", obj.TC) + ", " +
				"\"" + obj.caption + "\"" + ", " +
				onClickStr + ")" + string(13)
		case *tEdit:
			var onClickStr string = "nil"
			if obj.onClickStr != "" {
				onClickStr = obj.onClickStr
			}
			var onEnterStr string = "nil"
			if obj.onEnterStr != "" {
				onEnterStr = obj.onEnterStr
			}
			//editPortGet = CreateEdit(frmMain, 100, 68, 100, 20, 0xF8FCF8, 0x000000, "161", nil, nil)
			memCode.obj.(*tMemo).text += "    " + obj.name + " = CreateEdit(" + parentName + ", " +
				"\"" + obj.name + "\", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				fmt.Sprintf("0x%x", obj.BC) + ", " +
				fmt.Sprintf("0x%x", obj.TC) + ", " +
				"\"" + obj.text + "\"" + ", " +
				onClickStr + ", " + 
				onEnterStr + ")" + string(13)
		case *tPanel:
			var onClickStr string = "nil"
			if obj.onClickStr != "" {
				onClickStr = obj.onClickStr
			}
			//pnlTask = CreatePanel(frmDesktop, 0, frmDesktop.obj.(*tForm).sizeY - 28, BITMAP_WIDTH - 1, 28, 0x30B410, TASK, nil)
			memCode.obj.(*tMemo).text += "    " + obj.name + " = CreatePanel(" + parentName + ", " +
				"\"" + obj.name + "\", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				fmt.Sprintf("0x%x", obj.BC) + ", " +
				
				"NONE, " +
				onClickStr + ")" + string(13)
		case *tCheckBox:
			var onClickStr string = "nil"
			if obj.onClickStr != "" {
				onClickStr = obj.onClickStr
			}
			//cbxVersion1 = CreateCheckBox(frmMain, 430, 30, 140, 16, 0xD8DCC0, 0x000000, "Version 1", false, cbxVersion1Click)
			memCode.obj.(*tMemo).text += "    " + obj.name + " = CreateCheckBox(" + parentName + ", " +
				"\"" + obj.name + "\", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				fmt.Sprintf("0x%x", obj.BC) + ", " +
				fmt.Sprintf("0x%x", obj.TC) + ", " +
				"\"" + obj.caption + "\"" + ", " +
				"false, " +
				onClickStr + ")" + string(13)
		case *tCanvas:
			var onClickStr string = "nil"
			if obj.onClickStr != "" {
				onClickStr = obj.onClickStr
			}
			//cnvFlag = CreateCanvas(frmMain, 2, 17, 376, 321, nil)
			memCode.obj.(*tMemo).text += "    " + obj.name + " = CreateCanvas(" + parentName + ", " +
				"\"" + obj.name + "\", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				onClickStr + ")" + string(13)
		case *tBitBtn:

		case *tMemo:
			var onClickStr string = "nil"
			if obj.onClickStr != "" {
				onClickStr = obj.onClickStr
			}
			//memTest = CreateMemo(frmMain, 400, 30, 100, 100, 0x000000, 0xF8FCF8, "", nil)
			memCode.obj.(*tMemo).text += "    " + obj.name + " = CreateMemo(" + parentName + ", " +
				"\"" + obj.name + "\", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				fmt.Sprintf("0x%x", obj.BC) + ", " +
				fmt.Sprintf("0x%x", obj.TC) + ", " +
				"\"" + obj.text + "\"" + ", " +
				onClickStr + ")" + string(13)
		}
	}
	
	if node.children != nil {
			for i := 0; i < len(node.children); i++ { 
				PrintElementNode(node.children[i], node)
			}
	}
	return*/
}



