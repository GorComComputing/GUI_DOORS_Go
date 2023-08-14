package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tForm struct{
    x int
    y int
    sizeX int
    sizeY int
    BC int
    mode tMode
    caption string
    visible bool
    onClick func(*Node)
}

type tMode int

const (
    NONE tMode = iota
    WIN
)


func CreateForm(parent *Node, x int, y int, sizeX int, sizeY int, BC int, mode tMode, caption string, visible bool, onClick func(*Node)) *Node {
	obj := tForm{x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, mode: mode, caption: caption, visible: visible, onClick: onClick}
	node := Node{typ: FORM, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}



func (obj tForm) Draw(){
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
    
    if obj.mode == WIN {
    	p = nil
    	SetColor(0x0054E0);
    	p1 = tPoint{x: obj.x, y: obj.y}
	p = append(p, p1)
	
	p2 = tPoint{x: obj.x + obj.sizeX, y: obj.y}
	p = append(p, p2)
	
	p3 = tPoint{x: obj.x + obj.sizeX, y: obj.y + 17}
	p = append(p, p3)
	
	p4 = tPoint{x: obj.x, y: obj.y + 17}
	p = append(p, p4)
	
	FillPoly(4, p);

    	SetColor(0xF8FCF8);
    	SetBackColor(obj.BC);
    	TextOutgl(obj.caption, obj.x + 9, obj.y + 6, 1);
    
    	DrawBtnCls(obj.x + obj.sizeX - 17, obj.y+2, 15, 15, 0xD8DCC0, 0x000000, "X")
    }

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


func DrawBtnCls(x int, y int, sizeX int, sizeY int, BC int, TC int, caption string){
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

    SetColor(TC);
    SetBackColor(BC);
    TextOutgl(caption, x+4, y+4, 1);

    
    SetColor(0xF8FCF8);
    LinePP(x, y, x + sizeX, y);
    LinePP(x, y, x, y + sizeY);
    SetColor(0xE0E0E0);
    LinePP(x+1, y+1, x + sizeX - 2, y+1);
    LinePP(x+1, y+1, x+1, y + sizeY - 1);
    SetColor(0x787C78);
    LinePP(x+2, y + sizeY - 1, x + sizeX - 1, y + sizeY - 1);
    LinePP(x + sizeX - 1, y + 1, x + sizeX - 1, y + sizeY - 1);
    SetColor(0x000000);
    LinePP(x, y + sizeY, x + sizeX, y + sizeY);
    LinePP(x + sizeX, y, x + sizeX, y + sizeY);
}
