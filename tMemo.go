package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tMemo struct{
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
    curY int
    pos int
    line_start int
    onClick func(*Node)
    onClickStr string
}


func CreateMemo(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, text string, onClick func(*Node)) *Node {
	obj := tMemo{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, text: text, visible: true, enabled: true, curX: 0, curY: 0, pos: 0, line_start: 0, onClick: onClick}
	node := Node{typ: MEMO, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj *tMemo) Draw(parX int, parY int, parSizeX int, parSizeY int){
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
    TextOutgl(nil, obj.text, parX+obj.x + 4, parY+obj.y + 4, 1);
    if obj.focused && cursor {
    	TextOutgl(nil, "|", parX+obj.x + 4 + obj.curX*8, parY+obj.y + 4 + (7+2)*obj.curY, 1);
    }

    
    SetColor(0x000000);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);
}


func (obj *tMemo) RAD(x int, y int){
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
    		
			//frmProperties.obj.(*tForm).caption = "Properties: MEMO"
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


func (obj *tMemo) KeyDown(key int){
	if obj.enabled {
	if key == 8 {
    			if len(obj.text) > 0 {
    				obj.text = obj.text[:obj.pos-1] + obj.text[obj.pos:]
    				obj.curX--
    				obj.pos--
    			}
    		} else if key == 13 {
    			obj.text = obj.text[:obj.line_start + obj.pos] + string(key) + obj.text[obj.line_start + obj.pos:]
    			//obj.pos++
    			obj.line_start += obj.pos+1
    			obj.curY++
    			obj.curX = 0
    			obj.pos = 0
    		} else if key == 37 {
    			if obj.curX > 0 {
    				obj.curX--
    				obj.pos--
    			}
    		} else if key == 39 {
    			if ((obj.line_start + obj.pos) < (len(obj.text)-1)) && ((obj.line_start + obj.pos+1) != 13) {
    				obj.curX++
    				obj.pos++
    			}
    		} else if key == 38 {
    			oldNL := obj.line_start
    			if obj.line_start != 0 { 
    				obj.curY--
    				obj.line_start = findLeft(obj.text, obj.line_start)
    				if obj.curX > oldNL - obj.line_start - 1 {
    					obj.curX = oldNL - obj.line_start - 1
    					obj.pos = oldNL - obj.line_start - 1
    				}
    			}
    			
    		} else if key == 40 {
    			right := findRight(obj.text, obj.line_start)
    			if obj.line_start == right || right == len(obj.text)-1 {
    				return
    			}
    			obj.line_start = right
    			if obj.curX > len(obj.text)-1 - obj.line_start {
    				obj.curX = len(obj.text) - obj.line_start
    				obj.pos = len(obj.text) - obj.line_start
    			}
    			obj.curY++
			} else {
				obj.text = obj.text[:obj.line_start + obj.pos] + string(key) + obj.text[obj.line_start + obj.pos:]
				obj.pos++
				obj.curX++
			}
		}
}


func (obj *tMemo) Click(x int, y int){

}


func (obj *tMemo) MouseMove(x int, y int, Xl int, Yl int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && mouseIsDown && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tMemo) MouseDown(x int, y int){
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

