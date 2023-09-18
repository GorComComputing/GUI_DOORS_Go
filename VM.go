package main

import (
	"strconv"
	"fmt"
)

const (
   MemSize = 4096 

   cmStop   = 1

   cmAdd    = 2
   cmSub    = 3
   cmMult   = 4
   cmDiv    = 5
   cmMod    = 6
   cmNeg    = 7

   cmLoad   = 8
   cmSave  = 9

   cmDup    = 10
   cmDrop   = 11
   cmSwap   = 12
   cmOver   = 13

   cmJMP   = 14
   cmIfEQ   = 15
   cmIfNE   = 16
   cmIfLE   = 17
   cmIfLT   = 18
   cmIfGE   = 19
   cmIfGT   = 20

   cmIn     = 21
   cmOut    = 22
   cmOutLn  = 23
   
   cmSYSCALL  = 24
   
   cmPUSH = 25
   cmPUSHW = 26
   cmPUSHD = 27
   
)


var isRun bool
var IP uint32 = 0
var IR byte
var SP uint32
var FL [4]int       //CAEZ
var RegFile [4]int     //0-AX, 1-BX, 2-CX, 3-DX
var RAM []byte


func StepOVM(){
	var Buf byte
	
	IR = RAM[IP]

	switch IR {
    case cmPUSH:
    	RAM[SP-1] = RAM[IP+1]
    	SP--
    	IP++
    case cmPUSHW:
    	RAM[SP-1] = RAM[IP+1]
    	RAM[SP-2] = RAM[IP+2]
    	SP -= 2
    	IP += 2
    case cmPUSHD:
    	RAM[SP-1] = RAM[IP+1]
    	RAM[SP-2] = RAM[IP+2]
    	RAM[SP-3] = RAM[IP+3]
    	RAM[SP-4] = RAM[IP+4]
    	SP -= 4
    	IP += 4
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
    case cmJMP:
    	IP = uint32(RAM[IP+1]) + uint32(RAM[IP+2])*0xFF + uint32(RAM[IP+3])*0xFFFF + uint32(RAM[IP+4])*0xFFFFFF - 1
    case cmIfEQ:
    	if RAM[SP+2] == RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmIfNE:
    	if RAM[SP+2] != RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmIfLE:
    	if RAM[SP+2] <= RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmIfLT:
    	if RAM[SP+2] < RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmIfGE:
    	if RAM[SP+2] >= RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmIfGT:
    	if RAM[SP+2] > RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmIn:
    	SP--
        isRun = false
        lblIsRun.obj.(*tLabel).caption = "STOP"
        printTerminal("?")
    case cmOut:
        printTerminal(strconv.Itoa(int(uint32(RAM[IP+1]) + uint32(RAM[IP+2])*0xFF + uint32(RAM[IP+3])*0xFFFF + uint32(RAM[IP+4])*0xFFFFFF)))
        IP += 4
    case cmOutLn:
    	printTerminal("\n")
    case cmSYSCALL:
    	IP++
        switch RAM[IP]{
        case 0:		// execProcess
        	execProcess(int(RAM[SP]))
         	SP++
        case 1:		// CreateForm
        	var caption string = "" 
        	var i uint32
        	for i = uint32(RAM[SP+13])*0xFFFFFF + uint32(RAM[SP+14])*0xFFFF + uint32(RAM[SP+15])*0xFF + uint32(RAM[SP+16]); RAM[i] != 0; i++ {caption += string(RAM[i])}
        	//fmt.Println("Caption: " + caption)
        	//fmt.Println("i: " + strconv.Itoa(int(i)))
        	//fmt.Println(RAM)
        	var mode tMode
        	switch RAM[SP+12] {
        	case 1:
        		mode = WIN
        	}
         	CreateForm(&layout, "frm", nil, int(RAM[SP])*0xFF+int(RAM[SP+1]), int(RAM[SP+2])*0xFF+int(RAM[SP+3]), int(RAM[SP+4])*0xFF+int(RAM[SP+5]), int(RAM[SP+6])*0xFF+int(RAM[SP+7]), int(RAM[SP+8])*0xFFFFFF+int(RAM[SP+9])*0xFFFF+int(RAM[SP+10])*0xFF+int(RAM[SP+11]), mode, caption, (RAM[SP+17] != 0), nil)
         	SP += 17
         }
    default: 
         printTerminal("Error run time VM: Unrecognized command " + strconv.Itoa(int(RAM[IP])))
         fmt.Println(RAM)
         RAM[IP+1] = cmStop;
    }
    IP++	
}


func loadOVM(result string) {
	RAM = make([]byte, 0)
	for i := 0; i < len(result); i++ {
		RAM = append(RAM, byte(result[i]))
	}
	for i := 0; i < 4096; i++ {
		RAM = append(RAM, 0)
	}
	resetVM()
	fmt.Println(RAM)
}


func runVM() {
	isRun = true
	for ; isRun; {
		StepOVM()
	}
}


func resetVM() {
	isRun = false
	IP = 0
	SP = uint32(len(RAM))
}
