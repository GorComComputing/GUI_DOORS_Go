package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"
    "strings"
    //"reflect"
)


var frmDesktop *Node
var pnlTask *Node
var btnStart *Node
var lblTime *Node
var lsfDesktop *Node

var frmMenuStart *Node
var cnvMenuStart *Node
var menuStart *Node
var menuStartPrograms *Node
var cbxRAD *Node
var lblFPS *Node
var imgLogoMenu *Node

var cnvFlag *Node
var pnlFlag1 *Node
var pnlFlag2 *Node
var pnlFlag3 *Node
var lblLoginFlag *Node
var edtLoginFlag *Node
var lblPasswordFlag *Node
var edtPasswordFlag *Node
var btnLoginFlag *Node
var lblHiddenEnter *Node
var imgFieldDesktop *Node
//var imgSkyFlag *Node


func startDesktop(){
	frmDesktop = CreateForm(&layout, "frmDesktop", nil, 0, 0, BITMAP_WIDTH-1, BITMAP_HEIGHT-2, 0x0080C0, FLAT, "", true, nil)
	
	imgFieldDesktop = CreateImage(frmDesktop, "imgFieldDesktop", bmpFieldDesktop, frmDesktop.obj.(*tForm).sizeX/2 - 640, frmDesktop.obj.(*tForm).sizeY/2 - 400, 1280, 800, nil)

	pnlTask = CreatePanel(frmDesktop, "pnlTask", 0, frmDesktop.obj.(*tForm).sizeY - 28, BITMAP_WIDTH - 1, 28, 0x30B410, TASK, nil)
	btnStart = CreateBtn(pnlTask, "btnStart", 2, 2, 70, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick)

	listFileDesktop := GetCatalogList(DesktopDir)
    lsfDesktop = CreateListFileBox(frmDesktop, "lsfDesktop", 10, 30, 590, 300, 0xFF0080C0, 0xF8FCF8, listFileDesktop, BIGICON, lsfDesktopClick, nil)
	
	lblFPS = CreateLabel(frmDesktop, "lblFPS", 10, 10, 200, 20, 0xFF0080C0, 0xF8FCF8, "", nil)
	
	lblTime = CreateLabel(pnlTask, "lblTime", pnlTask.obj.(*tPanel).sizeX - 45, 6, 40, 20, 0x30B410, 0xF8FCF8, "", nil)

	
	listMenuStart := []tMenuList{{"Programs", bmpPrograms}, {"Settings", bmpSettings}, {"Log Out", bmpKey}}
	
	frmMenuStart = CreateForm(&layout, "frmMenuStart", nil, 2, BITMAP_HEIGHT-len(listMenuStart)*20-20-37-50+2-30, 147, len(listMenuStart)*20+26+50+30, 0xD8DCC0, NONE, "", false, nil)
	cnvMenuStart = CreateCanvas(frmMenuStart, "cnvMenuStart", 2, 2, 20, len(listMenuStart)*20+20+50+30, nil)
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
    imgLogoMenu = CreateImage(frmMenuStart, "imgLogoMenu", bmpLogo_menu, 6, 2+45+30, 12, 60, nil)
    
	menuStart = CreateMenu(frmMenuStart, "menuStart", 24, 3, 120, len(listMenuStart)*40, 0xd8dcc0, 0x0, FLAT, listMenuStart, menuStartClick, nil)
	menuStart.obj.(*tMenu).cellY = 40
	
	var listStartPrograms []tMenuList = make([]tMenuList, 0)
	for i := 0; i < len(programs); i++ {
		var pic []byte
		if programs[i].picture != nil {
			pic = *programs[i].picture
		}
		listItem := tMenuList{programs[i].name, pic}
		listStartPrograms = append(listStartPrograms, listItem)
	}
	menuStartPrograms = CreateMenu(frmMenuStart, "menuStartPrograms", 147, -143, 120, len(listStartPrograms)*20, 0xd8dcc0, 0x0, NONE, listStartPrograms, menuStartProgramsClick, nil)
	menuStartPrograms.obj.(*tMenu).visible = false
	
	cbxRAD = CreateCheckBox(frmMenuStart, "cbxRAD", 27, frmMenuStart.obj.(*tForm).sizeY - 25, 100, 16, 0xD8DCC0, 0x000000, "RAD", false, cbxRADClick)
	
	pnlFlag1 = CreatePanel(frmDesktop, "pnlFlag1", 0, 0, frmDesktop.obj.(*tForm).sizeX+1, frmDesktop.obj.(*tForm).sizeY+1, 0x0080C0, NONE, nil)
	//imgSkyFlag = CreateImage(frmDesktop, "imgSkyFlag", bmpSkyFlag, pnlFlag1.obj.(*tPanel).sizeX/2-400, pnlFlag1.obj.(*tPanel).sizeY/2-320, 800, 600, nil)
	pnlFlag2 = CreatePanel(pnlFlag1, "pnlFlag2", pnlFlag1.obj.(*tPanel).sizeX/2-188, pnlFlag1.obj.(*tPanel).sizeY/2-220, 480, 430, 0xFF0080C0, NONE, nil)
	cnvFlag = CreateCanvas(pnlFlag2, "cnvFlag", 0, 0, pnlFlag2.obj.(*tPanel).sizeX, pnlFlag2.obj.(*tPanel).sizeY, nil)
	pnlFlag3 = CreatePanel(pnlFlag1, "pnlFlag3", pnlFlag1.obj.(*tPanel).sizeX/2-120, pnlFlag1.obj.(*tPanel).sizeY/2+200, 240, 130, 0xd8dcc0, NONE, nil)
	lblLoginFlag = CreateLabel(pnlFlag3, "lblLoginFlag", 12, 14, 90, 20, 0xd8dcc0, 0x0, "Login:", nil)
  	edtLoginFlag = CreateEdit(pnlFlag3, "edtLoginFlag", 100, 14, 120, 20, 0xf8fcf8, 0x0, "", nil, nil)
  	lblPasswordFlag = CreateLabel(pnlFlag3, "lblPasswordFlag", 12, 60, 90, 20, 0xd8dcc0, 0x0, "Password:", nil)
  	edtPasswordFlag = CreatePassEdit(pnlFlag3, "edtPasswordFlag", 100, 60, 120, 20, 0xf8fcf8, 0x0, "", nil, nil)
  	btnLoginFlag = CreateBtn(pnlFlag3, "btnLoginFlag", 100, 94, 90, 24, 0xd8dcc0, 0x0, "Login", btnLoginFlagClick)
  	lblHiddenEnter = CreateLabel(pnlFlag3, "lblHiddenEnter", 12, 94, 60, 20, 0xd8dcc0, 0xd8dcc0, "HIDDEN ENTER", lblHiddenEnterClick)
}


