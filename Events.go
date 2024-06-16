package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"
    "strings"
    //"reflect"
	"time"
	//"math"
)


var mouseIsDown bool = false
var cursor bool


var layout_obj = tForm{x: 0, y: 0, sizeX: BITMAP_WIDTH-1, sizeY: BITMAP_HEIGHT-2, BC: 0x000000, mode: FLAT, caption: "", visible: true, onClick: nil}
var layout = Node{parent: nil, previous: nil, children: nil, obj: &layout_obj}
var list []*Node


type tComponents int

const (
    BUTTON tComponents = iota
    FORM
    EDIT
    LABEL
    PANEL
    CANVAS
    BIT_BUTTON
    MEMO
    CHECKBOX
    COMBOBOX
    LISTBOX
    TAB
    MENU
    LISTFILEBOX
    TABLE
    IMAGE
    PASSEDIT
    TRACKBAR
)

type tWinComponents interface {
   Draw(parX int, parY int, parSizeX int, parSizeY int)
   RAD(x int, y int)
   KeyDown(key int)
   Click(x int, y int)
   MouseMove(x int, y int, Xl int, Yl int)
   MouseDown(x int, y int)
}


type Node struct {
	typ tComponents 
    parent *Node
    previous *Node
    children []*Node
    obj tWinComponents 
}   
   

//export eventResizeWindow
func eventResizeWindow(width int, height int) {
	fmt.Println("Width: " + strconv.Itoa(width))
	fmt.Println("Height: " + strconv.Itoa(height))
	
	//SetBackColor(0xFFFF00)
	//ClearDevice(nil)
	
	BITMAP_WIDTH = width   
	BITMAP_HEIGHT = height
	
	SIZE = width*height 
	GETMAX_X = width - 1 
	GETMAX_Y = height - 1
	BUFFER_SIZE = width*height * 4
	
	fmt.Println("BITMAP_WIDTH: " + strconv.Itoa(BITMAP_WIDTH))
	fmt.Println("BITMAP_HEIGHT: " + strconv.Itoa(BITMAP_HEIGHT))
	fmt.Println("SIZE: " + strconv.Itoa(SIZE))
	
	// Desktop & Task Bar & Menu Start
	frmDesktop.obj.(*tForm).sizeY = BITMAP_HEIGHT - 2
	frmDesktop.obj.(*tForm).sizeX = BITMAP_WIDTH - 1
	pnlTask.obj.(*tPanel).y = frmDesktop.obj.(*tForm).sizeY - 27
	pnlTask.obj.(*tPanel).sizeX = BITMAP_WIDTH - 1
	lblTime.obj.(*tLabel).x = pnlTask.obj.(*tPanel).sizeX - 45
	frmMenuStart.obj.(*tForm).y = BITMAP_HEIGHT-len(menuStart.obj.(*tMenu).list)*20-20-37-50+2-29
	
	// Desktop Wallpaper
	if (frmDesktop.obj.(*tForm).sizeX/2 - 640) < 0 {
		imgFieldDesktop.obj.(*tImage).x = 0
	} else {
		imgFieldDesktop.obj.(*tImage).x = (frmDesktop.obj.(*tForm).sizeX/2 - 640)
	}
	if (frmDesktop.obj.(*tForm).sizeY/2 - 400) < 0 {
		imgFieldDesktop.obj.(*tImage).y = 0
	} else {
		imgFieldDesktop.obj.(*tImage).y = (frmDesktop.obj.(*tForm).sizeY/2 - 400)
	}
	
	// Log In Window
	pnlFlag1.obj.(*tPanel).sizeX = frmDesktop.obj.(*tForm).sizeX+1
	pnlFlag1.obj.(*tPanel).sizeY = frmDesktop.obj.(*tForm).sizeY+1
	pnlFlag2.obj.(*tPanel).x = pnlFlag1.obj.(*tPanel).sizeX/2 - 166
	pnlFlag2.obj.(*tPanel).y = BITMAP_HEIGHT - pnlFlag1.obj.(*tPanel).sizeY/2 - pnlFlag1.obj.(*tPanel).sizeY/4
	//cnvFlag.obj.(*tCanvas).sizeX = pnlFlag2.obj.(*tPanel).sizeX
	//cnvFlag.obj.(*tCanvas).sizeY = pnlFlag2.obj.(*tPanel).sizeY
	pnlFlag3.obj.(*tPanel).x = pnlFlag1.obj.(*tPanel).sizeX/2 - 120
	pnlFlag3.obj.(*tPanel).y = BITMAP_HEIGHT - pnlFlag1.obj.(*tPanel).sizeY/3
	
	//SetBackColor(0xFF0000)
	//FillLB(nil, 0, SIZE, BC)
	
	//graphicsBuffer = make([]uint8, BUFFER_SIZE)
	
	//SetViewPort(0, 0, GETMAX_X, GETMAX_Y)
	//SetLocalViewPort(0, 0, GETMAX_X, GETMAX_Y)
	
	//graphicsBuffer = make([]uint8, BUFFER_SIZE, BUFFER_SIZE)
}
      

