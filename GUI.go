package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"
    "strings"
    //"reflect"

)


var mouseIsDown bool = false
var cursor bool
var RAD bool = false
var RADElement *Node
var RADFormElement *Node

var layout_obj = tForm{x: 0, y: 0, sizeX: BITMAP_WIDTH-1, sizeY: BITMAP_HEIGHT-2, BC: 0x000000, mode: FLAT, caption: "", visible: true, onClick: nil}
var layout = Node{parent: nil, previous: nil, children: nil, obj: &layout_obj}
var list []*Node


type tComponents int

const (
    BUTTON tComponents = iota
    FORM
    EDIT
    LABEL
    PANEL
    CANVAS
    BIT_BUTTON
    MEMO
    CHECKBOX
)

type tWinComponents interface {
   Draw(parX int, parY int, parSizeX int, parSizeY int)
}


type Node struct {
	typ tComponents 
    parent *Node
    previous *Node
    children []*Node
    obj tWinComponents 
}

var lblPropTop *Node
var editPropTop *Node
var lblPropLeft *Node
var editPropLeft *Node
var lblPropCaption *Node
var editPropCaption *Node
var lblPropBC *Node
var editPropBC *Node
var lblPropWidth *Node
var editPropWidth *Node
var lblPropHeight *Node
var editPropHeight *Node
var lblPropTC *Node
var editPropTC *Node
var lblPropText *Node
var editPropText *Node
var lblPropName *Node
var editPropName *Node
var lblPropMode *Node
var editPropMode *Node
var lblPropVisible *Node
var editPropVisible *Node
var lblPropEnabled *Node
var editPropEnabled *Node
var lblPropChecked *Node
var editPropChecked *Node
var lblEvntClick *Node
var editEvntClick *Node
var lblEvntEnter *Node
var editEvntEnter *Node
         

func DrawNode(node *Node){
	
	var visible bool = false
	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tBtn:
			visible = obj.visible
		case *tForm:
			visible = obj.visible
		case *tPanel:
			visible = obj.visible
		case *tEdit:
			visible = obj.visible
		case *tLabel:
			visible = obj.visible
		case *tCanvas:
			visible = obj.visible
		case *tBitBtn:
			visible = obj.visible
		case *tMemo:
			visible = obj.visible
		case *tCheckBox:
			visible = obj.visible
		}
	}
	
	var parX int = 0
	var parY int = 0
	var parSizeX int = GETMAX_X
	var parSizeY int = GETMAX_Y
	if node.parent != nil && node.parent.obj != nil {
		switch obj := node.parent.obj.(type) {
		case *tBtn:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tForm:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tPanel:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tEdit:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tLabel:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tCanvas:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tBitBtn:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tMemo:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		case *tCheckBox:
			parX = obj.x
			parY = obj.y
			parSizeX = obj.sizeX
			parSizeY = obj.sizeY
		}
	}
	
	if node.obj != nil && visible  {
		node.obj.Draw(parX, parY, parSizeX, parSizeY)
	}
	
	if node.children != nil && visible {
			for i := 0; i < len(node.children); i++ { 
				DrawNode(node.children[i])
			}
	}
		
	return
}


//export eventClick
func eventClick(x int, y int)  {
	
		fmt.Println("Event: " + strconv.Itoa(x) + " " + strconv.Itoa(y))
		list = nil
		ClickRecurs(&layout, x, y, 0, 0)
		
		if !RAD || list[len(list)-1] == cbxRAD || layout.children[len(layout.children)-1] == frmProperties || layout.children[len(layout.children)-1] == frmRAD || layout.children[len(layout.children)-1] == frmCode {
		switch list[len(list)-1].obj.(type) {
		case *tBtn:
			fmt.Println("CLICKED: " + list[len(list)-1].obj.(*tBtn).caption)
			if list[len(list)-1].obj.(*tBtn).onClick != nil && list[len(list)-1].obj.(*tBtn).enabled {
				list[len(list)-1].obj.(*tBtn).onClick(list[len(list)-1])
			}
		case *tBitBtn:
			fmt.Println("CLICKED: " + list[len(list)-1].obj.(*tBitBtn).caption)
			if list[len(list)-1].obj.(*tBitBtn).onClick != nil && list[len(list)-1].obj.(*tBitBtn).enabled {
				list[len(list)-1].obj.(*tBitBtn).onClick(list[len(list)-1])
			}
		case *tCheckBox:
			fmt.Println("CLICKED: " + list[len(list)-1].obj.(*tCheckBox).caption)
			if list[len(list)-1].obj.(*tCheckBox).onClick != nil && list[len(list)-1].obj.(*tCheckBox).enabled {
				list[len(list)-1].obj.(*tCheckBox).onClick(list[len(list)-1])
			}
		case *tEdit:
			fmt.Println("CLICKED: " + list[len(list)-1].obj.(*tEdit).text)
			if list[len(list)-1].obj.(*tEdit).onClick != nil && list[len(list)-1].obj.(*tEdit).enabled {
				list[len(list)-1].obj.(*tEdit).onClick(list[len(list)-1])
			}
		}
	}
}


