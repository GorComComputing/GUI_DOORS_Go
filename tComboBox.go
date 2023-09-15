package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tComboBox struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    text string
    visible bool
    focused bool
    enabled bool
    curX int
    list []string
    selected int
    align tAlign
    onClick func(*Node)
    onClickStr string
    onEnter func(*Node)
    onEnterStr string
}


func CreateComboBox(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, text string, list []string, onClick func(*Node), onEnter func(*Node)) *Node {
	obj := tComboBox{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, text: text, visible: true, enabled: true, curX: 0, list: list, onClick: onClick, onEnter: onEnter}
	node := Node{typ: COMBOBOX, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	
	CreateBitBtn(&node, "bitBtnComboBox"+name, nil, obj.sizeX - 15, 1, 15, 15, 0xD8DCC0, 0x000000, "V", BORDER, btnComboBox) 
	
	CreateListBox(&node, "objListBox"+name, 0, obj.sizeY, obj.sizeX, len(obj.list)*20, 0xf8fcf8, 0x0, obj.list, objListBoxClick, objListBoxEnter)
	node.children[1].obj.(*tListBox).visible = false
	
	return &node
}


func btnComboBox(node *Node){
	if node.parent.obj.(*tComboBox).enabled && len(node.parent.obj.(*tComboBox).list) > 0 {
		node.parent.children[1].obj.(*tListBox).visible = !(node.parent.children[1].obj.(*tListBox).visible)
		if node.parent.children[1].obj.(*tListBox).visible {
			ToUpPlane(node.parent)
			SetFocus(node.parent.children[1])
		} else {
			SetFocus(nil)
		}
	}
}


func objListBoxEnter(node *Node){
	node.parent.children[1].obj.(*tListBox).visible = !(node.parent.children[1].obj.(*tListBox).visible)
	node.parent.obj.(*tComboBox).text = node.parent.children[1].obj.(*tListBox).list[node.parent.children[1].obj.(*tListBox).selected]
	node.parent.obj.(*tComboBox).onEnter(node.parent)
	SetFocus(nil)
}


func objListBoxClick(node *Node, x int, y int){
	node.parent.children[1].obj.(*tListBox).visible = !(node.parent.children[1].obj.(*tListBox).visible)
	node.parent.obj.(*tComboBox).text = node.parent.children[1].obj.(*tListBox).list[node.parent.children[1].obj.(*tListBox).selected]
	node.parent.obj.(*tComboBox).onEnter(node.parent)
	SetFocus(nil)
}


func (obj *tComboBox) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
	
	if obj.enabled {
    	SetColor(obj.BC);
    } else {
    	SetColor(0xBFBFBF);
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
    if obj.focused && cursor {
    	if (len(obj.text)+1)*7 > obj.sizeX {
    		TextOutgl(nil, "|", parX+obj.x + 4+obj.curX*8 + obj.sizeX - (len(obj.text)+1)*7, parY+obj.y + obj.sizeY/2-4, 1);
    	} else {
    		TextOutgl(nil, "|", parX+obj.x + 4+obj.curX*8, parY+obj.y + obj.sizeY/2-4, 1);
    	}
    }

    
    SetColor(0x000000);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    LinePP(nil, parX+obj.x, parY+obj.y+1, parX+obj.x + obj.sizeX, parY+obj.y+1);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);
}


func (obj *tComboBox) RAD(x int, y int){
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
    		
			//frmProperties.obj.(*tForm).caption = "Properties: COMBOBOX"
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
			
			lblPropList = CreateLabel(pnlProperties, "lblPropList", 5, 205, 95, 20, 0xD8DCC0, 0x000000, "List", nil)
    		cmbPropList = CreateComboBox(pnlProperties, "cmbPropList", 80, 205, 95, 16, 0xF8FCF8, 0x000000, obj.text, obj.list, nil, cmbPropListEnter)
			
			lblEvntClick = CreateLabel(pnlEvents, "lblEvntClick", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(pnlEvents, "editEvntClick", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
			lblEvntEnter = CreateLabel(pnlEvents, "lblEvntEnter", 5, 25, 95, 20, 0xD8DCC0, 0x000000, "Enter", nil)
			editEvntEnter = CreateEdit(pnlEvents, "editEvntEnter", 80, 25, 95, 20, 0xF8FCF8, 0x000000, obj.onEnterStr, nil, editEvntEnterEnter)
}


func (obj *tComboBox) KeyDown(key int){

}


func (obj *tComboBox) Click(x int, y int){
	fmt.Println("CLICKED: " + obj.text)
	if obj.onClick != nil && obj.enabled {
		obj.onClick(list[len(list)-1])
	}
}


func (obj *tComboBox) MouseMove(x int, y int, Xl int, Yl int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && mouseIsDown && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tComboBox) MouseDown(x int, y int){
	// RAD
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.RAD(x, y)
		} else {
			// Фокус
			if obj.enabled {
				obj.focused = true	
				obj.curX = len(obj.text)
			}
		}
}
