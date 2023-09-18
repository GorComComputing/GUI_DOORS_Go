package main

import (
	"strings"
	"strconv"
	"fmt"
)

var btnRunVM *Node
var btnEnterVM *Node
var btnResetVM *Node
var btnAsmVM *Node
var edtEnterVM *Node
var lblIsRun *Node


func startVM(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 290
    frmMain.obj.(*tForm).y = 240
    setSize(frmMain, 300, 170)
    
    frmMain.children[1].obj.(*tBitBtn).enabled = false
    
    frmMain.obj.(*tForm).BC = 0xd8dcff

    btnRunVM = CreateBtn(frmMain, "btnRunVM", 25, 115, 70, 24, 0xd8dcc0, 0x0, "Run", btnRunVMClick)
    btnEnterVM = CreateBtn(frmMain, "btnEnterVM", 105, 115, 70, 24, 0xd8dcc0, 0x0, "Enter", btnEnterVMClick)
    btnResetVM = CreateBtn(frmMain, "btnResetVM", 185, 115, 70, 24, 0xd8dcc0, 0x0, "Reset", btnResetVMClick)
    btnAsmVM = CreateBtn(frmMain, "btnAsmVM", 185, 33, 80, 24, 0xd8dcc0, 0x0, "Assembly", btnAsmVMClick)
    
    edtEnterVM = CreateEdit(frmMain, "edtEnterVM", 25, 75, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
    lblIsRun = CreateLabel(frmMain, "lblIsRun", 25, 33, 40, 20, 0xd8dcff, 0x0, "STOP", nil)
   
    
  
}


func btnAsmVMClick(node *Node){
	isRun = false
	lblIsRun.obj.(*tLabel).caption = "STOP"
	
	RAM = make([]int, 0) 
	
	textAsm = ReadFile(RootDir + "asm.asm")
	textAsm = strings.Replace(textAsm, "\r\n", string(10), -1)
	InitNameTable()
	Assemble()
	
	for i := 0; i < 4096; i++ {
		RAM = append(RAM, 0)
	}
	
	IP = 0
	SP = len(RAM)
	
	fmt.Println(RAM)
	
	var tmp string = ""
	for i := 0; i < PC; i++ {
		tmp += string(RAM[i])
	}
	
	WriteFile(RootDir + "asm.dor", tmp)
	
}


func btnRunVMClick(node *Node){
	isRun = true
	lblIsRun.obj.(*tLabel).caption = "RUN"
	for ; isRun; {
		StepOVM()
	}
}


func btnResetVMClick(node *Node){
	isRun = false
	lblIsRun.obj.(*tLabel).caption = "STOP"
	IP = 0
	//RAM = make([]int, MemSize, MemSize) 
	//copy(RAM, asmProgram1)
	SP = len(RAM)
}


func btnEnterVMClick(node *Node){
	RAM[SP], _ = strconv.Atoi(edtEnterVM.obj.(*tEdit).text);
	printTerminal(edtEnterVM.obj.(*tEdit).text + "\n")
	btnRunVMClick(btnRunVM)
}

/*
var asmProgram1 = []int{
	cmPUSH, 6,
	cmSYSCALL, 0,
	
	cmOut, 2,
	cmOutLn, 
	cmOut, 3, 
	
	cmIn, 
	cmSYSCALL, 0, 
	 
	cmPUSH, 0,	// nil
	cmPUSH, 1,	// true
	cmPUSH, 33,	// caption
	cmPUSH, 1, 	// WIN
	cmPUSH, 0xFFAA00, // BC
	cmPUSH, 200,
	cmPUSH, 200,
	cmPUSH, 300,
	cmPUSH, 600,
	cmSYSCALL, 1,
	   
	cmStop,
	
	int('H'),int('E'),int('L'),int('O'), 0,
}
*/

