package main

import (
    //"fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type tCanvas struct{
    x int
    y int
    sizeX int
    sizeY int
    visible bool
    buffer []uint8
    onClick func(*Node)
}


func CreateCanvas(parent *Node, x int, y int, sizeX int, sizeY int, onClick func(*Node)) (*Node) {
	buffer := make([]uint8, sizeX * sizeY * 4)
	obj := tCanvas{x: x, y: y, sizeX: sizeX, sizeY: sizeY, visible: true, buffer: buffer, onClick: onClick}
	node := Node{typ: CANVAS, parent: parent, previous: nil, children: nil, obj: &obj}
	parent.children = append(parent.children, &node)
	return &node
}


func (obj tCanvas) Draw(parX int, parY int){
	i := (((parY+obj.y+1) * BITMAP_WIDTH) + (parX+obj.x))*4

	for j := 0; j < obj.sizeX*obj.sizeY*4; j+=4 {
		if j >= obj.sizeX*4 && j%(obj.sizeX*4) == 0 {
			i += (BITMAP_WIDTH-obj.sizeX)*4
		}
      	graphicsBuffer[i + 0] = obj.buffer[j + 0] 	// Red
      	graphicsBuffer[i + 1] = obj.buffer[j + 1] 	// Green
      	graphicsBuffer[i + 2] = obj.buffer[j + 2]   // Blue
      	graphicsBuffer[i + 3] = 255 				// Alpha
      	i+=4
    }
}



