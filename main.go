package main

import (
	//"fmt"
	"syscall/js"
	//"strings"
)

// Work
var ServerIP string = "192.168.63.60" 
var RootDir string = "/home/gor/WORK/Go/projects/DOORS/Server/" 
var DesktopDir string = "/home/gor/WORK/Go/projects/DOORS/Server/files/" // "/home/gor/WORK/Go/projects/" 

// Remote
//var ServerIP string = "172.18.0.1" 
//var RootDir string = "/home/user/WORK/DOORS/" 
//var DesktopDir string = "/home/user/WORK/"

// Home
//var ServerIP string = "192.168.0.104"
//var RootDir string = "/home/jon/IT/WORK/Go/projects/DOORS/Server/"
//var DesktopDir string = "/home/jon/IT/WORK/Go/projects/"

// Список программ
var programs = []*tProgram{
	{"Files", startExplorer, &bmpFolder_small},
	{"Notepad", startNotepad, &bmpNotepad},
	{"Internet", startBrowser, &bmpBrowser},
	{"Flag", startFlag, &bmpProgram}, 
	{"SNMP", startSNMP, &bmpProgram},
	{"Dispatcher", startDispatch, &bmpProgram},
	{"Terminal", startTerminal, &bmpProgram},
	{"Virtual Machine", startVM, &bmpProgram},
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


















