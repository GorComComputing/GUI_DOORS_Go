package main

import (
    "fmt"
    //"math/rand"
    //"math"
    "syscall/js"
    //"time"
    "strconv"
    "strings"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    //"encoding/json"
    "time"
)


var frmDesktop *Node
var pnlTask *Node
var btnStart *Node
var lblTime *Node

var frmMenuStart *Node
var cnvMenuStart *Node
var btnMenuFlag *Node
var btnMenuTrap *Node
var btnMenuUsers *Node
var btnMenuEvents *Node
var cbxRAD *Node
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
var btnCodeGen *Node
var btnCodeSave *Node
var btnCodeOpen *Node


var process []*tProc

type tProc struct {
	name string 
    form *Node 
    btn  *Node
}




func main() {
	message := "👋 GUI started OK! 🌍"
  	fmt.Println(message)

	frmDesktop = CreateForm(&layout, "frmDesktop", 0, 0, BITMAP_WIDTH-1, BITMAP_HEIGHT-2, 0x0080C0, FLAT, "", true, nil)
	pnlTask = CreatePanel(frmDesktop, "pnlTask", 0, frmDesktop.obj.(*tForm).sizeY - 28, BITMAP_WIDTH - 1, 28, 0x30B410, TASK, nil)
	btnStart = CreateBtn(pnlTask, "btnStart", 2, 2, 70, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick)
	
	lblTime = CreateLabel(pnlTask, "lblTime", pnlTask.obj.(*tPanel).sizeX - 45, 6, 40, 20, 0x30B410, 0xF8FCF8, "", nil)
	
	frmMenuStart = CreateForm(&layout, "frmMenuStart", 0, BITMAP_HEIGHT-156, 127, 125, 0xD8DCC0, NONE, "", false, nil)
	cnvMenuStart = CreateCanvas(frmMenuStart, "cnvMenuStart", 2, 2, 20, 120, nil)
	for y := 0; y < cnvMenuStart.obj.(*tCanvas).sizeY; y++ {
    	for x := 0; x < cnvMenuStart.obj.(*tCanvas).sizeX; x++ {
    			squareNumber := (y * cnvMenuStart.obj.(*tCanvas).sizeX) + x;
      			squareRgbaIndex := squareNumber * 4;

      			cnvMenuStart.obj.(*tCanvas).buffer[squareRgbaIndex + 0] = 0; 	// Red
      			cnvMenuStart.obj.(*tCanvas).buffer[squareRgbaIndex + 1] = 0x54; 	// Green
      			cnvMenuStart.obj.(*tCanvas).buffer[squareRgbaIndex + 2] = 0xE0; 	// Blue
      			cnvMenuStart.obj.(*tCanvas).buffer[squareRgbaIndex + 3] = 255; 	// Alpha
    	}
    }
	
	
	btnMenuFlag = CreateBtn(frmMenuStart, "btnMenuFlag", 24, 3, 100, 20, 0xD8DCC0, 0x000000, "Flag", btnMenuFlagClick)
	btnMenuTrap = CreateBtn(frmMenuStart, "btnMenuTrap", 24, 3 + 20, 100, 20, 0xD8DCC0, 0x000000, "SNMP", btnMenuTrapClick)
	btnMenuUsers = CreateBtn(frmMenuStart, "btnMenuUsers", 24, 3 + 20 + 20, 100, 20, 0xD8DCC0, 0x000000, "Users", btnMenuUsersClick)
	btnMenuEvents = CreateBtn(frmMenuStart, "btnMenuEvents", 24, 3 + 20 + 20 + 20, 100, 20, 0xD8DCC0, 0x000000, "Events", btnMenuEventsClick)
	btnMenuEvents = CreateBtn(frmMenuStart, "btnMenuEvents", 24, 3 + 20 + 20 + 20 + 20, 100, 20, 0xD8DCC0, 0x000000, "Terminal", btnMenuTerminalClick)
	cbxRAD = CreateCheckBox(frmMenuStart, "cbxRAD", 24, 3 + 20 + 20 + 20 + 20 + 20, 100, 16, 0xD8DCC0, 0x000000, "RAD", false, cbxRADClick)
	
startProcess("Flag", startFlag)
startProcess("SNMP", startSNMP)
startProcess("Users", startUsers)
startProcess("Events", startEvents)
startProcess("Terminal", startTerminal)
startProcess("Test", startProc)

	frmRAD = CreateForm(&layout, "frmRAD", 0, 0, BITMAP_WIDTH-1, 59, 0xD8DCC0, WIN, "RAD", false, nil)
	btnAddBtn = CreateBtn(frmRAD, "btnAddBtn", 2, 18, 40, 40, 0xD8DCC0, 0x000000, "Btn", btnAddBtnClick)
	btnAddLabel = CreateBtn(frmRAD, "btnAddLabel", 2+41, 18, 40, 40, 0xD8DCC0, 0x000000, "Lbl", btnAddLabelClick)
	btnAddEdit = CreateBtn(frmRAD, "btnAddEdit", 2+41+41, 18, 40, 40, 0xD8DCC0, 0x000000, "Edt", btnAddEditClick)
	btnAddCheckBox = CreateBtn(frmRAD, "btnAddCheckBox", 2+41+41+41, 18, 40, 40, 0xD8DCC0, 0x000000, "Cbx", btnAddCheckBoxClick)
	btnAddMemo = CreateBtn(frmRAD, "btnAddMemo", 2+41+41+41+41, 18, 40, 40, 0xD8DCC0, 0x000000, "Mem", btnAddMemoClick)
	btnAddPanel = CreateBtn(frmRAD, "btnAddPanel", 2+41+41+41+41+41, 18, 40, 40, 0xD8DCC0, 0x000000, "Pnl", btnAddPanelClick)
	btnAddForm = CreateBtn(frmRAD, "btnAddForm", 2+41+41+41+41+41+41, 18, 40, 40, 0xD8DCC0, 0x000000, "Frm", btnAddFormClick)
	btnCodeGen = CreateBtn(frmRAD, "btnCodeGen", 2+41+41+41+41+41+41+41+41, 18, 50, 40, 0xD8DCC0, 0x000000, "Code", btnCodeGenClick)
	btnCodeSave = CreateBtn(frmRAD, "btnCodeSave", 2+41+41+41+41+41+41+41+41+51, 18, 50, 40, 0xD8DCC0, 0x000000, "Save", btnCodeSaveClick)
	btnCodeOpen = CreateBtn(frmRAD, "btnCodeOpen", 2+41+41+41+41+41+41+41+41+51+51, 18, 50, 40, 0xD8DCC0, 0x000000, "Open", btnCodeOpenClick)
	
	frmProperties = CreateForm(&layout, "frmProperties", 0, 60, 180, 400, 0xD8DCC0, WIN, "Properties", false, nil)
	
	frmCode = CreateForm(&layout, "frmCode", 181, 60, 900, 800, 0xD8DCC0, WIN, "Code", false, nil)
	memCode = CreateMemo(frmCode, "memCode", 2, 18, 900-4, 800-17-4, 0xF8FCF8, 0x000000, "", nil)
		

	
    <-make(chan bool)
}


