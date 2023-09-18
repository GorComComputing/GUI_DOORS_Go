package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tBtn struct{
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
    align tAlign
    onClick func(*Node)
    onClickStr string
}


func CreateBtn(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, caption string, onClick func(*Node)) *Node {
	obj := tBtn{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, pressed: false, enabled: true, onClick: onClick}
	node := Node{typ: BUTTON, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj *tBtn) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x+2, parY + obj.y+2, parX + obj.x + obj.sizeX-2, parY + obj.y + obj.sizeY-2)
	
	var x1,x2,x3,x4,y1,y2,y3,y4 int = 0,0,0,0,0,0,0,0
	var top,left,right,bottom bool = true,true,true,true

	if parX + obj.x < parX + 2 {
		x1 = parX + 2
		x4 = parX + 2
		left = false
	} else {
		x1 = obj.x
		x4 = obj.x
	}
	if obj.x + obj.sizeX > parSizeX - 2 {
		x2 = parSizeX - 2
		x3 = parSizeX - 2
		right = false
	} else {
		x2 = obj.x + obj.sizeX
		x3 = obj.x + obj.sizeX
	}
	if parY + obj.y < parY + 2 {
		y1 = parY + 2
		y2 = parY + 2
		top = false
	} else {
		y1 = obj.y
		y2 = obj.y
	}
	if obj.y + obj.sizeY > parSizeY - 2 {
		y3 = parSizeY - 2
		y4 = parSizeY - 2
		bottom = false
	} else {
		y3 = obj.y + obj.sizeY
		y4 = obj.y + obj.sizeY
	}
	
	SetColor(obj.BC); //obj.BC
    var p []tPoint

    p1 := tPoint{x: parX + x1, y: parY + y1}
	p = append(p, p1)
	
	p2 := tPoint{x: parX + x2, y: parY + y2}
	p = append(p, p2)
	
	p3 := tPoint{x: parX + x3, y: parY + y3}
	p = append(p, p3)
	
	p4 := tPoint{x: parX + x4, y: parY + y4}
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
		if top {
    		LinePP(nil, parX + x1, parY + y1, parX + x2, parY + y2);
    	}
    	if left {
    		LinePP(nil, parX + x1, parY + y1, parX + x4, parY + y4);
    	}
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + x1+1, parY + y1+1, parX + x2 - 2, parY + y2+1);
    	LinePP(nil, parX + x1+1, parY + y1+1, parX + x4 + 1, parY + y4 - 1);
    	SetColor(0x787C78);
    	if bottom {
    		LinePP(nil, parX + x4 + 2, parY + y4 - 1, parX + x3 - 1, parY + y3 - 1);
    	}
    	if right {
    		LinePP(nil, parX + x2 - 1, parY + y2 + 1, parX + x3 - 1, parY + y3 - 1);
    	}
    	SetColor(0x000000);
    	if obj.y + obj.sizeY - 1 < parSizeY - 2 {
    		LinePP(nil, parX + x4 + 1, parY + y4, parX + x3, parY + y3);
    	}
    	if obj.x + obj.sizeX - 1 < parSizeX - 2 {
    		LinePP(nil, parX + x2, parY + y2 + 1, parX + x3, parY + y3);
    	}
	}
	
	if obj.enabled {
    	SetColor(obj.TC);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX + obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2, parY + obj.y + obj.sizeY/2-4, 1);
    } else {
    	SetColor(0x787C78);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX + obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2, parY + obj.y + obj.sizeY/2-4, 1);
    }   
}


