package main

import (
	//"fmt"
	"syscall/js"
	//"strings"
	//"unsafe"
	//"strconv"
)


var ServerIP string
var ServerPort string
var ServerProtocol string

var RootDir string = "/DOORS/"
var DesktopDir string = "/DOORS/"

var mainUser string = ""


// Список программ
var programs = []*tProgram{
	{"Files", startExplorer, &bmpFolder_small},
	{"Notepad", startNotepad, &bmpNotepad},
	{"Internet", startBrowser, &bmpBrowser},
	//{"Flag", startFlag, &bmpProgram}, 
	//{"SNMP", startSNMP, &bmpProgram},
	//{"Dispatcher", startDispatch, &bmpProgram},
	{"Terminal", startTerminal, &bmpProgram},
	{"Virtual Machine", startVM, &bmpProgram},
	//{"Chrony", startChrony, &bmpProgram},
	{"COFFEE", startCOFFEE, &bmpProgram},
	{"PICO", startPICO, &bmpProgram},
	{"Camera 1", startCam1, &bmpCamera},
	{"Camera 2", startCam2, &bmpCamera},
	
	//{"ARM", startARM, &bmpProgram},
}
	

func main() {
	ServerIP = js.Global().Get("location").Get("host").String()
	ServerPort = js.Global().Get("location").Get("port").String()
	ServerProtocol = js.Global().Get("location").Get("protocol").String()
	
  	BITMAP_WIDTH = js.Global().Get("innerWidth").Int()
    BITMAP_HEIGHT = js.Global().Get("innerHeight").Int()
    	
    SIZE = BITMAP_WIDTH*BITMAP_HEIGHT 
	GETMAX_X = BITMAP_WIDTH - 1 
	GETMAX_Y = BITMAP_HEIGHT - 1
	BUFFER_SIZE = BITMAP_WIDTH*BITMAP_HEIGHT * 4
	
	initDOORS(programs)
	
	eventDraw()
  	
    <-make(chan bool)
}


















