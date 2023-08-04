package main

import (

    //"syscall/js"
    "math/rand"

)


var BITMAP_WIDTH int = 640   
var BITMAP_HEIGHT int = 480
var SIZE int = BITMAP_WIDTH*BITMAP_HEIGHT  
var GETMAX_X int = BITMAP_WIDTH - 1 
var GETMAX_Y int = BITMAP_HEIGHT - 1


var pBmp = make([]int, SIZE)


func FillLB(start int, count int, value int){
    for i := start; i <= start + count - 1; i++ {
        pBmp[i] = value
    }
}


func FillLBrnd(){
    for i := 0; i < SIZE; i++ {
        if pBmp[i] != 1 {
            pBmp[i] = rand.Intn(256)
        }
    }
}


func Draw(){
    
    //BitBlt(dc, 0, 0, BITMAP_WIDTH, BITMAP_HEIGHT, hBmpDC, 0, 0, SRCCOPY);
}




