package main

import (
	"strings"
	"fmt"
	"strconv"
)

var memNotepad *Node
var menuNotepad *Node
var menuFileNotepad *Node
var menuEditNotepad *Node
var menuRunNotepad *Node

var curFileNameNotepad string = ""
var syntaxLang tLang = langNONE


func startNotepad(frmMain *Node){ 
    setSize(frmMain, 400, 400)
    frmMain.obj.(*tForm).x = BITMAP_WIDTH/2 - frmMain.obj.(*tForm).sizeX/2
	frmMain.obj.(*tForm).y = BITMAP_HEIGHT/2 - frmMain.obj.(*tForm).sizeY/2
    
	memNotepad = CreateMemo(frmMain, "memNotepad", 2, 18+21, 400-4, 400-17-4-21, 0xF8FCF8, 0x000000, nil)
	memNotepad.obj.(*tMemo).list = []string{"#include <stdio.h>", "", "// Main function", "int main(){", "	printf(\"Hello %d\", 0x1A);", "", "	return 0;", "}"}
	memNotepad.obj.(*tMemo).align = CLIENT
	memNotepad.obj.(*tMemo).onKeyDown = memNotepadKeyDown
	
	listNotepad := []tMenuList{{"File", nil}, {"Syntax", nil}, {"Run", nil}}
	menuNotepad = CreateMenu(frmMain, "menuNotepad", 2, 18, 200, 20, 0xd8dcc0, 0x0, LINE, listNotepad, menuNotepadClick, nil)
	
	listFileNotepad := []tMenuList{{"New", bmpNew_file}, {"Open", bmpOpen_file}, {"Save", bmpSave_file}}
	menuFileNotepad = CreateMenu(frmMain, "menuFileNotepad", 2, 18+20, 100, len(listFileNotepad)*20, 0xd8dcc0, 0x0, NONE, listFileNotepad, menuFileNotepadClick, nil)
	menuFileNotepad.obj.(*tMenu).visible = false
	
	listEditNotepad := []tMenuList{{"C/C++", nil}, {"Go", nil}, {"Asm", nil}, {"HTML", nil}, {"CSS", nil}, {"SQL", nil}}
	menuEditNotepad = CreateMenu(frmMain, "menuEditNotepad", 2+60, 18+20, 100, len(listEditNotepad)*20, 0xd8dcc0, 0x0, NONE, listEditNotepad, menuEditNotepadClick, nil)
	menuEditNotepad.obj.(*tMenu).visible = false
	
	listRunNotepad := []tMenuList{{"Assemble", nil}, {"Run", nil}}
	menuRunNotepad = CreateMenu(frmMain, "menuRunNotepad", 2+60+60, 18+20, 120, len(listRunNotepad)*20, 0xd8dcc0, 0x0, NONE, listRunNotepad, menuRunNotepadClick, nil)
	menuRunNotepad.obj.(*tMenu).visible = false
}


func menuNotepadClick(node *Node, x int, y int){
	if node.obj.(*tMenu).selected == 0 {
		menuFileNotepad.obj.(*tMenu).visible = true
		menuEditNotepad.obj.(*tMenu).visible = false
		menuRunNotepad.obj.(*tMenu).visible = false
	} else if node.obj.(*tMenu).selected == 1 {
		menuFileNotepad.obj.(*tMenu).visible = false
		menuEditNotepad.obj.(*tMenu).visible = true
		menuRunNotepad.obj.(*tMenu).visible = false
	} else if node.obj.(*tMenu).selected == 2 {
		menuFileNotepad.obj.(*tMenu).visible = false
		menuEditNotepad.obj.(*tMenu).visible = false
		menuRunNotepad.obj.(*tMenu).visible = true
	} else {
		menuFileNotepad.obj.(*tMenu).visible = false
		menuEditNotepad.obj.(*tMenu).visible = false
		menuRunNotepad.obj.(*tMenu).visible = false
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
		SaveDialog(RootDir, memNotepad.obj.(*tMemo).list)
	}
}


func menuEditNotepadClick(node *Node, x int, y int){
	node.obj.(*tMenu).visible = false
	switch node.obj.(*tMenu).selected {
	case 0:
		syntax(langC)
	case 1:
		syntax(langGO)
	case 2:
		syntax(langASM)
	}
}


func menuRunNotepadClick(node *Node, x int, y int){
	node.obj.(*tMenu).visible = false
	switch node.obj.(*tMenu).selected {
	case 0:
		var tmp string = memNotepad.obj.(*tMemo).list[0]
		for i := 1; i < len(memNotepad.obj.(*tMemo).list); i++ {
			tmp += string(0x0D)+string(0x0A) + memNotepad.obj.(*tMemo).list[i] 
		}
		WriteFile(curFileNameNotepad, tmp, fUTF8)
		
		RAMasm = make([]byte, 0) 
		textAsm = ReadFileUTF8(curFileNameNotepad)
		textAsm = strings.Replace(textAsm, string(0x0D)+string(0x0A), string(10), -1)
		InitNameTable()
		Assemble()
	
		fmt.Println(RAMasm)
	
		tmp = string(RAMasm)
		var by []byte
		by = []byte(tmp)
		
		fmt.Println(by)
		/*tmp = ""
		for i := uint32(0); i < PC; i++ {
			tmp += string(RAMasm[i])
		}*/
		WriteFile(curFileNameNotepad[:len(curFileNameNotepad) - 4] + ".dor", tmp, fBIN)
	case 1:
		var tmp string = memNotepad.obj.(*tMemo).list[0]
		for i := 1; i < len(memNotepad.obj.(*tMemo).list); i++ {
			tmp += string(0x0D)+string(0x0A) + memNotepad.obj.(*tMemo).list[i] 
		}
		WriteFile(curFileNameNotepad, tmp, fUTF8)

		RAMasm = make([]byte, 0) 
		textAsm = strings.Replace(tmp, string(0x0D)+string(0x0A), string(10), -1)
		InitNameTable()
		Assemble()
	
		fmt.Println(RAMasm)
	
		tmp = ""
		for i := uint32(0); i < PC; i++ {
			tmp += string(RAMasm[i])
		}
		WriteFile(curFileNameNotepad[:len(curFileNameNotepad) - 4] + ".dor", tmp, fBIN)
		
		loadOVM(RAMasm)
		runVM()
		fmt.Println("Run")
	}
}


