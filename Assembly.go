package main

import (
	"strconv"
)


const (
   chSpace = byte(' ')   
   chTab   = byte(9) 
   chEOL   = byte(10) 
   chEOT   = byte(0)
   
   TabSize = 3
)
 
type tLexAsm int 
const (
	lexLabel  = iota
	lexOpCode
	lexNum
	lexName 
	lexEOL
	lexEOT
)
	

var Lex tLexAsm

var Ch byte
var Line int
var Pos int
var ChCount int

var PC int

var LexPos int


func Assemble(){

	Pass(LineFirst);
	Pass(LineSecond);

  	printTerminal("\n")
  	printTerminal("Assembled OK\n")

  	printTerminal("PC: " + strconv.Itoa(PC) + "\n")
  	printTerminal("\n")

  	btnResetVMClick(btnResetVM)
}


func Pass(Line func()) {

	ResetText();
	NextLex();
	PC = 0;
	Line();
	for Lex == lexEOL {
		NextLex();
		Line();
	}
	if Lex != lexEOT {
		//Error("Error")
		printTerminal("Error Pass\n")
	}
}


func ResetText() {
  	ChCount = 0;
  	Pos = 0;
  	Line = 1;
  	NextCh();
}


func NextCh() {
  	ChCount++

  	if len(memNotepad.obj.(*tMemo).list) == Line && len(memNotepad.obj.(*tMemo).list[Line - 1]) == ChCount{
    	Ch = chEOT 
	} else if len(memNotepad.obj.(*tMemo).list[Line - 1]) == ChCount {
    	ChCount = 0
    	Line++
    	Pos = 0
    	Ch = chEOL
    } else {
    	Ch = memNotepad.obj.(*tMemo).list[Line - 1][ChCount];
    	if Ch != chTab {
      		Pos++
      	} else {
      		for Pos++; Pos % TabSize == 0; {
      			Pos++
      		}
      	}
   }
}


func NextLex() {
	for ; Ch == ' ' || Ch == chTab; {
		NextCh()
	}
	LexPos = Pos;
	switch Ch {

	case 'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z','a','z':
		Ident();
	case '0','1','2','3','4','5','6','7','8','9':
		Number();
	case ';':
		for ; (Ch != chEOL) && (Ch != chEOT); {
				NextCh();
		}
		NextLex();
	case chEOL:
			Lex = lexEOL;
			NextCh();
	case chEOT:
		Lex = lexEOT;
	default:
		//Error('Íåäîïóñòèìûé ñèìâîë');
		printTerminal("Error NextLex\n")
	}
}


func LineFirst(){
	if Lex == lexLabel {
		NewName(PC)
		NextLex()
	}
	if Lex == lexName || Lex == lexNum || Lex == lexOpCode {
		PC++
		NextLex()
	}
}



func LineSecond(){
/*	var Addr int

	if Lex = lexLabel then
		NextLex;
	case Lex of
	lexName:
		begin
			Find(Addr);
			Gen(Addr);
			NextLex;
		end;
	lexNum:
		begin
			Gen(Num);
			NextLex;
		end;
	lexOpCode:
		begin
			Gen(OpCode);
			NextLex;
		end;
	end;*/
}


func NewName(Addr int){
/*	var Obj tObj

	Obj = Top;
	for ; (Obj != nil) && (Obj^.Name != Name); {
		Obj = Obj^.Prev;
	}
	if Obj == nil {
		New(Obj);
		Obj^.Name = Name;
		Obj^.Addr = Addr;
		Obj^.Prev = Top;
		Top = Obj;
	} else {
		//Error('Ïîâòîðíîå îáúÿâëåíèå èìåíè');
		printTerminal("Error NewName\n")
	}*/
}


func Ident() {
/*	var i int

	i := 0;
	Name := '';
	repeat
		if i < NameLen then begin
			i := i + 1;
			Name[i] := AnsiChar(Ch);
		end;
		NextCh;
	until not ( Ch in ['A'..'Z', 'a'..'z', '0'..'9'] );
	Name[0] := AnsiChar(chr(i));
	if Ch = ':' then begin
		Lex := lexLabel;
		NextCh;
		end
	else
		TestOpCode;*/
}


// ñîáèðàåì ÷èñëî
func Number(){
/*	var	d int

	Lex := lexNum;
	Num := 0;
	repeat
		d := ord(Ch)-ord('0');
		if ( Maxint - d ) div 10 >= Num then
			Num := 10*Num + d
		else
			Error('Ñëèøêîì áîëüøîå ÷èñëî');
		NextCh;
	until not ( Ch in ['0'..'9'] );*/
}
