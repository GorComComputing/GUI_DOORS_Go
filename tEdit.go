package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tEdit struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    text string
    visible bool
    focused bool
    enabled bool
    curX int
    onClick func(*Node)
    onClickStr string
    onEnter func(*Node)
    onEnterStr string
}


func CreateEdit(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, text string, onClick func(*Node), onEnter func(*Node)) *Node {
	obj := tEdit{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, text: text, visible: true, enabled: true, curX: 0, onClick: onClick, onEnter: onEnter}
	node := Node{typ: EDIT, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tEdit) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
	
	if obj.enabled {
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
    } else {
    SetColor(0xBFBFBF);
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
	}

    SetColor(obj.TC);
    SetBackColor(obj.BC);
    if (len(obj.text)+1)*7 > obj.sizeX {
    	TextOutgl(nil, obj.text, parX+obj.x + obj.sizeX - (len(obj.text)+1)*7, parY+obj.y + obj.sizeY/2-4, 1);
    } else {
    	TextOutgl(nil, obj.text, parX+obj.x + 4, parY+obj.y + obj.sizeY/2-4, 1);
    }
    if obj.focused && cursor {
    	if (len(obj.text)+1)*7 > obj.sizeX {
    		TextOutgl(nil, "|", parX+obj.x + 4+obj.curX*8 + obj.sizeX - (len(obj.text)+1)*7, parY+obj.y + obj.sizeY/2-4, 1);
    	} else {
    		TextOutgl(nil, "|", parX+obj.x + 4+obj.curX*8, parY+obj.y + obj.sizeY/2-4, 1);
    	}
    }

    
    SetColor(0x000000);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);
}