//export eventDraw
func eventDraw() {	
	start := time.Now()
	
	SetBackColor(0x000000) //0x111111 0xFFFFFF
	SetColor(0x000000)
	SetViewPort(0, 0, GETMAX_X, GETMAX_Y)
	ClearDevice(nil)
	DrawNode(&layout, 0, 0)
		//FillCircle(nil, 0, 100, 30)
		//Circle(nil, 0, 200, 30)
	onTimer()
	
	
	drawDo()
	
	
	t := time.Now()
	lblFPS.obj.(*tLabel).caption = strings.Split(strings.Split(t.Sub(start).String(), ".")[0], "ms")[0]
}


// Округление до целого
/*func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}*/


func DrawNode(node *Node, x int, y int){
	
	var visible bool = false
	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tBtn:
			visible = obj.visible
		case *tForm:
			visible = obj.visible
		case *tPanel:
			visible = obj.visible
		case *tEdit:
			visible = obj.visible
		case *tPassEdit:
			visible = obj.visible
		case *tTrackBar:
			visible = obj.visible
		case *tLabel:
			visible = obj.visible
		case *tCanvas:
			visible = obj.visible
		case *tBitBtn:
			visible = obj.visible
		case *tMemo:
			visible = obj.visible
		case *tCheckBox:
			visible = obj.visible
		case *tComboBox:
			visible = obj.visible
		case *tListBox:
			visible = obj.visible
		case *tTab:
			visible = obj.visible
		case *tMenu:
			visible = obj.visible
		case *tListFileBox:
			visible = obj.visible
		case *tTable:
			visible = obj.visible
		case *tImage:
			visible = obj.visible
		}
	}
	
	var parX int = 0
	var parY int = 0
	var parSizeX int = GETMAX_X
	var parSizeY int = GETMAX_Y
	if node.parent != nil && node.parent.obj != nil {
		switch obj := node.parent.obj.(type) {
		case *tBtn:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tForm:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tPanel:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tEdit:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tPassEdit:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tTrackBar:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tLabel:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tCanvas:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tBitBtn:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tMemo:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tCheckBox:
			parX = obj.x + x
			parY = obj.y + y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tComboBox:
			parX = obj.x + x 
			parY = obj.y + y 
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tListBox:
			parX = obj.x + x 
			parY = obj.y + y 
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tTab:
			parX = obj.x + x 
			parY = obj.y + y 
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tMenu:
			parX = obj.x + x 
			parY = obj.y + y 
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tListFileBox:
			parX = obj.x + x 
			parY = obj.y + y 
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tTable:
			parX = obj.x + x 
			parY = obj.y + y 
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tImage:
			parX = obj.x + x 
			parY = obj.y + y 
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		}
	}
	
	if node.obj != nil && visible  {
		node.obj.Draw(parX, parY, parSizeX, parSizeY)
	}
	
	if node.children != nil && visible {
			for i := 0; i < len(node.children); i++ { 
				DrawNode(node.children[i], parX, parY)
			}
	}
	return
}


