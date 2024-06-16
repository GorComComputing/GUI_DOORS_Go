package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tTable struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC uint32
    TC uint32
    visible bool
    focused bool
    enabled bool
	cols []string
	sizeCols []int
	rows []string
    list [][]string
    selectedX int
    selectedY int
    cellX int
    cellY int
    align tAlign
    onClick func(*Node, int, int)
    onClickStr string
    onEnter func(*Node)
    onEnterStr string
}


func CreateTable(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC uint32, TC uint32, cols []string, sizeCols []int, rows []string, list [][]string, cellX int, cellY int, onClick func(*Node, int, int), onEnter func(*Node)) *Node {
	obj := tTable{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, visible: true, enabled: true, cols: cols, sizeCols: sizeCols, rows: rows, list: list, selectedX: 0, selectedY: 0, cellX: cellX, cellY: cellY, onClick: onClick, onEnter: onEnter}
	node := Node{typ: TABLE, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj *tTable) Draw(parX int, parY int, parSizeX int, parSizeY int){
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
    
    SetColor(0x000000);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x + obj.sizeX, parY+obj.y);
    LinePP(nil, parX+obj.x, parY+obj.y, parX+obj.x, parY+obj.y + obj.sizeY);
    
    var col int = 0
	if obj.cols != nil {
		col = obj.cellY
	}
	var row int = 0
	if obj.rows != nil {
		row = obj.cellX
	}
    if obj.rows != nil && obj.cols != nil {
    	SetColor(0xd8dcc0);
    		var p []tPoint

   			p1 := tPoint{x: parX+obj.x +1+1, y: parY+obj.y+1}
			p = append(p, p1)
	
			p2 := tPoint{x: parX+obj.x + obj.cellX, y: parY+obj.y+1}
			p = append(p, p2)
	
			p3 := tPoint{x: parX+obj.x + obj.cellX, y: parY+obj.y + obj.cellY+1}
			p = append(p, p3)
	
			p4 := tPoint{x: parX+obj.x+1, y: parY+obj.y + obj.cellY+1}
			p = append(p, p4)

    		FillPoly(nil, 4, p);
    		SetColor(0x000000);
    	
    	SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x + obj.cellX, parY + obj.y+1);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1, parX + obj.x+1, parY + obj.y + obj.cellY+1);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+1+1, parY + obj.y+1+1, parX + obj.x + obj.cellX - 2+1, parY + obj.y+1+1);
    	LinePP(nil, parX + obj.x+1+1, parY + obj.y+1+1, parX + obj.x+1+1, parY + obj.y + obj.cellY - 1+1);
    	SetColor(0x787C78);
    	LinePP(nil, parX + obj.x+2+1, parY + obj.y + obj.cellY - 1+1, parX + obj.x + obj.cellX - 1+1, parY + obj.y + obj.cellY - 1+1);
    	LinePP(nil, parX + obj.x + obj.cellX - 1+1, parY + obj.y + 1+1, parX + obj.x + obj.cellX - 1+1, parY + obj.y + obj.cellY - 1+1);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+1+1, parY + obj.y + obj.cellY+1, parX + obj.x + obj.cellX, parY + obj.y + obj.cellY+1);
    	LinePP(nil, parX + obj.x + obj.cellX, parY + obj.y+1+1, parX + obj.x + obj.cellX, parY + obj.y + obj.cellY+1);
    }
	
	if obj.cols != nil {
	var x int = 0
	for j := 0; j < len(obj.cols); j++ {
    		SetColor(0xd8dcc0);
    		var p []tPoint

   			p1 := tPoint{x: parX+obj.x + x+1+row, y: parY+obj.y+1}
			p = append(p, p1)
	
			p2 := tPoint{x: parX+obj.x + x + obj.sizeCols[j]+1+row, y: parY+obj.y+1}
			p = append(p, p2)
	
			p3 := tPoint{x: parX+obj.x + x + obj.sizeCols[j]+1+row, y: parY+obj.y + obj.cellY+1}
			p = append(p, p3)
	
			p4 := tPoint{x: parX+obj.x + x+1+row, y: parY+obj.y + obj.cellY+1}
			p = append(p, p4)

    		FillPoly(nil, 4, p);
    		SetColor(0x000000);

    	TextOutgl(nil, obj.cols[j], parX+obj.x + 4 + x+1+row, parY+obj.y + 4+1, 1);
    	
    	SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x + x+1+row, parY + obj.y+1, parX + obj.x + obj.sizeCols[j] + x+1+row, parY + obj.y+1);
    	LinePP(nil, parX + obj.x + x+1+row, parY + obj.y+1, parX + obj.x + x+1+row, parY + obj.y + obj.cellY+1);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+1 + x+1+row, parY + obj.y+1+1, parX + obj.x + obj.sizeCols[j] + x - 2+1+row, parY + obj.y+1+1);
    	LinePP(nil, parX + obj.x+1 + x+1+row, parY + obj.y+1+1, parX + obj.x+1 + x+1+row, parY + obj.y + obj.cellY - 1+1);
    	SetColor(0x787C78);
    	LinePP(nil, parX + obj.x+2 + x+1+row, parY + obj.y + obj.cellY - 1+1, parX + obj.x + obj.sizeCols[j] + x - 1+1+row, parY + obj.y + obj.cellY - 1+1);
    	LinePP(nil, parX + obj.x + obj.sizeCols[j] + x - 1+1+row, parY + obj.y + 1+1, parX + obj.x + obj.sizeCols[j] + x - 1+1+row, parY + obj.y + obj.cellY - 1+1);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+1 + x+1+row, parY + obj.y + obj.cellY+1, parX + obj.x + obj.sizeCols[j] + x+1+row, parY + obj.y + obj.cellY+1);
    	LinePP(nil, parX + obj.x + obj.sizeCols[j] + x+1+row, parY + obj.y+1+1, parX + obj.x + obj.sizeCols[j] + x+1+row, parY + obj.y + obj.cellY+1);
    	
    	x += obj.sizeCols[j]
    }
    }
    
    
	if obj.rows != nil {
		for j := 0; j < len(obj.rows); j++ {
    		SetColor(0xd8dcc0);
    		var p []tPoint

   			p1 := tPoint{x: parX+obj.x +1+1, y: parY+obj.y+1 + j*obj.cellY+col}
			p = append(p, p1)
	
			p2 := tPoint{x: parX+obj.x + obj.cellX+1, y: parY+obj.y+1 + j*obj.cellY+col}
			p = append(p, p2)
	
			p3 := tPoint{x: parX+obj.x + obj.cellX+1, y: parY+obj.y + obj.cellY+1 + j*obj.cellY+col}
			p = append(p, p3)
	
			p4 := tPoint{x: parX+obj.x+1, y: parY+obj.y + obj.cellY+1 + j*obj.cellY+col}
			p = append(p, p4)

    		FillPoly(nil, 4, p);
    		SetColor(0x000000);

    	TextOutgl(nil, obj.rows[j], parX+obj.x + 4+1, parY+obj.y + 4+1 + j*obj.cellY+col, 1);
    	
    	SetColor(0xF8FCF8);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1 + j*obj.cellY+col, parX + obj.x + obj.cellX+1, parY + obj.y+1 + j*obj.cellY+col);
    	LinePP(nil, parX + obj.x+1, parY + obj.y+1 + j*obj.cellY+col, parX + obj.x+1, parY + obj.y + obj.cellY+1 + j*obj.cellY+col);
    	SetColor(0xE0E0E0);
    	LinePP(nil, parX + obj.x+1+1, parY + obj.y+1+1 + j*obj.cellY+col, parX + obj.x + obj.cellX - 2+1+1, parY + obj.y+1+1 + j*obj.cellY+col);
    	LinePP(nil, parX + obj.x+1+1, parY + obj.y+1+1 + j*obj.cellY+col, parX + obj.x+1+1, parY + obj.y + obj.cellY - 1+1 + j*obj.cellY+col);
    	SetColor(0x787C78);
    	LinePP(nil, parX + obj.x+2+1, parY + obj.y + obj.cellY - 1+1 + j*obj.cellY+col, parX + obj.x + obj.cellX - 1+1+1, parY + obj.y + obj.cellY - 1+1 + j*obj.cellY+col);
    	LinePP(nil, parX + obj.x + obj.cellX - 1+1+1, parY + obj.y + 1+1 + j*obj.cellY+col, parX + obj.x + obj.cellX - 1+1+1, parY + obj.y + obj.cellY - 1+1 + j*obj.cellY+col);
    	SetColor(0x000000);
    	LinePP(nil, parX + obj.x+1+1, parY + obj.y + obj.cellY+1 + j*obj.cellY+col, parX + obj.x + obj.cellX+1, parY + obj.y + obj.cellY+1 + j*obj.cellY+col);
    	LinePP(nil, parX + obj.x + obj.cellX+1, parY + obj.y+1+1 + j*obj.cellY+col, parX + obj.x + obj.cellX+1, parY + obj.y + obj.cellY+1 + j*obj.cellY+col);
    }
    }
    
    

    SetColor(obj.TC);
    SetBackColor(obj.BC);
    for i := 0; i < len(obj.list); i++ {
    	var x int = 0
    	for j := 0; j < len(obj.list[i]); j++ {
    	
    	SetLocalViewPort(parX + obj.x, parY + obj.y, parX + obj.x + obj.sizeX, parY + obj.y + obj.sizeY)
    	
    	if i == obj.selectedY && obj.enabled {  // && j == obj.selectedX
    		SetColor(0x0054E0);
    		var p []tPoint

   			p1 := tPoint{x: parX+obj.x + x+1 + row, y: parY+obj.y + i*obj.cellY + col+1}
			p = append(p, p1)
	
			p2 := tPoint{x: parX+obj.x + obj.sizeCols[j] + x+1 + row, y: parY+obj.y + i*obj.cellY + col+1}
			p = append(p, p2)
	
			p3 := tPoint{x: parX+obj.x + obj.sizeCols[j] + x+1 + row, y: parY+obj.y + obj.cellY + i*obj.cellY + col+1}
			p = append(p, p3)
	
			p4 := tPoint{x: parX+obj.x + x+1 + row, y: parY+obj.y + obj.cellY + i*obj.cellY + col+1}
			p = append(p, p4)

    		FillPoly(nil, 4, p);
    		SetColor(0xF8FCF8);
    	} else {
    		SetColor(obj.TC);
    	}
    	
    	SetLocalViewPort(parX+obj.x + x+1 + row, parY+obj.y + i*obj.cellY + col+1, parX+obj.x + x+1 + row + obj.sizeCols[j], parY+obj.y + i*obj.cellY + col+1 + obj.cellY)
    	TextOutgl(nil, obj.list[i][j], parX+obj.x + 4 + x+1 + row, parY+obj.y + 4 + i*obj.cellY + col+1, 1)
    	
    	SetColor(0x000000);
   		LinePP(nil, parX+obj.x + x + row, parY+obj.y + i*obj.cellY + col+1, parX+obj.x + obj.sizeCols[j] + x + row, parY+obj.y + i*obj.cellY + col+1)
    	LinePP(nil, parX+obj.x + x + row, parY+obj.y + i*obj.cellY + col+1, parX+obj.x + x + row, parY+obj.y + obj.cellY + i*obj.cellY + col+1)
    	
    	x += obj.sizeCols[j]
    	}
    	
    }
    
   /* if obj.focused && cursor {
    	TextOutgl(nil, "|", parX+obj.x + 4 + obj.curX*8, parY+obj.y + 4 + (7+2)*obj.curY, 1);
    }*/

    
    
}


