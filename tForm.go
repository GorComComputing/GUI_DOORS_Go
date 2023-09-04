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
    picture []byte
    onClick func(*Node)
    onClickStr string
}

type tMode int

const (
    NONE tMode = iota	//tForm, tPanel
    WIN			//tForm
    DIALOG		//tForm
    FLAT		//tBitBtn, tMenu, tForm, tPanel
    BORDER		//tBitBtn
    TASK		//tPanel
    LINE		//tMenu
    LISTICON	//tListFileBox
    BIGICON		//tListFileBox
)


func CreateForm(parent *Node, name string, picture []byte, x int, y int, sizeX int, sizeY int, BC int, mode tMode, caption string, visible bool, onClick func(*Node)) *Node {
	obj := tForm{name: name, picture: picture, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, mode: mode, caption: caption, visible: visible, focus: nil, isRAD: false, onClick: onClick}
	node := Node{typ: FORM, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	
	if obj.mode == WIN || obj.mode == DIALOG {
		CreateBitBtn(&node, "bitbtnClose"+name, nil, obj.sizeX - 17, 2, 15, 15, 0xD8DCC0, 0x000000, "X", BORDER, formClose)
	}
	return &node
}


func formClose(node *Node){
	node.parent.obj.(*tForm).visible = false
	
	var btn *Node
	var i int
	for i = 0; i < len(process); i++ {
		if node.parent == process[i].form {
			process[i].isRun = false
			btn = process[i].btn
			//copy(process[i:], process[i+1:])
			//process[len(process)-1] = nil
			//process = process[:len(process)-1]
						//process[i].form.obj.(*tForm).visible = false
			break
		}
	}
	
	// Удаляет кнопку
	for i = 0; i < len(pnlTask.children); i++ {
		if btn == pnlTask.children[i] {
			copy(pnlTask.children[i:], pnlTask.children[i+1:])
			pnlTask.children[len(pnlTask.children)-1] = nil
			pnlTask.children = pnlTask.children[:len(pnlTask.children)-1]
			break
		}
	}

	// Сдвигает кнопки
	xTask = 2-30 
	for i = 1; i < len(pnlTask.children); i++ {
		switch obj := pnlTask.children[i].obj.(type) {
		case *tBtn:
			obj.x = xTask
		case *tBitBtn:
			obj.x = xTask
		}
		xTask += 101
	}
	
	// Устанавливает фокус
	layout.children[len(layout.children)-3].obj.(*tForm).focused = true
	
	if node.parent.obj.(*tForm).mode == DIALOG {
	// Удаляет форму
	for i = 0; i < len(layout.children); i++ {
		if node.parent == layout.children[i] {
			copy(layout.children[i:], layout.children[i+1:])
			layout.children[len(layout.children)-1] = nil
			layout.children = layout.children[:len(layout.children)-1]
			break
		}
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
    
    if obj.mode == WIN || obj.mode == DIALOG {
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
    	if obj.picture != nil {
    		showBMP(nil, obj.picture, parX + obj.x + 7, parY + obj.y + 2)
    		TextOutgl(nil, obj.caption, parX + obj.x + 7 + 20, parY + obj.y + 6, 1);
    	} else {
    		TextOutgl(nil, obj.caption, parX + obj.x + 9, parY + obj.y + 6, 1);
    	}
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
    		} else if obj.mode == DIALOG {
    			mode = "DIALOG"
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
			lblPropWidth = CreateLabel(pnlProperties, "lblPropWidth", 5, 105, 95, 20, 0xD8DCC0, 0x000000, "Width", nil)
			editPropWidth = CreateEdit(pnlProperties, "editPropWidth", 80, 105, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeX), nil, editPropWidthEnter)
			lblPropHeight = CreateLabel(pnlProperties, "lblPropHeight", 5, 125, 95, 20, 0xD8DCC0, 0x000000, "Height", nil)
			editPropHeight = CreateEdit(pnlProperties, "editPropHeight", 80, 125, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeY), nil, editPropHeightEnter)
			lblPropMode = CreateLabel(pnlProperties, "lblPropMode", 5, 145, 95, 20, 0xD8DCC0, 0x000000, "Mode", nil)
			cmbPropMode = CreateComboBox(pnlProperties, "cmbPropMode", 80, 145, 95, 16, 0xF8FCF8, 0x000000, mode, listMode, nil, cmbPropModeEnter)
			lblPropVisible = CreateLabel(pnlProperties, "lblPropVisible", 5, 165, 95, 20, 0xD8DCC0, 0x000000, "Visible", nil)
			cmbPropVisible = CreateComboBox(pnlProperties, "cmbPropVisible", 80, 165, 95, 16, 0xF8FCF8, 0x000000, visible, listBool, nil, cmbPropVisibleEnter)
			
			lblEvntClick = CreateLabel(pnlEvents, "lblEvntClick", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(pnlEvents, "editEvntClick", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
}


func (obj *tForm) KeyDown(key int){

}


func (obj *tForm) Click(x int, y int){

}


func (obj *tForm) MouseMove(x int, y int, Xl int, Yl int){
	if !mouseIsDown {return}
	obj.x += x - downX
    obj.y += y - downY	
    	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG{
			editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)	
		}
}


func (obj *tForm) MouseDown(x int, y int){
	// Перенос окна
		if (obj.mode == WIN || obj.mode == DIALOG) &&
			(obj.x) < x && 
			(obj.x + obj.sizeX) > x && 
			(obj.y) < y && 
			(obj.y + 17) > y {
				downX = x 
    			downY = y 
    			mouseIsDown = true
    	}
    	// RAD
    	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
    		obj.RAD(x, y)
		}
}

/*func (obj tForm) SetSize(width int, height int){
	obj.sizeX = width
	obj.sizeY = height
	
}*/