//export eventClick
func eventClick(x int, y int)  {
		fmt.Println("Event: " + strconv.Itoa(x) + " " + strconv.Itoa(y))
		list = nil
		ClickRecurs(&layout, x, y, 0, 0)
		
		if !RAD || list[len(list)-1] == cbxRAD || layout.children[len(layout.children)-1] == frmProperties || layout.children[len(layout.children)-1] == frmRAD || layout.children[len(layout.children)-1] == frmCode || layout.children[len(layout.children)-1].obj.(*tForm).mode == DIALOG {
		switch obj := list[len(list)-1].obj.(type) {
		case *tBtn:
			obj.Click(x-X, y-Y)
		case *tBitBtn:
			obj.Click(x-X, y-Y)
		case *tCheckBox:
			obj.Click(x-X, y-Y)
		case *tEdit:
			obj.Click(x-X, y-Y)
		case *tPassEdit:
			obj.Click(x-X, y-Y)
		case *tTrackBar:
			obj.Click(x-X, y-Y)
		case *tComboBox:
			obj.Click(x-X, y-Y)
		case *tListBox:
			obj.Click(x-X, y-Y)
		case *tTab:
			obj.Click(x-X, y-Y)
		case *tMenu:
			obj.Click(x-X, y-Y)
		case *tListFileBox:
			obj.Click(x-X, y-Y)
		case *tTable:
			obj.Click(x-X, y-Y)
		case *tImage:
			obj.Click(x-X, y-Y)
		case *tLabel:
			obj.Click(x-X, y-Y)
		}
	}
	
	eventDraw()
}


