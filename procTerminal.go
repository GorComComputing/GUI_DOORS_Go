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
	arr := strings.Split(str, string(10))
	
	memTerminal.obj.(*tMemo).list[memTerminal.obj.(*tMemo).curYR + memTerminal.obj.(*tMemo).curY] += arr[0]
	if len(arr) > 0 {
		var i int 
		for i = 0; i < len(arr)-1; i++ {
			memTerminal.obj.(*tMemo).list = append(memTerminal.obj.(*tMemo).list, "")
		}
		copy(memTerminal.obj.(*tMemo).list[memTerminal.obj.(*tMemo).curYR + memTerminal.obj.(*tMemo).curY+1:], arr[1:])
		memTerminal.obj.(*tMemo).curY += i
	}
	
	if memTerminal.obj.(*tMemo).curY > memTerminal.obj.(*tMemo).sizeY/14-1 {
		memTerminal.obj.(*tMemo).curYR += memTerminal.obj.(*tMemo).curY - memTerminal.obj.(*tMemo).sizeY/14
		memTerminal.obj.(*tMemo).curY -= memTerminal.obj.(*tMemo).curY - memTerminal.obj.(*tMemo).sizeY/14
	}
	
	//memTerminal.obj.(*tMemo).list = strings.Split(str, string(10))
}


