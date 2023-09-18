package main

import (
	"strconv"
	"fmt"
	"strings"
)


const (
   chSpace = byte(' ')   
   chTab   = byte(9) 
   chEOL   = byte(10) 
   chEOT   = byte(0)
   
   TabSize = 4
)
 
// Lex 
const (
	lexLabel  = iota
	lexOpCode
	lexNum
	lexName 
	lexStr
	lexEOL
	lexEOT
	lexDir
)
	
var textAsm string

var Lex int
var Num int
var NumLen int
var OpCode byte
var Name string
var Str string

var Ch byte
var Line int
var Pos int
var ChCount int

var PC uint32

var LexPos int

var Top *tObjRec

var RAMasm []byte

type tOpCode struct{
    Code byte
    Mnemo string
    Len int
}

var Code = [44]tOpCode{
	{cmStop, "HLT", 0},
	{cmOut, "OUT", 4},
	{cmOutLn, "OUTLN", 0},
	{cmIn, "IN", 0},
	{cmAdd, "ADD", 0},
	{cmSub, "SUB", 0},
	{cmMult, "MUL", 0},
	{cmDiv, "DIV", 0},
	{cmMod, "MOD", 0},
	{cmNeg, "NEG", 0},
	{cmDup, "DUP", 0},
	{cmDrop, "POP", 0},
	{cmSwap, "SWAP", 0},
	{cmOver, "OVER", 0},
	{cmLoad, "LOAD", 0},
	{cmSave, "SAVE", 0},
	{cmJMP, "JMP", 4},
	{cmIfEQ, "JE", 4},
	{cmIfNE, "JNE", 4},
	{cmIfLE, "JLE", 4},
	{cmIfLT, "JL", 4},
	{cmIfGE, "JGE", 4},
	{cmIfGT, "JG", 4},
	{cmSYSCALL, "SYSCALL", 1},
	{cmPUSH, "PUSH", 1},
	{cmPUSHW, "PUSHW", 2},
	{cmPUSHD, "PUSHD", 4},
}

	
	/*"AND",
	"CALL",
	"CMP",
	"DEC",
	"INC",
	"NOT",
	"OR",
	"RET",
	"ROL",
	"ROR",
	"SHL",
	"SHR",
	"XOR",*/
	
type tDir struct{
    Name string
    Len int
}

var Dir = [4]tDir{
	{"DB", 1},
	{"DW", 2},
	{"DD", 4},
	{"DQ", 8},
}


func Assemble(){
	fmt.Println("Assemble...")

	Pass(LineFirst);
	Pass(LineSecond);

  	printTerminal("Assembled OK\n")
  	printTerminal("PC: " + strconv.Itoa(int(PC)) + "\n")

  	btnResetVMClick(btnResetVM)
}


func Pass(Line func()) {
	fmt.Println("Pass")
	
	ResetText()
	NextLex()
	PC = 0
	Line()
	for ; Lex == lexEOL; {
		NextLex()
		Line()
	}
	if Lex != lexEOT {
		printTerminal("Error Pass\n")
	}
	fmt.Println(" ")
}


func ResetText() {
  	ChCount = -1
  	Pos = 0;
  	Line = 1;
  	NextCh();
  	fmt.Println("Ch: " + string(Ch))
}


func NextCh() {
  	ChCount++
	if len(textAsm) >= ChCount {
  	if len(textAsm)-1 == ChCount{
    	Ch = chEOT 
	} else if textAsm[ChCount] == 10 {
    	ChCount++
    	Line++
    	Pos = 0
    	Ch = chEOL
    } else {
    	Ch = textAsm[ChCount]
    	if Ch != chTab {
      		Pos++
      	} else {
      		for Pos++; Pos % TabSize == 0; {
      			Pos++
      		}
      	}
   }
  }
}


func NextLex() {
	for ; Ch == ' ' || Ch == chTab || Ch == ','; {
		NextCh()
	}
	LexPos = Pos
	switch Ch {
	case 'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z','a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z':
		Ident()
	case '0','1','2','3','4','5','6','7','8','9':
		Number()
	case ';':
		for ; (Ch != chEOL) && (Ch != chEOT); {
				NextCh();
		}
		NextLex()
	case '"':
		String()
	case chEOL:
			Lex = lexEOL
			NextCh()
	case chEOT:
		Lex = lexEOT
	default:
		printTerminal("Error NextLex\n")
	}
	
	switch Lex {
	case 0: fmt.Println("Lex: LAB " + Name)
	case 1: fmt.Println("Lex: OP " + strconv.Itoa(int(OpCode)))
	case 2: fmt.Println("Lex: NUM " + strconv.Itoa(Num))
	case 3: fmt.Println("Lex: NAME "  + Name)
	case 4:	fmt.Println("Lex: STR " + Str)
	case 5: fmt.Println("Lex: EOL")
	case 6: fmt.Println("Lex: EOT")
	case 7: fmt.Println("Lex: DIR")
	}
}


