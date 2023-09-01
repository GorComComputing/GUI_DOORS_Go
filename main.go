package main

import ()


// Список программ
var programs = []*tProgram{
	{"Explorer", startExplorer},
	{"Notepad", startNotepad},
	{"Flag", startFlag}, 
	{"SNMP", startSNMP},
	{"Dispatch", startDispatch},
	{"Terminal", startTerminal},
	{"TMP", startTMP},
	{"Browser", startBrowser},
	}
	

func main() {
  	initDOORS(programs)
    <-make(chan bool)
}


















