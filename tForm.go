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
    isMaximized bool
    old_x int
    old_y int
    old_sizeX int
    old_sizeY int
    picture []byte
    onClick func(*Node)
    onClickStr string
}

type tMode int
const (
    NONE tMode = iota	//tForm, tPanel
    WIN			//tForm
    DIALOG		//tForm
    FIXED		//tForm
    FLAT		//tBitBtn, tMenu, tForm, tPanel
    BORDER		//tBitBtn
    TASK		//tPanel
    LINE		//tMenu
    LISTICON	//tListFileBox
    BIGICON		//tListFileBox
)

type tAlign int
const (
    CENTER tAlign = iota	
    LEFT
    RIGHT
    TOP
    BOTTOM
    RIGHT_TOP
    LEFT_TOP
    RIGHT_BOTTOM
    LEFT_BOTTOM
    CLIENT
)


func CreateForm(parent *Node, name string, picture []byte, x int, y int, sizeX int, sizeY int, BC int, mode tMode, caption string, visible bool, onClick func(*Node)) *Node {
	obj := tForm{name: name, picture: picture, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, mode: mode, caption: caption, visible: visible, focus: nil, isRAD: false, isMaximized: false, onClick: onClick}
	node := Node{typ: FORM, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	
	if obj.mode == WIN || obj.mode == DIALOG || obj.mode == FIXED {
		bitbtn := CreateBitBtn(&node, "bitbtnClose"+name, nil, obj.sizeX - 17, 2, 15, 15, 0xD8DCC0, 0x000000, "X", BORDER, formClose)
		bitbtn.obj.(*tBitBtn).align = RIGHT_TOP
	}
	if obj.mode == WIN {
		bitbtn := CreateBitBtn(&node, "bitbtnMax"+name, nil, obj.sizeX - 17-15, 2, 15, 15, 0xD8DCC0, 0x000000, string(0), BORDER, formMax)
		bitbtn.obj.(*tBitBtn).align = RIGHT_TOP
	} else if obj.mode == FIXED {
		bitbtn := CreateBitBtn(&node, "bitbtnMax"+name, nil, obj.sizeX - 17-15, 2, 15, 15, 0xD8DCC0, 0x000000, string(0), BORDER, formMax)
		bitbtn.obj.(*tBitBtn).align = RIGHT_TOP
		bitbtn.obj.(*tBitBtn).enabled = false
	}
	if obj.mode == WIN || obj.mode == FIXED {
		bitbtn := CreateBitBtn(&node, "bitbtnLine"+name, nil, obj.sizeX - 17-15-15, 2, 15, 15, 0xD8DCC0, 0x000000, "_", BORDER, formLine)
		bitbtn.obj.(*tBitBtn).align = RIGHT_TOP
	}
	return &node
}



func formLine(node *Node){
	node.parent.obj.(*tForm).visible = false
}


func formMax(node *Node){
	node.parent.obj.(*tForm).isMaximized = !(node.parent.obj.(*tForm).isMaximized)
	if 	node.parent.obj.(*tForm).isMaximized {
		node.obj.(*tBitBtn).caption = string(0x01)
		node.parent.obj.(*tForm).old_x = node.parent.obj.(*tForm).x
    	node.parent.obj.(*tForm).old_y = node.parent.obj.(*tForm).y
    	node.parent.obj.(*tForm).old_sizeX = node.parent.obj.(*tForm).sizeX
    	node.parent.obj.(*tForm).old_sizeY = node.parent.obj.(*tForm).sizeY
    	
    	node.parent.obj.(*tForm).x = 0
    	node.parent.obj.(*tForm).y = 0
    	setSize(node.parent, BITMAP_WIDTH-1, BITMAP_HEIGHT-2-28)
    	//node.parent.obj.(*tForm).sizeX = BITMAP_WIDTH-1
    	//node.parent.obj.(*tForm).sizeY = BITMAP_HEIGHT-2-28
    } else {
    	node.obj.(*tBitBtn).caption = string(0x00)
    	node.parent.obj.(*tForm).x = node.parent.obj.(*tForm).old_x
    	node.parent.obj.(*tForm).y = node.parent.obj.(*tForm).old_y
    	setSize(node.parent, node.parent.obj.(*tForm).old_sizeX, node.parent.obj.(*tForm).old_sizeY)
    	//node.parent.obj.(*tForm).sizeX = node.parent.obj.(*tForm).old_sizeX
    	//node.parent.obj.(*tForm).sizeY = node.parent.obj.(*tForm).old_sizeY
    }
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
	SetColor(obj.BC);
    var p []tPoint
    
    var sizeY int = obj.sizeY
    var sizeX int = obj.sizeX
    var startY int = parY + obj.y
    var startX int = parX + obj.x
    
    
	SetLocalViewPort(startX+2, startY+2, startX + sizeX-2, startY + sizeY-2)
    	
    p1 := tPoint{x: startX, y: startY}
	p = append(p, p1)
	
	p2 := tPoint{x: startX + sizeX, y: startY}
	p = append(p, p2)
	
	p3 := tPoint{x: startX + sizeX, y: startY + sizeY}
	p = append(p, p3)
	
	p4 := tPoint{x: startX, y: startY + sizeY}
	p = append(p, p4)
    FillPoly(nil, 4, p);
    
    if obj.isRAD {
    	SetColor(0xFF0000)
    	for i := 0; i < sizeY; i += 10 {
    		for j := 0; j < sizeX; j += 10 {
    			PutPixel(nil, startX + j, startY + i, 0x000000)
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

    	p1 := tPoint{x: startX, y: startY}
		p = append(p, p1)
	
		p2 := tPoint{x: startX + sizeX, y: startY}
		p = append(p, p2)
	
		p3 := tPoint{x: startX + sizeX, y: startY + 17}
		p = append(p, p3)
	
		p4 := tPoint{x: startX, y: startY + 17}
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
    		showBMP(nil, obj.picture, startX + 7, startY + 2)
    		TextOutgl(nil, obj.caption, startX + 7 + 20, startY + 6, 1);
    	} else {
    		TextOutgl(nil, obj.caption, startX + 9, startY + 6, 1);
    	}
    }

	if obj.mode != FLAT {
    	SetColor(0xF8FCF8);
    	LinePP(nil, startX, startY, startX + sizeX, startY);
    	LinePP(nil, startX, startY, startX, startY + sizeY);
    	SetColor(0xE0E0E0);
    	LinePP(nil, startX+1, startY+1, startX + sizeX - 2, startY+1);
    	LinePP(nil, startX+1, startY+1, startX+1, startY + sizeY - 1);
    	SetColor(0x787C78);
    	LinePP(nil, startX+2, startY + sizeY - 1, startX + sizeX - 1, startY + sizeY - 1);
    	LinePP(nil, startX + sizeX - 1, startY + 1, startX + sizeX - 1, startY + sizeY - 1);
    	SetColor(0x000000);
    	LinePP(nil, startX, startY + sizeY, startX + sizeX, startY + sizeY);
    	LinePP(nil, startX + sizeX, startY, startX + sizeX, startY + sizeY);  
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




