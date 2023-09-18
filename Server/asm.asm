; This is a test assembly program

		jmp Start

Title:  db "The My Test",0

Start:
		;push 6
		;syscall 0              ; Run Terminal

		;out 2
		;outln
		;out 3

		;in
		;syscall 0              ; Run any program

		out 120432
		outln
		push 0                  ; nil
		push 1                  ; true
		pushd Title              ; caption
		out 512
		push 1                 ; WIN
		pushd 11272427                ; BC
		pushw 300                ; sizeY
		pushw 400                ; sizeX
		pushw 300                ; y
		pushw 300                ; x
		syscall 1               ; Create window
Exit:
		hlt

Old:	dw 72,101,112,111,0        ; "Hepo",0