func (obj *tTable) RAD(x int, y int){
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
		editPropSelected = CreateEdit(pnlProperties, "editPropSelected", 80, 65, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.selectedX), nil, editPropSelectedEnter)
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
    	cmbPropList = CreateComboBox(pnlProperties, "cmbPropList", 80, 205, 95, 16, 0xF8FCF8, 0x000000, "", obj.list[0], nil, cmbPropListEnter)
			
		lblEvntClick = CreateLabel(pnlEvents, "lblEvntClick", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
		editEvntClick = CreateEdit(pnlEvents, "editEvntClick", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
}


func (obj *tTable) KeyDown(key int){
	if key == 8 {
    			
    } else if key == 13 {
    	if obj.onEnter != nil && obj.enabled {
    		obj.onEnter(layout.children[len(layout.children)-1].obj.(*tForm).focus)	
    	}
    } else if key == 37 {
    	if obj.selectedX >= 1 && obj.enabled && len(list) > 0 {
    				obj.selectedX--
    	}
    			
    } else if key == 39 {
   		if obj.selectedX < len(obj.list)-1  && obj.enabled && len(list) > 0 {
    		obj.selectedX++
    	}
    			
    } else if key == 38 {
    	if obj.selectedY >= 1 && obj.enabled && len(list) > 0 {
    				obj.selectedY--
    	}
    } else if key == 40 {
    	if obj.selectedY < len(obj.list)-1  && obj.enabled && len(list) > 0 {
    		obj.selectedY++
    	}
	} else {

	}
}


func (obj *tTable) Click(x int, y int){
	fmt.Println("CLICKED: ", strconv.Itoa(x), strconv.Itoa(y))
	var col int = 0
	if obj.cols != nil {
		col = obj.cellY
	}
	var row int = 0
	if obj.rows != nil {
		row = obj.cellX
	}
	if obj.enabled && len(list) > 0 && y > col && x > row {
		obj.selectedY = int((y-col)/obj.cellY)
		obj.selectedX = int((x-row)/obj.cellX)
	}
	
	if obj.onClick != nil && obj.enabled {
		obj.onClick(list[len(list)-1], x, y)
	}
}


func (obj *tTable) MouseMove(x int, y int, Xl int, Yl int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && mouseIsDown && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
		obj.x += x - downX
    	obj.y += y - downY
    	editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
		editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    }
}


func (obj *tTable) MouseDown(x int, y int){
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

