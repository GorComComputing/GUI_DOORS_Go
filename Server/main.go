package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	//"encoding/gob"
	
    "os"
    "bufio"
)


type Handler struct {
	fileServer http.Handler
}


func main() {
	var cmd_line string
	var words = make([]string, len(os.Args)-1)
	
	// Copy from the original map of command to the ptint map of command
	for key, value := range cmd {
  		cmd_print[key] = value
	}

	// pars command line args
	if len(os.Args) > 1 {

		// TUI mode
		if os.Args[1] == "-w"{
        	fmt.Println("Window")
            os.Exit(0)
        }

        copy(words[0:], os.Args[1:])
                
        out := interpretator(words)
		if len(out) > 0 {
			fmt.Print(out)
		}
        	
		os.Exit(0)
	}
	

	// start web server
	fmt.Println("Server DOORS started OK")
	fmt.Println("Port: 8081")
	
	go http.ListenAndServe(":8081", &Handler{
		fileServer: http.FileServer(http.Dir("www")),
	})
	
	// start shell
	for {  // exit_status
		fmt.Print("DOORS> ")
		// ввод строки с пробелами
    	scanner := bufio.NewScanner(os.Stdin)
    	scanner.Scan()
    	cmd_line = scanner.Text()
    	// разбиение на подстроки по пробелу
    	words = strings.Fields(cmd_line)
    		
    	isUnion := false
    	union := ""
    	var count []int
    	var length []int
    	var Unions []string
    	for i, val := range words {
    		if (val[0] == '"' && val[len(val)-1] != '"') || (val[0] == '"' && len(val) == 1) && isUnion != true {
    			isUnion = true
    			union += val + " "
    			count = append(count, i)
    			continue
    		}
    		if val[len(val)-1] != '"' && isUnion == true {
    			union += val + " "
    			continue
    		}
    		if val[len(val)-1] == '"' {
    			isUnion = false
    			union += val
    			length = append(length, i - count[len(count)-1])
    				
    			Unions = append(Unions, union)
    			union = ""
    			continue
    		}
    	}
    		
    		
    	x := 0
    	for i, val := range count {
    		words[val+x] = Unions[i] 
    		copy(words[val+x+1:], words[val+length[i]+1:])
    		x -= length[i]
		}
		words = words[:len(words)+x]
	   		
		out := interpretator(words)
		if len(out) > 0 {
			fmt.Print(out)
		}
	}
}


// Роутер
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v", r.Method, r.URL.Path)
	
	if strings.Trim(r.URL.Path, "/") == "" {
		index_page(w, r)
		return
	}
	// API from HTTP
	if strings.Trim(r.URL.Path, "/") == "api" {
		http_pars(w, r)
		return
	}
	// API from JSON 
	if strings.Trim(r.URL.Path, "/") == "json" {
		json_pars(w, r)
		return
	}
	// Long poll
	if strings.Trim(r.URL.Path, "/") == "poll" {
		PollResponse(w, r)
		return
	}

	if strings.Trim(r.URL.Path, "/") == "savebyte" {
		savebyte(w, r)
		return
	}
	
	if strings.Trim(r.URL.Path, "/") == "saveutf8" {
		saveutf8(w, r)
		return
	}
	
	// serve static assets from 'static' dir:
	h.fileServer.ServeHTTP(w, r)
}
