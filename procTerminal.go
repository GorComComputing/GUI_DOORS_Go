package main

import (
    //"fmt"
    //"math/rand"
    //"math"
    //"syscall/js"
    //"time"
    //"strconv"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    //"encoding/json"
    "strings"
)


var memTerminal *Node


func startTerminal(frmMain *Node){
	frmMain.obj.(*tForm).x = 200
	frmMain.obj.(*tForm).y = 700
	frmMain.obj.(*tForm).sizeX = 7*80+2*4
	frmMain.obj.(*tForm).sizeY = (7+2)*24 + 17 + 7
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	
	memTerminal = CreateMemo(frmMain, "memTerminal", 2, 18, 7*80+4, 24*9+4, 0x000000, 0xF8FCF8, nil)
}


func printTerminal(str string) {
	memTerminal.obj.(*tMemo).list = strings.Split(str, string(10))
}


