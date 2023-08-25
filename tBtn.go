package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tBtn struct{
	name string
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
    onClickStr string
}


func CreateBtn(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, caption string, onClick func(*Node)) *Node {
	obj := tBtn{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, pressed: false, enabled: true, onClick: onClick}
	node := Node{typ: BUTTON, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tBtn) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x+2, parY + obj.y+2, parX + obj.x + obj.sizeX-2, parY + obj.y + obj.sizeY-2)
	
	var x1,x2,x3,x4,y1,y2,y3,y4 int = 0,0,0,0,0,0,0,0
	var top,left,right,bottom bool = true,true,true,true

	if parX + obj.x < parX + 2 {
		x1 = parX + 2
		x4 = parX + 2
		left = false
	} else {
		x1 = obj.x
		x4 = obj.x
	}
	if obj.x + obj.sizeX > parSizeX - 2 {
		x2 = parSizeX - 2
		x3 = parSizeX - 2
		right = false
	} else {
		x2 = obj.x + obj.sizeX
		x3 = obj.x + obj.sizeX
	}
	if parY + obj.y < parY + 2 {
		y1 = parY + 2
		y2 = parY + 2
		top = false
	} else {
		y1 = obj.y
		y2 = obj.y
	}
	if obj.y + obj.sizeY > parSizeY - 2 {
		y3 = parSizeY - 2
		y4 = parSizeY - 2
		bottom = false
	} else {
		y3 = obj.y + obj.sizeY
		y4 = obj.y + obj.sizeY
	}
	
	SetColor(obj.BC); //obj.BC
    var p []tPoint

    p1 := tPoint{x: parX + x1, y: parY + y1}
	p = append(p, p1)
	
	p2 := tPoint{x: parX + x2, y: parY + y2}
	p = append(p, p2)
	
	p3 := tPoint{x: parX + x3, y: parY + y3}
	p = append(p, p3)
	
	p4 := tPoint{x: parX + x4, y: parY + y4}
	p = append(p, p4)

    FillPoly(nil, 4, p);
    	
	if obj.pressed {
		SetColor(0x787C78);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x, parY + obj.y + obj.sizeY);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x + obj.sizeX - 2, parY + obj.y+1);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x+1, parY + obj.y + obj.sizeY - 1);
    	SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x+2, parY + obj.y + obj.sizeY - 1, parX + obj.x + obj.sizeX - 1, parY + obj.y + obj.sizeY - 1);
    	LinePP(nil, parX + obj.x + obj.sizeX - 1, parY + obj.y + 1, parX + obj.x + obj.sizeX - 1, parY + obj.y + obj.sizeY - 1);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+1, parY + obj.y + obj.sizeY, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY);
    	LinePP(nil, parX + obj.x + obj.sizeX, parY + obj.y+1, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY);
	} else {
		SetColor(0xF8FCF8);
		if top {
    		LinePP(nil, parX + x1, parY + y1, parX + x2, parY + y2);
    	}
    	if left {
    		LinePP(nil, parX + x1, parY + y1, parX + x4, parY + y4);
    	}
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + x1+1, parY + y1+1, parX + x2 - 2, parY + y2+1);
    	LinePP(nil, parX + x1+1, parY + y1+1, parX + x4 + 1, parY + y4 - 1);
    	SetColor(0x787C78);
    	if bottom {
    		LinePP(nil, parX + x4 + 2, parY + y4 - 1, parX + x3 - 1, parY + y3 - 1);
    	}
    	if right {
    		LinePP(nil, parX + x2 - 1, parY + y2 + 1, parX + x3 - 1, parY + y3 - 1);
    	}
    	SetColor(0x000000);
    	if obj.y + obj.sizeY - 1 < parSizeY - 2 {
    		LinePP(nil, parX + x4 + 1, parY + y4, parX + x3, parY + y3);
    	}
    	if obj.x + obj.sizeX - 1 < parSizeX - 2 {
    		LinePP(nil, parX + x2, parY + y2 + 1, parX + x3, parY + y3);
    	}
	}
	
	if obj.enabled {
    	SetColor(obj.TC);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX + obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2, parY + obj.y + obj.sizeY/2-4, 1);
    } else {
    	SetColor(0x787C78);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX + obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2, parY + obj.y + obj.sizeY/2-4, 1);
    }
    
	//SetViewPort(0, 0, GETMAX_X, GETMAX_Y)    
}
