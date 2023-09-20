package main

import (
	"strconv"
	"fmt"
	//"unsafe"
)

var frmVM *Node

const (
   MemSize = 4096 

   cmHLT   = 1

   cmADD    = 2
   cmSUB    = 3
   cmMUL   = 4
   cmDIV    = 5
   cmMOD    = 6
   cmNEG    = 7

   cmLD   = 8
   cmST  = 9

   cmDUP    = 10
   cmPOP   = 11
   cmSWAP   = 12
   cmOVER   = 13

   cmJMP   = 14
   cmJE   = 15
   cmJNE   = 16
   cmJLE   = 17
   cmJL   = 18
   cmJGE   = 19
   cmJG   = 20

   cmIN     = 21
   cmOUT    = 22
   cmOUTLN  = 23
   
   cmSYSCALL  = 24
   
   cmPUSH = 25
   cmPUSHW = 26
   cmPUSHD = 27
   cmOUTW  = 28
   cmOUTD  = 29
   cmADDW  = 30
   cmADDD  = 31
   cmSUBW  = 32
   cmSUBD  = 33
   cmMULW  = 34
   cmMULD  = 35
   cmDIVW  = 36
   cmDIVD  = 37
   cmMODW  = 38
   cmMODD  = 39
   cmNEGW  = 40
   cmNEGD  = 41
   cmPOPW  = 42
   cmPOPD  = 43
   cmSTW  = 44
   cmSTD  = 45
   cmLDW  = 46
   cmLDD  = 47
   
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
    case cmHLT:
    	isRun = false;
    case cmADD:
    	RAM[SP+1] = RAM[SP] + RAM[SP+1]
    	SP++
  	case cmADDW:
    	//RAM[SP+1] = RAM[SP] + RAM[SP+1]
    	//SP--
    case cmADDD:
    	//RAM[SP+1] = RAM[SP] + RAM[SP+1]
    	//SP--
    case cmSUB:
    	RAM[SP+1] = RAM[SP+1] - RAM[SP]
    	SP++
    case cmMUL:
    	RAM[SP+1] = RAM[SP+1] * RAM[SP]
    	SP++
    case cmDIV:
    	RAM[SP+1] = RAM[SP+1] / RAM[SP]
    	SP++
    case cmMOD:
    	RAM[SP+1] = RAM[SP+1] % RAM[SP]
    	SP++
    case cmNEG:
    	RAM[SP] = -RAM[SP]
    case cmLD:
    	RAM[SP+3] = RAM[int((uint32(RAM[SP]) << 24) + (uint32(RAM[SP+1]) << 16) + (uint32(RAM[SP+2]) << 8) + uint32(RAM[SP+3]))]
    	SP += 3
    case cmST:
    	RAM[int((uint32(RAM[SP+1]) << 24) + (uint32(RAM[SP+2]) << 16) + (uint32(RAM[SP+3]) << 8) + uint32(RAM[SP+4]))] = RAM[SP]
    	SP += 4
    case cmDUP:
    	SP--
    	RAM[SP] = RAM[SP+1];
    case cmPOP:
    	SP++
    case cmPOPW:
    	SP += 2
    case cmPOPD:
    	SP += 4
    case cmSWAP:
    	Buf = RAM[SP];
    	RAM[SP] = RAM[SP+1];
    	RAM[SP+1] = Buf;
    case cmOVER:
    	SP--
    	RAM[SP] = RAM[SP+2];
    case cmJMP:
    	IP = uint32(RAM[IP+1]) + uint32(RAM[IP+2])*0xFF + uint32(RAM[IP+3])*0xFFFF + uint32(RAM[IP+4])*0xFFFFFF - 1
    case cmJE:
    	if RAM[SP+2] == RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmJNE:
    	if RAM[SP+2] != RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmJLE:
    	if RAM[SP+2] <= RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmJL:
    	if RAM[SP+2] < RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmJGE:
    	if RAM[SP+2] >= RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmJG:
    	if RAM[SP+2] > RAM[SP+1] {
        	IP = uint32(RAM[SP])
        }
        SP += 3
    case cmIN:
    	SP--
        isRun = false
        lblIsRun.obj.(*tLabel).caption = "STOP"
        printTerminal("?")
    case cmOUT:
        printTerminal(strconv.Itoa(int(uint32(RAM[SP]))))
        SP++
    case cmOUTW:
        printTerminal(strconv.Itoa(int((uint32(RAM[SP]) << 8) + uint32(RAM[SP+1]))))
        SP += 2
    case cmOUTD:
        printTerminal(strconv.Itoa(int((uint32(RAM[SP]) << 24) + (uint32(RAM[SP+1]) << 16) + (uint32(RAM[SP+2]) << 8) + uint32(RAM[SP+3]))))
        SP += 4
    case cmOUTLN:
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
        	for i = (uint32(RAM[SP+13]) << 24) + (uint32(RAM[SP+14]) << 16) + (uint32(RAM[SP+15]) << 8) + uint32(RAM[SP+16]); RAM[i] != 0; i++ {caption += string(RAM[i])}
        	var mode tMode
        	switch RAM[SP+12] {
        	case 1:
        		mode = WIN
        	}
         	frmVM = CreateForm(&layout, "frm", nil, int(RAM[SP])*0xFF+int(RAM[SP+1]), int(RAM[SP+2])*0xFF+int(RAM[SP+3]), int(RAM[SP+4])*0xFF+int(RAM[SP+5]), int(RAM[SP+6])*0xFF+int(RAM[SP+7]), int((uint32(RAM[SP+8]) << 24)+(uint32(RAM[SP+9]) << 16)+(uint32(RAM[SP+10]) << 8)+uint32(RAM[SP+11])), mode, caption, (RAM[SP+17] != 0), nil)
         	

         	/*RAM[SP+17] = byte(uint32(unsafe.Pointer(frmVM)))
			RAM[SP+16] = byte(uint32(unsafe.Pointer(frmVM)) >> 8)
			RAM[SP+15] = byte(uint32(unsafe.Pointer(frmVM)) >> 16)
			RAM[SP+14] = byte(uint32(unsafe.Pointer(frmVM)) >> 24)*/
			
		//	fmt.Println("Pointer:")
		//	s := fmt.Sprintf("%d", unsafe.Pointer(frmVM))
		//	fmt.Println(s)
		//	fmt.Println(frmVM)
		//	a,_ := strconv.Atoi(s)
			
		//	frmVM = (*Node)(unsafe.Pointer(a))
			//frmVM = uintptr(a)
			//frmVM = (*Node)(unsafe.Pointer(frmVM))
			//out := (*two)(unsafe.Pointer(&in))
			
			//uintptr(unsafe.Pointer(frmVM))
			
			/*var tmp [4]byte
    		f := 123 
    		binary.LittleEndian.PutUint32(tmp[:], math.Int32bits(f))
    		i := int32(binary.LittleEndian.Uint32(tmp[:]))
    		fmt.Println(f, i)*/
				
         	SP += 17 //13
         case 2:		// CreateLabel
         	//fmt.Println("SYSCALL LABEL:")
         	var caption string = "" 
        	var i uint32
        	
        	/*fmt.Println(int(RAM[SP])*0xFF+int(RAM[SP+1]))
        	fmt.Println(int(RAM[SP+2])*0xFF+int(RAM[SP+3]))
        	fmt.Println(int(RAM[SP+4])*0xFF+int(RAM[SP+5]))
        	fmt.Println(int(RAM[SP+6])*0xFF+int(RAM[SP+7]))
        	fmt.Println(int(RAM[SP+8])*0xFFFFFF+int(RAM[SP+9])*0xFFFF+int(RAM[SP+10])*0xFF+int(RAM[SP+11]))
        	fmt.Println(int(RAM[SP+12])*0xFFFFFF+int(RAM[SP+13])*0xFFFF+int(RAM[SP+14])*0xFF+int(RAM[SP+15]))
        	fmt.Println((uint32(RAM[SP+16]) << 24) + (uint32(RAM[SP+17]) << 16) + (uint32(RAM[SP+18]) << 8) + uint32(RAM[SP+19]))*/
        	for i = (uint32(RAM[SP+16]) << 24) + (uint32(RAM[SP+17]) << 16) + (uint32(RAM[SP+18]) << 8) + uint32(RAM[SP+19]); RAM[i] != 0; i++ {caption += string(RAM[i])}
        	
        	//fmt.Println(caption)
        	//fmt.Println(RAM)
        	
        	CreateLabel(frmVM, "lbl", int(RAM[SP])*0xFF+int(RAM[SP+1]), int(RAM[SP+2])*0xFF+int(RAM[SP+3]), int(RAM[SP+4])*0xFF+int(RAM[SP+5]), int(RAM[SP+6])*0xFF+int(RAM[SP+7]), int((uint32(RAM[SP+8]) << 24)+(uint32(RAM[SP+9]) << 16)+(uint32(RAM[SP+10]) << 8)+uint32(RAM[SP+11])), int((uint32(RAM[SP+12]) << 24)+(uint32(RAM[SP+13]) << 16)+(uint32(RAM[SP+14]) << 8)+uint32(RAM[SP+15])), caption, nil)
         	SP += 19
         case 3:		// CreateBtn
         	var caption string = "" 
        	var i uint32
        	for i = (uint32(RAM[SP+16]) << 24) + (uint32(RAM[SP+17]) << 16) + (uint32(RAM[SP+18]) << 8) + uint32(RAM[SP+19]); RAM[i] != 0; i++ {caption += string(RAM[i])}
        	CreateBtn(frmVM, "btn", int(RAM[SP])*0xFF+int(RAM[SP+1]), int(RAM[SP+2])*0xFF+int(RAM[SP+3]), int(RAM[SP+4])*0xFF+int(RAM[SP+5]), int(RAM[SP+6])*0xFF+int(RAM[SP+7]), int((uint32(RAM[SP+8]) << 24)+(uint32(RAM[SP+9]) << 16)+(uint32(RAM[SP+10]) << 8)+uint32(RAM[SP+11])), int((uint32(RAM[SP+12]) << 24)+(uint32(RAM[SP+13]) << 16)+(uint32(RAM[SP+14]) << 8)+uint32(RAM[SP+15])), caption, nil)
         	SP += 19
         case 4:		// CreateEdit
         	var text string = "" 
        	var i uint32
        	for i = (uint32(RAM[SP+16]) << 24) + (uint32(RAM[SP+17]) << 16) + (uint32(RAM[SP+18]) << 8) + uint32(RAM[SP+19]); RAM[i] != 0; i++ {text += string(RAM[i])}
        	CreateEdit(frmVM, "edt", int(RAM[SP])*0xFF+int(RAM[SP+1]), int(RAM[SP+2])*0xFF+int(RAM[SP+3]), int(RAM[SP+4])*0xFF+int(RAM[SP+5]), int(RAM[SP+6])*0xFF+int(RAM[SP+7]), int((uint32(RAM[SP+8]) << 24)+(uint32(RAM[SP+9]) << 16)+(uint32(RAM[SP+10]) << 8)+uint32(RAM[SP+11])), int((uint32(RAM[SP+12]) << 24)+(uint32(RAM[SP+13]) << 16)+(uint32(RAM[SP+14]) << 8)+uint32(RAM[SP+15])), text, nil, nil)
         	SP += 19
         case 5:		// CreateCheckBox
         	var caption string = "" 
        	var i uint32
        	for i = (uint32(RAM[SP+16]) << 24) + (uint32(RAM[SP+17]) << 16) + (uint32(RAM[SP+18]) << 8) + uint32(RAM[SP+19]); RAM[i] != 0; i++ {caption += string(RAM[i])}
        	CreateCheckBox(frmVM, "cbx", int(RAM[SP])*0xFF+int(RAM[SP+1]), int(RAM[SP+2])*0xFF+int(RAM[SP+3]), int(RAM[SP+4])*0xFF+int(RAM[SP+5]), int(RAM[SP+6])*0xFF+int(RAM[SP+7]), int((uint32(RAM[SP+8]) << 24)+(uint32(RAM[SP+9]) << 16)+(uint32(RAM[SP+10]) << 8)+uint32(RAM[SP+11])), int((uint32(RAM[SP+12]) << 24)+(uint32(RAM[SP+13]) << 16)+(uint32(RAM[SP+14]) << 8)+uint32(RAM[SP+15])), caption, (RAM[SP+20] != 0), nil)
         	SP += 20
         case 6:		// CreateMemo
        	CreateMemo(frmVM, "mem", int(RAM[SP])*0xFF+int(RAM[SP+1]), int(RAM[SP+2])*0xFF+int(RAM[SP+3]), int(RAM[SP+4])*0xFF+int(RAM[SP+5]), int(RAM[SP+6])*0xFF+int(RAM[SP+7]), int((uint32(RAM[SP+8]) << 24)+(uint32(RAM[SP+9]) << 16)+(uint32(RAM[SP+10]) << 8)+uint32(RAM[SP+11])), int((uint32(RAM[SP+12]) << 24)+(uint32(RAM[SP+13]) << 16)+(uint32(RAM[SP+14]) << 8)+uint32(RAM[SP+15])), nil)
         	SP += 15
         }
    default: 
         printTerminal("Error run time VM: Unrecognized command " + strconv.Itoa(int(RAM[IP])))
         fmt.Println(RAM)
         RAM[IP+1] = cmHLT;
    }
    IP++	
}


func loadOVM(result []byte) {
	RAM = make([]byte, 0)
	for i := 0; i < len(result); i++ {
		RAM = append(RAM, result[i])
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