func ClickRecurs(node *Node, x int, y int, parX int, parY int) {
	
	var visible bool = false
	if node.obj != nil {
		switch obj := node.obj.(type) {
		case *tBtn:
			visible = obj.visible
		case *tForm:
			visible = obj.visible
		case *tPanel:
			visible = obj.visible
		case *tEdit:
			visible = obj.visible
		case *tLabel:
			visible = obj.visible
		case *tCanvas:
			visible = obj.visible
		case *tBitBtn:
			visible = obj.visible
		case *tMemo:
			visible = obj.visible
		case *tCheckBox:
			visible = obj.visible
		}
	}
	
	if node.parent != nil && node.parent.obj != nil {
		switch obj := node.parent.obj.(type) {
		case *tBtn:
			parX = obj.x
			parY = obj.y
		case *tForm:
			parX = obj.x
			parY = obj.y
		case *tPanel:
			parX = obj.x
			parY = obj.y
		case *tEdit:
			parX = obj.x
			parY = obj.y
		case *tLabel:
			parX = obj.x
			parY = obj.y
		case *tCanvas:
			parX = obj.x
			parY = obj.y
		case *tBitBtn:
			parX = obj.x
			parY = obj.y
		case *tMemo:
			parX = obj.x
			parY = obj.y
		case *tCheckBox:
			parX = obj.x
			parY = obj.y
		}
	}
	
	if node.obj != nil && visible {
		switch node.obj.(type) {
		case *tBtn:
			if (parX+node.obj.(*tBtn).x) < x && 
			(parX+node.obj.(*tBtn).x + node.obj.(*tBtn).sizeX) > x && 
			(parY+node.obj.(*tBtn).y) < y && 
			(parY+node.obj.(*tBtn).y + node.obj.(*tBtn).sizeY) > y {
				list = append(list, node)
			}
		case *tForm:
			if (parX+node.obj.(*tForm).x) < x && 
			(parX+node.obj.(*tForm).x + node.obj.(*tForm).sizeX) > x && 
			(parY+node.obj.(*tForm).y) < y && 
			(parY+node.obj.(*tForm).y + node.obj.(*tForm).sizeY) > y {
				list = append(list, node)
			}
		case *tEdit:
			if (parX+node.obj.(*tEdit).x) < x && 
			(parX+node.obj.(*tEdit).x + node.obj.(*tEdit).sizeX) > x && 
			(parY+node.obj.(*tEdit).y) < y && 
			(parY+node.obj.(*tEdit).y + node.obj.(*tEdit).sizeY) > y {
				list = append(list, node)
			}
		case *tLabel:
			if (parX+node.obj.(*tLabel).x) < x && 
			(parX+node.obj.(*tLabel).x + node.obj.(*tLabel).sizeX) > x && 
			(parY+node.obj.(*tLabel).y) < y && 
			(parY+node.obj.(*tLabel).y + node.obj.(*tLabel).sizeY) > y {
				list = append(list, node)
			}
		case *tPanel:
			if (parX+node.obj.(*tPanel).x) < x && 
			(parX+node.obj.(*tPanel).x + node.obj.(*tPanel).sizeX) > x && 
			(parY+node.obj.(*tPanel).y) < y && 
			(parY+node.obj.(*tPanel).y + node.obj.(*tPanel).sizeY) > y {
				list = append(list, node)
			}
		case *tCanvas:
			if (parX+node.obj.(*tCanvas).x) < x && 
			(parX+node.obj.(*tCanvas).x + node.obj.(*tCanvas).sizeX) > x && 
			(parY+node.obj.(*tCanvas).y) < y && 
			(parY+node.obj.(*tCanvas).y + node.obj.(*tCanvas).sizeY) > y {
				list = append(list, node)
			}
		case *tBitBtn:
			if (parX+node.obj.(*tBitBtn).x) < x && 
			(parX+node.obj.(*tBitBtn).x + node.obj.(*tBitBtn).sizeX) > x && 
			(parY+node.obj.(*tBitBtn).y) < y && 
			(parY+node.obj.(*tBitBtn).y + node.obj.(*tBitBtn).sizeY) > y {
				list = append(list, node)
			}
		case *tMemo:
			if (parX+node.obj.(*tMemo).x) < x && 
			(parX+node.obj.(*tMemo).x + node.obj.(*tMemo).sizeX) > x && 
			(parY+node.obj.(*tMemo).y) < y && 
			(parY+node.obj.(*tMemo).y + node.obj.(*tMemo).sizeY) > y {
				list = append(list, node)
			}
		case *tCheckBox:
			if (parX+node.obj.(*tCheckBox).x) < x && 
			(parX+node.obj.(*tCheckBox).x + node.obj.(*tCheckBox).sizeX) > x && 
			(parY+node.obj.(*tCheckBox).y) < y && 
			(parY+node.obj.(*tCheckBox).y + node.obj.(*tCheckBox).sizeY) > y {
				list = append(list, node)
			}
		}
	}
			
	if node.children != nil && visible {
		for i := 0; i < len(node.children); i++ { 
			ClickRecurs(node.children[i], x, y, parX, parY)
		}
	}
	return
}


