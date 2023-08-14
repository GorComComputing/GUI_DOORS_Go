package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tLabel struct{
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


func CreateLabel(parent *Node, x int, y int, sizeX int, sizeY int, BC int, TC int, caption string, onClick func(*Node)) *Node {
	obj := tLabel{x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, onClick: onClick}
	node := Node{typ: LABEL, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tLabel) Draw(){
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
    TextOutgl(obj.caption, obj.x, obj.y, 1);
}



