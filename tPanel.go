package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tPanel struct{
    x int
    y int
    sizeX int
    sizeY int
    BC int
    visible bool
    onClick func(*Node)
}


func CreatePanel(parent *Node, x int, y int, sizeX int, sizeY int, BC int, onClick func(*Node)) *Node {
	obj := tPanel{x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, visible: true, onClick: onClick}
	node := Node{typ: PANEL, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tPanel) Draw(){
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


/*func DrawTaskbar(x int, y int, sizeX int, sizeY int, BC int){
    SetColor(BC);
    var p []tPoint

    p1 := tPoint{x: x, y: y}
	p = append(p, p1)
	
	p2 := tPoint{x: x + sizeX, y: y}
	p = append(p, p2)
	
	p3 := tPoint{x: x + sizeX, y: y + sizeY}
	p = append(p, p3)
	
	p4 := tPoint{x: x, y: y + sizeY}
	p = append(p, p4)

    FillPoly(4, p);

    SetColor(0xA0DC88);
    LinePP(x, y, x + sizeX, y);
    LinePP(x, y, x, y + sizeY);
    SetColor(0x80C848);
    LinePP(x+1, y+1, x + sizeX - 2, y+1);
    LinePP(x+1, y+1, x+1, y + sizeY - 1);
    SetColor(0x089000);
    LinePP(x+2, y + sizeY - 1, x + sizeX - 1, y + sizeY - 1);
    LinePP(x + sizeX - 1, y + 1, x + sizeX - 1, y + sizeY - 1);
    SetColor(0x005C00);
    LinePP(x, y + sizeY, x + sizeX, y + sizeY);
    LinePP(x + sizeX, y, x + sizeX, y + sizeY);
}*/

