; This is a first assembly program (SNMP)

		jmp Start

Head:  db "SNMP by assembly",0

Start:		
		push 0                   ; nil
		push 1                   ; true
		pushd Head               ; caption
		push 1                   ; WIN
		pushd 14212288           ; BC
		pushw 440                ; sizeY
		pushw 568                ; sizeX
		pushw 300                ; y
		pushw 300                ; x
		syscall 1                ; Create window

		pushd lblIPaddr
		pushd 0
		pushd 14212288
		pushw 20
		pushw 120
		pushw 32
		pushw 12
		syscall 2                ; Create Label
		
		pushd lblPortGet
		pushd 0
		pushd 14212288
		pushw 20
		pushw 120
		pushw 70
		pushw 12
		syscall 2                ; Create Label
		
		pushd lblPortTrap
		pushd 0
		pushd 14212288
		pushw 20
		pushw 120
		pushw 70
		pushw 220
		syscall 2                ; Create Label
		
		pushd lblOID
		pushd 0
		pushd 14212288
		pushw 20
		pushw 120
		pushw 170
		pushw 165
		syscall 2                ; Create Label
		
		pushd lblValue
		pushd 0
		pushd 14212288
		pushw 20
		pushw 120
		pushw 200
		pushw 150
		syscall 2                ; Create Label
		
		pushd editIPaddr
		pushd 0
		pushd 16317688
		pushw 20
		pushw 100
		pushw 30
		pushw 100
		syscall 4                ; Create Edit
		
		pushd editPortGet
		pushd 0
		pushd 16317688
		pushw 20
		pushw 100
		pushw 68
		pushw 100
		syscall 4                ; Create Edit
		
		pushd editPortTrap
		pushd 0
		pushd 16317688
		pushw 20
		pushw 100
		pushw 68
		pushw 300
		syscall 4                ; Create Edit
		
		pushd editOID
		pushd 0
		pushd 16317688
		pushw 20
		pushw 200
		pushw 170
		pushw 200
		syscall 4                ; Create Edit
		
		pushd editValue
		pushd 0
		pushd 16317688
		pushw 20
		pushw 200
		pushw 200
		pushw 200
		syscall 4                ; Create Edit
		
		pushd btnTrapServer
		pushd 0
		pushd 14212288
		pushw 24
		pushw 100
		pushw 30
		pushw 300
		syscall 3                ; Create Btn
		
		pushd btnSendGet
		pushd 0
		pushd 14212288
		pushw 24
		pushw 70
		pushw 100
		pushw 50
		syscall 3                ; Create Btn
		
		pushd btnSet
		pushd 0
		pushd 14212288
		pushw 24
		pushw 70
		pushw 100
		pushw 130
		syscall 3                ; Create Btn
		
		pushd btnSendTrap
		pushd 0
		pushd 14212288
		pushw 24
		pushw 70
		pushw 100
		pushw 300
		syscall 3                ; Create Btn
		
		push 0
		pushd cbxVersion1
		pushd 0
		pushd 14212288
		pushw 16
		pushw 140
		pushw 30
		pushw 430
		syscall 5                ; Create CheckBox
		
		push 1
		pushd cbxVersion2
		pushd 0
		pushd 14212288
		pushw 16
		pushw 140
		pushw 60
		pushw 430
		syscall 5                ; Create CheckBox
		
		push 0
		pushd cbxVersion3
		pushd 0
		pushd 14212288
		pushw 16
		pushw 140
		pushw 90
		pushw 430
		syscall 5                ; Create CheckBox
		
		pushd 0
		pushd 14212288
		pushw 208
		pushw 564
		pushw 230
		pushw 2
		syscall 6                ; Create Memo
		


Exit:
		hlt

 lblIPaddr:   db "IP address",0
 lblPortGet:  db "Port Get",0
 lblPortTrap: db "Port Trap",0
 lblOID:      db "OID",0
 lblValue:    db "Value",0
 
  editIPaddr:   db "127.0.0.1",0
  editPortGet:  db "161",0
  editPortTrap: db "9161",0
  editOID:      db "1.3.6.1.2.1.1.4.0",0
  editValue:    db "",0
  
  btnTrapServer:  db "Run Server",0
  btnSendGet:     db "Get",0
  btnSet:         db "Set",0
  btnSendTrap:    db "Trap",0
  
  cbxVersion1:    db "Version 1",0
  cbxVersion2:    db "Version 2",0
  cbxVersion3:    db "Version 3",0
