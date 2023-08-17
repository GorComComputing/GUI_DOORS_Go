package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tBtn struct{
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    caption string
    visible bool
    pressed bool
    enabled bool
    onClick func(*Node)
}


func CreateBtn(parent *Node, x int, y int, sizeX int, sizeY int, BC int, TC int, caption string, onClick func(*Node)) *Node {
	obj := tBtn{x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, pressed: false, enabled: true, onClick: onClick}
	node := Node{typ: BUTTON, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tBtn) Draw(parX int, parY int){
	if obj.pressed {
		SetColor(0x0000FF);
    	var p []tPoint

    	p1 := tPoint{x: parX+obj.x, y: parY+obj.y}
		p = append(p, p1)
	
		p2 := tPoint{x: parX+obj.x + obj.sizeX, y: parY+obj.y}
		p = append(p, p2)
	
		p3 := tPoint{x: parX+obj.x + obj.sizeX, y: parY+obj.y + obj.sizeY}
		p = append(p, p3)
	
		p4 := tPoint{x: parX+obj.x, y: parY+obj.y + obj.sizeY}
		p = append(p, p4)

    	FillPoly(nil, 4, p);

    	SetColor(obj.TC);
    	SetBackColor(0x0000FF);
    	TextOutgl(nil, obj.caption, parX+obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2, parY+obj.y + obj.sizeY/2-4, 1);	
	} else if obj.enabled {

    	SetColor(obj.BC);
    	var p []tPoint

    	p1 := tPoint{x: parX+obj.x, y: parY+obj.y}
		p = append(p, p1)
	
		p2 := tPoint{x: parX+obj.x + obj.sizeX, y: parY+obj.y}
		p = append(p, p2)
	
		p3 := tPoint{x: parX+obj.x + obj.sizeX, y: parY+obj.y + obj.sizeY}
		p = append(p, p3)
	
		p4 := tPoint{x: parX+obj.x, y: parY+obj.y + obj.sizeY}
		p = append(p, p4)

    	FillPoly(nil, 4, p);

    	SetColor(obj.TC);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX+obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2, parY+obj.y + obj.sizeY/2-4, 1);

    
    	SetColor(0xF8FCF8);
    	LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    	LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX+obj.x+1, parY+obj.y+1, parX+obj.x + obj.sizeX - 2, parY+obj.y+1);
    	LinePP(nil, parX+obj.x+1, parY+obj.y+1, parX+obj.x+1, parY+obj.y + obj.sizeY - 1);
    	SetColor(0x787C78);
    	LinePP(nil, parX+obj.x+2, parY+obj.y + obj.sizeY - 1, parX+obj.x + obj.sizeX - 1, parY+obj.y + obj.sizeY - 1);
    	LinePP(nil, parX+obj.x + obj.sizeX - 1, parY+obj.y + 1, parX+obj.x + obj.sizeX - 1, parY+obj.y + obj.sizeY - 1);
    	SetColor(0x000000);
    	LinePP(nil, parX+obj.x, parY+obj.y + obj.sizeY, parX+obj.x + obj.sizeX, parY+obj.y + obj.sizeY);
    	LinePP(nil, parX+obj.x + obj.sizeX, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y + obj.sizeY);
    } else if !(obj.enabled) {
    	SetColor(0x333333);
    	var p []tPoint

    	p1 := tPoint{x: parX+obj.x, y: parY+obj.y}
		p = append(p, p1)
	
		p2 := tPoint{x: parX+obj.x + obj.sizeX, y: parY+obj.y}
		p = append(p, p2)
	
		p3 := tPoint{x: parX+obj.x + obj.sizeX, y: parY+obj.y + obj.sizeY}
		p = append(p, p3)
	
		p4 := tPoint{x: parX+obj.x, y: parY+obj.y + obj.sizeY}
		p = append(p, p4)

    	FillPoly(nil, 4, p);

    	SetColor(obj.TC);
    	SetBackColor(0x333333);
    	TextOutgl(nil, obj.caption, parX+obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2, parY+obj.y + obj.sizeY/2-4, 1);
    }
    
    
}
