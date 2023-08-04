package main

import (

    //"syscall/js"

)


var CC int 
var BC int 

var Black int = 0x000000
var White int = 0xFFFFFF
var Red int = 0xFF0000
var Green int = 0x00FF00
var Blue int = 0x0000FF
var Yellow int = 0xFFFF00
var Magenta int = 0xFF00FF
var Cyan int = 0x00FFFF


func SetColor(Color int){
    CC = Color
}


func SetBackColor(Color int){
    BC = Color
}


func ClearDevice(){
    FillLB(0, SIZE, BC)
}







