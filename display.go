package main

import (
    //"fmt"
    //"syscall/js"
    "math/rand"
    //"strconv"

)


// Объявление глобальных переменных
var (
	// Размеры экрана и параметры графики
	BITMAP_WIDTH  int = 1920 // 1920 //1600 //1024 //640 
	BITMAP_HEIGHT int = 1080 // 1080 //900 //768 //480
	SIZE          int = BITMAP_WIDTH*BITMAP_HEIGHT
	GETMAX_X      int = BITMAP_WIDTH - 1
	GETMAX_Y      int = BITMAP_HEIGHT - 1
	BUFFER_SIZE   int = SIZE * 4 *2
)

var graphicsBuffer []uint8 = make([]uint8, BUFFER_SIZE, BUFFER_SIZE)


func FillLB(buffer []uint8, start int, count int, value uint32){
	if buffer == nil {
    	for i := start*4; i <= start*4 + count*4 - 1; i+=4 {
      		graphicsBuffer[i + 0] = uint8(255 & (value >> 16)); 	// Red
      		graphicsBuffer[i + 1] = uint8(255 & (value >> 8)); 		// Green
      		graphicsBuffer[i + 2] = uint8(255 & (value)); 			// Blue
      		graphicsBuffer[i + 3] = 255; 							// Alpha
    	}
    } else {
    	for i := start*4; i <= start*4 + count*4 - 1; i+=4 {
      		buffer[i + 0] = uint8(255 & (value >> 16)); 	// Red
      		buffer[i + 1] = uint8(255 & (value >> 8)); 		// Green
      		buffer[i + 2] = uint8(255 & (value)); 			// Blue
      		buffer[i + 3] = 255;							// Alpha
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


func GetPixelgl(buffer []uint8, x int, y int) uint32 {
    var val uint32 = 0
    squareNumber := (y * BITMAP_WIDTH) + x;
    squareRgbaIndex := squareNumber * 4;
    
    val += uint32(graphicsBuffer[squareRgbaIndex + 0]) << 16	// Red
    val += uint32(graphicsBuffer[squareRgbaIndex + 1]) << 8		// Green
    val += uint32(graphicsBuffer[squareRgbaIndex + 2])			// Blue
    val += uint32(graphicsBuffer[squareRgbaIndex + 3]) << 24 	// Alpha
      	
    return val;
}


// Function to return a pointer (Index) to our buffer in wasm memory
//export getGraphicsBufferPointer
func getGraphicsBufferPointer() *uint8 {		//*[BUFFER_SIZE]uint8
	return &graphicsBuffer[0]
}


// Function to return the size of our buffer in wasm memory
//export getGraphicsBufferSize
func getGraphicsBufferSize() int {
	return BUFFER_SIZE
}



