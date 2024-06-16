package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    "strconv"

)


type tCanvas struct{
	name string
    x int
    y int
    sizeX int
    sizeY int
    visible bool
    buffer []uint8
    align tAlign
    onClick func(*Node)
    onClickStr string
}


func CreateCanvas(parent *Node, name string, x int, y int, sizeX int, sizeY int, onClick func(*Node)) (*Node) {
	buffer := make([]uint8, sizeX * sizeY * 4 + sizeX * sizeY)
	obj := tCanvas{name: name, x: x, y: y, sizeX: sizeX, sizeY: sizeY, visible: true, buffer: buffer, onClick: onClick}
	node := Node{typ: CANVAS, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj *tCanvas) Draw(parX int, parY int, parSizeX int, parSizeY int){
	//SetViewPort(parX, parY, parX + parSizeX, parY + parSizeY)
	i := (((parY+obj.y+1) * BITMAP_WIDTH) + (parX+obj.x))*4

	for j := 0; j < obj.sizeX*obj.sizeY*4; j+=4 {
		if j >= obj.sizeX*4 && j%(obj.sizeX*4) == 0 {
			i += (BITMAP_WIDTH-obj.sizeX)*4
		}
		if obj.buffer[j + 3] == 0xFF {
      	graphicsBuffer[i + 0] = obj.buffer[j + 0] 	// Red
      	graphicsBuffer[i + 1] = obj.buffer[j + 1] 	// Green
      	graphicsBuffer[i + 2] = obj.buffer[j + 2]   // Blue
      	graphicsBuffer[i + 3] = 255 				// Alpha
      	}
      	i+=4
    }
    
    //SetViewPort(0, 0, GETMAX_X, GETMAX_Y)
}


func (obj *tCanvas) RAD(x int, y int){

}


func (obj *tCanvas) KeyDown(key int){

}


func (obj *tCanvas) Click(x int, y int){

}


func (obj *tCanvas) MouseMove(x int, y int, Xl int, Yl int){
	if RAD && layout.children[len(layout.children)-1] != frmProperties && layout.children[len(layout.children)-1] != frmRAD && layout.children[len(layout.children)-1] != frmCode && mouseIsDown && layout.children[len(layout.children)-1].obj.(*tForm).mode != DIALOG {
			obj.x += x - downX
    		obj.y += y - downY
    		editPropLeft.obj.(*tEdit).text = strconv.Itoa(obj.x)
			editPropTop.obj.(*tEdit).text = strconv.Itoa(obj.y)
    	}
}


func (obj *tCanvas) MouseDown(x int, y int){

}



