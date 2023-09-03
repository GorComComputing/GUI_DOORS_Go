package main

import ()

var ServerIP string = "192.168.0.104"
var RootDir string = "/home/jon/IT/WORK/Go/projects/DOORS/Server/"
var DesktopDir string = "/home/jon/IT/WORK/Go/projects/"

// Список программ
var programs = []*tProgram{
	{"Explorer", startExplorer, &bmpFolder_small},
	{"Notepad", startNotepad, &bmpNotepad},
	{"Flag", startFlag, &bmpProgram}, 
	{"SNMP", startSNMP, &bmpProgram},
	{"Dispatch", startDispatch, &bmpProgram},
	{"Terminal", startTerminal, &bmpProgram},
	{"TMP", startTMP, nil},
	{"Browser", startBrowser, nil},
	}
	

func main() {
  	initDOORS(programs)
    <-make(chan bool)
}


















