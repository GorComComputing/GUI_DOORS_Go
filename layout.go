package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"
    //"reflect"

)


var mouseIsDown bool = false


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
)

type tWinComponents interface {
   Draw(parX int, parY int)
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
		}
	}
	
	var parX int = 0
	var parY int = 0
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
		}
	}
	
	if node.obj != nil && visible  {
		node.obj.Draw(parX, parY)
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
		if list[len(list)-1].obj.(*tBtn).onClick != nil {
			list[len(list)-1].obj.(*tBtn).onClick(list[len(list)-1])
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
}


var downX int = 0
var downY int = 0

//export eventMouseDown
func eventMouseDown(x int, y int)  {
	list = nil
	ClickRecurs(&layout, x, y, 0, 0)
	
	i := findNode(list[len(list)-1])
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
	
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
	}
}


//export eventMouseUp
func eventMouseUp(x int, y int)  {
	//if(mouseIsDown) mouseClick(e)
    mouseIsDown = false
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