var xTask int = 2 + 71

func startProcess(name string, onStart func(*Node)){
	obj := tBtn{name: "btnTask"+name, x: xTask, y: 2, sizeX: 80, sizeY: 28 - 4, BC: 0xD8DCC0, TC: 0x000000, caption: name, visible: true, pressed: false, enabled: true, onClick: btnTaskClick}
	node := Node{typ: BUTTON, parent: pnlTask, previous: nil, children: nil, obj: &obj}
	pnlTask.children = append(pnlTask.children, &node)
	
	frm := CreateForm(&layout, "frm" + name, 400, 400, 200, 130, 0xD8DCC0, WIN, name, true, nil)
	
	proc := tProc{name: name, form: frm, btn: &node}
	process = append(process, &proc)
	xTask += 81
	layout.children[len(layout.children)-2].obj.(*tForm).focused = false
	layout.children[len(layout.children)-1].obj.(*tForm).focused = true
	onStart(frm)
}


func cbxRADClick(node *Node){
	node.obj.(*tCheckBox).checked = !(node.obj.(*tCheckBox).checked)
	RAD = !(RAD)
	frmCode.obj.(*tForm).visible = !(frmCode.obj.(*tForm).visible)
	frmRAD.obj.(*tForm).visible = !(frmRAD.obj.(*tForm).visible)
	i := findNode(frmProperties)
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
	frmProperties.obj.(*tForm).visible = !(frmProperties.obj.(*tForm).visible)
	i = findNode(frmProperties)
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
			
	if !RAD {
		for i := 0; i < len(layout.children); i++ {
			layout.children[i].obj.(*tForm).RAD = false
		}
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
	CreateMemo(RADFormElement, "memMemo", 10, 24, 100, 100, 0xF8FCF8, 0x000000, "Memo", nil)
}

func btnAddPanelClick(node *Node){
	CreatePanel(RADFormElement, "pnlPanel", 10, 24, 100, 100, 0xD8DCC0, NONE, nil)
}

func btnAddFormClick(node *Node){
	obj := CreateForm(&layout, "frmForm", 190, 70, 500, 300, 0xD8DCC0, WIN, "Form", true, nil)
	i := findNode(obj)
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
}

func btnCodeGenClick(node *Node){
	memCode.obj.(*tMemo).text = ""	
	//package main
	//import ()
	memCode.obj.(*tMemo).text += "package main" + string(13) + string(13) + "import ()" + string(13) + string(13)
	PrintVarNode(RADFormElement)
	memCode.obj.(*tMemo).text += string(13) + string(13)
	//func startSNMP(frmMain *Node){
	memCode.obj.(*tMemo).text += "func startProc(frmMain *Node){ " + string(13)
	
	//frmMain.obj.(*tForm).x = 190
	//frmMain.obj.(*tForm).y = 70
	//frmMain.obj.(*tForm).sizeX = 550
	//frmMain.obj.(*tForm).sizeY = 300
	//frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	memCode.obj.(*tMemo).text += "    frmMain.obj.(*tForm).x = " + strconv.Itoa(RADFormElement.obj.(*tForm).x) + string(13)
	memCode.obj.(*tMemo).text += "    frmMain.obj.(*tForm).y = " + strconv.Itoa(RADFormElement.obj.(*tForm).y) + string(13)
	memCode.obj.(*tMemo).text += "    frmMain.obj.(*tForm).sizeX = " + strconv.Itoa(RADFormElement.obj.(*tForm).sizeX) + string(13)
	memCode.obj.(*tMemo).text += "    frmMain.obj.(*tForm).sizeY = " + strconv.Itoa(RADFormElement.obj.(*tForm).sizeY) + string(13)
	memCode.obj.(*tMemo).text += "    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17" + string(13) + string(13)
	
	PrintElementNode(RADFormElement, &layout)	
	// }		
	memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
	
	PrintFuncNode(RADFormElement)
}

func PrintFuncNode(node *Node){
	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tForm:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tBtn:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tLabel:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tEdit:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tPanel:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tCheckBox:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tCanvas:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tBitBtn:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
			}
		case *tMemo:
			if obj.onClickStr != ""{
				memCode.obj.(*tMemo).text += "func " + obj.onClickStr +"(node *Node){" + string(13) + string(13)
				memCode.obj.(*tMemo).text += "}" + string(13) + string(13) + string(13)
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
	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tForm:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tBtn:
			memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tLabel:
			memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tEdit:
			memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tPanel:
			memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tCheckBox:
			memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tCanvas:
			memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tBitBtn:
			//memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		case *tMemo:
			memCode.obj.(*tMemo).text += "var " + obj.name + " *Node" + string(13)
		}
	}
	
	if node.children != nil {
			for i := 0; i < len(node.children); i++ { 
				PrintVarNode(node.children[i])
			}
	}
	return
}


