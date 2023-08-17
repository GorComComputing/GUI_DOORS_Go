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
var btnFlag *Node
var btnTrap *Node
var btnUsers *Node
var btnEvents *Node
var lblTime *Node

var frmMenuStart *Node
var btnMenuFlag *Node
var btnMenuTrap *Node
var btnMenuUsers *Node
var btnMenuEvents *Node






func main() {
	message := "👋 Wasm started OK! 🌍"
  	fmt.Println(message)

	frmDesktop = CreateForm(&layout, 0, 0, BITMAP_WIDTH-1, BITMAP_HEIGHT-2, 0x0080C0, NONE, "", true, nil)
	pnlTask = CreatePanel(frmDesktop, 0, frmDesktop.obj.(*tForm).sizeY - 28, BITMAP_WIDTH - 1, 28, 0x30B410, nil)
	btnStart = CreateBtn(pnlTask, 2, 2, 50, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick)
	btnFlag = CreateBtn(pnlTask, 2 + 52, 2, 50, 28 - 4, 0xD8DCC0, 0x000000, "FLAG", btnFlagClick)
	btnTrap = CreateBtn(pnlTask, 2 + 52 + 52, 2, 50, 28 - 4, 0xD8DCC0, 0x000000, "Trap", btnTrapClick)
	btnUsers = CreateBtn(pnlTask, 2 + 52 + 52 + 52, 2, 50, 28 - 4, 0xD8DCC0, 0x000000, "Users", btnUsersClick)
	btnEvents = CreateBtn(pnlTask, 2 + 52 + 52 + 52 + 52, 2, 50, 28 - 4, 0xD8DCC0, 0x000000, "Events", btnEventsClick)
	
	lblTime = CreateLabel(pnlTask, pnlTask.obj.(*tPanel).sizeX - 45, 6, 40, 20, 0x30B410, 0xF8FCF8, "", nil)
	
	frmMenuStart = CreateForm(&layout, 0, BITMAP_HEIGHT-116, 107, 85, 0xD8DCC0, NONE, "", false, nil)
	btnMenuFlag = CreateBtn(frmMenuStart, 4, 3, 100, 20, 0xD8DCC0, 0x000000, "Flag", btnMenuFlagClick)
	btnMenuTrap = CreateBtn(frmMenuStart, 4, 3 + 20, 100, 20, 0xD8DCC0, 0x000000, "Trap", btnMenuTrapClick)
	btnMenuUsers = CreateBtn(frmMenuStart, 4, 3 + 20 + 20, 100, 20, 0xD8DCC0, 0x000000, "Users", btnMenuUsersClick)
	btnMenuEvents = CreateBtn(frmMenuStart, 4, 3 + 20 + 20 + 20, 100, 20, 0xD8DCC0, 0x000000, "Events", btnMenuEventsClick)
	
	
	startTrap()
	startFlag()
	startUsers()
	startEvents()
	
    <-make(chan bool)
}





func btnStartClick(node *Node){
	frmMenuStart.obj.(*tForm).visible = true	
}


func btnMenuFlagClick(node *Node){
	frmMenuStart.obj.(*tForm).visible = false
	i := findNode(frmFlag)
	if i > 0 {
		sortChildren(i)
	}
	frmFlag.obj.(*tForm).visible = true
}


func btnMenuTrapClick(node *Node){
	frmMenuStart.obj.(*tForm).visible = false
	i := findNode(frmTrap)
	if i > 0 {
		sortChildren(i)
	}
	frmTrap.obj.(*tForm).visible = true
}


func btnMenuUsersClick(node *Node){
	frmMenuStart.obj.(*tForm).visible = false
	i := findNode(frmUsers)
	if i > 0 {
		sortChildren(i)
	}
	frmUsers.obj.(*tForm).visible = true
}


func btnMenuEventsClick(node *Node){
	frmMenuStart.obj.(*tForm).visible = false
	i := findNode(frmEvents)
	if i > 0 {
		sortChildren(i)
	}
	frmEvents.obj.(*tForm).visible = true
}


func btnFlagClick(node *Node){
	i := findNode(frmFlag)
	if i > 0 {
		sortChildren(i)
	}
	frmFlag.obj.(*tForm).visible = !(frmFlag.obj.(*tForm).visible)
}


func btnTrapClick(node *Node){
	i := findNode(frmTrap)
	if i > 0 {
		sortChildren(i)
	}
	frmTrap.obj.(*tForm).visible = !(frmTrap.obj.(*tForm).visible)
}


func btnUsersClick(node *Node){
	i := findNode(frmUsers)
	if i > 0 {
		sortChildren(i)
	}
	frmUsers.obj.(*tForm).visible = !(frmUsers.obj.(*tForm).visible)
}


func btnEventsClick(node *Node){
	i := findNode(frmEvents)
	if i > 0 {
		sortChildren(i)
	}
	frmEvents.obj.(*tForm).visible = !(frmEvents.obj.(*tForm).visible)
}


func onTimer() {
	flagDraw(cnvFlag.obj.(*tCanvas).x+50, cnvFlag.obj.(*tCanvas).y+50)
	
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