func (obj *tBtn) RAD(x int, y int){
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
    		
			//frmProperties.obj.(*tForm).caption = "Properties: BUTTON"
			downX = x 
    		downY = y 
    		mouseIsDown = true
    		lblPropName = CreateLabel(pnlProperties, "lblPropName", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Name", nil)
			editPropName = CreateEdit(pnlProperties, "editPropName", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.name, nil, editPropNameEnter)
			lblPropLeft = CreateLabel(pnlProperties, "lblPropLeft", 5, 25, 95, 20, 0xD8DCC0, 0x000000, "Left", nil)
			editPropLeft = CreateEdit(pnlProperties, "editPropLeft", 80, 25, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.x), nil, editPropLeftEnter)
			lblPropTop = CreateLabel(pnlProperties, "lblPropTop", 5, 45, 95, 20, 0xD8DCC0, 0x000000, "Top", nil)
			editPropTop = CreateEdit(pnlProperties, "editPropTop", 80, 45, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.y), nil, editPropTopEnter)
			lblPropCaption = CreateLabel(pnlProperties, "lblPropCaption", 5, 65, 95, 20, 0xD8DCC0, 0x000000, "Caption", nil)
			editPropCaption = CreateEdit(pnlProperties, "editPropCaption", 80, 65, 95, 20, 0xF8FCF8, 0x000000, obj.caption, nil, editPropCaptionEnter)
			lblPropBC = CreateLabel(pnlProperties, "lblPropBC", 5, 85, 95, 20, 0xD8DCC0, 0x000000, "BC", nil)
			editPropBC = CreateEdit(pnlProperties, "editPropBC", 80, 85, 95, 20, 0xF8FCF8, 0x000000, fmt.Sprintf("%x", obj.BC), nil, editPropBCEnter)
			lblPropTC = CreateLabel(pnlProperties, "lblPropTC", 5, 105, 95, 20, 0xD8DCC0, 0x000000, "TC", nil)
			editPropTC = CreateEdit(pnlProperties, "editPropTC", 80, 105, 95, 20, 0xF8FCF8, 0x000000, fmt.Sprintf("%x", obj.TC), nil, editPropTCEnter)
			lblPropWidth = CreateLabel(pnlProperties, "lblPropWidth", 5, 125, 95, 20, 0xD8DCC0, 0x000000, "Width", nil)
			editPropWidth = CreateEdit(pnlProperties, "editPropWidth", 80, 125, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeX), nil, editPropWidthEnter)
			lblPropHeight = CreateLabel(pnlProperties, "lblPropHeight", 5, 145, 95, 20, 0xD8DCC0, 0x000000, "Height", nil)
			editPropHeight = CreateEdit(pnlProperties, "editPropHeight", 80, 145, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeY), nil, editPropHeightEnter)
			lblPropVisible = CreateLabel(pnlProperties, "lblPropVisible", 5, 165, 95, 20, 0xD8DCC0, 0x000000, "Visible", nil)
			cmbPropVisible = CreateComboBox(pnlProperties, "cmbPropVisible", 80, 165, 95, 16, 0xF8FCF8, 0x000000, visible, listBool, nil, cmbPropVisibleEnter)
			
			lblPropEnabled = CreateLabel(pnlProperties, "lblPropEnabled", 5, 185, 95, 20, 0xD8DCC0, 0x000000, "Enabled", nil)
			cmbPropEnabled = CreateComboBox(pnlProperties, "cmbPropEnabled", 80, 185, 95, 16, 0xF8FCF8, 0x000000, enabled, listBool, nil, cmbPropEnabledEnter)
			
			lblEvntClick = CreateLabel(pnlEvents, "lblEvntClick", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(pnlEvents, "editEvntClick", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
}


func (obj *tBtn) KeyDown(key int){

}


func (obj *tBtn) Click(x int, y int){
	fmt.Println("CLICKED: " + obj.caption)
	if obj.onClick != nil && obj.enabled {
		obj.onClick(list[len(list)-1])
	}
}


func (obj *tBtn) MouseMove(x int, y int, Xl int, Yl int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && mouseIsDown && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tBtn) MouseDown(x int, y int){
	// RAD
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.RAD(x, y)
		} else {
			// Нажатие на кнопку
			if obj.enabled {
    			btnPressed = obj
				obj.pressed = true	
			}
		}
}
