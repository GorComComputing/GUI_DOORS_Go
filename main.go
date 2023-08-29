package main

import ()


// Список программ
var programs = []*tProgram{
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


















