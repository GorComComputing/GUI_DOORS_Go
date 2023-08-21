package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tMemo struct{
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
    curY int
    pos int
    line_start int
    onClick func(*Node)
}


func CreateMemo(parent *Node, x int, y int, sizeX int, sizeY int, BC int, TC int, text string, onClick func(*Node)) *Node {
	obj := tMemo{x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, text: text, visible: true, enabled: true, curX: 0, curY: 0, pos: 0, line_start: 0, onClick: onClick}
	node := Node{typ: MEMO, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tMemo) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
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
    TextOutgl(nil, obj.text, parX+obj.x + 4, parY+obj.y + 4, 1);
    if obj.focused && cursor {
    	TextOutgl(nil, "|", parX+obj.x + 4 + obj.curX*7, parY+obj.y + 4 + (7+2)*obj.curY, 1);
    }

    
    SetColor(0x000000);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);
}



