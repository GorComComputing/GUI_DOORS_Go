; This is a test assembly program
	
		jmp Start

Title:  "My First Assembly Program",0

Start:
		;push 6
		;syscall 0              ; Run Terminal

		;out 2
		;outln
		;out 3

		;in
		;syscall 0              ; Run any program

		push 0                  ; nil
		push 1                  ; true
		push Title              ; caption
		push 1                  ; WIN
		push 127           ; BC
		push 127                ; sizeY
		push 127                ; sizeX
		push 127                ; y
		push 127                ; x
		syscall 1               ; Create window
Exit:
		hlt
	
Old:	72,101,112,111,0        ; "Hepo",0