func findNode(node *Node) int {
	var i int
	if node.typ == FORM {
		for i := 0; i < len(layout.children); i++ {
			if node == layout.children[i] {
				return i
			}
		}
	} else {
		if node.parent != nil {
			i = findNode(node.parent)
		} else {
			return -1
		}
	}
	return i
}


func sortChildren(i int) {
	tmp := layout.children[i]
	copy(layout.children[i:], layout.children[i+1:])
	layout.children[len(layout.children)-1] = tmp
	layout.children[len(layout.children)-2].obj.(*tForm).focused = false
	layout.children[len(layout.children)-1].obj.(*tForm).focused = true
}





var downX int = 0
var downY int = 0
var btnPressed *tBtn = nil
var bitbtnPressed *tBitBtn = nil

//export eventMouseDown
func eventMouseDown(x int, y int)  {
	list = nil
	ClickRecurs(&layout, x, y, 0, 0)
	
	i := findNode(list[len(list)-1])
	fmt.Println(i)
	if i > 0 {
		sortChildren(i)
	}
	if layout.children[len(layout.children)-1].obj.(*tForm).focus != nil {
		switch obj := layout.children[len(layout.children)-1].obj.(*tForm).focus.obj.(type) {
    	case *tEdit:
			obj.focused = false
		case *tMemo:
			obj.focused = false
		}
	}
	layout.children[len(layout.children)-1].obj.(*tForm).focus = list[len(list)-1]
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			RADElement = list[len(list)-1]
			for i := 0; i < len(layout.children); i++ {
				layout.children[i].obj.(*tForm).RAD = false
			}
			layout.children[len(layout.children)-1].obj.(*tForm).RAD = true
			RADFormElement = layout.children[len(layout.children)-1]
			frmProperties.children = nil
	}
	
	switch obj := list[len(list)-1].obj.(type) {
	case *tForm:
		if (obj.mode == WIN) &&
			(obj.x) < x && 
			(obj.x + obj.sizeX) > x && 
			(obj.y) < y && 
			(obj.y + 17) > y {
				downX = x 
    			downY = y 
    			mouseIsDown = true
    	}
    	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
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
    case *tBtn:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
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
    		
			frmProperties.obj.(*tForm).caption = "Properties: BUTTON"
			downX = x 
    		downY = y 
    		mouseIsDown = true
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
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 240, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 240, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
			
		} else {
			if obj.enabled {
    			btnPressed = obj
				obj.pressed = true	
			}
		}
	case *tLabel:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			var visible string
    		if obj.visible {
    			visible = "true"
    		} else {
    			visible = "false"
    		}
    		
			frmProperties.obj.(*tForm).caption = "Properties: LABEL"
			downX = x 
    		downY = y 
    		mouseIsDown = true
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
			lblPropTC = CreateLabel(frmProperties, "lblPropTC", 5, 120, 95, 20, 0xD8DCC0, 0x000000, "TC", nil)
			editPropTC = CreateEdit(frmProperties, "editPropTC", 80, 120, 95, 20, 0xF8FCF8, 0x000000, fmt.Sprintf("%x", obj.TC), nil, editPropTCEnter)
			lblPropWidth = CreateLabel(frmProperties, "lblPropWidth", 5, 140, 95, 20, 0xD8DCC0, 0x000000, "Width", nil)
			editPropWidth = CreateEdit(frmProperties, "editPropWidth", 80, 140, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeX), nil, editPropWidthEnter)
			lblPropHeight = CreateLabel(frmProperties, "lblPropHeight", 5, 160, 95, 20, 0xD8DCC0, 0x000000, "Height", nil)
			editPropHeight = CreateEdit(frmProperties, "editPropHeight", 80, 160, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeY), nil, editPropHeightEnter)
			lblPropVisible = CreateLabel(frmProperties, "lblPropVisible", 5, 180, 95, 20, 0xD8DCC0, 0x000000, "Visible", nil)
			editPropVisible = CreateEdit(frmProperties, "editPropVisible", 80, 180, 95, 20, 0xF8FCF8, 0x000000, visible, nil, editPropVisibleEnter)
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 220, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 220, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
		} 
	case *tEdit:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
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
    		
			frmProperties.obj.(*tForm).caption = "Properties: EDIT"
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
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 240, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 240, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
			lblEvntEnter = CreateLabel(frmProperties, "lblEvntEnter", 5, 260, 95, 20, 0xD8DCC0, 0x000000, "Enter", nil)
			editEvntEnter = CreateEdit(frmProperties, "editEvntEnter", 80, 260, 95, 20, 0xF8FCF8, 0x000000, obj.onEnterStr, nil, editEvntEnterEnter)
		} else {
			if obj.enabled {
				obj.focused = true	
				obj.curX = len(obj.text)
			}
		}
	case *tBitBtn:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
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
			
			frmProperties.obj.(*tForm).caption = "Properties: BITBTN"
			downX = x 
    		downY = y 
    		mouseIsDown = true
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
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 240, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 240, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
		} else {
			if obj.enabled {
    			bitbtnPressed = obj
				obj.pressed = true	
			}
		}
	case *tMemo:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
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
    		
			frmProperties.obj.(*tForm).caption = "Properties: MEMO"
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
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 240, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 240, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
		} else {
			if obj.enabled {
				obj.focused = true	
			}
		}
	case *tCheckBox:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
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
    		var checked string
    		if obj.checked {
    			checked = "true"
    		} else {
    			checked = "false"
    		}
    		
			frmProperties.obj.(*tForm).caption = "Properties: CHECKBOX"
			downX = x 
    		downY = y 
    		mouseIsDown = true
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
			lblPropChecked = CreateLabel(frmProperties, "lblPropChecked", 5, 220, 95, 20, 0xD8DCC0, 0x000000, "Checked", nil)
			editPropChecked = CreateEdit(frmProperties, "editPropChecked", 80, 220, 95, 20, 0xF8FCF8, 0x000000, checked, nil, editPropCheckedEnter)
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 260, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 260, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
		}
	case *tPanel:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
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
    		
			frmProperties.obj.(*tForm).caption = "Properties: PANEL"
			downX = x 
    		downY = y 
    		mouseIsDown = true
    		lblPropName = CreateLabel(frmProperties, "lblPropName", 5, 20, 95, 20, 0xD8DCC0, 0x000000, "Name", nil)
			editPropName = CreateEdit(frmProperties, "editPropName", 80, 20, 95, 20, 0xF8FCF8, 0x000000, obj.name, nil, editPropNameEnter)
			lblPropLeft = CreateLabel(frmProperties, "lblPropLeft", 5, 40, 95, 20, 0xD8DCC0, 0x000000, "Left", nil)
			editPropLeft = CreateEdit(frmProperties, "editPropLeft", 80, 40, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.x), nil, editPropLeftEnter)
			lblPropTop = CreateLabel(frmProperties, "lblPropTop", 5, 60, 95, 20, 0xD8DCC0, 0x000000, "Top", nil)
			editPropTop = CreateEdit(frmProperties, "editPropTop", 80, 60, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.y), nil, editPropTopEnter)
			lblPropBC = CreateLabel(frmProperties, "lblPropBC", 5, 80, 95, 20, 0xD8DCC0, 0x000000, "BC", nil)
			editPropBC = CreateEdit(frmProperties, "editPropBC", 80, 80, 95, 20, 0xF8FCF8, 0x000000, fmt.Sprintf("%x", obj.BC), nil, editPropBCEnter)
			lblPropWidth = CreateLabel(frmProperties, "lblPropWidth", 5, 100, 95, 20, 0xD8DCC0, 0x000000, "Width", nil)
			editPropWidth = CreateEdit(frmProperties, "editPropWidth", 80, 100, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeX), nil, editPropWidthEnter)
			lblPropHeight = CreateLabel(frmProperties, "lblPropHeight", 5, 120, 95, 20, 0xD8DCC0, 0x000000, "Height", nil)
			editPropHeight = CreateEdit(frmProperties, "editPropHeight", 80, 120, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.sizeY), nil, editPropHeightEnter)
			lblPropMode = CreateLabel(frmProperties, "lblPropMode", 5, 140, 95, 20, 0xD8DCC0, 0x000000, "Mode", nil)
			editPropMode = CreateEdit(frmProperties, "editPropMode", 80, 140, 95, 20, 0xF8FCF8, 0x000000, mode, nil, editPropModeEnter)
			lblPropVisible = CreateLabel(frmProperties, "lblPropVisible", 5, 160, 95, 20, 0xD8DCC0, 0x000000, "Visible", nil)
			editPropVisible = CreateEdit(frmProperties, "editPropVisible", 80, 160, 95, 20, 0xF8FCF8, 0x000000, visible, nil, editPropVisibleEnter)
			
			lblEvntClick = CreateLabel(frmProperties, "lblEvntClick", 5, 200, 95, 20, 0xD8DCC0, 0x000000, "Click", nil)
			editEvntClick = CreateEdit(frmProperties, "editEvntClick", 80, 200, 95, 20, 0xF8FCF8, 0x000000, obj.onClickStr, nil, editEvntClickEnter)
		} 
	}
}


func editPropNameEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.name = node.obj.(*tEdit).text
	case *tBtn:
		obj.name  = node.obj.(*tEdit).text
	case *tEdit:
		obj.name  = node.obj.(*tEdit).text
	case *tLabel:
		obj.name  = node.obj.(*tEdit).text
	case *tPanel:
		obj.name  = node.obj.(*tEdit).text
	case *tMemo:
		obj.name  = node.obj.(*tEdit).text
	case *tBitBtn:
		obj.name  = node.obj.(*tEdit).text
	case *tCheckBox:
		obj.name  = node.obj.(*tEdit).text
	case *tCanvas:
		obj.name  = node.obj.(*tEdit).text
	}
}


func editPropLeftEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		obj.x, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropTopEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropCaptionEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.caption = node.obj.(*tEdit).text
	case *tBtn:
		obj.caption = node.obj.(*tEdit).text
	case *tEdit:
		//obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.caption = node.obj.(*tEdit).text
	case *tPanel:
		//obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		//obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.caption = node.obj.(*tEdit).text
	case *tCheckBox:
		obj.caption = node.obj.(*tEdit).text
	case *tCanvas:
		//obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropTextEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		//obj.text = node.obj.(*tEdit).text
	case *tBtn:
		//obj.caption = node.obj.(*tEdit).text
	case *tEdit:
		obj.text = node.obj.(*tEdit).text
	case *tLabel:
		//obj.caption = node.obj.(*tEdit).text
	case *tPanel:
		//obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.text = node.obj.(*tEdit).text
	case *tBitBtn:
		//obj.caption = node.obj.(*tEdit).text
	case *tCheckBox:
		//obj.caption = node.obj.(*tEdit).text
	case *tCanvas:
		//obj.y, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropWidthEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		obj.sizeX, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropHeightEnter(node *Node){
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropBCEnter(node *Node){
	val, _ := strconv.ParseInt(node.obj.(*tEdit).text, 16, 32)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.BC = int(val)
	case *tBtn:
		obj.BC = int(val)
	case *tEdit:
		obj.BC = int(val)
	case *tLabel:
		obj.BC = int(val)
	case *tPanel:
		obj.BC = int(val)
	case *tMemo:
		obj.BC = int(val)
	case *tBitBtn:
		obj.BC = int(val)
	case *tCheckBox:
		obj.BC = int(val)
	case *tCanvas:
		//obj.BC, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}


func editPropTCEnter(node *Node){
	val, _ := strconv.ParseInt(node.obj.(*tEdit).text, 16, 32)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		//val, _ := strconv.ParseInt(node.obj.(*tEdit).text, 16, 32)
		//obj.TC = int(val)
	case *tBtn:
		obj.TC = int(val)
	case *tEdit:
		obj.TC = int(val)
	case *tLabel:
		obj.TC = int(val)
	case *tPanel:
		//obj.BC, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		obj.TC = int(val)
	case *tBitBtn:
		obj.TC = int(val)
	case *tCheckBox:
		obj.TC = int(val)
	case *tCanvas:
		//obj.BC, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}

func editPropModeEnter(node *Node){
	//node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		if node.obj.(*tEdit).text == "NONE" {
			obj.mode = NONE	
		} else if node.obj.(*tEdit).text == "WIN" {
			obj.mode = WIN
		} else if node.obj.(*tEdit).text == "FLAT" {
			obj.mode = FLAT
		} else if node.obj.(*tEdit).text == "TASK" {
			obj.mode = TASK
		}
	case *tBtn:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		if node.obj.(*tEdit).text == "NONE" {
			obj.mode = NONE	
		} else if node.obj.(*tEdit).text == "WIN" {
			obj.mode = WIN
		} else if node.obj.(*tEdit).text == "FLAT" {
			obj.mode = FLAT
		} else if node.obj.(*tEdit).text == "TASK" {
			obj.mode = TASK
		}
	case *tMemo:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCanvas:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}

func editPropVisibleEnter(node *Node){
	node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	case *tBtn:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	case *tEdit:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	case *tLabel:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	case *tPanel:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	case *tMemo:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	case *tBitBtn:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	case *tCheckBox:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	case *tCanvas:
		if node.obj.(*tEdit).text == "true" {
			obj.visible = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.visible = false
		}
	}
}

