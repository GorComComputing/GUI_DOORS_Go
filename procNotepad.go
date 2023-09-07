package main

import (
	//"strings"
	"fmt"
	"strconv"
)

var memNotepad *Node
var menuNotepad *Node
var menuFileNotepad *Node
var menuEditNotepad *Node


func startNotepad(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 300
    frmMain.obj.(*tForm).y = 200
    frmMain.obj.(*tForm).sizeX = 900
    frmMain.obj.(*tForm).sizeY = 800
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
    
	memNotepad = CreateMemo(frmMain, "memNotepad", 2, 18+21, 900-4, 800-17-4-21, 0xF8FCF8, 0x000000, nil)
	memNotepad.obj.(*tMemo).list = []string{"#include <stdio.h>", "", "// Main function", "int main(){", "	printf(\"Hello %d\", 0x1A);", "", "	return 0;", "}"}
	
	listNotepad := []tMenuList{{"File", nil}, {"Syntax", nil}}
	menuNotepad = CreateMenu(frmMain, "menuNotepad", 2, 18, 900-1-4, 20, 0xd8dcc0, 0x0, LINE, listNotepad, menuNotepadClick, nil)
	
	listFileNotepad := []tMenuList{{"New", bmpNew_file}, {"Open", bmpOpen_file}, {"Save", bmpSave_file}}
	menuFileNotepad = CreateMenu(frmMain, "menuFileNotepad", 2, 18+20, 100, len(listFileNotepad)*20, 0xd8dcc0, 0x0, NONE, listFileNotepad, menuFileNotepadClick, nil)
	menuFileNotepad.obj.(*tMenu).visible = false
	
	listEditNotepad := []tMenuList{{"C/C++", nil}, {"Go", nil}, {"HTML", nil}, {"CSS", nil}, {"SQL", nil}}
	menuEditNotepad = CreateMenu(frmMain, "menuEditNotepad", 2+60, 18+20, 100, len(listEditNotepad)*20, 0xd8dcc0, 0x0, NONE, listEditNotepad, menuEditNotepadClick, nil)
	menuEditNotepad.obj.(*tMenu).visible = false
}


func menuNotepadClick(node *Node, x int, y int){
	if node.obj.(*tMenu).selected == 0 {
		menuFileNotepad.obj.(*tMenu).visible = true
		menuEditNotepad.obj.(*tMenu).visible = false
	} else if node.obj.(*tMenu).selected == 1 {
		menuFileNotepad.obj.(*tMenu).visible = false
		menuEditNotepad.obj.(*tMenu).visible = true
	} else {
		menuFileNotepad.obj.(*tMenu).visible = false
		menuEditNotepad.obj.(*tMenu).visible = false
	}
}


func menuFileNotepadClick(node *Node, x int, y int){
	node.obj.(*tMenu).visible = false
	switch node.obj.(*tMenu).selected {
	case 0:
		memNotepad.obj.(*tMemo).list = nil
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
	case 1:	
		OpenDialog(RootDir, &(memNotepad.obj.(*tMemo).list))
		memNotepad.obj.(*tMemo).curX = 0
		memNotepad.obj.(*tMemo).curY = 0
		memNotepad.obj.(*tMemo).curXR = 0
		memNotepad.obj.(*tMemo).curYR = 0
	case 2:
		//SaveDialog(RootDir, &(memNotepad.obj.(*tMemo).text))
	}
}


func menuEditNotepadClick(node *Node, x int, y int){
	node.obj.(*tMenu).visible = false
	switch node.obj.(*tMenu).selected {
	case 0:
		syntax(keyWordsC1, keyWordsC2)
	case 1:
		syntax(keyWordsGo1, keyWordsGo2)
	}
}


type tLex int
const (
    NORMAL tLex = iota	
    NUM
    STR
    KEYWORD1	
    KEYWORD2
    COMMENT
)

var keyWordsC1 = []string{"break", "continue", "else", "for", "switch", "case", "goto", "sizeof", "default", "do", "while", "return", "if"}
var keyWordsC2 = []string{"double", "float", "int", "short", "unsigned", "long", "signed", "void", "char"}

