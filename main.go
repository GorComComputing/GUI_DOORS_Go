package main

import ()


// Список программ
var programs = []*tProgram{
	{"Flag", startFlag}, 
	{"SNMP", startSNMP},
	{"Users", startUsers},
	{"Events", startEvents},
	{"Terminal", startTerminal},
	{"TMP", startTMP},
	{"Web Browser", startBrowser},
	}
	

func main() {
  	initDOORS(programs)
    <-make(chan bool)
}


