func editPropEnabledEnter(node *Node){
	node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		if node.obj.(*tEdit).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.enabled = false
		}
	case *tEdit:
		if node.obj.(*tEdit).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.enabled = false
		}
	case *tLabel:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		if node.obj.(*tEdit).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.enabled = false
		}
	case *tBitBtn:
		if node.obj.(*tEdit).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.enabled = false
		}
	case *tCheckBox:
		if node.obj.(*tEdit).text == "true" {
			obj.enabled = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.enabled = false
		}
	case *tCanvas:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}

func editPropCheckedEnter(node *Node){
	node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBtn:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tEdit:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tLabel:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tPanel:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tMemo:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tBitBtn:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	case *tCheckBox:
		if node.obj.(*tEdit).text == "true" {
			obj.checked = true	
		} else if node.obj.(*tEdit).text == "false" {
			obj.checked = false
		}
	case *tCanvas:
		//obj.sizeY, _ = strconv.Atoi(node.obj.(*tEdit).text)
	}
}

func editEvntClickEnter(node *Node){
	node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tBtn:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tEdit:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tLabel:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tPanel:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tMemo:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tBitBtn:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tCheckBox:
		obj.onClickStr = node.obj.(*tEdit).text
	case *tCanvas:
		obj.onClickStr = node.obj.(*tEdit).text
	}
}