func LineFirst(){

	if Lex == lexLabel {
		NewName(PC)
		NextLex()
	}
	for ; Lex != lexEOL && Lex != lexEOT; {
		if Lex == lexOpCode {
			PC++
			NextLex()
		} else if Lex == lexStr {
			PC += uint32(len(Str))
			NextLex()
		} else if Lex == lexName {
			PC += 4
			NextLex()
		} else if Lex == lexNum {
			switch NumLen {
			case 1:
				PC++
			case 2:
				PC += 2
			case 4:
				PC += 4
			}
			NextLex()
		} else if Lex == lexDir {
			NextLex()
		}
	}
}


func LineSecond(){
	if Lex == lexLabel {
		NextLex()
	}
	for ; Lex != lexEOL && Lex != lexEOT; {
		switch Lex {
		case lexName:
			Addr :=	FindName()
			Gen(byte(Addr))
			Gen(byte(Addr/0xFF))
			Gen(byte(Addr/0xFFFF))
			Gen(byte(Addr/0xFFFFFF))
			NextLex()
		case lexNum:
			switch NumLen {
			case 1:
				Gen(byte(Num))
			case 2:
				Gen(byte(Num))
				Gen(byte(Num/0xFF))
			case 4:
				Gen(byte(Num))
				Gen(byte(Num/0xFF))
				Gen(byte(Num/0xFFFF))
				Gen(byte(Num/0xFFFFFF))
			}
			NextLex()
		case lexOpCode:
			Gen(OpCode)
			NextLex()
		case lexStr:
			for i := 0; i < len(Str); i++ {
				Gen(Str[i])
			}
			NextLex()
		case lexDir:
			NextLex()
		}	
	}
}


type tObjRec struct{
    Name string
    Addr uint32
    Prev *tObjRec
}


func InitNameTable() {
	Top = nil
}


func NewName(Addr uint32){
	var Obj *tObjRec

	Obj = Top
	for ; (Obj != nil) && (Obj.Name != Name); {
		Obj = Obj.Prev
	}
	if Obj == nil {
		Obj = &tObjRec{Name: Name, Addr: Addr, Prev: Top}
		Top = Obj
	} else {
		printTerminal("Error NewName\n")
	}
}


func FindName() uint32 {
	var Obj *tObjRec

	Obj = Top
	for ; (Obj != nil) && (Obj.Name != Name); {
		Obj = Obj.Prev
	}
	if Obj == nil {
		printTerminal("Error Name\n")
		return 0xFFFFFFFF
	} else {
		Addr := Obj.Addr
		return Addr
	}
}


func Ident() {
	Name = ""
	Name += string(Ch)
	NextCh()
	
	for ; (Ch >= 0x30 && Ch <= 0x39) || (Ch >= 0x41 && Ch <= 0x5A) || (Ch >= 0x61 && Ch <= 0x7A); {
		Name += string(Ch)
		NextCh()
    }
	
	if Ch == ':' {
		Lex = lexLabel
		NextCh()
	} else {
		TestOpCode()
	}
}


func TestOpCode(){
	NameUpper := strings.ToUpper(Name)
	for i := 0; i < len(Code); i++ {
		if i < len(Code) && NameUpper == Code[i].Mnemo {
			Lex = lexOpCode
			OpCode = Code[i].Code
			NumLen = Code[i].Len
			return
		} else if i < len(Dir) && NameUpper == Dir[i].Name {
			Lex = lexDir
			NumLen = Dir[i].Len
			return
		}
	}
	Lex = lexName
}


func Number(){
	Lex = lexNum
	Num = 0
	for ;(Ch >= 0x30 && Ch <= 0x39); {
		//if ( Maxint - d ) div 10 >= Num {
			Num = 10*Num + int(Ch) - 0x30
		//} else {
		//	printTerminal("Error Number\n")
		//}
		NextCh()
	}
}


func String() {
	Lex = lexStr
	Str = ""
	NextCh()
	for ; (Ch != chEOL) && (Ch != chEOT) && (Ch != 0x22); {
		Str += string(Ch)
		NextCh()
	}
	if Ch == 0x22 {
		NextCh()
	}
}


func Gen(Cmd byte) {
	//fmt.Println("PC: " + strconv.Itoa(PC))
	if PC < MemSize {
		//fmt.Println("Gen: " + strconv.Itoa(Cmd))
		RAMasm = append(RAMasm, Cmd)
		PC++
	} else {
		printTerminal("Error Memory\n")
	}
}



