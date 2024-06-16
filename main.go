package main

import (
	//"fmt"
	"syscall/js"
	//"strings"
	//"unsafe"
	//"strconv"
)

// Work
//var ServerIP string = "192.168.63.60" 
//var RootDir string = "/home/gor/WORK/Go/projects/DOORS/Server/" 
//var DesktopDir string = "/home/gor/WORK/Go/projects/DOORS/Server/files/" // "/home/gor/WORK/Go/projects/" 

// Remote
//var ServerIP string = "172.18.0.1" 
//var RootDir string = "/home/user/WORK/DOORS/" 
//var DesktopDir string = "/home/user/WORK/"

// Home
//var ServerIP string = "192.168.0.10"
//var ServerIP string = "100.97.214.127"
//var ServerIP string = "85.113.55.171"
//var ServerPort string = "8081"
var ServerIP string = "www.gorcom.online"
var ServerPort string = "80"
//var RootDir string = "/home/jon/IT/WORK/Go/projects/DOORS/Server/"
//var DesktopDir string = "/home/jon/IT/WORK/Go/projects/"

// OpenWrt
//var ServerIP string = "192.168.0.254"
var RootDir string = "/Server/"
var DesktopDir string = "/Server/"

// STV
//var ServerIP string = "192.168.1.136"
//var RootDir string = "/root/DOORS/"
//var DesktopDir string = "/root/DOORS/files/"

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


















