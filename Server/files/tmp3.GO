package main

import ()

var lblIPaddr *Node
var editIPaddr *Node
var lblPortGet *Node
var editPortGet *Node
var lblPortTrap *Node
var editPortTrap *Node
var btnTrapServer *Node
var btnSendGet *Node
var btnSet *Node
var btnSendTrap *Node
var lblOID *Node
var editOID *Node
var lblValue *Node
var editValue *Node
var btnSendHelp *Node
var btnBrowser *Node
var cbxVersion1 *Node
var cbxVersion2 *Node
var cbxVersion3 *Node


func startProc(frmMain *Node){ 
    frmMain.obj.(*tForm).x = 190
    frmMain.obj.(*tForm).y = 70
    frmMain.obj.(*tForm).sizeX = 550
    frmMain.obj.(*tForm).sizeY = 300
    frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17

    lblIPaddr = CreateLabel(frmMain, "lblIPaddr", 12, 32, 120, 20, 0xd8dcc0, 0x0, "IP address", nil)
    editIPaddr = CreateEdit(frmMain, "editIPaddr", 100, 30, 100, 20, 0xf8fcf8, 0x0, "127.0.0.1", nil, nil)
    lblPortGet = CreateLabel(frmMain, "lblPortGet", 12, 70, 120, 20, 0xd8dcc0, 0x0, "Port Get", nil)
    editPortGet = CreateEdit(frmMain, "editPortGet", 100, 68, 100, 20, 0xf8fcf8, 0x0, "161", nil, nil)
    lblPortTrap = CreateLabel(frmMain, "lblPortTrap", 220, 70, 120, 20, 0xd8dcc0, 0x0, "Port Trap", nil)
    editPortTrap = CreateEdit(frmMain, "editPortTrap", 300, 68, 100, 20, 0xf8fcf8, 0x0, "9161", nil, nil)
    btnTrapServer = CreateBtn(frmMain, "btnTrapServer", 300, 30, 100, 24, 0xd8dcc0, 0x0, "Run Server", nil)
    btnSendGet = CreateBtn(frmMain, "btnSendGet", 50, 100, 70, 24, 0xd8dcc0, 0x0, "Get", nil)
    btnSet = CreateBtn(frmMain, "btnSet", 130, 100, 70, 24, 0xd8dcc0, 0x0, "Set", nil)
    btnSendTrap = CreateBtn(frmMain, "btnSendTrap", 300, 100, 70, 24, 0xd8dcc0, 0x0, "Trap", nil)
    lblOID = CreateLabel(frmMain, "lblOID", 165, 170, 120, 20, 0xd8dcc0, 0x0, "OID", nil)
    editOID = CreateEdit(frmMain, "editOID", 200, 170, 200, 20, 0xf8fcf8, 0x0, "1.3.6.1.2.1.1.4.0", nil, nil)
    lblValue = CreateLabel(frmMain, "lblValue", 150, 200, 120, 20, 0xd8dcc0, 0x0, "Value", nil)
    editValue = CreateEdit(frmMain, "editValue", 200, 200, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
    btnSendHelp = CreateBtn(frmMain, "btnSendHelp", 470, 270, 70, 24, 0xd8dcc0, 0x0, "Help", nil)
    btnBrowser = CreateBtn(frmMain, "btnBrowser", 12, 270, 90, 24, 0xd8dcc0, 0x0, "Browser", nil)
    cbxVersion1 = CreateCheckBox(frmMain, "cbxVersion1", 430, 30, 140, 16, 0xd8dcc0, 0x0, "Version 1", false, nil)
    cbxVersion2 = CreateCheckBox(frmMain, "cbxVersion2", 430, 60, 140, 16, 0xd8dcc0, 0x0, "Version 2", false, nil)
    cbxVersion3 = CreateCheckBox(frmMain, "cbxVersion3", 430, 90, 140, 16, 0xd8dcc0, 0x0, "Version 3", false, nil)
}