func editEvntEnterEnter(node *Node){
	node.obj.(*tEdit).text = strings.ToLower(node.obj.(*tEdit).text)
	switch obj := RADElement.obj.(type) {
	case *tForm:
		//obj.onClickStr = node.obj.(*tEdit).text
	case *tBtn:
		//obj.onClickStr = node.obj.(*tEdit).text
	case *tEdit:
		obj.onEnterStr = node.obj.(*tEdit).text
	case *tLabel:
		//obj.onClickStr = node.obj.(*tEdit).text
	case *tPanel:
		//obj.onClickStr = node.obj.(*tEdit).text
	case *tMemo:
		//obj.onClickStr = node.obj.(*tEdit).text
	case *tBitBtn:
		//obj.onClickStr = node.obj.(*tEdit).text
	case *tCheckBox:
		//obj.onClickStr = node.obj.(*tEdit).text
	case *tCanvas:
		//obj.onClickStr = node.obj.(*tEdit).text
	}
}




//export eventMouseUp
func eventMouseUp(x int, y int)  {
	//if(mouseIsDown) mouseClick(e)
    mouseIsDown = false

	if btnPressed != nil {
		btnPressed.pressed = false
		btnPressed = nil
		bitbtnPressed = nil
	}

}


//export eventMouseMove
func eventMouseMove(x int, y int)  {
	if !mouseIsDown {return}
	
	switch obj := list[len(list)-1].obj.(type) {
	case *tBtn:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
    case *tLabel:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
    case *tEdit:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
    case *tMemo:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
    case *tCanvas:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
    case *tCheckBox:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
    case *tPanel:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
    case *tBitBtn:
		if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
    case *tForm:
		obj.x += x - downX
    	obj.y += y - downY	
    	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode {
			editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)	
		}		
	}
	
    downX = x 
    downY = y
    return
}


