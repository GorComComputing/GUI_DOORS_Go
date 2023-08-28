package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tBitBtn struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    caption string
    visible bool
    pressed bool
    enabled bool
    onClick func(*Node)
    onClickStr string
}


func CreateBitBtn(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, caption string, onClick func(*Node)) *Node {
	obj := tBitBtn{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, pressed: false, enabled: true, onClick: onClick}
	node := Node{typ: BIT_BUTTON, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj *tBitBtn) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
	SetColor(obj.BC);
    	var p []tPoint

    	p1 := tPoint{x: parX + obj.x, y: parY + obj.y}
		p = append(p, p1)
	
		p2 := tPoint{x: parX + obj.x + obj.sizeX, y: parY + obj.y}
		p = append(p, p2)
	
		p3 := tPoint{x: parX + obj.x + obj.sizeX, y: parY + obj.y + obj.sizeY}
		p = append(p, p3)
	
		p4 := tPoint{x: parX + obj.x, y: parY + obj.y + obj.sizeY}
		p = append(p, p4)

    	FillPoly(nil, 4, p);
    	
    	
	if obj.pressed {
		SetColor(0x787C78);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x, parY + obj.y + obj.sizeY);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x + obj.sizeX - 2, parY + obj.y+1);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x+1, parY + obj.y + obj.sizeY - 1);
    	SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x+2, parY + obj.y + obj.sizeY - 1, parX + obj.x + obj.sizeX - 1, parY + obj.y + obj.sizeY - 1);
    	LinePP(nil, parX + obj.x + obj.sizeX - 1, parY + obj.y + 1, parX + obj.x + obj.sizeX - 1, parY + obj.y + obj.sizeY - 1);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+1, parY + obj.y + obj.sizeY, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY);
    	LinePP(nil, parX + obj.x + obj.sizeX, parY + obj.y+1, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY);
	} else {
		SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x, parY + obj.y + obj.sizeY);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x + obj.sizeX - 2, parY + obj.y+1);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x+1, parY + obj.y + obj.sizeY - 1);
    	SetColor(0x787C78);
    	LinePP(nil, parX + obj.x+2, parY + obj.y + obj.sizeY - 1, parX + obj.x + obj.sizeX - 1, parY + obj.y + obj.sizeY - 1);
    	LinePP(nil, parX + obj.x + obj.sizeX - 1, parY + obj.y + 1, parX + obj.x + obj.sizeX - 1, parY + obj.y + obj.sizeY - 1);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+1, parY + obj.y + obj.sizeY, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY);
    	LinePP(nil, parX + obj.x + obj.sizeX, parY + obj.y+1, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY);
	}
	
	if obj.enabled {
    	SetColor(obj.TC);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX+obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2 - 3, parY+obj.y + obj.sizeY/2-4, 1);
    } else {
    	SetColor(0x787C78);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX+obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2 - 3, parY+obj.y + obj.sizeY/2-4, 1);
    }
}


func (obj *tBitBtn) RAD(x int, y int){
	var visible string
    		if obj.visible {
    			visible = "true"
    		} else {
    			visible = "false"
    		}
			var enabled string
    		if obj.enabled {
    			enabled = "true"
    		} else {
    			enabled = "false"
    		}
			
			frmProperties.obj.(*tForm).caption = "Properties: BITBTN"
			downX = x 
    		downY = y 
    		mouseIsDown = true
    		lblPropName = CreateLabel(frmProperties, "lblPropName", 5, 20, 95, 20, 0xD8DCC0, 0x000000, "Name", nil)
			editPropName = CreateEdit(frmProperties, "editPropName", 80, 20, 95, 20, 0xF8FCF8, 0x000000, obj.name, nil, editPropNameEnter)
			lblPropLeft = CreateLabel(frmProperties, "lblPropLeft", 5, 40, 95, 20, 0xD8DCC0, 0x000000, "Left", nil)
			editPropLeft = CreateEdit(frmProperties, "editPropLeft", 80, 40, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.x), nil, editPropLeftEnter)
			lblPropTop = CreateLabel(frmProperties, "lblPropTop", 5, 60, 95, 20, 0xD8DCC0, 0x000000, "Top", nil)
			editPropTop = CreateEdit(frmProperties, "editPropTop", 80, 60, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.y), nil, editPropTopEnter)
			lblPropCaption = CreateLabel(frmProperties, "lblPropCaption", 5, 80, 95, 20, 0xD8DCC0, 0x000000, "Caption", nil)
			editPropCaption = CreateEdit(frmProperties, "editPropCaption", 80, 80, 95, 20, 0xF8FCF8, 0x000000, obj.caption, nil, editPropCaptionEnter)
			lblPropBC = CreateLabel(frmProperties, "lblPropBC", 5, 100, 95, 20, 0xD8DCC0, 0x000000, "BC", nil)
			editPropBC = CreateEdit(frmProperties, "editPropBC", 80, 100, 95, 20, 0xF8FCF8, 0x000000, fmt.Sprintf("%x", obj.BC), nil, editPropBCEnter)
			lblPropTC = CreateLabel(frmProperties, "lblPropTC", 5, 120, 95, 20, 0xD8DCC0, 0x000000, "TC", nil)
			editPropTC = CreateEdit(frmProperties, "editPropTC", 80, 120, 95, 20, 0xF8FCF8, 0x000000, fmt.Sprintf("%x", obj.TC), nil, editPropTCEnter)
			lblPropWidth = CreateLabel(frmProperties, "lblPropWidth", 5, 140, 95, 20, 0xD8DCC0, 0x000000, "Width", nil)
			editPropWidth = CreateEdit(frmProperties, "editPropWidth", 80, 140, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeX), nil, editPropWidthEnter)
			lblPropHeight = CreateLabel(frmProperties, "lblPropHeight", 5, 160, 95, 20, 0xD8DCC0, 0x000000, "Height", nil)
			editPropHeight = CreateEdit(frmProperties, "editPropHeight", 80, 160, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeY), nil, editPropHeightEnter)
			lblPropVisible = CreateLabel(frmProperties, "lblPropVisible", 5, 180, 95, 20, 0xD8DCC0, 0x000000, "Visible", nil)
			editPropVisible = CreateEdit(frmProperties, "editPropVisible", 80, 180, 95, 20, 0xF8FCF8, 0x000000, visible, nil, editPropVisibleEnter)
			lblPropEnabled = CreateLabel(frmProperties, "lblPropEnabled", 5, 200, 95, 20, 0xD8DCC0, 0x000000, "Enabled", nil)
			editPropEnabled = CreateEdit(frmProperties, "editPropEnabled", 80, 200, 95, 20, 0xF8FCF8, 0x000000, enabled, nil, editPropEnabledEnter)
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 240, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 240, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
}


func (obj *tBitBtn) KeyDown(key int){

}


func (obj *tBitBtn) Click(){
	fmt.Println("CLICKED: " + obj.caption)
			if obj.onClick != nil && obj.enabled {
				obj.onClick(list[len(list)-1])
			}
}


func (obj *tBitBtn) MouseMove(x int, y int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tBitBtn) MouseDown(x int, y int){
	// RAD
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.RAD(x, y)
		} else {
			// Нажатие на кнопку
			if obj.enabled {
    			bitbtnPressed = obj
				obj.pressed = true	
			}
		}
}
