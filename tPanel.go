package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tPanel struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    visible bool
    mode tMode
    onClick func(*Node)
    onClickStr string
}


func CreatePanel(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, mode tMode, onClick func(*Node)) *Node {
	obj := tPanel{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, visible: true, mode: mode, onClick: onClick}
	node := Node{typ: PANEL, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tPanel) Draw(parX int, parY int, parSizeX int, parSizeY int){
	//SetViewPort(parX, parY, parX + parSizeX, parY + parSizeY)
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
    
    if obj.mode == TASK {
    	SetColor(0xA0DC88);
    	LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    	LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);
    	SetColor(0x80C848);
    	LinePP(nil, parX+obj.x+1, parY+obj.y+1, parX+obj.x + obj.sizeX - 2, parY+obj.y+1);
    	LinePP(nil, parX+obj.x+1, parY+obj.y+1, parX+obj.x+1, parY+obj.y + obj.sizeY - 1);
    	SetColor(0x089000);
    	LinePP(nil, parX+obj.x+2, parY+obj.y + obj.sizeY - 1, parX+obj.x + obj.sizeX - 1, parY+obj.y + obj.sizeY - 1);
    	LinePP(nil, parX+obj.x + obj.sizeX - 1, parY+obj.y + 1, parX+obj.x + obj.sizeX - 1, parY+obj.y + obj.sizeY - 1);
    	SetColor(0x005C00);
    	LinePP(nil, parX+obj.x, parY+obj.y + obj.sizeY, parX+obj.x + obj.sizeX, parY+obj.y + obj.sizeY);
    	LinePP(nil, parX+obj.x + obj.sizeX, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y + obj.sizeY);
    } else {
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
    }
    
    
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

