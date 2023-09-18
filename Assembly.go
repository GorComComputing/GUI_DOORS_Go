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
)
	
var textAsm string

var NameLen int = 31

var Lex int
var Num int
var OpCode int
var Name string
var Str string

var Ch byte
var Line int
var Pos int
var ChCount int

var PC int

var LexPos int

var Top *tObjRec

var Code = [44]int{
	cmStop,
	cmOut,
	cmOutLn,
	cmIn,
   	cmAdd,
   	cmSub,
   	cmMult,
   	cmDiv,
   	cmMod,
   	cmNeg,
   	cmDup,
   	cmDrop,
   	cmSwap,
   	cmOver,
   	cmLoad,
  	cmSave,
   	cmGOTO,
   	cmIfEQ,
   	cmIfNE,
  	cmIfLE,
   	cmIfLT,
   	cmIfGE,
   	cmIfGT,
   	cmSYSCALL,
   	cmPUSH,
} 
   
var Mnemo = [44]string{
	"HLT",
	"OUT",
	"OUTLN",
	"IN",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"NEG",
	"DUP",
	"POP",
	"SWAP",
	"OVER",
	"LOAD",
	"SAVE",
	"JMP",	// GOTO
	"JE",   // IFEQ
	"JNE", 	// IFNE
	"JLE",	// IFLE
	"JL",	// IFLT
	"JGE",	// IFGE
	"JG",	// IFGT
	"SYSCALL",
	"PUSH",
	
	"AND",
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
	"XOR",
}


func Assemble(){
	fmt.Println("Assemble...")

	Pass(LineFirst);
	Pass(LineSecond);

  	printTerminal("Assembled OK\n")
  	printTerminal("PC: " + strconv.Itoa(PC) + "\n")

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
  // ChCount++
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
	case 1: fmt.Println("Lex: OP " + strconv.Itoa(OpCode))
	case 2: fmt.Println("Lex: NUM " + strconv.Itoa(Num))
	case 3: fmt.Println("Lex: NAME "  + Name)
	case 4:	fmt.Println("Lex: STR " + Str)
	case 5: fmt.Println("Lex: EOL")
	case 6: fmt.Println("Lex: EOT")
	}
}


func LineFirst(){

	if Lex == lexLabel {
		NewName(PC)
		NextLex()
	}
	for ; Lex != lexEOL && Lex != lexEOT; {
		if Lex == lexName || Lex == lexNum || Lex == lexOpCode {
			PC++
			NextLex()
		} else if Lex == lexStr {
			PC += len(Str)
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
			Gen(Addr)
			NextLex()
		case lexNum:
			Gen(Num)
			NextLex()
		case lexOpCode:
			Gen(OpCode)
			NextLex()
		case lexStr:
			for i := 0; i < len(Str); i++ {
				Gen(int(Str[i]))
			}
			NextLex()
		}	
	}
}


type tObjRec struct{
    Name string
    Addr int
    Prev *tObjRec
}


func InitNameTable() {
	Top = nil
}


func NewName(Addr int){
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


func FindName() int {
	var Obj *tObjRec

	Obj = Top
	for ; (Obj != nil) && (Obj.Name != Name); {
		Obj = Obj.Prev
	}
	if Obj == nil {
		printTerminal("Error Name\n")
		return -1
	} else {
		Addr := Obj.Addr
		return Addr
	}
}


func Ident() {
	//var i int = 0
	Name = ""
	
	//if i < NameLen {
	//	i++
		Name += string(Ch)
	//}
	NextCh()
	
	for ; (Ch >= 0x30 && Ch <= 0x39) || (Ch >= 0x41 && Ch <= 0x5A) || (Ch >= 0x61 && Ch <= 0x7A); {
      	//if i < NameLen {
		//	i++
			Name += string(Ch)
		//}
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
	for i := 0; i < len(Mnemo); i++ {
		if NameUpper == Mnemo[i] {
			Lex = lexOpCode
			OpCode = Code[i]
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


func Gen(Cmd int) {
	fmt.Println("PC: " + strconv.Itoa(PC))
	if PC < MemSize {
		fmt.Println("Gen: " + strconv.Itoa(Cmd))
		RAM = append(RAM, Cmd)
		PC++
	} else {
		printTerminal("Error Memory\n")
	}
}



