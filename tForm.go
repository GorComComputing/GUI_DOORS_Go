package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tForm struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    mode tMode
    caption string
    visible bool
    focused bool
    focus *Node
    isRAD bool
    onClick func(*Node)
    onClickStr string
}

type tMode int

const (
    NONE tMode = iota
    WIN
    FLAT
    TASK
)


func CreateForm(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, mode tMode, caption string, visible bool, onClick func(*Node)) *Node {
	obj := tForm{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, mode: mode, caption: caption, visible: visible, focus: nil, isRAD: false, onClick: onClick}
	node := Node{typ: FORM, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	
	if obj.mode == WIN {
		CreateBitBtn(&node, "bitbtnClose"+name, obj.sizeX - 17, 2, 15, 15, 0xD8DCC0, 0x000000, "X", formClose)
	}
	return &node
}


func formClose(node *Node){
	//node.parent.obj.(*tForm).visible = false
	
	
	var btn *Node
	var i int
	for i = 0; i < len(process); i++ {
		if node.parent == process[i].form {
			btn = process[i].btn
			copy(process[i:], process[i+1:])
			process[len(process)-1] = nil
			process = process[:len(process)-1]
			//process[i].form.obj.(*tForm).visible = !(process[i].form.obj.(*tForm).visible)
			break
		}
	}
	
	for i = 0; i < len(pnlTask.children); i++ {
		if btn == pnlTask.children[i] {
			copy(pnlTask.children[i:], pnlTask.children[i+1:])
			pnlTask.children[len(pnlTask.children)-1] = nil
			pnlTask.children = pnlTask.children[:len(pnlTask.children)-1]
			break
		}
	}

	xTask = 2-10 
	for i = 1; i < len(pnlTask.children); i++ {
		switch obj := pnlTask.children[i].obj.(type) {
		case *tBtn:
			obj.x = xTask
		}
		xTask += 81
	}
	
	layout.children[len(layout.children)-3].obj.(*tForm).focused = true
	for i = 0; i < len(layout.children); i++ {
		if node.parent == layout.children[i] {
			copy(layout.children[i:], layout.children[i+1:])
			layout.children[len(layout.children)-1] = nil
			layout.children = layout.children[:len(layout.children)-1]
			break
		}
	}
	
	
	

}


func (obj *tForm) Draw(parX int, parY int, parSizeX int, parSizeY int){
	SetLocalViewPort(parX + obj.x+2, parY + obj.y+2, parX + obj.x + obj.sizeX-2, parY + obj.y + obj.sizeY-2)
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
    
    if obj.isRAD {
    	SetColor(0xFF0000)
    	for i := 0; i < obj.sizeY; i += 10 {
    		for j := 0; j < obj.sizeX; j += 10 {
    			PutPixel(nil, parX + obj.x + j, parY + obj.y + i, 0x000000)
    		}
    	}
    }
    
    if obj.mode == WIN {
    	if obj.focused {
    		SetColor(0x0054E0)
    	} else {
    		SetColor(0x80A8E8)
    	}
    	p = nil
    	p1 = tPoint{x: parX + obj.x, y: parY + obj.y}
		p = append(p, p1)
	
		p2 = tPoint{x: parX + obj.x + obj.sizeX, y: parY + obj.y}
		p = append(p, p2)
	
		p3 = tPoint{x: parX + obj.x + obj.sizeX, y: parY + obj.y + 17}
		p = append(p, p3)
	
		p4 = tPoint{x: parX + obj.x, y: parY + obj.y + 17}
		p = append(p, p4)
	
		FillPoly(nil, 4, p);
		
		if obj.focused {
    		SetColor(0xF8FCF8)
    		SetBackColor(0x0054E0);
    	} else {
    		SetColor(0x787C78)
    		SetBackColor(0x80A8E8);
    	}
    	TextOutgl(nil, obj.caption, parX + obj.x + 9, parY + obj.y + 6, 1);
    }

	if obj.mode != FLAT {
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
    	LinePP(nil, parX + obj.x, parY + obj.y + obj.sizeY, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY);
    	LinePP(nil, parX + obj.x + obj.sizeX, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY);  
    }  
}


func (obj *tForm) RAD(x int, y int){
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

    		frmProperties.obj.(*tForm).caption = "Properties: FORM"
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
			lblPropWidth = CreateLabel(frmProperties, "lblPropWidth", 5, 120, 95, 20, 0xD8DCC0, 0x000000, "Width", nil)
			editPropWidth = CreateEdit(frmProperties, "editPropWidth", 80, 120, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeX), nil, editPropWidthEnter)
			lblPropHeight = CreateLabel(frmProperties, "lblPropHeight", 5, 140, 95, 20, 0xD8DCC0, 0x000000, "Height", nil)
			editPropHeight = CreateEdit(frmProperties, "editPropHeight", 80, 140, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeY), nil, editPropHeightEnter)
			lblPropMode = CreateLabel(frmProperties, "lblPropMode", 5, 160, 95, 20, 0xD8DCC0, 0x000000, "Mode", nil)
			editPropMode = CreateEdit(frmProperties, "editPropMode", 80, 160, 95, 20, 0xF8FCF8, 0x000000, mode, nil, editPropModeEnter)
			lblPropVisible = CreateLabel(frmProperties, "lblPropVisible", 5, 180, 95, 20, 0xD8DCC0, 0x000000, "Visible", nil)
			editPropVisible = CreateEdit(frmProperties, "editPropVisible", 80, 180, 95, 20, 0xF8FCF8, 0x000000, visible, nil, editPropVisibleEnter)
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 220, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 220, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
}


func (obj *tForm) KeyDown(key int){

}


func (obj *tForm) Click(){

}


func (obj *tForm) MouseMove(x int, y int){
	obj.x += x - downX
    	obj.y += y - downY	
    	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)	
		}
}


func (obj *tForm) MouseDown(x int, y int){
	// Перенос окна
		if (obj.mode == WIN) &&
			(obj.x) < x && 
			(obj.x + obj.sizeX) > x && 
			(obj.y) < y && 
			(obj.y + 17) > y {
				downX = x 
    			downY = y 
    			mouseIsDown = true
    	}
    	// RAD
    	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
    		obj.RAD(x, y)
		}
}

/*func (obj tForm) SetSize(width int, height int){
	obj.sizeX = width
	obj.sizeY = height
	
}*/