func btnLoginFlagClick(node *Node){
	if edtLoginFlag.obj.(*tEdit).text == loginAdmin && edtPasswordFlag.obj.(*tPassEdit).text == passwordAdmin {
		mainUser = edtLoginFlag.obj.(*tEdit).text
		edtLoginFlag.obj.(*tEdit).text = ""
		edtPasswordFlag.obj.(*tPassEdit).text = ""
		fmt.Println(mainUser)
		pnlFlag1.obj.(*tPanel).visible = false
	} else if edtLoginFlag.obj.(*tEdit).text == loginUser && edtPasswordFlag.obj.(*tPassEdit).text == passwordUser {
		mainUser = edtLoginFlag.obj.(*tEdit).text
		edtLoginFlag.obj.(*tEdit).text = ""
		edtPasswordFlag.obj.(*tPassEdit).text = ""
		pnlFlag1.obj.(*tPanel).visible = false
		fmt.Println(mainUser)
	} else {
		mainUser = edtLoginFlag.obj.(*tEdit).text
		edtLoginFlag.obj.(*tEdit).text = ""
		edtPasswordFlag.obj.(*tPassEdit).text = ""
		pnlFlag1.obj.(*tPanel).visible = false
		fmt.Println(mainUser)
	}
	
}


func lblHiddenEnterClick(node *Node){
		pnlFlag1.obj.(*tPanel).visible = false
}


