package main

import (
    "fmt"
    //"math/rand"
    //"math"
    //"syscall/js"
    //"time"
    "strconv"
    //"strings"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    //"encoding/json"
    "time"
)


type tProgram struct {
	name string 
    start func(*Node)
}


var process []*tProc

type tProc struct {
	name string 
    form *Node 
    btn  *Node
    isRun bool
}


// Инициализация DOORS 
func initDOORS(programs []*tProgram) {
	fmt.Println("👋 DOORS started OK! 🌍")
	startDesktop()
	startRAD()
	// Запуск программ
	for i := 0; i < len(programs); i++ {
		startProcess(programs[i].name, programs[i].start) 
	}
}


var xTask int = 2 + 71
func startProcess(name string, onStart func(*Node)){
	/*obj := tBtn{name: "btnTask"+name, x: xTask, y: 2, sizeX: 80, sizeY: 28 - 4, BC: 0xD8DCC0, TC: 0x000000, caption: name, visible: true, pressed: false, enabled: true, onClick: btnTaskClick}
	node := Node{typ: BUTTON, parent: pnlTask, previous: nil, children: nil, obj: &obj}
	pnlTask.children = append(pnlTask.children, &node)*/
	
	frm := CreateForm(&layout, "frm" + name, 400, 400, 200, 130, 0xD8DCC0, WIN, name, true, nil)
	frm.obj.(*tForm).visible = false
	
	proc := tProc{name: name, form: frm, btn: nil, isRun: false} //, btn: &node
	process = append(process, &proc)
	//xTask += 81
	//layout.children[len(layout.children)-2].obj.(*tForm).focused = false
	//layout.children[len(layout.children)-1].obj.(*tForm).focused = true
	onStart(frm)
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


func execProcess(num int) {
	if !(process[num].isRun) {
		process[num].isRun = true
		process[num].form.obj.(*tForm).visible = true
	
		obj := tBtn{name: "btnTask"+process[num].name, x: xTask, y: 2, sizeX: 80, sizeY: 28 - 4, BC: 0xD8DCC0, TC: 0x000000, caption: process[num].name, visible: true, pressed: false, enabled: true, onClick: btnTaskClick}
		node_new := Node{typ: BUTTON, parent: pnlTask, previous: nil, children: nil, obj: &obj}
		pnlTask.children = append(pnlTask.children, &node_new)
		//obj.pressed = true
	
		process[num].btn = &node_new
		xTask += 81
		layout.children[len(layout.children)-2].obj.(*tForm).focused = false
		process[num].form.obj.(*tForm).focused = true
	
		i := findNode(process[num].form)
		fmt.Println(i)
		if i > 0 {
			sortChildren(i)
		}	
	}
}
