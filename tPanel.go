package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

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


func (obj *tPanel) Draw(parX int, parY int, parSizeX int, parSizeY int){
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


func (obj *tPanel) RAD(x int, y int){
	var mode string
    		if obj.mode == NONE {
    			mode = "NONE"
    		} else if obj.mode == WIN {
    			mode = "WIN"
    		} else if obj.mode == FLAT {
    			mode = "FLAT"
    		} else if obj.mode == TASK {
    			mode = "TASK"
    		} 
    		var visible string
    		if obj.visible {
    			visible = "true"
    		} else {
    			visible = "false"
    		} 
    		
			frmProperties.obj.(*tForm).caption = "Properties: PANEL"
			downX = x 
    		downY = y 
    		mouseIsDown = true
    		lblPropName = CreateLabel(frmProperties, "lblPropName", 5, 20, 95, 20, 0xD8DCC0, 0x000000, "Name", nil)
			editPropName = CreateEdit(frmProperties, "editPropName", 80, 20, 95, 20, 0xF8FCF8, 0x000000, obj.name, nil, editPropNameEnter)
			lblPropLeft = CreateLabel(frmProperties, "lblPropLeft", 5, 40, 95, 20, 0xD8DCC0, 0x000000, "Left", nil)
			editPropLeft = CreateEdit(frmProperties, "editPropLeft", 80, 40, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.x), nil, editPropLeftEnter)
			lblPropTop = CreateLabel(frmProperties, "lblPropTop", 5, 60, 95, 20, 0xD8DCC0, 0x000000, "Top", nil)
			editPropTop = CreateEdit(frmProperties, "editPropTop", 80, 60, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.y), nil, editPropTopEnter)
			lblPropBC = CreateLabel(frmProperties, "lblPropBC", 5, 80, 95, 20, 0xD8DCC0, 0x000000, "BC", nil)
			editPropBC = CreateEdit(frmProperties, "editPropBC", 80, 80, 95, 20, 0xF8FCF8, 0x000000, fmt.Sprintf("%x", obj.BC), nil, editPropBCEnter)
			lblPropWidth = CreateLabel(frmProperties, "lblPropWidth", 5, 100, 95, 20, 0xD8DCC0, 0x000000, "Width", nil)
			editPropWidth = CreateEdit(frmProperties, "editPropWidth", 80, 100, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeX), nil, editPropWidthEnter)
			lblPropHeight = CreateLabel(frmProperties, "lblPropHeight", 5, 120, 95, 20, 0xD8DCC0, 0x000000, "Height", nil)
			editPropHeight = CreateEdit(frmProperties, "editPropHeight", 80, 120, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeY), nil, editPropHeightEnter)
			lblPropMode = CreateLabel(frmProperties, "lblPropMode", 5, 140, 95, 20, 0xD8DCC0, 0x000000, "Mode", nil)
			editPropMode = CreateEdit(frmProperties, "editPropMode", 80, 140, 95, 20, 0xF8FCF8, 0x000000, mode, nil, editPropModeEnter)
			lblPropVisible = CreateLabel(frmProperties, "lblPropVisible", 5, 160, 95, 20, 0xD8DCC0, 0x000000, "Visible", nil)
			editPropVisible = CreateEdit(frmProperties, "editPropVisible", 80, 160, 95, 20, 0xF8FCF8, 0x000000, visible, nil, editPropVisibleEnter)
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 200, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 200, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
}


func (obj *tPanel) KeyDown(key int){

}


func (obj *tPanel) Click(){

}


func (obj *tPanel) MouseMove(x int, y int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tPanel) MouseDown(x int, y int){
	// RAD
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.RAD(x, y)
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