func lsfDesktopClick(node *Node, x int, y int){
	switch node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ {
	case "F":
		execProcess(1)  // Run Notepad
		result := ReadFileUTF8(DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		memNotepad.obj.(*tMemo).list = strings.Split(result, string(0x0D)+string(0x0A))
		curFileNameNotepad = DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
		memNotepad.obj.(*tMemo).color = nil
	case "D":
		execProcess(0)  // Run Explorer
		edtExplorerPath.obj.(*tEdit).text = DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name + "/"
		lsfExplorer.obj.(*tListFileBox).list = GetCatalogList(edtExplorerPath.obj.(*tEdit).text)
	case ".dor":
		result := ReadFileByte(DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		loadOVM(result)
		runVM()
	case ".go":
		execProcess(1)  // Run Notepad
		result := ReadFileUTF8(DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		memNotepad.obj.(*tMemo).list = strings.Split(result, string(0x0D)+string(0x0A))
		curFileNameNotepad = edtExplorerPath.obj.(*tEdit).text + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
		syntax(langGO)
	case ".c":
		execProcess(1)  // Run Notepad
		result := ReadFileUTF8(DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		memNotepad.obj.(*tMemo).list = strings.Split(result, string(0x0D)+string(0x0A))
		curFileNameNotepad = edtExplorerPath.obj.(*tEdit).text + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
		syntax(langC)
	case ".asm":
		execProcess(1)  // Run Notepad
		result := ReadFileUTF8(DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		memNotepad.obj.(*tMemo).list = strings.Split(result, string(0x0D)+string(0x0A))
		curFileNameNotepad = edtExplorerPath.obj.(*tEdit).text + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
		syntax(langASM)
	case ".html":
		execProcess(2)  // Run Browser
		result := ReadFileUTF8(DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		memWebPage.obj.(*tMemo).list = strings.Split(result, string(0x0D)+string(0x0A))
		//curFileNameNotepad = edtExplorerPath.obj.(*tEdit).text + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name
		memWebPage.obj.(*tMemo).curX = 0
		memWebPage.obj.(*tMemo).curY = 0
		memWebPage.obj.(*tMemo).curXR = 0
		memWebPage.obj.(*tMemo).curYR = 0
	}
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
			layout.children[i].obj.(*tForm).isRAD = false
		}
	}
}


func btnStartClick(node *Node){
	frmMenuStart.obj.(*tForm).visible = !(frmMenuStart.obj.(*tForm).visible)
	node.obj.(*tBtn).pressed = frmMenuStart.obj.(*tForm).visible
	i := findNode(frmMenuStart)
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
	//if frmMenuStart.obj.(*tForm).visible {
		SetFocus(menuStart)
	//} else {
	//	SetFocus(nil)
	//}
}


func menuStartClick(node *Node, x int, y int){
	if node.obj.(*tMenu).selected == 0 {
		menuStartPrograms.obj.(*tMenu).visible = true
	} else {
		menuStartPrograms.obj.(*tMenu).visible = false
	}
	if node.obj.(*tMenu).selected == 2 {
		pnlFlag1.obj.(*tPanel).visible = true
		frmMenuStart.obj.(*tForm).visible = false
		btnStart.obj.(*tBtn).pressed = false
	} 
}


func menuStartProgramsClick(node *Node, x int, y int){
	menuStartPrograms.obj.(*tMenu).visible = false
	frmMenuStart.obj.(*tForm).visible = false
	btnStart.obj.(*tBtn).pressed = false
	
	execProcess(node.obj.(*tMenu).selected)
	
}



