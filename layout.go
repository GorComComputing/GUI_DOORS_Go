package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"
    //"reflect"

)


var mouseIsDown bool = false
var cursor bool

var layout_obj = tForm{x: 0, y: 0, sizeX: BITMAP_WIDTH-1, sizeY: BITMAP_HEIGHT-2, BC: 0x0080C0, mode: NONE, caption: "", visible: true, onClick: nil}
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
)

type tWinComponents interface {
   Draw(parX int, parY int, parSizeX int, parSizeY int)
}


type Node struct {
	typ tComponents 
    parent *Node
    previous *Node
    children []*Node
    obj tWinComponents 
}
         

func DrawNode(node *Node){
	
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
		}
	}
	
	var parX int = 0
	var parY int = 0
	var parSizeX int = GETMAX_X
	var parSizeY int = GETMAX_Y
	if node.parent != nil && node.parent.obj != nil {
		switch obj := node.parent.obj.(type) {
		case *tBtn:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tForm:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tPanel:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tEdit:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tLabel:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tCanvas:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tBitBtn:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tMemo:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		}
	}
	
	if node.obj != nil && visible  {
		node.obj.Draw(parX, parY, parSizeX, parSizeY)
	}
	
	if node.children != nil && visible {
			for i := 0; i < len(node.children); i++ { 
				DrawNode(node.children[i])
			}
	}
		
	return
}


