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


func startDesktop(){
	frmDesktop = CreateForm(&layout, "frmDesktop", nil, 0, 0, BITMAP_WIDTH-1, BITMAP_HEIGHT-2, 0x0080C0, FLAT, "", true, nil)
	pnlTask = CreatePanel(frmDesktop, "pnlTask", 0, frmDesktop.obj.(*tForm).sizeY - 28, BITMAP_WIDTH - 1, 28, 0x30B410, TASK, nil)
	btnStart = CreateBtn(pnlTask, "btnStart", 2, 2, 70, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick)

	listFileDesktop := GetCatalogList(DesktopDir)
    lsfDesktop = CreateListFileBox(frmDesktop, "lsfDesktop", 10, 30, 590, 300, 0x0080C0, 0xF8FCF8, listFileDesktop, BIGICON, lsfDesktopClick, nil)
	
	lblFPS = CreateLabel(frmDesktop, "lblFPS", 10, 10, 200, 20, 0x0080C0, 0xF8FCF8, "", nil)
	
	lblTime = CreateLabel(pnlTask, "lblTime", pnlTask.obj.(*tPanel).sizeX - 45, 6, 40, 20, 0x30B410, 0xF8FCF8, "", nil)

	
	listMenuStart := []tMenuList{{"Programs", bmpPrograms}, {"Settings", bmpSettings}}
	
	frmMenuStart = CreateForm(&layout, "frmMenuStart", nil, 2, BITMAP_HEIGHT-len(listMenuStart)*20-20-37-50+2, 147, len(listMenuStart)*20+26+50, 0xD8DCC0, NONE, "", false, nil)
	cnvMenuStart = CreateCanvas(frmMenuStart, "cnvMenuStart", 2, 2, 20, len(listMenuStart)*20+20+50, nil)
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
    imgLogoMenu = CreateImage(frmMenuStart, "imgLogoMenu", bmpLogo_menu, 6, 2+45, 12, 60, nil)
    
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
}


func lsfDesktopClick(node *Node, x int, y int){
	if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "F" {
		execProcess(1)  // Run Notepad
		result := ReadFile(DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name)
		memNotepad.obj.(*tMemo).list = strings.Split(result, string(10))
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
	} else if node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].typ == "D" {
		execProcess(0)  // Run Explorer
		edtExplorerPath.obj.(*tEdit).text = DesktopDir + node.obj.(*tListFileBox).list[node.obj.(*tListFileBox).selected].name + "/"
		lsfExplorer.obj.(*tListFileBox).list = GetCatalogList(edtExplorerPath.obj.(*tEdit).text)
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
}


func menuStartProgramsClick(node *Node, x int, y int){
	menuStartPrograms.obj.(*tMenu).visible = false
	frmMenuStart.obj.(*tForm).visible = false
	btnStart.obj.(*tBtn).pressed = false
	
	execProcess(node.obj.(*tMenu).selected)
	
}



