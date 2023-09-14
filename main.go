package main

import (
	//"syscall/js"
)

// Work
var ServerIP string = "192.168.63.60" 
var RootDir string = "/home/gor/WORK/Go/projects/DOORS/Server/" 
var DesktopDir string = "/home/gor/WORK/Go/projects/" 

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
	{"Dispatch", startDispatch, &bmpProgram},
	{"Terminal", startTerminal, &bmpProgram},
	{"VM", startVM, &bmpProgram},
	}
	

func main() {

  	initDOORS(programs)
  	
  	/*BITMAP_WIDTH := js.Global().Get("innerWidth").Float()
    	BITMAP_HEIGHT := js.Global().Get("innerHeight").Float()
    	
    	SIZE = BITMAP_WIDTH*BITMAP_HEIGHT 
	GETMAX_X = BITMAP_WIDTH - 1 
	GETMAX_Y = BITMAP_HEIGHT - 1
	BUFFER_SIZE = BITMAP_WIDTH*BITMAP_HEIGHT * 4
	

	
	
	frmDesktop.obj.(*tForm).sizeY = BITMAP_HEIGHT-2
	frmDesktop.obj.(*tForm).sizeX = BITMAP_WIDTH-1
	pnlTask.obj.(*tPanel).y = frmDesktop.obj.(*tForm).sizeY - 28
	pnlTask.obj.(*tPanel).sizeX = BITMAP_WIDTH - 1
	lblTime.obj.(*tLabel).x = pnlTask.obj.(*tPanel).sizeX - 45
	frmMenuStart.obj.(*tForm).y = BITMAP_HEIGHT-len(menuStart.obj.(*tMenu).list)*20-20-37-50+2
	
	eventDraw()*/
  	
    <-make(chan bool)
}


















