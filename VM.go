package main

import (
	"strconv"
)

const (
   MemSize = 256; // 8*1024; 

   cmStop   = -1;

   cmAdd    = -2;
   cmSub    = -3;
   cmMult   = -4;
   cmDiv    = -5;
   cmMod    = -6;
   cmNeg    = -7;

   cmLoad   = -8;
   cmSave  = -9;

   cmDup    = -10;
   cmDrop   = -11;
   cmSwap   = -12;
   cmOver   = -13;

   cmGOTO   = -14;
   cmIfEQ   = -15;
   cmIfNE   = -16;
   cmIfLE   = -17;
   cmIfLT   = -18;
   cmIfGE   = -19;
   cmIfGT   = -20;

   cmIn     = -21;
   cmOut    = -22;
   cmOutLn  = -23;
   
   cmSYSCALL  = -24;
   
   cmPUSH = -25
   
)


var isRun bool
var IP int = 0
var IR int
var SP int
var FL [4]int       //CAEZ
var RegFile [4]int     //0-AX, 1-BX, 2-CX, 3-DX
var RAM []int


func StepOVM(){
	var Buf int
	
	IR = RAM[IP]

	switch IR {
    case cmPUSH:
    	SP--
    	IP++
    	RAM[SP] = RAM[IP]
    case cmStop:
    	isRun = false;
    	lblIsRun.obj.(*tLabel).caption = "STOP"
    case cmAdd:
    	SP++
    	RAM[SP] = RAM[SP] + RAM[SP-1];
    case cmSub:
    	SP++
    	RAM[SP] = RAM[SP] - RAM[SP-1];
    case cmMult:
    	SP++
    	RAM[SP] = RAM[SP]*RAM[SP-1];
    case cmDiv:
    	SP++
    	RAM[SP] = RAM[SP] / RAM[SP-1];
    case cmMod:
    	SP++
    	RAM[SP] = RAM[SP] % RAM[SP-1];
    case cmNeg:
    	RAM[SP] = -RAM[SP];
    case cmLoad:
    	RAM[SP] = RAM[RAM[SP]];
    case cmSave:
    	RAM[RAM[SP+1]] = RAM[SP];
    	SP += 2
    case cmDup:
    	SP--
    	RAM[SP] = RAM[SP+1];
    case cmDrop:
    	SP++
    case cmSwap:
    	Buf = RAM[SP];
    	RAM[SP] = RAM[SP+1];
    	RAM[SP+1] = Buf;
    case cmOver:
    	SP--
    	RAM[SP] = RAM[SP+2];
    case cmGOTO:
    	IP = RAM[SP];
    	SP++
    case cmIfEQ:
    	if RAM[SP+2] == RAM[SP+1] {
        	IP = RAM[SP];
        }
        SP += 3
    case cmIfNE:
    	if RAM[SP+2] != RAM[SP+1] {
        	IP = RAM[SP];
        }
        SP += 3
    case cmIfLE:
    	if RAM[SP+2] <= RAM[SP+1] {
        	IP = RAM[SP];
        }
        SP += 3
    case cmIfLT:
    	if RAM[SP+2] < RAM[SP+1] {
        	IP = RAM[SP];
        }
        SP += 3
    case cmIfGE:
    	if RAM[SP+2] >= RAM[SP+1] {
        	IP = RAM[SP];
        }
        SP += 3
    case cmIfGT:
    	if RAM[SP+2] > RAM[SP+1] {
        	IP = RAM[SP];
        }
        SP += 3
    case cmIn:
    	SP--
        isRun = false
        lblIsRun.obj.(*tLabel).caption = "STOP"
        printTerminal("?")
    case cmOut:
    	IP++
        printTerminal(strconv.Itoa(RAM[IP]))
    case cmOutLn:
    	printTerminal("\n")
    case cmSYSCALL:
    	IP++
        switch RAM[IP]{
        case 0:
        	execProcess(RAM[SP])
         	SP++
        case 1:
        	var caption string = "" 
        	for i := RAM[SP+6]; RAM[i] != 0; i++ {caption += string(RAM[i])}
        	var mode tMode
        	switch RAM[SP+5] {
        	case 1:
        		mode = WIN
        	}
         	CreateForm(&layout, "frm", nil, RAM[SP], RAM[SP+1], RAM[SP+2], RAM[SP+3], RAM[SP+4], mode, caption, (RAM[SP+7] != 0), nil)
         	SP += 7
         }
    default: 
         printTerminal("Error run time VM: Unrecognized command")
         RAM[IP+1] = cmStop;
    }
    IP++	
}
