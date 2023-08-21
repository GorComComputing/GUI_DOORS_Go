package main

import (
    //"fmt"
    //"syscall/js"
    "math/rand"
    //"strconv"

)


const BITMAP_WIDTH int = 1024 //640   
const BITMAP_HEIGHT int = 768 //480
const SIZE int = BITMAP_WIDTH*BITMAP_HEIGHT  
const GETMAX_X int = BITMAP_WIDTH - 1 
const GETMAX_Y int = BITMAP_HEIGHT - 1


const BUFFER_SIZE int = SIZE * 4;
var graphicsBuffer [BUFFER_SIZE]uint8;
//var graphicsBuffer []uint8 = make([]uint8, BUFFER_SIZE)


func FillLB(buffer []uint8, start int, count int, value int){
	
	if buffer == nil {
    	for i := start*4; i <= start*4 + count*4 - 1; i+=4 {
      		graphicsBuffer[i + 0] = uint8(255 & (value >> 16)); 	// Red
      		graphicsBuffer[i + 1] = uint8(255 & (value >> 8)); 	// Green
      		graphicsBuffer[i + 2] = uint8(255 & (value)); 		// Blue
      		graphicsBuffer[i + 3] = 255 				// Alpha
    	}
    } else {
    	for i := start*4; i <= start*4 + count*4 - 1; i+=4 {
      		buffer[i + 0] = uint8(255 & (value >> 16)); 	// Red
      		buffer[i + 1] = uint8(255 & (value >> 8)); 	// Green
      		buffer[i + 2] = uint8(255 & (value)); 		// Blue
      		buffer[i + 3] = 255 				// Alpha
    	}
    }
}


func FillLBrnd(){
    for i := 0; i < BUFFER_SIZE; i++ {
        if graphicsBuffer[i] != 1 {
        	graphicsBuffer[i] = uint8(rand.Intn(255))
        }
    }
}


func GetPixelgl(buffer []uint8, x int, y int) int {
    var val int = 0
    squareNumber := (y * BITMAP_WIDTH) + x;
    squareRgbaIndex := squareNumber * 4;
    
    val += int(graphicsBuffer[squareRgbaIndex + 0]) << 16	// Red
    val += int(graphicsBuffer[squareRgbaIndex + 1]) << 8	// Green
    val += int(graphicsBuffer[squareRgbaIndex + 2])			// Blue
    //val += int(graphicsBuffer[squareRgbaIndex + 3]) << 24 // Alpha
      	
    return val;
}


//export Draw
func Draw() {
	SetBackColor(0xFFFFFF) //0x111111
	SetColor(0xFF0000)
	SetViewPort(0, 0, GETMAX_X, GETMAX_Y)
	//ClearDevice()
	DrawNode(&layout)
	//FillCircle(nil, 0, 100, 30)
	//Circle(nil, 0, 200, 30)
	onTimer()
	
	/*SetColor(0xFF00FF)
    	var p []tPoint

    	p1 := tPoint{x: 100, y: 100}
		p = append(p, p1)
	
		p2 := tPoint{x: 400, y: 100}
		p = append(p, p2)
	
		p3 := tPoint{x: 400, y: 400}
		p = append(p, p3)
	
		p4 := tPoint{x: 100, y: 400}
		p = append(p, p4)

    	FillPoly(nil, 4, p);*/
}



// Function to return a pointer (Index) to our buffer in wasm memory
//export getGraphicsBufferPointer
func getGraphicsBufferPointer() *[BUFFER_SIZE]uint8 {
  return &graphicsBuffer
}


// Function to return the size of our buffer in wasm memory
//export getGraphicsBufferSize
func getGraphicsBufferSize() int {
  return BUFFER_SIZE;
}