func PrintElementNode(node *Node, parent *Node){
	if node.obj != nil {
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
			//frm := CreateForm(&layout, 400, 400, 200, 130, 0xD8DCC0, WIN, name, true, nil)
			/*memCode.obj.(*tMemo).text += "    " + obj.name + " := CreateForm(" + parentName + ", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				fmt.Sprintf("0x%x", obj.BC) + ", " +
				"WIN, " +
				"\"" + obj.caption + "\"" + ", " +
				"true, nil)" + string(13)*/
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
			/*var onClickStr string = "nil"
			if obj.onClickStr != "" {
				onClickStr = obj.onClickStr
			}
			//CreateBitBtn(&node, obj.sizeX - 17, 2, 15, 15, 0xD8DCC0, 0x000000, "X", formClose)
			memCode.obj.(*tMemo).text += "    " + obj.name + " = CreateBitBtn(" + parentName + ", " +
				"\"" + obj.name + "\", " +
				strconv.Itoa(obj.x) + ", " +
				strconv.Itoa(obj.y) + ", " +
				strconv.Itoa(obj.sizeX) + ", " +
				strconv.Itoa(obj.sizeY) + ", " +
				fmt.Sprintf("0x%x", obj.BC) + ", " +
				fmt.Sprintf("0x%x", obj.TC) + ", " +
				"\"" + obj.caption + "\"" + ", " +
				onClickStr + ")" + string(13)*/
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
	return
}


func btnCodeSaveClick(node *Node){
	tmp := strings.Replace(memCode.obj.(*tMemo).text, string(13), "\r\n", -1)
	result := js.Global().Call("HttpRequest", "http://localhost:8085/save", tmp).Get("response").String() 
	fmt.Println("Responsed: ", result)
	
	memTerminal.obj.(*tMemo).text = result	
}

func btnCodeOpenClick(node *Node){	
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=read ./files/tmp.go", "").Get("response").String()  //?" + memCode.obj.(*tMemo).text
	fmt.Println("Responsed: ", result)
	
	result = strings.Replace(result, "\r\n", string(13), -1)
	result = strings.Replace(result, "\t", string(0x20) + string(0x20) + string(0x20) + string(0x20), -1)
	memCode.obj.(*tMemo).text = result
	
}


func btnStartClick(node *Node){
	frmMenuStart.obj.(*tForm).visible = !(frmMenuStart.obj.(*tForm).visible)
}


func btnMenuFlagClick(node *Node){
	startProcess("Flag", startFlag)
}


func btnMenuTrapClick(node *Node){
	startProcess("SNMP", startSNMP)
}


func btnMenuUsersClick(node *Node){
	startProcess("Users", startUsers)
}


func btnMenuEventsClick(node *Node){
	startProcess("Events", startEvents)
}

func btnMenuTerminalClick(node *Node){
	startProcess("Terminal", startTerminal)
}


func btnTaskClick(node *Node){
	var i int = 0
	for ; i < len(process); i++ {
		if node == process[i].btn {
			process[i].form.obj.(*tForm).visible = !(process[i].form.obj.(*tForm).visible)
			i := findNode(process[i].form)
			fmt.Println(i)
			if i > 0 {
				sortChildren(i)
			}
			break
		}
	}
}











func onTimer() {
	if cnvFlag != nil {
		flagDraw(cnvFlag.obj.(*tCanvas).x+50, cnvFlag.obj.(*tCanvas).y+50)
	}
	
	cursor = !(cursor)
	
	t := time.Now()
	lblTime.obj.(*tLabel).caption = strconv.Itoa(t.Hour()) + ":" + fmt.Sprintf("%02d", t.Minute())
	

	
	//SetColor(0xFFFF00)
	//LinePP(cnvFlag.obj, 10, 10, 100, 100)
	//Circle(cnvFlag.obj, 50, 50, 30)
	
	
    
/*
	canvas := js.Global().Get("document").Call("getElementById", "cnvs1")

	context := canvas.Call("getContext", "2d")
	
	context.Set("fillStyle", "#ffc107")
	//context.Call("clearRect", 50, 50, 100, 100)
	context.Call("fillRect", 50, 50, 50, 50)

	ctx.Set("fillStyle", grd)
	ctx.Set("strokeStyle", "#FF0000")
	ctx.Call("fillRect", 0, 0, 40, 80)
	ctx.Call("fillText", "Fill text", 20, 50)
*/
}















