package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tTrackBar struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC uint32
    TC uint32
    text string
    visible bool
    focused bool
    enabled bool
    //curX int
    align tAlign
    onClick func(*Node)
    onClickStr string
    //onEnter func(*Node)
    //onEnterStr string
}


func CreateTrackBar(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC uint32, TC uint32, text string, onClick func(*Node)) *Node {
	obj := tTrackBar{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, text: text, visible: true, enabled: true,  onClick: onClick}
	node := Node{typ: TRACKBAR, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj *tTrackBar) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
	
	if obj.enabled {
    	SetColor(obj.BC);
    } else {
    	SetColor(0x7B7D7B);
    	//SetColor(0xBFBFBF);
	}
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
    if (len(obj.text)+1)*7 > obj.sizeX {
    	TextOutgl(nil, obj.text, parX+obj.x + obj.sizeX - (len(obj.text)+1)*7, parY+obj.y + obj.sizeY/2-4, 1);
    } else {
    	TextOutgl(nil, obj.text, parX+obj.x + 4, parY+obj.y + obj.sizeY/2-4, 1);
    }
 
    SetColor(0x000000);
    //LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    //LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);
    
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
}


func (obj *tTrackBar) RAD(x int, y int){
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
    		
			//frmProperties.obj.(*tForm).caption = "Properties: EDIT"
			downX = x 
    		downY = y 
    		mouseIsDown = true
    		lblPropName = CreateLabel(pnlProperties, "lblPropName", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Name", nil)
			editPropName = CreateEdit(pnlProperties, "editPropName", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.name, nil, editPropNameEnter)
			lblPropLeft = CreateLabel(pnlProperties, "lblPropLeft", 5, 25, 95, 20, 0xD8DCC0, 0x000000, "Left", nil)
			editPropLeft = CreateEdit(pnlProperties, "editPropLeft", 80, 25, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.x), nil, editPropLeftEnter)
			lblPropTop = CreateLabel(pnlProperties, "lblPropTop", 5, 45, 95, 20, 0xD8DCC0, 0x000000, "Top", nil)
			editPropTop = CreateEdit(pnlProperties, "editPropTop", 80, 45, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.y), nil, editPropTopEnter)
			lblPropText = CreateLabel(pnlProperties, "lblPropText", 5, 65, 95, 20, 0xD8DCC0, 0x000000, "Text", nil)
			editPropText = CreateEdit(pnlProperties, "editPropText", 80, 65, 95, 20, 0xF8FCF8, 0x000000, obj.text, nil, editPropTextEnter)
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


func (obj *tTrackBar) KeyDown(key int){

}


func (obj *tTrackBar) Click(x int, y int){
	fmt.Println("CLICKED: " + obj.text)
			if obj.onClick != nil && obj.enabled {
				obj.onClick(list[len(list)-1])
			}
}


func (obj *tTrackBar) MouseMove(x int, y int, Xl int, Yl int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && mouseIsDown && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tTrackBar) MouseDown(x int, y int){
	// RAD
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.RAD(x, y)
		} else {
			// Фокус
			if obj.enabled {
				obj.focused = true	
			}
		}
}

