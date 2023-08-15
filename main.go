package main

import (
    "fmt"
    //"math/rand"
    //"math"
    //"syscall/js"
    //"time"
    //"strconv"
    //"net/http"
    //"io"
)

var frmFlag *Node
var frmDesktop *Node
var pnlTask *Node
var btnStart *Node
var btnFlag *Node
var btnWin1 *Node
var frmWin1 *Node
var btnEnter *Node
var btnCancel *Node
var btnOther *Node
var lblName *Node
var lblPswd *Node
var editName *Node
var editPswd *Node

var cnvFlag *Node



func main() {
	message := "👋 Wasm started OK! 🌍"
  	fmt.Println(message)
  	
	//res, _ := http.DefaultClient.Get("http://localhost:8000")
	//if err != nil {
	//	fmt.Println("error making http request: \n")
	//}

	//fmt.Println("client: got response!\n")
	//fmt.Println("client: status code: " + strconv.Itoa(res.StatusCode))
	
/*js.FuncOf(func(this js.Value, args []js.Value) interface{} {

	go func(){

			res, _ := http.DefaultClient.Get("http://localhost:8000")
			defer res.Body.Close()
			

			b, _ := io.ReadAll(res.Body)

			fmt.Println("client: got response!")
			fmt.Println(string(b))
	}()
	
	return nil
})*/


	frmDesktop = CreateForm(&layout, 0, 0, BITMAP_WIDTH-1, BITMAP_HEIGHT-2, 0x0080C0, NONE, "", true, nil)
	pnlTask = CreatePanel(frmDesktop, 0, frmDesktop.obj.(*tForm).sizeY - 28, BITMAP_WIDTH - 1, 28, 0x30B410, nil)
	btnStart = CreateBtn(pnlTask, 2, 2, 50, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick)
	btnFlag = CreateBtn(pnlTask, 2 + 52, 2, 50, 28 - 4, 0xD8DCC0, 0x000000, "FLAG", btnFlagClick)
	btnWin1 = CreateBtn(pnlTask, 2 + 54 + 50, 2, 50, 28 - 4, 0xD8DCC0, 0x000000, "WIN 1", btnWin1Click)
	
	frmWin1 = CreateForm(&layout, 100, 100, 300, 200, 0xD8DCC0, WIN, "WINDOW 1", true, nil)
	btnEnter = CreateBtn(frmWin1, 40, 20 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "ENTER", nil)
	btnCancel = CreateBtn(frmWin1, 40 + 70, 20 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "CANCEL", nil)
	btnOther = CreateBtn(frmWin1, 80 + 60, 20 + 30 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "OTHER", nil)
	
	lblName = CreateLabel(frmWin1, 12, 22, 80, 20, 0xD8DCC0, 0x000000, "USER NAME", nil)
	lblPswd = CreateLabel(frmWin1, 12, 22 + 30, 80, 20, 0xD8DCC0, 0x000000, "PASSWORD", nil)
	
	editName = CreateEdit(frmWin1, 80, 20, 80, 20, 0xF8FCF8, 0x000000, "MYUSERNAME", nil)
	editPswd = CreateEdit(frmWin1, 80, 20 + 30, 80, 20, 0xF8FCF8, 0x000000, "PSWD", nil)
	
	
	frmFlag = CreateForm(&layout, 50, 50, 380, 340, 0x000000, WIN, "FLAGS", false, nil)
	cnvFlag = CreateCanvas(frmFlag, 2, 17, 376, 321, nil)


 
    <-make(chan bool)
}


func btnStartClick(node *Node){
	node.obj.(*tBtn).visible = false
}


func btnFlagClick(node *Node){
	i := findNode(frmFlag)
	if i > 0 {
		sortChildren(i)
	}
	frmFlag.obj.(*tForm).visible = !(frmFlag.obj.(*tForm).visible)
}


func btnWin1Click(node *Node){
	i := findNode(frmWin1)
	if i > 0 {
		sortChildren(i)
	}
	frmWin1.obj.(*tForm).visible = !(frmWin1.obj.(*tForm).visible)
}


func onTimer() {
	flagDraw(cnvFlag.obj.(*tCanvas).x+50, cnvFlag.obj.(*tCanvas).y+50)
	//SetColor(0xFFFF00)
	//LinePP(cnvFlag.obj, 10, 10, 100, 100)
	//Circle(cnvFlag.obj, 50, 50, 30)
}










