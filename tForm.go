package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tForm struct{
    x int
    y int
    sizeX int
    sizeY int
    BC int
    mode tMode
    caption string
    visible bool
    focused bool
    onClick func(*Node)
}

type tMode int

const (
    NONE tMode = iota
    WIN
)


func CreateForm(parent *Node, x int, y int, sizeX int, sizeY int, BC int, mode tMode, caption string, visible bool, onClick func(*Node)) *Node {
	obj := tForm{x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, mode: mode, caption: caption, visible: visible, onClick: onClick}
	node := Node{typ: FORM, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	
	if obj.mode == WIN {
		CreateBtn(&node, obj.sizeX - 17, 2, 15, 15, 0xD8DCC0, 0x000000, "X", formClose)
	}
	return &node
}


func formClose(node *Node){
	node.parent.obj.(*tForm).visible = false
}


func (obj tForm) Draw(parX int, parY int){
    SetColor(obj.BC);
    var p []tPoint

    p1 := tPoint{x: obj.x, y: obj.y}
	p = append(p, p1)
	
	p2 := tPoint{x: obj.x + obj.sizeX, y: obj.y}
	p = append(p, p2)
	
	p3 := tPoint{x: obj.x + obj.sizeX, y: obj.y + obj.sizeY}
	p = append(p, p3)
	
	p4 := tPoint{x: obj.x, y: obj.y + obj.sizeY}
	p = append(p, p4)

    FillPoly(nil, 4, p);
    
    if obj.mode == WIN {
    	if obj.focused {
    		SetColor(0x0054E0)
    	} else {
    		SetColor(0x80A8E8)
    	}
    	p = nil
    	p1 = tPoint{x: obj.x, y: obj.y}
		p = append(p, p1)
	
		p2 = tPoint{x: obj.x + obj.sizeX, y: obj.y}
		p = append(p, p2)
	
		p3 = tPoint{x: obj.x + obj.sizeX, y: obj.y + 17}
		p = append(p, p3)
	
		p4 = tPoint{x: obj.x, y: obj.y + 17}
		p = append(p, p4)
	
		FillPoly(nil, 4, p);
		
		if obj.focused {
    		SetColor(0xF8FCF8)
    		SetBackColor(0x0054E0);
    	} else {
    		SetColor(0x787C78)
    		SetBackColor(0x80A8E8);
    	}
    	TextOutgl(nil, obj.caption, obj.x + 9, obj.y + 6, 1);
    }

    SetColor(0xF8FCF8);
    LinePP(nil, obj.x, obj.y, obj.x + obj.sizeX, obj.y);
    LinePP(nil, obj.x, obj.y, obj.x, obj.y + obj.sizeY);
    SetColor(0xE0E0E0);
    LinePP(nil, obj.x+1, obj.y+1, obj.x + obj.sizeX - 2, obj.y+1);
    LinePP(nil, obj.x+1, obj.y+1, obj.x+1, obj.y + obj.sizeY - 1);
    SetColor(0x787C78);
    LinePP(nil, obj.x+2, obj.y + obj.sizeY - 1, obj.x + obj.sizeX - 1, obj.y + obj.sizeY - 1);
    LinePP(nil, obj.x + obj.sizeX - 1, obj.y + 1, obj.x + obj.sizeX - 1, obj.y + obj.sizeY - 1);
    SetColor(0x000000);
    LinePP(nil, obj.x, obj.y + obj.sizeY, obj.x + obj.sizeX, obj.y + obj.sizeY);
    LinePP(nil, obj.x + obj.sizeX, obj.y, obj.x + obj.sizeX, obj.y + obj.sizeY);    
}

