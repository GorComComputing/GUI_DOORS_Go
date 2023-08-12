package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


func DrawDesktop(BC int){
    SetColor(BC);
    var p []tPoint
    
	p1 := tPoint{x: 0, y: 0}
	p = append(p, p1)
	
	p2 := tPoint{x: BITMAP_WIDTH-1, y: 0}
	p = append(p, p2)
	
	p3 := tPoint{x: BITMAP_WIDTH-1, y: BITMAP_HEIGHT-2}
	p = append(p, p3)
	
	p4 := tPoint{x: 0, y: BITMAP_HEIGHT-2}
	p = append(p, p4)

    FillPoly(4, p);
}


func DrawTaskbar(x int, y int, sizeX int, sizeY int, BC int){
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
}


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
	node := Node{parent: parent, previous: nil, children: nil, obj: &obj}
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


func DrawEdit(x int, y int, sizeX int, sizeY int, BC int, TC int, text string){
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
    TextOutgl(text, x + 4, y + 8, 1);

    
    SetColor(0x000000);
    LinePP(x, y, x + sizeX, y);
    LinePP(x, y, x, y + sizeY);
}


func DrawLabel(x int, y int, sizeX int, sizeY int, BC int, TC int, text string){
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
    TextOutgl(text, x, y, 1);
}



func DrawWindow(x int, y int, sizeX int, sizeY int, BC int, caption string){
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
    
    p = nil
    SetColor(0x0054E0);
    p1 = tPoint{x: x, y: y}
	p = append(p, p1)
	
	p2 = tPoint{x: x + sizeX, y: y}
	p = append(p, p2)
	
	p3 = tPoint{x: x + sizeX, y: y + 17}
	p = append(p, p3)
	
	p4 = tPoint{x: x, y: y + 17}
	p = append(p, p4)
	
	FillPoly(4, p);

    SetColor(0xF8FCF8);
    SetBackColor(BC);
    TextOutgl(caption, x + 9, y + 6, 1);
    
    DrawBtnCls(x + sizeX - 17, y+2, 15, 15, BC, 0x000000, "X")

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