type tLang int
const (
	langNONE tLang = iota
    langC 	
    langASM
    langGO
    langHTML	
    langCSS
    langSQL
)


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
    
var keyWordsAsm1 = []string{"hlt", "out", "outln",	"in", "add", "sub", "mul", "div", "mod", "neg", "dup", "pop",	"swap",	"over",	"load",	"save",	"jmp",	"je",	"jne",	"jle",	"jl",	"jge",	"jg",	"syscall", "push", "pushw", "pushd", "outw", "outd", "addw", "addd", "subw", "subd", "mulw", "muld", "divw", "divd", "modw", "modd", "negw", "negd", "popw", "popd", "stw", "std", "ldw", "ldd"}
var keyWordsAsm2 = []string{"db", "dw", "dd", "dq"}


func syntax(lang tLang) {
	syntaxLang = lang
	switch lang {
	case langC, langGO:
		memNotepad.obj.(*tMemo).BC = 0xF8FCF8
	case langASM:
		memNotepad.obj.(*tMemo).BC = 0x293134
	}	
	
	memNotepad.obj.(*tMemo).color = make([][]uint32, len(memNotepad.obj.(*tMemo).list))
	for i := 0; i < len(memNotepad.obj.(*tMemo).list); i++ {
		syntaxString(memNotepad.obj.(*tMemo).list[i], i, lang)
	}
}


func syntaxString(str string, i int, lang tLang) {
	var keyWords1 []string
	var keyWords2 []string
	switch lang {
	case langC:
		keyWords1 = keyWordsC1
		keyWords2 = keyWordsC2
	case langGO:
		keyWords1 = keyWordsGo1
		keyWords2 = keyWordsGo2
	case langASM:
		keyWords1 = keyWordsAsm1
		keyWords2 = keyWordsAsm2
	}
	
	memNotepad.obj.(*tMemo).color[i] = make([]uint32, len(memNotepad.obj.(*tMemo).list[i]))

	var c int = 0
	var r int = 0
	for ; c < len(str); c++ {
			fmt.Println("bef sp " + strconv.Itoa(r))
			c = skipSpace(r, str)
			if c == len(str) {break}
			fmt.Println("af sp " + strconv.Itoa(c))
			begin := c
			c, Lex := getLex(c, str, keyWords1, keyWords2, lang)
			fmt.Println("af lex " + strconv.Itoa(c))
			fmt.Println("color " + strconv.Itoa(begin) + " " + strconv.Itoa(c) + " " + strconv.Itoa(i) + " " + strconv.Itoa(int(Lex)))
			setColorLex(begin, c, i, Lex, lang)
			if Lex == COMMENT {break}
			r = c
			r++
			if r >= len(str) {break}
	}
}


func skipSpace(c int, str string) int {
	for ; c < len(str); c++ { 
		if str[c] != ' ' && str[c] != '\t' {return c}
	}
	return c
}


func getLex(c int, str string, keyWords1 []string, keyWords2 []string, lang tLang) (int, tLex) {
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
	} else if str[c] == '/' && len(str)-1 > c && str[c+1] == '/' && (lang == langC || lang == langGO) {	// COMMENT
		return len(str) - 1, COMMENT
	} else if str[c] == ';' && lang == langASM {	// COMMENT
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


func setColorLex(begin int, c int, i int, lex tLex, lang tLang) {
	var color uint32
	switch lang {
	case langC, langGO:
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
	case langASM:
		switch lex {
		case NORMAL:
			color = 0xE0E2E4
		case NUM:
			color = 0xFFCD22
		case KEYWORD1:
			color = 0x93C754
		case KEYWORD2:
			color = 0x007EC3
		case COMMENT:
			color = 0x66747B	
		case STR:
			color = 0x914925
		}
	}
	
	for ; begin <= c; begin++ {
		memNotepad.obj.(*tMemo).color[i][begin] = color
	}
}


func memNotepadKeyDown(node *Node, key int) {
	syntaxString(memNotepad.obj.(*tMemo).list[memNotepad.obj.(*tMemo).curYR + memNotepad.obj.(*tMemo).curY], memNotepad.obj.(*tMemo).curYR + memNotepad.obj.(*tMemo).curY, syntaxLang)
	if key == 13 {
		memNotepad.obj.(*tMemo).color = append(memNotepad.obj.(*tMemo).color, make([]uint32, 0))
	}
	switch key {
	case 13, 46, 8:
		for i := memNotepad.obj.(*tMemo).curYR + memNotepad.obj.(*tMemo).curY-1; i < len(memNotepad.obj.(*tMemo).list); i++ {
			syntaxString(memNotepad.obj.(*tMemo).list[i], i, syntaxLang)
		}	
	}
}









