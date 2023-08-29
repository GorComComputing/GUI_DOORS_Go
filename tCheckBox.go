package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tCheckBox struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    caption string
    visible bool
    checked bool
    enabled bool
    onClick func(*Node)
    onClickStr string
}


func CreateCheckBox(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, caption string, checked bool, onClick func(*Node)) *Node {
	obj := tCheckBox{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, checked: checked, enabled: true, onClick: onClick}
	node := Node{typ: CHECKBOX, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj *tCheckBox) Draw(parX int, parY int, parSizeX int, parSizeY int){
	const size int = 16
	const size_sm int = size - 8
	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
	SetColor(obj.BC);
    	var p []tPoint

    	p1 := tPoint{x: parX + obj.x, y: parY + obj.y}
		p = append(p, p1)
	
		p2 := tPoint{x: parX + obj.x + size, y: parY + obj.y}
		p = append(p, p2)
	
		p3 := tPoint{x: parX + obj.x + size, y: parY + obj.y + size}
		p = append(p, p3)
	
		p4 := tPoint{x: parX + obj.x, y: parY + obj.y + size}
		p = append(p, p4)

    	FillPoly(nil, 4, p);
    	
    if obj.enabled {
    	SetColor(0x787C78);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x + size, parY + obj.y);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x, parY + obj.y + size);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x + size - 2, parY + obj.y+1);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x+1, parY + obj.y + size - 1);
    	SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x+2, parY + obj.y + size - 1, parX + obj.x + size - 1, parY + obj.y + size - 1);
    	LinePP(nil, parX + obj.x + size - 1, parY + obj.y + 1, parX + obj.x + size - 1, parY + obj.y + size - 1);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+1, parY + obj.y + size, parX + obj.x + size, parY + obj.y + size);
    	LinePP(nil, parX + obj.x + size, parY + obj.y+1, parX + obj.x + size, parY + obj.y + size);
    } else {
    	SetColor(0x777777);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x + size, parY + obj.y);
    	LinePP(nil, parX + obj.x, parY + obj.y, parX + obj.x, parY + obj.y + size);
    	SetColor(0x777777);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x + size - 2, parY + obj.y+1);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x+1, parY + obj.y + size - 1);
    	SetColor(0x777B77);
    	LinePP(nil, parX + obj.x+2, parY + obj.y + size - 1, parX + obj.x + size - 1, parY + obj.y + size - 1);
    	LinePP(nil, parX + obj.x + size - 1, parY + obj.y + 1, parX + obj.x + size - 1, parY + obj.y + size - 1);
    	SetColor(0x777B77);
    	LinePP(nil, parX + obj.x+1, parY + obj.y + size, parX + obj.x + size, parY + obj.y + size);
    	LinePP(nil, parX + obj.x + size, parY + obj.y+1, parX + obj.x + size, parY + obj.y + size);
    }
    	
	if obj.checked {
		SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x+4, parY + obj.y+4, parX + obj.x + size_sm+4, parY + obj.y+4);
    	LinePP(nil, parX + obj.x+4, parY + obj.y+4, parX + obj.x+4, parY + obj.y + size_sm+4);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+4+1, parY + obj.y+4+1, parX + obj.x + size_sm+4 - 2, parY + obj.y+4+1);
    	LinePP(nil, parX + obj.x+4+1, parY + obj.y+4+1, parX + obj.x+4+1, parY + obj.y + size_sm+4 - 1);
    	SetColor(0x787C78);
    	LinePP(nil, parX + obj.x+4+2, parY + obj.y + size_sm+4 - 1, parX + obj.x + size_sm+4 - 1, parY + obj.y + size_sm+4 - 1);
    	LinePP(nil, parX + obj.x + size_sm+4 - 1, parY + obj.y+4 + 1, parX + obj.x + size_sm+4 - 1, parY + obj.y + size_sm+4 - 1);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+4+1, parY + obj.y + size_sm+4, parX + obj.x + size_sm+4, parY + obj.y + size_sm+4);
    	LinePP(nil, parX + obj.x + size_sm+4, parY + obj.y+4+1, parX + obj.x + size_sm+4, parY + obj.y + size_sm+4);
	} 
	
	if obj.enabled {
    	SetColor(obj.TC);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX+obj.x + 25, parY+obj.y + obj.sizeY/2-4, 1);
    } else {
    	SetColor(0x777B77);  // 0x787C78
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, parX+obj.x + 25, parY+obj.y + obj.sizeY/2-4, 1);
    	//TextOutgl(nil, obj.caption, parX+obj.x + obj.sizeX/2-((len(obj.caption)-1)*8)/2 - 3, parY+obj.y + obj.sizeY/2-4, 1);
    }    
}


func (obj *tCheckBox) RAD(x int, y int){
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
    		var checked string
    		if obj.checked {
    			checked = "true"
    		} else {
    			checked = "false"
    		}
    		
			//frmProperties.obj.(*tForm).caption = "Properties: CHECKBOX"
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
			lblPropChecked = CreateLabel(pnlProperties, "lblPropChecked", 5, 205, 95, 20, 0xD8DCC0, 0x000000, "Checked", nil)
			cmbPropChecked = CreateComboBox(pnlProperties, "cmbPropChecked", 80, 205, 95, 16, 0xF8FCF8, 0x000000, checked, listBool, nil, cmbPropCheckedEnter)
			
			lblEvntClick = CreateLabel(pnlEvents, "lblEvntClick", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(pnlEvents, "editEvntClick", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
}


func (obj *tCheckBox) KeyDown(key int){

}


func (obj *tCheckBox) Click(x int, y int){
	fmt.Println("CLICKED: " + obj.caption)
			if obj.onClick != nil && obj.enabled {
				obj.onClick(list[len(list)-1])
			}
}


func (obj *tCheckBox) MouseMove(x int, y int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tCheckBox) MouseDown(x int, y int){
	// RAD
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.RAD(x, y)
		}
}