var X, Y int = 0, 0
func ClickRecurs(node *Node, x int, y int, parX int, parY int) {
	 
	var visible bool = false
	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tBtn:
			visible = obj.visible
		case *tForm:
			visible = obj.visible
		case *tPanel:
			visible = obj.visible
		case *tEdit:
			visible = obj.visible
		case *tTrackBar:
			visible = obj.visible
		case *tPassEdit:
			visible = obj.visible
		case *tLabel:
			visible = obj.visible
		case *tCanvas:
			visible = obj.visible
		case *tBitBtn:
			visible = obj.visible
		case *tMemo:
			visible = obj.visible
		case *tCheckBox:
			visible = obj.visible
		case *tComboBox:
			visible = obj.visible
		case *tListBox:
			visible = obj.visible
		case *tTab:
			visible = obj.visible
		case *tMenu:
			visible = obj.visible
		case *tListFileBox:
			visible = obj.visible
		case *tTable:
			visible = obj.visible
		case *tImage:
			visible = obj.visible
		}
	}
	
	if node.parent != nil && node.parent.obj != nil {
		switch obj := node.parent.obj.(type) {
		case *tBtn:
			parX += obj.x
			parY += obj.y
		case *tForm:
			parX += obj.x
			parY += obj.y
		case *tPanel:
			parX += obj.x
			parY += obj.y
		case *tEdit:
			parX += obj.x
			parY += obj.y
		case *tPassEdit:
			parX += obj.x
			parY += obj.y
		case *tTrackBar:
			parX += obj.x
			parY += obj.y
		case *tLabel:
			parX += obj.x
			parY += obj.y
		case *tCanvas:
			parX += obj.x
			parY += obj.y
		case *tBitBtn:
			parX += obj.x
			parY += obj.y
		case *tMemo:
			parX += obj.x
			parY += obj.y
		case *tCheckBox:
			parX += obj.x
			parY += obj.y
		case *tComboBox:
			parX += obj.x
			parY += obj.y
		case *tListBox:
			parX += obj.x
			parY += obj.y
		case *tTab:
			parX += obj.x
			parY += obj.y
		case *tMenu:
			parX += obj.x
			parY += obj.y
		case *tListFileBox:
			parX += obj.x
			parY += obj.y
		case *tTable:
			parX += obj.x
			parY += obj.y
		case *tImage:
			parX += obj.x
			parY += obj.y
		}
	}
	
	if node.obj != nil && visible {
		switch node.obj.(type) {
		case *tBtn:
			if (parX+node.obj.(*tBtn).x) < x && 
			(parX+node.obj.(*tBtn).x + node.obj.(*tBtn).sizeX) > x && 
			(parY+node.obj.(*tBtn).y) < y && 
			(parY+node.obj.(*tBtn).y + node.obj.(*tBtn).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tBtn).x
				Y = parY+node.obj.(*tBtn).y
			}
		case *tForm:
			if (parX+node.obj.(*tForm).x) < x && 
			(parX+node.obj.(*tForm).x + node.obj.(*tForm).sizeX) > x && 
			(parY+node.obj.(*tForm).y) < y && 
			(parY+node.obj.(*tForm).y + node.obj.(*tForm).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tForm).x
				Y = parY+node.obj.(*tForm).y
			}
		case *tEdit:
			if (parX+node.obj.(*tEdit).x) < x && 
			(parX+node.obj.(*tEdit).x + node.obj.(*tEdit).sizeX) > x && 
			(parY+node.obj.(*tEdit).y) < y && 
			(parY+node.obj.(*tEdit).y + node.obj.(*tEdit).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tEdit).x
				Y = parY+node.obj.(*tEdit).y
			}
		case *tPassEdit:
			if (parX+node.obj.(*tPassEdit).x) < x && 
			(parX+node.obj.(*tPassEdit).x + node.obj.(*tPassEdit).sizeX) > x && 
			(parY+node.obj.(*tPassEdit).y) < y && 
			(parY+node.obj.(*tPassEdit).y + node.obj.(*tPassEdit).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tPassEdit).x
				Y = parY+node.obj.(*tPassEdit).y
			}
		case *tTrackBar:
			if (parX+node.obj.(*tEdit).x) < x && 
			(parX+node.obj.(*tEdit).x + node.obj.(*tEdit).sizeX) > x && 
			(parY+node.obj.(*tEdit).y) < y && 
			(parY+node.obj.(*tEdit).y + node.obj.(*tEdit).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tEdit).x
				Y = parY+node.obj.(*tEdit).y
			}
		case *tLabel:
			if (parX+node.obj.(*tLabel).x) < x && 
			(parX+node.obj.(*tLabel).x + node.obj.(*tLabel).sizeX) > x && 
			(parY+node.obj.(*tLabel).y) < y && 
			(parY+node.obj.(*tLabel).y + node.obj.(*tLabel).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tLabel).x
				Y = parY+node.obj.(*tLabel).y
			}
		case *tPanel:
			if (parX+node.obj.(*tPanel).x) < x && 
			(parX+node.obj.(*tPanel).x + node.obj.(*tPanel).sizeX) > x && 
			(parY+node.obj.(*tPanel).y) < y && 
			(parY+node.obj.(*tPanel).y + node.obj.(*tPanel).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tPanel).x
				Y = parY+node.obj.(*tPanel).y
			}
		case *tCanvas:
			if (parX+node.obj.(*tCanvas).x) < x && 
			(parX+node.obj.(*tCanvas).x + node.obj.(*tCanvas).sizeX) > x && 
			(parY+node.obj.(*tCanvas).y) < y && 
			(parY+node.obj.(*tCanvas).y + node.obj.(*tCanvas).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tCanvas).x
				Y = parY+node.obj.(*tCanvas).y
			}
		case *tBitBtn:
			if (parX+node.obj.(*tBitBtn).x) < x && 
			(parX+node.obj.(*tBitBtn).x + node.obj.(*tBitBtn).sizeX) > x && 
			(parY+node.obj.(*tBitBtn).y) < y && 
			(parY+node.obj.(*tBitBtn).y + node.obj.(*tBitBtn).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tBitBtn).x
				Y = parY+node.obj.(*tBitBtn).y
			}
		case *tMemo:
			if (parX+node.obj.(*tMemo).x) < x && 
			(parX+node.obj.(*tMemo).x + node.obj.(*tMemo).sizeX) > x && 
			(parY+node.obj.(*tMemo).y) < y && 
			(parY+node.obj.(*tMemo).y + node.obj.(*tMemo).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tMemo).x
				Y = parY+node.obj.(*tMemo).y
			}
		case *tCheckBox:
			if (parX+node.obj.(*tCheckBox).x) < x && 
			(parX+node.obj.(*tCheckBox).x + node.obj.(*tCheckBox).sizeX) > x && 
			(parY+node.obj.(*tCheckBox).y) < y && 
			(parY+node.obj.(*tCheckBox).y + node.obj.(*tCheckBox).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tCheckBox).x
				Y = parY+node.obj.(*tCheckBox).y
			}
		case *tComboBox:
			if (parX+node.obj.(*tComboBox).x) < x && 
			(parX+node.obj.(*tComboBox).x + node.obj.(*tComboBox).sizeX) > x && 
			(parY+node.obj.(*tComboBox).y) < y && 
			(parY+node.obj.(*tComboBox).y + node.obj.(*tComboBox).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tComboBox).x
				Y = parY+node.obj.(*tComboBox).y
			}
		case *tListBox:
			if (parX+node.obj.(*tListBox).x) < x && 
			(parX+node.obj.(*tListBox).x + node.obj.(*tListBox).sizeX) > x && 
			(parY+node.obj.(*tListBox).y) < y && 
			(parY+node.obj.(*tListBox).y + node.obj.(*tListBox).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tListBox).x
				Y = parY+node.obj.(*tListBox).y
			}
		case *tTab:
			if (parX+node.obj.(*tTab).x) < x && 
			(parX+node.obj.(*tTab).x + node.obj.(*tTab).sizeX*len(node.obj.(*tTab).list)) > x && 
			(parY+node.obj.(*tTab).y) < y && 
			(parY+node.obj.(*tTab).y + node.obj.(*tTab).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tTab).x
				Y = parY+node.obj.(*tTab).y
			}
		case *tMenu:
			if (parX+node.obj.(*tMenu).x) < x && 
			(parX+node.obj.(*tMenu).x + node.obj.(*tMenu).sizeX*len(node.obj.(*tMenu).list)) > x && 
			(parY+node.obj.(*tMenu).y) < y && 
			(parY+node.obj.(*tMenu).y + node.obj.(*tMenu).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tMenu).x
				Y = parY+node.obj.(*tMenu).y
			}
		case *tListFileBox:
			if (parX+node.obj.(*tListFileBox).x) < x && 
			(parX+node.obj.(*tListFileBox).x + node.obj.(*tListFileBox).sizeX) > x && 
			(parY+node.obj.(*tListFileBox).y) < y && 
			(parY+node.obj.(*tListFileBox).y + node.obj.(*tListFileBox).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tListFileBox).x
				Y = parY+node.obj.(*tListFileBox).y
			}
		case *tTable:
			if (parX+node.obj.(*tTable).x) < x && 
			(parX+node.obj.(*tTable).x + node.obj.(*tTable).sizeX) > x && 
			(parY+node.obj.(*tTable).y) < y && 
			(parY+node.obj.(*tTable).y + node.obj.(*tTable).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tTable).x
				Y = parY+node.obj.(*tTable).y
			}
		case *tImage:
			if (parX+node.obj.(*tImage).x) < x && 
			(parX+node.obj.(*tImage).x + node.obj.(*tImage).sizeX) > x && 
			(parY+node.obj.(*tImage).y) < y && 
			(parY+node.obj.(*tImage).y + node.obj.(*tImage).sizeY) > y {
				list = append(list, node)
				X = parX+node.obj.(*tImage).x
				Y = parY+node.obj.(*tImage).y
			}
		}
	}
			
	if node.children != nil && visible {
		for i := 0; i < len(node.children); i++ { 
			ClickRecurs(node.children[i], x, y, parX, parY)
		}
	}
	return
}


