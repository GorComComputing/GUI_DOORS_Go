package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tEdit struct{
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    text string
    visible bool
    onClick func(*Node)
}


func CreateEdit(parent *Node, x int, y int, sizeX int, sizeY int, BC int, TC int, text string, onClick func(*Node)) *Node {
	obj := tEdit{x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, text: text, visible: true, onClick: onClick}
	node := Node{typ: EDIT, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tEdit) Draw(){
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
    TextOutgl(obj.text, obj.x + 4, obj.y + 8, 1);

    
    SetColor(0x000000);
    LinePP(obj.x, obj.y, obj.x + obj.sizeX, obj.y);
    LinePP(obj.x, obj.y, obj.x, obj.y + obj.sizeY);
}


