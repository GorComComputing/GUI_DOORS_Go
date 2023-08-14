package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"
    //"reflect"

)

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
)

type tWinComponents interface {
   Draw()
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
		}
	}
	
	if node.obj != nil && visible  {
		node.obj.Draw()
		
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
	ClickRecurs(&layout, x, y)
	switch list[len(list)-1].obj.(type) {
	case *tBtn:
		fmt.Println("CLICKED: " + list[len(list)-1].obj.(*tBtn).caption)
		if list[len(list)-1].obj.(*tBtn).onClick != nil {
			list[len(list)-1].obj.(*tBtn).onClick(list[len(list)-1])
		}
		
	}
}


func ClickRecurs(node *Node, x int, y int) {
	
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
		}
	}
	
	if node.obj != nil && visible {
		switch node.obj.(type) {
		case *tBtn:
			if node.obj.(*tBtn).x < x && (node.obj.(*tBtn).x + node.obj.(*tBtn).sizeX) > x && node.obj.(*tBtn).y < y && (node.obj.(*tBtn).y + node.obj.(*tBtn).sizeY) > y {
				list = append(list, node)
			}
		}
	}
			
	if node.children != nil && visible {
		for i := 0; i < len(node.children); i++ { 
			ClickRecurs(node.children[i], x, y)
		}
	}
	return
}