//export eventClick
func eventClick(x int, y int)  {
	fmt.Println("Event: " + strconv.Itoa(x) + " " + strconv.Itoa(y))
	list = nil
	ClickRecurs(&layout, x, y, 0, 0)
	
	switch list[len(list)-1].obj.(type) {
	case *tBtn:
		fmt.Println("CLICKED: " + list[len(list)-1].obj.(*tBtn).caption)
		if list[len(list)-1].obj.(*tBtn).onClick != nil && list[len(list)-1].obj.(*tBtn).enabled {
			list[len(list)-1].obj.(*tBtn).onClick(list[len(list)-1])
		}
	case *tBitBtn:
		fmt.Println("CLICKED: " + list[len(list)-1].obj.(*tBitBtn).caption)
		if list[len(list)-1].obj.(*tBitBtn).onClick != nil && list[len(list)-1].obj.(*tBitBtn).enabled {
			list[len(list)-1].obj.(*tBitBtn).onClick(list[len(list)-1])
		}
		
	}
}


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
		}
	}
	
	if node.parent != nil && node.parent.obj != nil {
		switch obj := node.parent.obj.(type) {
		case *tBtn:
			parX = obj.x
			parY = obj.y
		case *tForm:
			parX = obj.x
			parY = obj.y
		case *tPanel:
			parX = obj.x
			parY = obj.y
		case *tEdit:
			parX = obj.x
			parY = obj.y
		case *tLabel:
			parX = obj.x
			parY = obj.y
		case *tCanvas:
			parX = obj.x
			parY = obj.y
		case *tBitBtn:
			parX = obj.x
			parY = obj.y
		case *tMemo:
			parX = obj.x
			parY = obj.y
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
			}
		case *tForm:
			if (parX+node.obj.(*tForm).x) < x && 
			(parX+node.obj.(*tForm).x + node.obj.(*tForm).sizeX) > x && 
			(parY+node.obj.(*tForm).y) < y && 
			(parY+node.obj.(*tForm).y + node.obj.(*tForm).sizeY) > y {
				list = append(list, node)
			}
		case *tEdit:
			if (parX+node.obj.(*tEdit).x) < x && 
			(parX+node.obj.(*tEdit).x + node.obj.(*tEdit).sizeX) > x && 
			(parY+node.obj.(*tEdit).y) < y && 
			(parY+node.obj.(*tEdit).y + node.obj.(*tEdit).sizeY) > y {
				list = append(list, node)
			}
		case *tLabel:
			if (parX+node.obj.(*tLabel).x) < x && 
			(parX+node.obj.(*tLabel).x + node.obj.(*tLabel).sizeX) > x && 
			(parY+node.obj.(*tLabel).y) < y && 
			(parY+node.obj.(*tLabel).y + node.obj.(*tLabel).sizeY) > y {
				list = append(list, node)
			}
		case *tPanel:
			if (parX+node.obj.(*tPanel).x) < x && 
			(parX+node.obj.(*tPanel).x + node.obj.(*tPanel).sizeX) > x && 
			(parY+node.obj.(*tPanel).y) < y && 
			(parY+node.obj.(*tPanel).y + node.obj.(*tPanel).sizeY) > y {
				list = append(list, node)
			}
		case *tCanvas:
			if (parX+node.obj.(*tCanvas).x) < x && 
			(parX+node.obj.(*tCanvas).x + node.obj.(*tCanvas).sizeX) > x && 
			(parY+node.obj.(*tCanvas).y) < y && 
			(parY+node.obj.(*tCanvas).y + node.obj.(*tCanvas).sizeY) > y {
				list = append(list, node)
			}
		case *tBitBtn:
			if (parX+node.obj.(*tBitBtn).x) < x && 
			(parX+node.obj.(*tBitBtn).x + node.obj.(*tBitBtn).sizeX) > x && 
			(parY+node.obj.(*tBitBtn).y) < y && 
			(parY+node.obj.(*tBitBtn).y + node.obj.(*tBitBtn).sizeY) > y {
				list = append(list, node)
			}
		case *tMemo:
			if (parX+node.obj.(*tMemo).x) < x && 
			(parX+node.obj.(*tMemo).x + node.obj.(*tMemo).sizeX) > x && 
			(parY+node.obj.(*tMemo).y) < y && 
			(parY+node.obj.(*tMemo).y + node.obj.(*tMemo).sizeY) > y {
				list = append(list, node)
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


func sortChildren(i int) {
	tmp := layout.children[i]
	copy(layout.children[i:], layout.children[i+1:])
	layout.children[len(layout.children)-1] = tmp
	layout.children[len(layout.children)-2].obj.(*tForm).focused = false
	layout.children[len(layout.children)-1].obj.(*tForm).focused = true
}


var downX int = 0
var downY int = 0
var btnPressed *tBtn = nil
var bitbtnPressed *tBitBtn = nil

//export eventMouseDown
func eventMouseDown(x int, y int)  {
	list = nil
	ClickRecurs(&layout, x, y, 0, 0)
	
	i := findNode(list[len(list)-1])
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
	if layout.children[len(layout.children)-1].obj.(*tForm).focus != nil {
		switch obj := layout.children[len(layout.children)-1].obj.(*tForm).focus.obj.(type) {
    	case *tEdit:
			obj.focused = false
		case *tMemo:
			obj.focused = false
		}
	}
	layout.children[len(layout.children)-1].obj.(*tForm).focus = list[len(list)-1]
	
	switch obj := list[len(list)-1].obj.(type) {
	case *tForm:
		if (obj.mode == WIN) &&
			(obj.x) < x && 
			(obj.x + obj.sizeX) > x && 
			(obj.y) < y && 
			(obj.y + 17) > y {
				downX = x 
    			downY = y 
    			mouseIsDown = true
    		}
    case *tBtn:
    	if obj.enabled {
    		btnPressed = obj
			obj.pressed = true	
		}
	case *tEdit:
    	if obj.enabled {
			obj.focused = true	
		}
	case *tBitBtn:
    	if obj.enabled {
    		bitbtnPressed = obj
			obj.pressed = true	
		}
	case *tMemo:
    	if obj.enabled {
			obj.focused = true	
		}
	}
}


//export eventMouseUp
func eventMouseUp(x int, y int)  {
	//if(mouseIsDown) mouseClick(e)
    mouseIsDown = false

	if btnPressed != nil {
		btnPressed.pressed = false
		btnPressed = nil
		bitbtnPressed = nil
	}

}


//export eventMouseMove
func eventMouseMove(x int, y int)  {
	if !mouseIsDown {return}
	
	switch list[len(list)-1].obj.(type) {
	case *tBtn:
		list[len(list)-1].obj.(*tBtn).x += x - downX
    	list[len(list)-1].obj.(*tBtn).y += y - downY
    case *tForm:
		list[len(list)-1].obj.(*tForm).x += x - downX
    	list[len(list)-1].obj.(*tForm).y += y - downY				
	}
	
    downX = x 
    downY = y
    return
}


//export keyDown
func keyDown(key int){
	if layout.children[len(layout.children)-1].obj.(*tForm).focus != nil {
		switch obj := layout.children[len(layout.children)-1].obj.(*tForm).focus.obj.(type) {
    	case *tEdit:
    		if key == 8 {
    			if len(obj.text) > 0 {
    				obj.text = obj.text[:len(obj.text)-1]
    				obj.curX--
    			}
    		} else if key == 37 {
    			if obj.curX > 0 {
    				obj.curX--
    			}
    		} else if key == 39 {
    			if obj.curX < len(obj.text) {
    				obj.curX++
    			}
    		} else if key == 36 {
    				obj.curX = 0
    		} else if key == 35 {
    				obj.curX = len(obj.text)
    		} else {
    			obj.text = obj.text[:obj.curX] + string(key) + obj.text[obj.curX:]
				obj.curX++
			}
		case *tMemo:
    		if key == 8 {
    			if len(obj.text) > 0 {
    				obj.text = obj.text[:obj.pos-1] + obj.text[obj.pos:]
    				obj.curX--
    				obj.pos--
    			}
    		} else if key == 13 {
    			obj.text = obj.text[:obj.line_start + obj.pos] + string(key) + obj.text[obj.line_start + obj.pos:]
    			//obj.pos++
    			obj.line_start += obj.pos+1
    			obj.curY++
    			obj.curX = 0
    			obj.pos = 0
    		} else if key == 37 {
    			if obj.curX > 0 {
    				obj.curX--
    				obj.pos--
    			}
    		} else if key == 39 {
    			if ((obj.line_start + obj.pos) < (len(obj.text)-1)) && ((obj.line_start + obj.pos+1) != 13) {
    				obj.curX++
    				obj.pos++
    			}
    		} else if key == 38 {
    			oldNL := obj.line_start
    			if obj.line_start != 0 { 
    				obj.curY--
    				obj.line_start = findLeft(obj.text, obj.line_start)
    				if obj.curX > oldNL - obj.line_start - 1 {
    					obj.curX = oldNL - obj.line_start - 1
    					obj.pos = oldNL - obj.line_start - 1
    				}
    			}
    			
    		} else if key == 40 {
    			right := findRight(obj.text, obj.line_start)
    			if obj.line_start == right || right == len(obj.text)-1 {
    				return
    			}
    			obj.line_start = right
    			if obj.curX > len(obj.text)-1 - obj.line_start {
    				obj.curX = len(obj.text) - obj.line_start
    				obj.pos = len(obj.text) - obj.line_start
    			}
    			obj.curY++
			} else {
				obj.text = obj.text[:obj.line_start + obj.pos] + string(key) + obj.text[obj.line_start + obj.pos:]
				obj.pos++
				obj.curX++
			}
			
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


/*func findNodeByObj(obj tWinComponents) *Node {
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
}*/