func findNode(node *Node) int {
	var i int
	if node.typ == FORM {
		for i := 0; i < len(layout.children); i++ {
			if node == layout.children[i] {
				return i
			}
		}
	} else {
		if node.parent != nil {
			i = findNode(node.parent)
		} else {
			return -1
		}
	}
	return i
}


// Установить фокус на окно
func sortChildren(i int) {
	tmp := layout.children[i]
	copy(layout.children[i:], layout.children[i+1:])
	layout.children[len(layout.children)-1] = tmp
	layout.children[len(layout.children)-2].obj.(*tForm).focused = false
	layout.children[len(layout.children)-1].obj.(*tForm).focused = true
}


// Вывести элемент на передний план
func ToUpPlane(node *Node) {
	var i int
	for i = 0; i < len(node.parent.children); i++ {
		if node == node.parent.children[i] {
			break
		}
	}	
	tmp := node.parent.children[i]
	copy(node.parent.children[i:], node.parent.children[i+1:])
	node.parent.children[len(node.parent.children)-1] = tmp
}


// Установитьь фокус на элемент
func SetFocus(node *Node) {
	// Сброс фокуса с прежнего элемента
	if layout.children[len(layout.children)-1].obj.(*tForm).focus != nil {
		switch obj := layout.children[len(layout.children)-1].obj.(*tForm).focus.obj.(type) {
    	case *tEdit:
			obj.focused = false
		case *tPassEdit:
			obj.focused = false
		case *tTrackBar:
			obj.focused = false
		case *tMemo:
			obj.focused = false
		case *tComboBox:
			obj.focused = false
		case *tListBox:
			obj.focused = false
		case *tMenu:
			obj.focused = false
		case *tListFileBox:
			obj.focused = false
		case *tTable:
			obj.focused = false
		}
	}
	// Установка нового фокуса элемента
	layout.children[len(layout.children)-1].obj.(*tForm).focus = node
	//list = append(list, node)
}


