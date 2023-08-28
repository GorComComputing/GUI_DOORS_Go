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
var btnMenuFlag *Node
var btnMenuTrap *Node
var btnMenuUsers *Node
var btnMenuEvents *Node
var cbxRAD *Node


func startDesktop(){
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
