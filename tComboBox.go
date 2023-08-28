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
    onClick func(*Node)
    onClickStr string
    onEnter func(*Node)
    onEnterStr string
}


func CreateComboBox(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, text string, list []string, onClick func(*Node), onEnter func(*Node)) *Node {
	obj := tComboBox{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, text: text, visible: true, enabled: true, curX: 0, list: list, onClick: onClick, onEnter: onEnter}
	node := Node{typ: COMBOBOX, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	
	CreateBitBtn(&node, "bitBtnComboBox"+name, obj.sizeX - 15, 1, 15, 15, 0xD8DCC0, 0x000000, "v", btnComboBox) 
	objPnl := CreatePanel(&node, "pnlComboBox"+name, 0, obj.sizeY+1, obj.sizeX, 16, 0xd8dcc0, FLAT, nil)
	CreateBtn(objPnl, "listComboBox"+name, 2, 2, obj.sizeX-4, 16-4, 0xd8dcc0, 0x0, "true", nil)
	node.children[1].obj.(*tPanel).visible = false
	
	return &node
}


func btnComboBox(node *Node){
	node.parent.children[1].obj.(*tPanel).visible = !(node.parent.children[1].obj.(*tPanel).visible)
}


func (obj *tComboBox) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
	
	if obj.enabled {
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
    } else {
    	SetColor(0xBFBFBF);
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
	}

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
    		
			frmProperties.obj.(*tForm).caption = "Properties: COMBOBOX"
			downX = x 
    		downY = y 
    		mouseIsDown = true
    		lblPropName = CreateLabel(frmProperties, "lblPropName", 5, 20, 95, 20, 0xD8DCC0, 0x000000, "Name", nil)
			editPropName = CreateEdit(frmProperties, "editPropName", 80, 20, 95, 20, 0xF8FCF8, 0x000000, obj.name, nil, editPropNameEnter)
			lblPropLeft = CreateLabel(frmProperties, "lblPropLeft", 5, 40, 95, 20, 0xD8DCC0, 0x000000, "Left", nil)
			editPropLeft = CreateEdit(frmProperties, "editPropLeft", 80, 40, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.x), nil, editPropLeftEnter)
			lblPropTop = CreateLabel(frmProperties, "lblPropTop", 5, 60, 95, 20, 0xD8DCC0, 0x000000, "Top", nil)
			editPropTop = CreateEdit(frmProperties, "editPropTop", 80, 60, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.y), nil, editPropTopEnter)
			lblPropText = CreateLabel(frmProperties, "lblPropText", 5, 80, 95, 20, 0xD8DCC0, 0x000000, "Text", nil)
			editPropText = CreateEdit(frmProperties, "editPropText", 80, 80, 95, 20, 0xF8FCF8, 0x000000, obj.text, nil, editPropTextEnter)
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
			
			lblPropList = CreateLabel(frmProperties, "lblPropList", 5, 220, 95, 20, 0xD8DCC0, 0x000000, "List", nil)
    		cmbPropList = CreateComboBox(frmProperties, "cmbPropList", 80, 220, 95, 20, 0xF8FCF8, 0x000000, obj.text, obj.list, nil, cmbPropListEnter)
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 260, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 260, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
			lblEvntEnter = CreateLabel(frmProperties, "lblEvntEnter", 5, 280, 95, 20, 0xD8DCC0, 0x000000, "Enter", nil)
			editEvntEnter = CreateEdit(frmProperties, "editEvntEnter", 80, 280, 95, 20, 0xF8FCF8, 0x000000, obj.onEnterStr, nil, editEvntEnterEnter)
}


func (obj *tComboBox) KeyDown(key int){

}


func (obj *tComboBox) Click(){
	fmt.Println("CLICKED: " + obj.text)
			if obj.onClick != nil && obj.enabled {
				obj.onClick(list[len(list)-1])
			}
}


func (obj *tComboBox) MouseMove(x int, y int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tComboBox) MouseDown(x int, y int){
	// RAD
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.RAD(x, y)
		} else {
			// Фокус
			if obj.enabled {
				obj.focused = true	
				obj.curX = len(obj.text)
			}
		}
}
