package main

import (
    //"fmt"
    "os"
)

	
type Cmd struct {
	addr 	func([]string) string
	descr string     
}

// Command list for interpretator
var cmd =  map[string]Cmd{ 
	"tst": Cmd{addr: cmd_tst, descr: "Test command"},
	"ls": Cmd{addr: cmd_ls, descr: "Test command: print all file names from catalog (ls)"},

	"run": Cmd{addr: cmd_run, descr: "Test command: run process "},
	
	"curl": Cmd{addr: cmd_curl, descr: "Make JSON request to the device and return struct"},
	"curls": Cmd{addr: cmd_curls, descr: "Make JSON request to the device and return string"},
	"curlj": Cmd{addr: cmd_curl_json, descr: "Make JSON request to the device and return JSON"},
	"curl_get": Cmd{addr: cmd_curl_get, descr: "Make GET request to the device and return response"},
	
	
	".quit": Cmd{addr: cmd_quit, descr: "Exit from this program"},
	".help": Cmd{addr: cmd_help, descr: "Print this Help"},
	
	"read": Cmd{addr: cmd_read, descr: "Read file"},
}



// Interpretator 
func interpretator(words []string) string {
	if _, ok := cmd[words[0]]; ok {
		return cmd[words[0]].addr(words)
	} else{
		return "Unknown command: " + words[0] + "\n"
	}
}


// HELP - Print command list
var cmd_print = make(map[string]Cmd)
func cmd_help(words []string) string {
	var output string
	for key, val := range cmd_print {
		output += key 
		for i := len(key); i < 10; i++ {
			output += " "
		} 
		output += " - " + val.descr + "\n"
	}
	return output
}


// Exit from this program
func cmd_quit(words []string) string {
	os.Exit(0)
	return ""
}


