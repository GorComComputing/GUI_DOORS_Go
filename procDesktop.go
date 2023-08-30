package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"
    //"strings"
    //"reflect"

)


var frmDesktop *Node
var pnlTask *Node
var btnStart *Node
var lblTime *Node

var frmMenuStart *Node
var cnvMenuStart *Node
var menuStart *Node
var menuStartPrograms *Node
var cbxRAD *Node

var lblFPS *Node


func startDesktop(){
	frmDesktop = CreateForm(&layout, "frmDesktop", 0, 0, BITMAP_WIDTH-1, BITMAP_HEIGHT-2, 0x0080C0, FLAT, "", true, nil)
	pnlTask = CreatePanel(frmDesktop, "pnlTask", 0, frmDesktop.obj.(*tForm).sizeY - 28, BITMAP_WIDTH - 1, 28, 0x30B410, TASK, nil)
	btnStart = CreateBtn(pnlTask, "btnStart", 2, 2, 70, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick)
	
	lblFPS = CreateLabel(frmDesktop, "lblFPS", 10, 10, 200, 20, 0x0080C0, 0xF8FCF8, "", nil)
	
	lblTime = CreateLabel(pnlTask, "lblTime", pnlTask.obj.(*tPanel).sizeX - 45, 6, 40, 20, 0x30B410, 0xF8FCF8, "", nil)

	
	listMenuStart := []string{"Programs", "Settings"}
	
	frmMenuStart = CreateForm(&layout, "frmMenuStart", 0, BITMAP_HEIGHT-len(listMenuStart)*20-20-37, 127, len(listMenuStart)*20+26, 0xD8DCC0, NONE, "", false, nil)
	cnvMenuStart = CreateCanvas(frmMenuStart, "cnvMenuStart", 2, 2, 20, len(listMenuStart)*20+20, nil)
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
	menuStart = CreateMenu(frmMenuStart, "menuStart", 24, 3, 100, len(listMenuStart)*20, 0xd8dcc0, 0x0, FLAT, listMenuStart, menuStartClick, nil)
	
	var listStartPrograms []string = make([]string, 0)
	for i := 0; i < len(programs); i++ {
		listStartPrograms = append(listStartPrograms, programs[i].name)
	}
	menuStartPrograms = CreateMenu(frmMenuStart, "menuStartPrograms", 127, -103, 100, len(listStartPrograms)*20, 0xd8dcc0, 0x0, NONE, listStartPrograms, menuStartProgramsClick, nil)
	menuStartPrograms.obj.(*tMenu).visible = false
	
	cbxRAD = CreateCheckBox(frmMenuStart, "cbxRAD", 24, menuStart.obj.(*tMenu).y + menuStart.obj.(*tMenu).sizeY + 2, 100, 16, 0xD8DCC0, 0x000000, "RAD", false, cbxRADClick)
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
	
	if !(process[node.obj.(*tMenu).selected].isRun) {
	process[node.obj.(*tMenu).selected].isRun = true
	process[node.obj.(*tMenu).selected].form.obj.(*tForm).visible = true
	
	obj := tBtn{name: "btnTask"+process[node.obj.(*tMenu).selected].name, x: xTask, y: 2, sizeX: 80, sizeY: 28 - 4, BC: 0xD8DCC0, TC: 0x000000, caption: process[node.obj.(*tMenu).selected].name, visible: true, pressed: false, enabled: true, onClick: btnTaskClick}
	node_new := Node{typ: BUTTON, parent: pnlTask, previous: nil, children: nil, obj: &obj}
	pnlTask.children = append(pnlTask.children, &node_new)
	//obj.pressed = true
	
	process[node.obj.(*tMenu).selected].btn = &node_new
	xTask += 81
	layout.children[len(layout.children)-2].obj.(*tForm).focused = false
	process[node.obj.(*tMenu).selected].form.obj.(*tForm).focused = true
	
	i := findNode(process[node.obj.(*tMenu).selected].form)
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
	}
}