var downX int = 0
var downY int = 0
var btnPressed *tBtn = nil
var bitbtnPressed *tBitBtn = nil

//export eventMouseDown
func eventMouseDown(x int, y int) {
	list = nil
	ClickRecurs(&layout, x, y, 0, 0)
	
	// Смена фокуса окна
	i := findNode(list[len(list)-1])
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
	SetFocus(list[len(list)-1])
	
	/*for j := 0; j < len(process); j++ {
		if process[j].form == layout.children[i] {
			process[j].btn.obj.(*tBtn).pressed = true
			
		} else {
			process[j].btn.obj.(*tBtn).pressed = false
		}
	}*/

	
	
	// RAD
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			RADElement = list[len(list)-1]
			for i := 0; i < len(layout.children); i++ {
				layout.children[i].obj.(*tForm).isRAD = false
			}
			layout.children[len(layout.children)-1].obj.(*tForm).isRAD = true
			RADFormElement = layout.children[len(layout.children)-1]
			pnlProperties.children = nil
			pnlEvents.children = nil
	}
	
	switch obj := list[len(list)-1].obj.(type) {
	case *tForm:
		obj.MouseDown(x, y)
    case *tBtn:
    	obj.MouseDown(x, y)
	case *tLabel:
		obj.MouseDown(x, y)
	case *tEdit:
		obj.MouseDown(x, y)
	case *tPassEdit:
		obj.MouseDown(x, y)
	case *tTrackBar:
		obj.MouseDown(x, y)
	case *tBitBtn:
		obj.MouseDown(x, y)
	case *tMemo:
		obj.MouseDown(x, y)
	case *tCheckBox:
		obj.MouseDown(x, y)
	case *tPanel:
		obj.MouseDown(x, y) 
	case *tComboBox:
		obj.MouseDown(x, y)
	case *tListBox:
		obj.MouseDown(x, y)
	case *tTab:
		obj.MouseDown(x, y)
	case *tMenu:
		obj.MouseDown(x, y)
	case *tListFileBox:
		obj.MouseDown(x, y)
	case *tTable:
		obj.MouseDown(x, y)
	case *tImage:
		obj.MouseDown(x, y)
	}
	
	eventDraw()
}


//export eventMouseUp
func eventMouseUp(x int, y int)  {
	//if(mouseIsDown) mouseClick(e)
    mouseIsDown = false

	// Отжатие Btn
	if btnPressed != nil {
		btnPressed.pressed = false
		btnPressed = nil
	}
	// Отжатие BitBtn
	if bitbtnPressed != nil {
		bitbtnPressed.pressed = false
		bitbtnPressed = nil
	}
	
	eventDraw()
}


//export eventMouseMove
func eventMouseMove(x int, y int)  {
	if !mouseIsDown {
		list = nil
		ClickRecurs(&layout, x, y, 0, 0)
	}
	
	
	
	if len(list) > 0 {
	switch obj := list[len(list)-1].obj.(type) {
	case *tBtn:
		obj.MouseMove(x, y, x-X, y-Y)
    case *tLabel:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tEdit:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tPassEdit:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tTrackBar:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tMemo:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tCanvas:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tCheckBox:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tPanel:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tBitBtn:
    	obj.MouseMove(x, y, x-X, y-Y)
    case *tForm:
		obj.MouseMove(x, y, x-X, y-Y)
	case *tComboBox:
		obj.MouseMove(x, y, x-X, y-Y)	
	case *tListBox:
		obj.MouseMove(x, y, x-X, y-Y)	
	case *tTab:
		obj.MouseMove(x, y, x-X, y-Y)	
	case *tMenu:
		obj.MouseMove(x, y, x-X, y-Y)
	case *tListFileBox:
		obj.MouseMove(x, y, x-X, y-Y)
	case *tTable:
		obj.MouseMove(x, y, x-X, y-Y)
	case *tImage:
		obj.MouseMove(x, y, x-X, y-Y)
	}
	}
    downX = x 
    downY = y
    
    eventDraw()
}


