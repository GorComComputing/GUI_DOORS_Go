package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tCheckBox struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    caption string
    visible bool
    checked bool
    enabled bool
    onClick func(*Node)
}


func CreateCheckBox(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, caption string, checked bool, onClick func(*Node)) *Node {
	obj := tCheckBox{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, checked: checked, enabled: true, onClick: onClick}
	node := Node{typ: CHECKBOX, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tCheckBox) Draw(parX int, parY int, parSizeX int, parSizeY int){
	const size int = 16
	const size_sm int = size - 8
	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
	SetColor(obj.BC);
    	var p []tPoint

    	p1 := tPoint{x: parX + obj.x, y: parY + obj.y}
		p = append(p, p1)
	
		p2 := tPoint{x: parX + obj.x + size, y: parY + obj.y}
		p = append(p, p2)
	
		p3 := tPoint{x: parX + obj.x + size, y: parY + obj.y + size}
		p = append(p, p3)
	
		p4 := tPoint{x: parX + obj.x, y: parY + obj.y + size}
		p = append(p, p4)

    	FillPoly(nil, 4, p);
    	
    	
    	SetColor(0x787C78);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x + size, parY + obj.y);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x, parY + obj.y + size);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x + size - 2, parY + obj.y+1);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x+1, parY + obj.y + size - 1);
    	SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x+2, parY + obj.y + size - 1, parX + obj.x + size - 1, parY + obj.y + size - 1);
    	LinePP(nil, parX + obj.x + size - 1, parY + obj.y + 1, parX + obj.x + size - 1, parY + obj.y + size - 1);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+1, parY + obj.y + size, parX + obj.x + size, parY + obj.y + size);
    	LinePP(nil, parX + obj.x + size, parY + obj.y+1, parX + obj.x + size, parY + obj.y + size);
    	
    	
	if obj.checked {
		SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x+4, parY + obj.y+4, parX + obj.x + size_sm+4, parY + obj.y+4);
    	LinePP(nil, parX + obj.x+4, parY + obj.y+4, parX + obj.x+4, parY + obj.y + size_sm+4);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+4+1, parY + obj.y+4+1, parX + obj.x + size_sm+4 - 2, parY + obj.y+4+1);
    	LinePP(nil, parX + obj.x+4+1, parY + obj.y+4+1, parX + obj.x+4+1, parY + obj.y + size_sm+4 - 1);
    	SetColor(0x787C78);
    	LinePP(nil, parX + obj.x+4+2, parY + obj.y + size_sm+4 - 1, parX + obj.x + size_sm+4 - 1, parY + obj.y + size_sm+4 - 1);
    	LinePP(nil, parX + obj.x + size_sm+4 - 1, parY + obj.y+4 + 1, parX + obj.x + size_sm+4 - 1, parY + obj.y + size_sm+4 - 1);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+4+1, parY + obj.y + size_sm+4, parX + obj.x + size_sm+4, parY + obj.y + size_sm+4);
    	LinePP(nil, parX + obj.x + size_sm+4, parY + obj.y+4+1, parX + obj.x + size_sm+4, parY + obj.y + size_sm+4);
	} 
	
	if obj.enabled {
    	SetColor(obj.TC);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX+obj.x + 25, parY+obj.y + obj.sizeY/2-4, 1);
    } else {
    	SetColor(0x787C78);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX+obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2 - 3, parY+obj.y + obj.sizeY/2-4, 1);
    }    
}
