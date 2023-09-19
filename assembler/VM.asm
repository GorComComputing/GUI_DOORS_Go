; This is a second assembly program (Virtual Machine)

		jmp Start

Head:  db "Virtual Machine by assembly",0     

Start:		
		push 0                   ; nil
		push 1                   ; true
		pushd Head               ; caption
		push 1                   ; WIN
		pushd 14212351           ; BC
		pushw 170                ; sizeY
		pushw 300                ; sizeX
		pushw 240                ; y
		pushw 290                ; x
		syscall 1                ; Create window

		pushd lblIsRun
		pushd 0
		pushd 14212351
		pushw 20
		pushw 40
		pushw 33
		pushw 25
		syscall 2                ; Create Label
		
		pushd edtEnterVM
		pushd 0
		pushd 16317688
		pushw 20
		pushw 200
		pushw 75
		pushw 25
		syscall 4                ; Create Edit
		
		pushd btnRunVM
		pushd 0
		pushd 14212351
		pushw 24
		pushw 70
		pushw 115
		pushw 25
		syscall 3                ; Create Btn
		
		pushd btnEnterVM
		pushd 0
		pushd 14212351
		pushw 24
		pushw 70
		pushw 115
		pushw 105
		syscall 3                ; Create Btn
		
		pushd btnResetVM
		pushd 0
		pushd 14212351
		pushw 24
		pushw 70
		pushw 115
		pushw 185
		syscall 3                ; Create Btn
		
		pushd btnAsmVM
		pushd 0
		pushd 14212351
		pushw 24
		pushw 80
		pushw 33
		pushw 185
		syscall 3                ; Create Btn


 Exit:
		hlt
		
 lblIsRun:    db "STOP",0
 edtEnterVM:  db "",0
  
 btnRunVM:    db "Run",0
 btnEnterVM:  db "Enter",0
 btnResetVM:  db "Reset",0
 btnAsmVM:    db "Assembly",0