var isShiftKeyDown bool = false
var isCtrlKeyDown bool = false
var isAltKeyDown bool = false

//export eventKeyDown
func eventKeyDown(key int){
	if key == 16 {
		isShiftKeyDown = true
	} else if key == 17 {
		isCtrlKeyDown = true
	} else if key == 18 {
		isAltKeyDown = true
	}
	
	// Удаление элемента RAD (Del)
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && key == 46 && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
		for i := 0; i < len(RADElement.parent.children); i++ {
			if RADElement.parent.children[i] == RADElement {
				copy(RADElement.parent.children[i:], RADElement.parent.children[i+1:])
				RADElement.parent.children[len(RADElement.parent.children)-1] = nil
				RADElement.parent.children = RADElement.parent.children[:len(RADElement.parent.children)-1]
				pnlProperties.children = nil
				pnlEvents.children = nil
			}		
		}
	}
	// Если в фокусе
	if layout.children[len(layout.children)-1].obj.(*tForm).focus != nil {
		switch obj := layout.children[len(layout.children)-1].obj.(*tForm).focus.obj.(type) {
    	case *tEdit:
    		obj.KeyDown(key)
    	case *tPassEdit:
    		obj.KeyDown(key)
    	case *tTrackBar:
    		obj.KeyDown(key)
		case *tMemo:
			obj.KeyDown(key)
		case *tBtn:
			obj.KeyDown(key)
		case *tListBox:
			obj.KeyDown(key)
		case *tMenu:
			obj.KeyDown(key)
		case *tListFileBox:
			obj.KeyDown(key)
		case *tTable:
			obj.KeyDown(key)
		}
		fmt.Println(key)
		fmt.Println(string(key))
	}
	
	eventDraw()
}


//export eventKeyUp
func eventKeyUp(key int){
	if key == 16 {
		isShiftKeyDown = false
	} else if key == 17 {
		isCtrlKeyDown = false
	} else if key == 18 {
		isAltKeyDown = false
	}
	
	eventDraw()
}


func setSize(node *Node, sizeX int, sizeY int){
	var sizeXpar int
	var sizeYpar int
	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tForm:
			sizeXpar = obj.sizeX
			sizeYpar = obj.sizeY
			obj.sizeX = sizeX
			obj.sizeY = sizeY
		case *tPanel:
			sizeXpar = obj.sizeX
			sizeYpar = obj.sizeY
			obj.sizeX = sizeX
			obj.sizeY = sizeY
		}
	}
	
	
	for i := 0; i < len(node.children); i++ {
		if node.children[i].obj != nil {
		switch obj := node.children[i].obj.(type) {
		case *tBtn:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			case BOTTOM:
				deltaY := sizeYpar - obj.y
				obj.y = sizeY - deltaY
			case CLIENT:
				obj.sizeX = sizeX - obj.x - 2
				obj.sizeY = sizeY - obj.y - 2
			} 
		case *tPanel:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			case CLIENT:
				setSize(node.children[i], sizeX - obj.x - 2, sizeY - obj.y - 2)
			} 
		case *tEdit:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tPassEdit:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			}
		case *tTrackBar:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tLabel:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tCanvas:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tBitBtn:
			/*if obj.alRight {
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX
			}*/
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tMemo:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			case CLIENT:
				obj.sizeX = sizeX - obj.x - 2
				obj.sizeY = sizeY - obj.y - 2
			} 
		case *tCheckBox:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tComboBox:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tListBox:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tTab:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tMenu:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tListFileBox:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			case CLIENT:
				obj.sizeX = sizeX - obj.x - 2
				obj.sizeY = sizeY - obj.y - 2
			} 
		case *tTable:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		case *tImage:
			switch obj.align {
			case RIGHT_TOP:
				deltaX := sizeXpar - obj.x
				obj.x = sizeX - deltaX 
			} 
		}
	}
	}	
}


