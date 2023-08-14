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
    onClick func(*Node)
}


func CreateBtn(parent *Node, x int, y int, sizeX int, sizeY int, BC int, TC int, caption string, onClick func(*Node)) *Node {
	obj := tBtn{x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, onClick: onClick}
	node := Node{typ: BUTTON, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tBtn) Draw(){
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

    FillPoly(4, p);

    SetColor(obj.TC);
    SetBackColor(obj.BC);
    TextOutgl(obj.caption, obj.x + 8, obj.y + 8, 1);

    
    SetColor(0xF8FCF8);
    LinePP(obj.x, obj.y, obj.x + obj.sizeX, obj.y);
    LinePP(obj.x, obj.y, obj.x, obj.y + obj.sizeY);
    SetColor(0xE0E0E0);
    LinePP(obj.x+1, obj.y+1, obj.x + obj.sizeX - 2, obj.y+1);
    LinePP(obj.x+1, obj.y+1, obj.x+1, obj.y + obj.sizeY - 1);
    SetColor(0x787C78);
    LinePP(obj.x+2, obj.y + obj.sizeY - 1, obj.x + obj.sizeX - 1, obj.y + obj.sizeY - 1);
    LinePP(obj.x + obj.sizeX - 1, obj.y + 1, obj.x + obj.sizeX - 1, obj.y + obj.sizeY - 1);
    SetColor(0x000000);
    LinePP(obj.x, obj.y + obj.sizeY, obj.x + obj.sizeX, obj.y + obj.sizeY);
    LinePP(obj.x + obj.sizeX, obj.y, obj.x + obj.sizeX, obj.y + obj.sizeY);
}
