package main

import (
    //"fmt"
    //"os"
)

	
type Cmd struct {
	addr 	func([]string) string
	descr string     
}

// Command list for interpretator
var cmd =  map[string]Cmd{ 
	"get_cam": Cmd{addr: cmd_get_cam, descr: "Get image fron IP cam"},
	"ls": Cmd{addr: cmd_ls, descr: "Print all file names from catalog (ls)"},
	"curl": Cmd{addr: cmd_curl, descr: "Make JSON request to the device and return struct"},
	"curls": Cmd{addr: cmd_curls, descr: "Make JSON request to the device and return string"},
	"curlj": Cmd{addr: cmd_curl_json, descr: "Make JSON request to the device and return JSON"},
	"curl_get": Cmd{addr: cmd_curl_get, descr: "Make GET request to the device and return response"},
	"read_byte": Cmd{addr: cmd_read_byte, descr: "Read file binary"},
	"read_utf8": Cmd{addr: cmd_read_utf8, descr: "Read file UTF-8"},
}


// Interpretator 
func interpretator(words []string) string {
	if _, ok := cmd[words[0]]; ok {
		return cmd[words[0]].addr(words)
	} else{
		return "Unknown command: " + words[0] + "\n"
	}
}






