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
    BC uint32
    TC uint32
    caption string
    visible bool
    pressed bool
    enabled bool
    picture []byte
    mode tMode
    inalign tAlign
    align tAlign
    alTop bool
    alBottom bool
    alLeft bool
    alRight bool
    onClick func(*Node)
    onClickStr string
}


func CreateBitBtn(parent *Node, name string, picture []byte, x int, y int, sizeX int, sizeY int, BC uint32, TC uint32, caption string, mode tMode, onClick func(*Node)) *Node {
	obj := tBitBtn{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, caption: caption, visible: true, pressed: false, enabled: true, picture: picture, mode: mode, inalign: CENTER, alTop: true, alLeft: true, alBottom: false, alRight: false, onClick: onClick}
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
		if obj.mode == BORDER {
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
    	
	}
	
	
	var add int = 0
    if obj.picture != nil {
    	add = 20
    }
	
	
	var startYpic int = parY + obj.y + obj.sizeY/2-4
    var startXpic int = parX + obj.x + 8
    var startYtex int = parY + obj.y + obj.sizeY/2-4
    var startXtex int = parX + obj.x + 8
    switch obj.inalign {
    case LEFT:
    	startYpic = parY + obj.y + obj.sizeY/2-4
    	startXpic = parX + obj.x + 8
    	startYtex = parY + obj.y + obj.sizeY/2-4
    	startXtex = parX + obj.x + 8
    case CENTER:
    	startYpic = parY + obj.y + obj.sizeY/2-4
    	startXpic = parX + obj.x + obj.sizeX/2-((len(obj.caption)-1)*8+20)/2
    	startYtex = parY + obj.y + obj.sizeY/2-4
    	startXtex = parX+obj.x + obj.sizeX/2-((len(obj.caption)-1)*8 + add)/2
    }
	
	if obj.picture != nil {
    	showBMP(nil, obj.picture, startXpic - 3, startYpic-3)
    }
	
	if obj.enabled {
    	SetColor(obj.TC);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, startXtex - 3 + add, startYtex, 1)
    } else {
    	SetColor(0x787C78);
    	SetBackColor(obj.BC);
    	TextOutgl(nil, obj.caption, startXtex - 3 + add, startYtex, 1)
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
			
			//frmProperties.obj.(*tForm).caption = "Properties: BITBTN"
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


func (obj *tBitBtn) KeyDown(key int){

}


func (obj *tBitBtn) Click(x int, y int){
	fmt.Println("CLICKED: " + obj.caption)
			if obj.onClick != nil && obj.enabled {
				obj.onClick(list[len(list)-1])
			}
}


func (obj *tBitBtn) MouseMove(x int, y int, Xl int, Yl int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && mouseIsDown && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tBitBtn) MouseDown(x int, y int){
	// RAD
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.RAD(x, y)
		} else {
			// Нажатие на кнопку
			if obj.enabled {
    			bitbtnPressed = obj
				obj.pressed = true	
			}
		}
}
