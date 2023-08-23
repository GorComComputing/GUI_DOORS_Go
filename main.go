package main

import (
    "fmt"
    //"math/rand"
    //"math"
    //"syscall/js"
    //"time"
    "strconv"
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
/*var btnFlag *Node
var btnTrap *Node
var btnUsers *Node
var btnEvents *Node*/
var lblTime *Node

var frmMenuStart *Node
var cnvMenuStart *Node
var btnMenuFlag *Node
var btnMenuTrap *Node
var btnMenuUsers *Node
var btnMenuEvents *Node


var process []*tProc

type tProc struct {
	name string 
    form *Node 
    btn  *Node
}




func main() {
	message := "👋 GUI started OK! 🌍"
  	fmt.Println(message)

	frmDesktop = CreateForm(&layout, 0, 0, BITMAP_WIDTH-1, BITMAP_HEIGHT-2, 0x0080C0, FLAT, "", true, nil)
	pnlTask = CreatePanel(frmDesktop, 0, frmDesktop.obj.(*tForm).sizeY - 28, BITMAP_WIDTH - 1, 28, 0x30B410, TASK, nil)
	btnStart = CreateBtn(pnlTask, 2, 2, 70, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick)
	
	lblTime = CreateLabel(pnlTask, pnlTask.obj.(*tPanel).sizeX - 45, 6, 40, 20, 0x30B410, 0xF8FCF8, "", nil)
	
	frmMenuStart = CreateForm(&layout, 0, BITMAP_HEIGHT-136, 127, 105, 0xD8DCC0, NONE, "", false, nil)
	cnvMenuStart = CreateCanvas(frmMenuStart, 2, 2, 20, 100, nil)
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
	
	
	btnMenuFlag = CreateBtn(frmMenuStart, 24, 3, 100, 20, 0xD8DCC0, 0x000000, "Flag", btnMenuFlagClick)
	btnMenuTrap = CreateBtn(frmMenuStart, 24, 3 + 20, 100, 20, 0xD8DCC0, 0x000000, "SNMP", btnMenuTrapClick)
	btnMenuUsers = CreateBtn(frmMenuStart, 24, 3 + 20 + 20, 100, 20, 0xD8DCC0, 0x000000, "Users", btnMenuUsersClick)
	btnMenuEvents = CreateBtn(frmMenuStart, 24, 3 + 20 + 20 + 20, 100, 20, 0xD8DCC0, 0x000000, "Events", btnMenuEventsClick)
	btnMenuEvents = CreateBtn(frmMenuStart, 24, 3 + 20 + 20 + 20 + 20, 100, 20, 0xD8DCC0, 0x000000, "Terminal", btnMenuTerminalClick)
	
startProcess("Flag", startFlag)
startProcess("SNMP", startSNMP)
startProcess("Users", startUsers)
startProcess("Events", startEvents)
startProcess("Terminal", startTerminal)

		

	
    <-make(chan bool)
}


var xTask int = 2 + 71

func startProcess(name string, onStart func(*Node)){
	obj := tBtn{x: xTask, y: 2, sizeX: 80, sizeY: 28 - 4, BC: 0xD8DCC0, TC: 0x000000, caption: name, visible: true, pressed: false, enabled: true, onClick: btnTaskClick}
	node := Node{typ: BUTTON, parent: pnlTask, previous: nil, children: nil, obj: &obj}
	pnlTask.children = append(pnlTask.children, &node)
	
	frm := CreateForm(&layout, 400, 400, 200, 130, 0xD8DCC0, WIN, name, true, nil)
	
	proc := tProc{name: name, form: frm, btn: &node}
	process = append(process, &proc)
	xTask += 81
	layout.children[len(layout.children)-2].obj.(*tForm).focused = false
	layout.children[len(layout.children)-1].obj.(*tForm).focused = true
	onStart(frm)
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















