package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"
    //"strings"
    //"reflect"
	"time"
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
	
	t := time.Now()
	lblFPS.obj.(*tLabel).caption = t.Sub(start).String()
}


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
		
		if !RAD || list[len(list)-1] == cbxRAD || layout.children[len(layout.children)-1] == frmProperties || layout.children[len(layout.children)-1] == frmRAD || layout.children[len(layout.children)-1] == frmCode {
		switch obj := list[len(list)-1].obj.(type) {
		case *tBtn:
			obj.Click(x-X, y-Y)
		case *tBitBtn:
			obj.Click(x-X, y-Y)
		case *tCheckBox:
			obj.Click(x-X, y-Y)
		case *tEdit:
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
		}
	}
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
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			RADElement = list[len(list)-1]
			for i := 0; i < len(layout.children); i++ {
				layout.children[i].obj.(*tForm).isRAD = false
			}
			layout.children[len(layout.children)-1].obj.(*tForm).isRAD = true
			RADFormElement = layout.children[len(layout.children)-1]
			pnlProperties.children = nil
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
	}
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
	}
	}
    downX = x 
    downY = y
}


//export eventKeyDown
func eventKeyDown(key int){
	// Удаление элемента RAD (Del)
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && key == 46 {
		for i := 0; i < len(RADElement.parent.children); i++ {
			if RADElement.parent.children[i] == RADElement {
				copy(RADElement.parent.children[i:], RADElement.parent.children[i+1:])
				RADElement.parent.children[len(RADElement.parent.children)-1] = nil
				RADElement.parent.children = RADElement.parent.children[:len(RADElement.parent.children)-1]
				frmProperties.children = nil
			}		
		}
	}
	// Если в фокусе
	if layout.children[len(layout.children)-1].obj.(*tForm).focus != nil {
		switch obj := layout.children[len(layout.children)-1].obj.(*tForm).focus.obj.(type) {
    	case *tEdit:
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
		}
		fmt.Println(key)
		fmt.Println(string(key))
	}
}


func findLeft(str string, current int) int {
	if current == 0 {
		return 0
	}
	for i := current - 2; i >= 0; i-- {
		if str[i] == 13 {
			return i + 1
		}
	}
	return 0
}


func findRight(str string, current int) int {
	if current == len(str)-1 {
		return current
	}
	for i := current; i < len(str)-1; i++ {
		if str[i] == 13 {
			return i+1
		}
	}
	return len(str)-1
}


