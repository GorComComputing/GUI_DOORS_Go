package main

import (
	//"fmt"
	"syscall/js"
	//"strings"
	//"unsafe"
	//"strconv"
)


// Home
//var ServerIP string = "85.113.55.171"
//var ServerIP string = "localhost"
//var ServerPort string = "8081"
var ServerIP string = "www.gorcom.online"
var ServerPort string = "80"
//var RootDir string = "/home/jon/IT/WORK/Go/projects/DOORS/Server/"
//var DesktopDir string = "/home/jon/IT/WORK/Go/projects/"

// OpenWrt
//var ServerIP string = "192.168.0.254"
var RootDir string = "/Server/"
var DesktopDir string = "/Server/"

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


















