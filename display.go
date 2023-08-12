package main

import (
    //"fmt"
    //"syscall/js"
    "math/rand"
    //"strconv"

)


const BITMAP_WIDTH int = 640   
const BITMAP_HEIGHT int = 480
const SIZE int = BITMAP_WIDTH*BITMAP_HEIGHT  
const GETMAX_X int = BITMAP_WIDTH - 1 
const GETMAX_Y int = BITMAP_HEIGHT - 1


//var pBmp = make([]int, SIZE)


func FillLB(start int, count int, value int){
    for i := start*4; i <= start*4 + count*4 - 1; i+=4 {
      	graphicsBuffer[i + 0] = uint8(255 & (value >> 16)); 	// Red
      	graphicsBuffer[i + 1] = uint8(255 & (value >> 8)); 	// Green
      	graphicsBuffer[i + 2] = uint8(255 & (value)); 		// Blue
      	graphicsBuffer[i + 3] = 255 				// Alpha
    }
}


func FillLBrnd(){
    for i := 0; i < BUFFER_SIZE; i++ {
        if graphicsBuffer[i] != 1 {
        	graphicsBuffer[i] = uint8(rand.Intn(255))
        }
    }
}


func GetPixelgl(x int, y int) int {
    var val int = 0
    squareNumber := (y * BITMAP_WIDTH) + x;
    squareRgbaIndex := squareNumber * 4;
    
    val += int(graphicsBuffer[squareRgbaIndex + 0]) << 16	// Red
    val += int(graphicsBuffer[squareRgbaIndex + 1]) << 8	// Green
    val += int(graphicsBuffer[squareRgbaIndex + 2])			// Blue
    //val += int(graphicsBuffer[squareRgbaIndex + 3]) << 24 // Alpha
      	
    return val;
}


func Draw(){
    
    //BitBlt(dc, 0, 0, BITMAP_WIDTH, BITMAP_HEIGHT, hBmpDC, 0, 0, SRCCOPY);
}