//export keyDown
func keyDown(key int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && key == 46 {
		for i := 0; i < len(RADElement.parent.children); i++ {
			if RADElement.parent.children[i] == RADElement {
				copy(RADElement.parent.children[i:], RADElement.parent.children[i+1:])
				RADElement.parent.children[len(RADElement.parent.children)-1] = nil
				RADElement.parent.children = RADElement.parent.children[:len(RADElement.parent.children)-1]
				frmProperties.children = nil
			}		
		}
	}
	if layout.children[len(layout.children)-1].obj.(*tForm).focus != nil {
		switch obj := layout.children[len(layout.children)-1].obj.(*tForm).focus.obj.(type) {
    	case *tEdit:
    		if (RAD && (layout.children[len(layout.children)-1] == frmProperties || layout.children[len(layout.children)-1] == frmRAD || layout.children[len(layout.children)-1] == frmCode) && obj.enabled) || (!RAD && obj.enabled) { 
    		if key == 8 {
    			if len(obj.text) > 0 {
    				obj.text = obj.text[:len(obj.text)-1]
    				obj.curX--
    			}
    		} else if key == 37 {
    			if obj.curX > 0 {
    				obj.curX--
    			}
    		} else if key == 39 {
    			if obj.curX < len(obj.text) {
    				obj.curX++
    			}
    		} else if key == 36 {
    				obj.curX = 0
    		} else if key == 35 {
    				obj.curX = len(obj.text)
    		} else if key == 13 {		
    			obj.onEnter(layout.children[len(layout.children)-1].obj.(*tForm).focus)				
    		} else {
    			var char string
    			switch key {
    			case 190:
    				char = "."
    			default:
    				char = string(key)
    			}
    			obj.text = obj.text[:obj.curX] + char + obj.text[obj.curX:]
				obj.curX++
			}
			}
		case *tMemo:
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
		case *tBtn:
			
		}
		fmt.Println(key)
		fmt.Println(string(key))
	}
}


func findLeft(str string, current int) int {
	if current == 0 {
		return 0
	}
	for i := current - 2; i >= 0; i-- {
		if str[i] == 13 {
			return i + 1
		}
	}
	return 0
}


func findRight(str string, current int) int {
	if current == len(str)-1 {
		return current
	}
	for i := current; i < len(str)-1; i++ {
		if str[i] == 13 {
			return i+1
		}
	}
	return len(str)-1
}


/*func findNodeByObj(obj tWinComponents) *Node {
	var i int
	if node.typ == FORM {
		for i := 0; i < len(layout.children); i++ {
			if node == layout.children[i] {
				return i
			}
		}
	} else {
		if node.parent != nil {
			i = findNode(node.parent)
		} else {
			return -1
		}
	}
	return i
}*/