var keyWordsGo1 = []string{"break", "case", "const", "continue", "default", "defer", "else", "fallthrough", "for", "func", "go", "goto", "if", "import", "interface", "package", "range", "return", "select", "struct", "switch", "type", "var"}
var keyWordsGo2 = []string{"chan", "map", "bool", "string", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune", "float32", "float64", "complex64", "complex128"}
    

func syntax(keyWords1 []string, keyWords2 []string) {
	memNotepad.obj.(*tMemo).color = make([][]int, len(memNotepad.obj.(*tMemo).list))
	for i := 0; i < len(memNotepad.obj.(*tMemo).list); i++ {
		memNotepad.obj.(*tMemo).color[i] = make([]int, len(memNotepad.obj.(*tMemo).list[i]))
	
		var c int = 0
		var r int = 0
		for ; c < len(memNotepad.obj.(*tMemo).list[i]); c++ {
			fmt.Println("bef sp " + strconv.Itoa(r))
			c = skipSpace(r, memNotepad.obj.(*tMemo).list[i])
			if c == len(memNotepad.obj.(*tMemo).list[i]) {break}
			fmt.Println("af sp " + strconv.Itoa(c))
			begin := c
			c, Lex := getLex(c, memNotepad.obj.(*tMemo).list[i], keyWords1, keyWords2)
			fmt.Println("af lex " + strconv.Itoa(c))
			fmt.Println("color " + strconv.Itoa(begin) + " " + strconv.Itoa(c) + " " + strconv.Itoa(i) + " " + strconv.Itoa(int(Lex)))
			setColorLex(begin, c, i, Lex)
			if Lex == COMMENT {break}
			r = c
			r++
			if r >= len(memNotepad.obj.(*tMemo).list[i]) {break}
		}
	}
}


func skipSpace(c int, str string) int {
	for ; c < len(str); c++ { 
		if str[c] != ' ' && str[c] != '\t' {return c}
	}
	return c
}


func getLex(c int, str string, keyWords1 []string, keyWords2 []string) (int, tLex) {
	if str[c] >= 0x30 && str[c] <= 0x39 {			// NUM
		if c < len(str) - 1 {
			c++
			for ; c < len(str); c++ {
				if !((str[c] >= 0x30 && str[c] <= 0x39) || 
			      	(str[c] >= 0x41 && str[c] <= 0x46) ||
				  	(str[c] >= 0x61 && str[c] <= 0x66) ||
				   	 str[c] == 0x58 || str[c] == 0x78) {
				return c-1, NUM
				} 
			}	
			return len(str)-1, NUM
		} else {return c, NUM}
	} else if str[c] == '/' && len(str)-1 > c && str[c+1] == '/' {	// COMMENT
		return len(str) - 1, COMMENT
	} else if str[c] == 0x22 {						// STR
		if c < len(str) - 1 {
			c++
			for ; c < len(str); c++ {
				if str[c] == 0x22 {return c, STR} 
			}
			return len(str)-1, STR	
		} else {return c, STR}
	} else if (str[c] >= 0x41 && str[c] <= 0x5A) ||
			  (str[c] >= 0x61 && str[c] <= 0x7A) || 
			   str[c] == 0x5F {
			if c < len(str) - 1 {
				begin := c
				c++
				for ; c < len(str); c++ {
					if !((str[c] >= 0x41 && str[c] <= 0x5A) ||
			  		   (str[c] >= 0x61 && str[c] <= 0x7A) ||
			  		   (str[c] >= 0x30 && str[c] <= 0x39) ||
			            str[c] == 0x5F) {
			            for _, val := range keyWords1 {
							if str[begin:c] == val {return c-1, KEYWORD1}
						}
						for _, val := range keyWords2 {
							if str[begin:c] == val {return c-1, KEYWORD2}
						}
					 	return c-1, NORMAL
					} 		
	
				}
				for _, val := range keyWords1 {
					if str[begin:c] == val {return len(str)-1, KEYWORD1}
				}
				for _, val := range keyWords2 {
					if str[begin:c] == val {return len(str)-1, KEYWORD2}
				}
				return len(str)-1, NORMAL
			}
	}
	return c, NORMAL
}


func setColorLex(begin int, c int, i int, lex tLex) {
	var color int
	switch lex {
	case NORMAL:
		color = 0x000000
	case NUM:
		color = 0x9E519D
	case KEYWORD1:
		color = 0x128421
	case KEYWORD2:
		color = 0x007EC3
	case COMMENT:
		color = 0xA95292	
	case STR:
		color = 0xF7B41A
	}
	for ; begin <= c; begin++ {
		memNotepad.obj.(*tMemo).color[i][begin] = color
	}
}









