package main

import ()

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
	{"Temp", startTMP, nil},
	}
	

func main() {
  	initDOORS(programs)
    <-make(chan bool)
}


















