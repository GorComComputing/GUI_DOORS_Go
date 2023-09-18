package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"
    //"strings"

)


type tMemo struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    BC int
    TC int
    list []string
    color [][]int
    visible bool
    focused bool
    enabled bool
    curX int
    curY int
    curXR int
    curYR int
    align tAlign
    onClick func(*Node)
    onClickStr string
    onKeyDown func(*Node, int)
    onKeyDownStr string
}


func CreateMemo(parent *Node, name string, x int, y int, sizeX int, sizeY int, BC int, TC int, onClick func(*Node)) *Node {
	obj := tMemo{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, BC: BC, TC: TC, list: []string{""}, visible: true, enabled: true, curX: 0, curY: 0, curXR: 0, curYR: 0, onClick: onClick, onKeyDown: nil}
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
   

	var end_row int = 0
	if len(obj.list)-1 < obj.sizeY/14 {
		end_row = len(obj.list)-1
	} else {
		end_row = obj.sizeY/14-1
	}

	var offset_num int = len(strconv.Itoa(len(obj.list)))*8+5
    //SetColor(obj.TC);
    SetBackColor(obj.BC);
    if obj.list != nil {
    	for i := 0; i <= end_row; i++ {
    		SetColor(0xCDBAB5)
    		TextOutgl(nil, strconv.Itoa(i + obj.curYR), parX+obj.x + 4 + offset_num - len(strconv.Itoa(i + obj.curYR))*8 - 5, parY+obj.y + 4 + 14*i, 1)
    		
    		if (i + obj.curYR) < len(obj.list) && obj.curXR < len(obj.list[i+obj.curYR]) {
    			var end int
    			if len(obj.list[i+obj.curYR]) > obj.curXR + 80 {
    				end = obj.curXR + 80
    			} else {
    				end = len(obj.list[i+obj.curYR])
    			}
    			
    			var offset int = 0
    			for c := obj.curXR; c < end; c++ {
    				if obj.color != nil && len(obj.color) > (i+obj.curYR) && len(obj.color[i+obj.curYR]) > c {
    					SetColor(obj.color[i+obj.curYR][c])
    				} else {
    					SetColor(obj.TC)
    				}
    				TextOutgl(nil, string(obj.list[i+obj.curYR][c]), parX+obj.x + 4 + offset + offset_num, parY+obj.y + 4 + 14*i, 1)
    				
    				switch obj.list[i+obj.curYR][c] { //obj.curYR + obj.curY    obj.curXR + i
    				case 'A': offset += 8
    				case 'B': offset += 8
    				case 'C': offset += 8
    				case 'D': offset += 8
    				case 'E': offset += 8
    				case 'F': offset += 8
    				case 'G': offset += 8
    				case 'H': offset += 8
    				case 'I': offset += 8
    				case 'J': offset += 8
    				case 'K': offset += 8
    				case 'L': offset += 8
    				case 'M': offset += 9
    				case 'N': offset += 8
    				case 'O': offset += 8
    				case 'P': offset += 8
    				case 'Q': offset += 8
    				case 'R': offset += 8
    				case 'S': offset += 8
    				case 'T': offset += 9
    				case 'U': offset += 8
    				case 'V': offset += 9
    				case 'W': offset += 9
    				case 'X': offset += 9
    				case 'Y': offset += 9
    				case 'Z': offset += 9
    		
    				case ' ': offset += 8
    				case '-': offset += 7
    				case '0': offset += 8
    				case '1': offset += 8
    				case '2': offset += 8
    				case '3': offset += 8
    				case '4': offset += 8
    				case '5': offset += 8
    				case '6': offset += 8
    				case '7': offset += 8
    				case '8': offset += 8
    				case '9': offset += 8
    				case '|': offset += 7
    		
    				case 'a': offset += 8
    				case 'b': offset += 8
    				case 'c': offset += 8
    				case 'd': offset += 8
    				case 'e': offset += 8
    				case 'f': offset += 8
    				case 'g': offset += 8
    				case 'h': offset += 8
    				case 'i': offset += 8
    				case 'j': offset += 8
    				case 'k': offset += 8
    				case 'l': offset += 8
    				case 'm': offset += 9
    				case 'n': offset += 8
    				case 'o': offset += 8
    				case 'p': offset += 8
    				case 'q': offset += 8
    				case 'r': offset += 8
    				case 's': offset += 8
    				case 't': offset += 8
    				case 'u': offset += 8
    				case 'v': offset += 9
    				case 'w': offset += 9
    				case 'x': offset += 8
    				case 'y': offset += 8
    				case 'z': offset += 8
    		
    				case '{': offset += 8
       				case '}': offset += 8
       				case '~': offset += 8
					case '!': offset += 5
					case 0x22: offset += 8
					case '#': offset += 8
					case '$': offset += 8
					case '%': offset += 8
					case '&': offset += 8
					case '\'': offset += 5
					case '(': offset += 5
					case ')': offset += 5
					case '*': offset += 9
					case '+': offset += 9
					case ',': offset += 5
        			case '.': offset += 5
        			case '/': offset += 8
        			case ':': offset += 5
        			case ';': offset += 5
        			case '<': offset += 8
        			case '=': offset += 8
        			case '>': offset += 8
        			case '?': offset += 8
        			case '@': offset += 8
        			case '[': offset += 5
					case '\\': offset += 8
					case ']': offset += 5
					case '^': offset += 8
					case '_': offset += 8
					case '`': offset += 5
					case 0x09: offset += 32 // TAB

					default: offset += 7
					}
    			}
    			//TextOutgl(nil, obj.list[i+obj.curYR][obj.curXR:end], parX+obj.x + 4, parY+obj.y + 4 + 14*i, 1)	
    		}
    	}
    }
    if obj.focused && cursor {
    	var offset int = 0
    	for i := 0; i < obj.curX; i++ {
    		if len(obj.list[obj.curYR + obj.curY]) > (obj.curXR + i) {
    		switch obj.list[obj.curYR + obj.curY][obj.curXR + i] {
    		case 'A': offset += 8
    		case 'B': offset += 8
    		case 'C': offset += 8
    		case 'D': offset += 8
    		case 'E': offset += 8
    		case 'F': offset += 8
    		case 'G': offset += 8
    		case 'H': offset += 8
    		case 'I': offset += 8
    		case 'J': offset += 8
    		case 'K': offset += 8
    		case 'L': offset += 8
    		case 'M': offset += 9
    		case 'N': offset += 8
    		case 'O': offset += 8
    		case 'P': offset += 8
    		case 'Q': offset += 8
    		case 'R': offset += 8
    		case 'S': offset += 8
    		case 'T': offset += 9
    		case 'U': offset += 8
    		case 'V': offset += 9
    		case 'W': offset += 9
    		case 'X': offset += 9
    		case 'Y': offset += 9
    		case 'Z': offset += 9
    		
    		case ' ': offset += 8
    		case '-': offset += 7
    		case '0': offset += 8
    		case '1': offset += 8
    		case '2': offset += 8
    		case '3': offset += 8
    		case '4': offset += 8
    		case '5': offset += 8
    		case '6': offset += 8
    		case '7': offset += 8
    		case '8': offset += 8
    		case '9': offset += 8
    		case '|': offset += 7
    		
    		case 'a': offset += 8
    		case 'b': offset += 8
    		case 'c': offset += 8
    		case 'd': offset += 8
    		case 'e': offset += 8
    		case 'f': offset += 8
    		case 'g': offset += 8
    		case 'h': offset += 8
    		case 'i': offset += 8
    		case 'j': offset += 8
    		case 'k': offset += 8
    		case 'l': offset += 8
    		case 'm': offset += 9
    		case 'n': offset += 8
    		case 'o': offset += 8
    		case 'p': offset += 8
    		case 'q': offset += 8
    		case 'r': offset += 8
    		case 's': offset += 8
    		case 't': offset += 8
    		case 'u': offset += 8
    		case 'v': offset += 9
    		case 'w': offset += 9
    		case 'x': offset += 8
    		case 'y': offset += 8
    		case 'z': offset += 8
    		
    		case '{': offset += 8
       		case '}': offset += 8
       		case '~': offset += 8
			case '!': offset += 5
			case 0x22: offset += 8
			case '#': offset += 8
			case '$': offset += 8
			case '%': offset += 8
			case '&': offset += 8
			case '\'': offset += 5
			case '(': offset += 5
			case ')': offset += 5
			case '*': offset += 9
			case '+': offset += 9
			case ',': offset += 5
        	case '.': offset += 5
        	case '/': offset += 8
        	case ':': offset += 5
        	case ';': offset += 5
        	case '<': offset += 8
        	case '=': offset += 8
        	case '>': offset += 8
        	case '?': offset += 8
        	case '@': offset += 8
        	case '[': offset += 5
			case '\\': offset += 8
			case ']': offset += 5
			case '^': offset += 8
			case '_': offset += 8
			case '`': offset += 5
			case 0x09: offset += 32 // TAB

			default: offset += 7
			}
			}
    	} 
    	TextOutgl(nil, "|", parX+obj.x + 4 + offset + offset_num, parY+obj.y + 4 + 14*obj.curY, 1);
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
    		
			downX = x 
    		downY = y 
    		mouseIsDown = true
    		lblPropName = CreateLabel(pnlProperties, "lblPropName", 5, 5, 95, 20, 0xD8DCC0, 0x000000, "Name", nil)
			editPropName = CreateEdit(pnlProperties, "editPropName", 80, 5, 95, 20, 0xF8FCF8, 0x000000, obj.name, nil, editPropNameEnter)
			lblPropLeft = CreateLabel(pnlProperties, "lblPropLeft", 5, 25, 95, 20, 0xD8DCC0, 0x000000, "Left", nil)
			editPropLeft = CreateEdit(pnlProperties, "editPropLeft", 80, 25, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.x), nil, editPropLeftEnter)
			lblPropTop = CreateLabel(pnlProperties, "lblPropTop", 5, 45, 95, 20, 0xD8DCC0, 0x000000, "Top", nil)
			editPropTop = CreateEdit(pnlProperties, "editPropTop", 80, 45, 95, 20, 0xF8FCF8, 0x000000, strconv.Itoa(obj.y), nil, editPropTopEnter)
			//lblPropText = CreateLabel(pnlProperties, "lblPropText", 5, 65, 95, 20, 0xD8DCC0, 0x000000, "Text", nil)
			//editPropText = CreateEdit(pnlProperties, "editPropText", 80, 65, 95, 20, 0xF8FCF8, 0x000000, obj.text, nil, editPropTextEnter)
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
			if key == 8 {	// BACKSPACE
    			if len(obj.list[obj.curYR + obj.curY]) > 0 && (obj.curXR + obj.curX) > 0 {
    				obj.list[obj.curYR + obj.curY] = obj.list[obj.curYR + obj.curY][:obj.curXR + obj.curX-1] + obj.list[obj.curYR + obj.curY][obj.curXR + obj.curX:]
    				obj.curX--
    			} else if (obj.curXR + obj.curX) == 0 && (obj.curYR + obj.curY) > 0 {
    				obj.curX = len(obj.list[obj.curYR + obj.curY-1])
    				obj.list[obj.curYR + obj.curY-1] += obj.list[obj.curYR + obj.curY]
    				copy(obj.list[obj.curYR + obj.curY:], obj.list[obj.curYR + obj.curY+1:])
    				obj.list[len(obj.list)-1] = ""
    				obj.list = obj.list[:len(obj.list)-1]
    				obj.curY--
    			}
    		} else if key == 13 {	// ENTER
    			begin := obj.list[obj.curYR + obj.curY][:obj.curXR + obj.curX]
    			end := obj.list[obj.curYR + obj.curY][obj.curXR + obj.curX:]
    			
    			obj.list = append(obj.list, "")
    			copy(obj.list[obj.curYR + obj.curY+1:], obj.list[obj.curYR + obj.curY:])
    			obj.list[obj.curYR + obj.curY] = begin
    			obj.list[obj.curYR + obj.curY+1] = end
    			
    			obj.curY++
    			obj.curX = 0
				obj.curXR = 0
			} else if key == 33 {	// PG_UP
				if obj.curYR > obj.sizeY/14-1 {
					obj.curYR -= obj.sizeY/14-1
				} else {
					obj.curYR = 0
				}
				obj.curY = 0
				if (obj.curX + obj.curXR) > len(obj.list[obj.curYR + obj.curY]) {
    				obj.curX = len(obj.list[obj.curYR + obj.curY]) - obj.curXR
    			}
			} else if key == 34 {	// PG_DOWN
				if obj.curYR + (obj.sizeY/14-1)*2 < len(obj.list) {
					obj.curYR += obj.sizeY/14-1
					obj.curY = obj.sizeY/14-1
				} else if obj.curYR == 0 && obj.curYR + (obj.sizeY/14-1) > len(obj.list) {
					obj.curY = len(obj.list) - 1
				} else {
					obj.curYR = len(obj.list) - obj.sizeY/14-1
					obj.curY = obj.sizeY/14-1
				}
				if (obj.curX + obj.curXR) > len(obj.list[obj.curYR + obj.curY]) {
    				obj.curX = len(obj.list[obj.curYR + obj.curY]) - obj.curXR
    			}
			} else if key == 35 {	// END
				if len(obj.list[obj.curYR + obj.curY]) > obj.curXR + 79 {
					obj.curXR = len(obj.list[obj.curYR + obj.curY]) - 79
					obj.curX = 79
				} else {
    				obj.curX = len(obj.list[obj.curYR + obj.curY]) - obj.curXR
    			}
    		} else if key == 36 {	// HOME
    			obj.curXR = 0
    			obj.curX = 0
    		} else if key == 37 {	// ARROW_LEFT
    			if obj.curX > 0 {
    				obj.curX--
    			} else {
    				if obj.curXR > 0 {
    					obj.curXR--
    				}
    			}
    		} else if key == 39 {	// ARROW_RIGHT
    			var end_col int = 0
				if len(obj.list[obj.curYR + obj.curY])-1 < 80 {
					end_col = len(obj.list[obj.curYR + obj.curY])
				} else {
					end_col = 80-1
				}
				
    			if obj.curX < end_col {
    				obj.curX++
    			} else {
    				if obj.curXR + obj.curX < len(obj.list[obj.curYR + obj.curY])-1 {
    					obj.curXR++
    				}
    			}
    		} else if key == 38 {	// ARROW_UP
    			if obj.curY > 0 {	
    				obj.curY--
    				if (obj.curX + obj.curXR) > len(obj.list[obj.curYR + obj.curY]) {
    					obj.curX = len(obj.list[obj.curYR + obj.curY]) - obj.curXR
    				}
    			} else {
    				if obj.curYR > 0 {
    					obj.curYR--
    					if (obj.curX + obj.curXR) > len(obj.list[obj.curYR + obj.curY]) {
    						obj.curX = len(obj.list[obj.curYR + obj.curY]) - obj.curXR
    					}
    				}
    			}
    		} else if key == 40 {	// ARROW_DOWN
    			var end_row int = 0
				if len(obj.list)-1 < obj.sizeY/14 {
					end_row = len(obj.list)-1
				} else {
					end_row = obj.sizeY/14-1
				}
	
    			if obj.curY < end_row {
    				obj.curY++
    				if (obj.curX + obj.curXR) > len(obj.list[obj.curYR + obj.curY]) {
    					obj.curX = len(obj.list[obj.curYR + obj.curY]) - obj.curXR
    				}
    			} else {
    				if obj.curYR + obj.curY < len(obj.list)-1 {
    					obj.curYR++
    					if (obj.curX + obj.curXR) > len(obj.list[obj.curYR + obj.curY]) {
    						obj.curX = len(obj.list[obj.curYR + obj.curY]) - obj.curXR
    					}
    				}
    			}
    		} else if key == 46 {	// DEL
    			if (obj.curXR + obj.curX) < len(obj.list[obj.curYR + obj.curY]) {
    				obj.list[obj.curYR + obj.curY] = obj.list[obj.curYR + obj.curY][:obj.curXR + obj.curX] + obj.list[obj.curYR + obj.curY][obj.curXR + obj.curX+1:]
    			} else if (obj.curXR + obj.curX) == len(obj.list[obj.curYR + obj.curY]) && (obj.curYR + obj.curY) < len(obj.list)-1 {
    				obj.list[obj.curYR + obj.curY] += obj.list[obj.curYR + obj.curY+1]
    				copy(obj.list[obj.curYR + obj.curY+1:], obj.list[obj.curYR + obj.curY+2:])
    				obj.list[len(obj.list)-1] = ""
    				obj.list = obj.list[:len(obj.list)-1]
    			}
			} else {	// OTHER
				var end_col int = 0
				if len(obj.list[obj.curYR + obj.curY])-1 < 80 {
					end_col = len(obj.list[obj.curYR + obj.curY])
				} else {
					end_col = 80-1
				}
				
				//if key >= 0x20 && key <= 0x7E {
					
					
	
				switch key {
				case 0x20: key = ' '
    			case 0x30: if isShiftKeyDown {key = ')'}
    			case 0x31: if isShiftKeyDown {key = '!'}
    			case 0x32: if isShiftKeyDown {key = '@'}
    			case 0x33: if isShiftKeyDown {key = '#'}
    			case 0x34: if isShiftKeyDown {key = '$'}
    			case 0x35: if isShiftKeyDown {key = '%'}
    			case 0x36: if isShiftKeyDown {key = '^'}
    			case 0x37: if isShiftKeyDown {key = '&'}
    			case 0x38: if isShiftKeyDown {key = '*'}
    			case 0x39: if isShiftKeyDown {key = '('}
    			case 220: if isShiftKeyDown {key = '|'} else {key = '\\'}
    			case 219: if isShiftKeyDown {key = 0x7B} else {key = 0x5B}
    			case 221: if isShiftKeyDown {key = 0x7D} else {key = 0x5D}
    			case 192: if isShiftKeyDown {key = 0x7E} else {key = 0x60}
    			case 222: if isShiftKeyDown {key = 0x22} else {key = 0x27}
    			case 187: if isShiftKeyDown {key = '+'} else {key = '='}
    			case 188: if isShiftKeyDown {key = '<'} else {key = ','}
    			case 190: if isShiftKeyDown {key = '>'} else {key = '.'}
    			case 191: if isShiftKeyDown {key = '?'} else {key = '/'}
    			case 186: if isShiftKeyDown {key = 0x3A} else {key = 0x3B}
    			case 189: if isShiftKeyDown {key = 0x5F} else {key = 0x2D}
				case 0x09: key = '\t' 
				case 96: key = '0'
				case 97: key = '1'
				case 98: key = '2'
				case 99: key = '3'
				case 100: key = '4'
				case 101: key = '5'
				case 102: key = '6'
				case 103: key = '7'
				case 104: key = '8'
				case 105: key = '9'
				case 106: key = '*'
				case 107: key = '+'
				case 109: key = 0x2D
				case 110: key = ','
				case 111: key = '/'
				
				default: 
					if !(isShiftKeyDown) && key >= 0x41 && key <= 0x5A {
						key += 0x20
					} else if !(key >= 0x41 && key <= 0x5A) {return}
    			}
    		
    		//case '|': 220 s
    		
   
    		
    		//case '{': 219 s
       		//case '}': 221 s
       		//case '~': 192
			//case '!': 0x31 s
			//case 0x22: 222 s
			//case '#': 51 s
			//case '$': 52 s
			//case '%': 53 s
			//case '&': 55 s
			//case '\'': 222
			//case '(': 57 s
			//case ')': 48 s
			//case '*': 56 s
			//case '+': 187 s
			//case ',': 188
        	//case '.': 190
        	//case '/': 191
        	//case ':': 186 s
        	//case ';': 186
        	//case '<': 188 s
        	//case '=': 187
        	//case '>': 190 s
        	//case '?': 191 s
        	//case '@': 50 s
        	//case '[': 219
			//case '\\': 220
			//case ']': 221
			//case '^': 54 s
			//case '_': 189 s
			//case '`': 192
			//case 0x09: 0x9 // TAB
    		
		
					
				
				
    				if obj.curX <= end_col {
    					obj.list[obj.curYR + obj.curY] = obj.list[obj.curYR + obj.curY][:obj.curXR + obj.curX] + string(key) + obj.list[obj.curYR + obj.curY][obj.curXR + obj.curX:]
    					obj.curX++
    				} else {
    					if obj.curXR + obj.curX < len(obj.list[obj.curYR + obj.curY])-1 {
    						obj.list[obj.curYR + obj.curY] = obj.list[obj.curYR + obj.curY][:obj.curXR + obj.curX] + string(key) + obj.list[obj.curYR + obj.curY][obj.curXR + obj.curX:]
    						obj.curXR++
    					}
    				}
    			//}
			}
		if obj.onKeyDown != nil && obj.enabled {
			obj.onKeyDown(nil, key)
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

