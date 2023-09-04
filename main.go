package main

import ()

var ServerIP string = "192.168.0.104" // "172.18.0.1" // "192.168.63.60" // "192.168.0.104"
var RootDir string = "/home/jon/IT/WORK/Go/projects/DOORS" // "/home/user/WORK/DOORS/" // "/home/gor/WORK/Go/projects/DOORS/Server/" // "/home/jon/IT/WORK/Go/projects/DOORS/Server/"
var DesktopDir string = "/home/jon/IT/WORK/Go/projects/" // "/home/user/WORK/" // "/home/gor/WORK/Go/projects/" 		// "/home/jon/IT/WORK/Go/projects/"

// Список программ
var programs = []*tProgram{
	{"Explorer", startExplorer, &bmpFolder_small},
	{"Notepad", startNotepad, &bmpNotepad},
	{"Browser", startBrowser, &bmpBrowser},
	{"Flag", startFlag, &bmpProgram}, 
	{"SNMP", startSNMP, &bmpProgram},
	{"Dispatch", startDispatch, &bmpProgram},
	{"Terminal", startTerminal, &bmpProgram},
	{"TMP", startTMP, nil},
	}
	

func main() {
  	initDOORS(programs)
    <-make(chan bool)
}


















