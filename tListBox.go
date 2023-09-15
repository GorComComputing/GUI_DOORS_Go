package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tListBox struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    visible bool
    focused bool
    enabled bool
    list []string
    selected int
    align tAlign
    onClick func(*Node, int, int)
    onClickStr string
    onEnter func(*Node)
    onEnterStr string
}


func CreateListBox(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, list []string, onClick func(*Node, int, int), onEnter func(*Node)) *Node {
	obj := tListBox{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, visible: true, enabled: true, list: list, selected: 0, onClick: onClick, onEnter: onEnter}
	node := Node{typ: LISTBOX, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj *tListBox) Draw(parX int, parY int, parSizeX int, parSizeY int){
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
    for i := 0; i < len(obj.list); i++ {
    	if i == obj.selected && obj.enabled {
    		SetColor(0x0054E0);
    		var p []tPoint

   			p1 := tPoint{x: parX+obj.x, y: parY+obj.y + i*20}
			p = append(p, p1)
	
			p2 := tPoint{x: parX+obj.x + obj.sizeX, y: parY+obj.y + i*20}
			p = append(p, p2)
	
			p3 := tPoint{x: parX+obj.x + obj.sizeX, y: parY+obj.y + 20 + i*20}
			p = append(p, p3)
	
			p4 := tPoint{x: parX+obj.x, y: parY+obj.y + 20 + i*20}
			p = append(p, p4)

    		FillPoly(nil, 4, p);
    		SetColor(0xF8FCF8);
    	} else {
    		SetColor(obj.TC);
    	}
    	TextOutgl(nil, obj.list[i], parX+obj.x + 4, parY+obj.y + 4 + i*20, 1);
    	
    }
    
   /* if obj.focused && cursor {
    	TextOutgl(nil, "|", parX+obj.x + 4 + obj.curX*8, parY+obj.y + 4 + (7+2)*obj.curY, 1);
    }*/

    
    /*SetColor(0x000000);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);*/
}


func (obj *tListBox) RAD(x int, y int){
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
    		
		//frmProperties.obj.(*tForm).caption = "Properties: LISTBOX"
		downX = x 
    	downY = y 
    	mouseIsDown = true
    	lblPropName = CreateLabel(pnlProperties, "lblPropName", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Name", nil)
		editPropName = CreateEdit(pnlProperties, "editPropName", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.name, nil, editPropNameEnter)
		lblPropLeft = CreateLabel(pnlProperties, "lblPropLeft", 5, 25, 95, 20, 0xD8DCC0, 0x000000, "Left", nil)
		editPropLeft = CreateEdit(pnlProperties, "editPropLeft", 80, 25, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.x), nil, editPropLeftEnter)
		lblPropTop = CreateLabel(pnlProperties, "lblPropTop", 5, 45, 95, 20, 0xD8DCC0, 0x000000, "Top", nil)
		editPropTop = CreateEdit(pnlProperties, "editPropTop", 80, 45, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.y), nil, editPropTopEnter)
		lblPropSelected = CreateLabel(pnlProperties, "lblPropSelected", 5, 65, 95, 20, 0xD8DCC0, 0x000000, "Selected", nil)
		editPropSelected = CreateEdit(pnlProperties, "editPropSelected", 80, 65, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.selected), nil, editPropSelectedEnter)
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
    	cmbPropList = CreateComboBox(pnlProperties, "cmbPropList", 80, 205, 95, 16, 0xF8FCF8, 0x000000, "", obj.list, nil, cmbPropListEnter)
			
		lblEvntClick = CreateLabel(pnlEvents, "lblEvntClick", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
		editEvntClick = CreateEdit(pnlEvents, "editEvntClick", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
}


func (obj *tListBox) KeyDown(key int){
	if key == 8 {
    			
    } else if key == 13 {
    	if obj.onEnter != nil && obj.enabled {
    		obj.onEnter(layout.children[len(layout.children)-1].obj.(*tForm).focus)	
    	}
    } else if key == 37 {
    			
    } else if key == 39 {
    			
    } else if key == 38 {
    	if obj.selected >= 1 && obj.enabled && len(list) > 0 {
    				obj.selected--
    	}
    } else if key == 40 {
    	if obj.selected < len(obj.list)-1  && obj.enabled && len(list) > 0 {
    		obj.selected++
    	}
	} else {

	}
}


func (obj *tListBox) Click(x int, y int){
	fmt.Println("CLICKED: ", strconv.Itoa(x), strconv.Itoa(y))
	if obj.enabled && len(list) > 0 {
		obj.selected = int(y/20)
	}
	
	if obj.onClick != nil && obj.enabled {
		obj.onClick(list[len(list)-1], x, y)
	}
}


func (obj *tListBox) MouseMove(x int, y int, Xl int, Yl int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && mouseIsDown && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
		obj.x += x - downX
    	obj.y += y - downY
    	editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
		editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    }
}


func (obj *tListBox) MouseDown(x int, y int){
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

