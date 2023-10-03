package main

import (
    //"fmt"
    //"math/rand"
    //"math"
    //"syscall/js"
    //"time"
    //"strconv"
    //"strings"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    //"encoding/json"
)


var pnlChrony *Node
var pnlSTVChrony *Node
var pnlSNMPChrony *Node
var pnlSystemChrony *Node
var pnlWebChrony *Node
var pnlWorkChrony *Node
var tabChrony *Node

// pnlSystemChrony
var memSystemChrony *Node
var lblPID *Node
var lblNameProc *Node
var lblCommand *Node
var edtPID *Node
var edtNameProc *Node
var edtCommand *Node
var btnPIDKill *Node
var btnNameKill *Node
var btnRunCmd *Node
var btnMonitor *Node
var btnNetstat *Node
var btnNtpq *Node

// pnlChrony
var memChrony *Node
var lblServer *Node
var lblAllow *Node
var lblDeny *Node
var edtServer *Node
var edtAllow *Node
var edtDeny *Node
var cbxIburst *Node
var cbxAllAllow *Node
var cbxAllDeny *Node
var btnStartChrony *Node
var btnStopChrony *Node
var btnRestartChrony *Node
var btnActivity *Node
var btnTracking *Node
var btnSources *Node
var btnSourceStats *Node
var btnClients *Node
var btnConfig *Node
var btnSaveConfig *Node
var btnRestore *Node

var cbxRTCsync *Node
var lblLeapsectz *Node
var lblDriftfile *Node
var lblMakestep *Node
var lblLogdir *Node
var lblLocalstratum *Node
var edtLeapsectz *Node
var edtDriftfile *Node
var edtMakestep *Node
var edtMakestep2 *Node
var edtLogdir *Node
var edtLocalstratum *Node


func startChrony(frmMain *Node){
	setSize(frmMain, 934, 600)
	frmMain.obj.(*tForm).x = BITMAP_WIDTH/2 - frmMain.obj.(*tForm).sizeX/2
	frmMain.obj.(*tForm).y = BITMAP_HEIGHT/2 - frmMain.obj.(*tForm).sizeY/2
	
	
	listTabChrony := []string{"Chrony", "STV", "SNMP", "System", "Web & SSH", "Working"} 
	pnlChrony = CreatePanel(frmMain, "pnlChrony", 2, 41, 930, 557, 0xd8dcc0, NONE, nil)
    pnlSTVChrony = CreatePanel(frmMain, "pnlSTVChrony", 2, 41, 930, 557, 0xd8dcc0, NONE, nil)
    pnlSTVChrony.obj.(*tPanel).visible = false
    pnlSNMPChrony = CreatePanel(frmMain, "pnlSNMPChrony", 2, 41, 930, 557, 0xd8dcc0, NONE, nil)
    pnlSNMPChrony.obj.(*tPanel).visible = false
    pnlSystemChrony = CreatePanel(frmMain, "pnlSystemChrony", 2, 41, 930, 557, 0xd8dcc0, NONE, nil)
    pnlSystemChrony.obj.(*tPanel).visible = false
    pnlWebChrony = CreatePanel(frmMain, "pnlWebChrony", 2, 41, 930, 557, 0xd8dcc0, NONE, nil)
    pnlWebChrony.obj.(*tPanel).visible = false
    pnlWorkChrony = CreatePanel(frmMain, "pnlWorkChrony", 2, 41, 930, 557, 0xd8dcc0, NONE, nil)
    pnlWorkChrony.obj.(*tPanel).visible = false
	tabChrony = CreateTab(frmMain, "tabChrony", 2, 20, 90, 20, 0xd8dcc0, 0x0, listTabChrony, tabChronyClick, nil)
	
	// pnlSystemChrony
	memSystemChrony = CreateMemo(pnlSystemChrony, "memSystemChrony", 7, 8, 800, 430, 0x2A242D, 0x000000, nil)
	
	lblPID = CreateLabel(pnlSystemChrony, "lblPID", 16, 457, 90, 20, 0xd8dcc0, 0x0, "PID process", nil)
	lblNameProc = CreateLabel(pnlSystemChrony, "lblNameProc", 16, 485, 90, 20, 0xd8dcc0, 0x0, "Name process", nil)
	lblCommand = CreateLabel(pnlSystemChrony, "lblCommand", 16, 517, 90, 20, 0xd8dcc0, 0x0, "Command", nil)
	
	edtPID = CreateEdit(pnlSystemChrony, "edtPID", 112, 457, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
	edtNameProc = CreateEdit(pnlSystemChrony, "edtNameProc", 112, 485, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
	edtCommand = CreateEdit(pnlSystemChrony, "edtCommand", 112, 517, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
	
	btnPIDKill = CreateBtn(pnlSystemChrony, "btnPIDKill", 320, 457, 70, 24, 0xd8dcc0, 0x0, "Kill", btnPIDKillClick)
	btnNameKill = CreateBtn(pnlSystemChrony, "btnNameKill", 320, 485, 70, 24, 0xd8dcc0, 0x0, "Kill", btnNameKillClick)
	btnRunCmd = CreateBtn(pnlSystemChrony, "btnRunCmd", 320, 517, 70, 24, 0xd8dcc0, 0x0, "Run", btnRunCmdClick)
	
	btnMonitor = CreateBtn(pnlSystemChrony, "btnMonitor", 818, 9, 100, 24, 0xd8dcc0, 0x0, "Monitor", btnMonitorClick)
	btnNetstat = CreateBtn(pnlSystemChrony, "btnNetstat", 818, 45, 100, 24, 0xd8dcc0, 0x0, "Netstat -a", btnNetstatClick)
	btnNtpq = CreateBtn(pnlSystemChrony, "btnNtpq", 818, 81, 100, 24, 0xd8dcc0, 0x0, "ntpq -p", btnNtpqClick)
	
	// pnlChrony
	memChrony = CreateMemo(pnlChrony, "memChrony", 405, 9, 400, 544, 0x2A242D, 0x000000, nil)
	
	lblServer = CreateLabel(pnlChrony, "lblServer", 16, 54, 90, 20, 0xd8dcc0, 0x0, "server", nil)
	lblAllow = CreateLabel(pnlChrony, "lblAllow", 16, 82, 90, 20, 0xd8dcc0, 0x0, "allow", nil)
	lblDeny = CreateLabel(pnlChrony, "lblDeny", 16, 110, 90, 20, 0xd8dcc0, 0x0, "deny", nil)
	
	edtServer = CreateEdit(pnlChrony, "edtServer", 112, 54, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
	edtAllow = CreateEdit(pnlChrony, "edtAllow", 112, 82, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
	edtDeny = CreateEdit(pnlChrony, "edtDeny", 112, 110, 200, 20, 0xf8fcf8, 0x0, "", nil, nil)
	
	cbxIburst = CreateCheckBox(pnlChrony, "cbxIburst", 320, 54, 140, 16, 0xD8DCC0, 0x000000, "iburst", false, cbxIburstClick)
	cbxAllAllow = CreateCheckBox(pnlChrony, "cbxAllAllow", 320, 82, 140, 16, 0xD8DCC0, 0x000000, "all", false, cbxAllAllowClick)
	cbxAllDeny = CreateCheckBox(pnlChrony, "cbxAllDeny", 320, 110, 140, 16, 0xD8DCC0, 0x000000, "all", false, cbxAllDenyClick)
	
	cbxRTCsync = CreateCheckBox(pnlChrony, "cbxRTCsync", 16, 138, 140, 16, 0xD8DCC0, 0x000000, "rtcsync", false, cbxRTCsyncClick)
	
	lblLeapsectz = CreateLabel(pnlChrony, "lblLeapsectz", 16, 166, 90, 20, 0xd8dcc0, 0x0, "leapsectz", nil)
	lblDriftfile = CreateLabel(pnlChrony, "lblDriftfile", 16, 194, 90, 20, 0xd8dcc0, 0x0, "driftfile", nil)
	lblMakestep = CreateLabel(pnlChrony, "lblMakestep", 16, 222, 90, 20, 0xd8dcc0, 0x0, "makestep", nil)
	lblLogdir = CreateLabel(pnlChrony, "lblLogdir", 16, 250, 90, 20, 0xd8dcc0, 0x0, "logdir", nil)
	lblLocalstratum = CreateLabel(pnlChrony, "lblLocalstratum", 16, 278, 120, 20, 0xd8dcc0, 0x0, "local stratum", nil)
	
	edtLeapsectz = CreateEdit(pnlChrony, "edtLeapsectz", 112, 166, 200, 20, 0xf8fcf8, 0x0, "right/UTC", nil, nil)
	edtDriftfile = CreateEdit(pnlChrony, "edtDriftfile", 112, 194, 200, 20, 0xf8fcf8, 0x0, "/var/lib/chrony/drift", nil, nil)
	edtMakestep = CreateEdit(pnlChrony, "edtMakestep", 112, 222, 100, 20, 0xf8fcf8, 0x0, "1.0", nil, nil)
	edtMakestep2 = CreateEdit(pnlChrony, "edtMakestep2", 217, 222, 96, 20, 0xf8fcf8, 0x0, "3", nil, nil)
	edtLogdir = CreateEdit(pnlChrony, "edtLogdir", 112, 250, 200, 20, 0xf8fcf8, 0x0, "/var/log/chrony", nil, nil)
	edtLocalstratum = CreateEdit(pnlChrony, "edtLocalstratum", 217, 278, 96, 20, 0xf8fcf8, 0x0, "8", nil, nil)
	
	
	
	
	btnStartChrony = CreateBtn(pnlChrony, "btnStartChrony", 70, 11, 70, 24, 0xd8dcc0, 0x0, "Start", btnStartChronyClick)
	btnStopChrony = CreateBtn(pnlChrony, "btnStopChrony", 160, 11, 70, 24, 0xd8dcc0, 0x0, "Stop", btnStopChronyClick)
	btnRestartChrony = CreateBtn(pnlChrony, "btnRestartChrony", 250, 11, 70, 24, 0xd8dcc0, 0x0, "Restart", btnRestartChronyClick)
	
	btnActivity = CreateBtn(pnlChrony, "btnActivity", 818, 9, 100, 24, 0xd8dcc0, 0x0, "Activity", btnActivityClick)
	btnTracking = CreateBtn(pnlChrony, "btnTracking", 818, 45, 100, 24, 0xd8dcc0, 0x0, "Tracking", btnTrackingClick)
	btnSources = CreateBtn(pnlChrony, "btnSources", 818, 81, 100, 24, 0xd8dcc0, 0x0, "Sources", btnSourcesClick)
	btnSourceStats = CreateBtn(pnlChrony, "btnSourceStats", 818, 117, 100, 24, 0xd8dcc0, 0x0, "SourceStats", btnSourceStatsClick)
	btnClients = CreateBtn(pnlChrony, "btnClients", 818, 153, 100, 24, 0xd8dcc0, 0x0, "Clients", btnClientsClick)
	btnConfig = CreateBtn(pnlChrony, "btnConfig", 818, 189, 100, 24, 0xd8dcc0, 0x0, "Config", btnConfigClick)
	btnSaveConfig = CreateBtn(pnlChrony, "btnSaveConfig", 818, 225, 100, 24, 0xd8dcc0, 0x0, "Save Config", btnSaveConfigClick)
	btnRestore = CreateBtn(pnlChrony, "btnRestore", 818, 261, 100, 24, 0xd8dcc0, 0x0, "Restore", btnRestoreClick)

}


func tabChronyClick(node *Node, x int, y int) {
	if node.obj.(*tTab).selected == 0 {
		pnlChrony.obj.(*tPanel).visible = true
		pnlSTVChrony.obj.(*tPanel).visible = false
		pnlSNMPChrony.obj.(*tPanel).visible = false
		pnlSystemChrony.obj.(*tPanel).visible = false
		pnlWebChrony.obj.(*tPanel).visible = false
		pnlWorkChrony.obj.(*tPanel).visible = false
	} else if node.obj.(*tTab).selected == 1 {
		pnlChrony.obj.(*tPanel).visible = false
		pnlSTVChrony.obj.(*tPanel).visible = true
		pnlSNMPChrony.obj.(*tPanel).visible = false
		pnlSystemChrony.obj.(*tPanel).visible = false
		pnlWebChrony.obj.(*tPanel).visible = false
		pnlWorkChrony.obj.(*tPanel).visible = false
	} else if node.obj.(*tTab).selected == 2 {
		pnlChrony.obj.(*tPanel).visible = false
		pnlSTVChrony.obj.(*tPanel).visible = false
		pnlSNMPChrony.obj.(*tPanel).visible = true
		pnlSystemChrony.obj.(*tPanel).visible = false
		pnlWebChrony.obj.(*tPanel).visible = false
		pnlWorkChrony.obj.(*tPanel).visible = false
	} else if node.obj.(*tTab).selected == 3 {
		pnlChrony.obj.(*tPanel).visible = false
		pnlSTVChrony.obj.(*tPanel).visible = false
		pnlSNMPChrony.obj.(*tPanel).visible = false
		pnlSystemChrony.obj.(*tPanel).visible = true
		pnlWebChrony.obj.(*tPanel).visible = false
		pnlWorkChrony.obj.(*tPanel).visible = false
	} else if node.obj.(*tTab).selected == 4 {
		pnlChrony.obj.(*tPanel).visible = false
		pnlSTVChrony.obj.(*tPanel).visible = false
		pnlSNMPChrony.obj.(*tPanel).visible = false
		pnlSystemChrony.obj.(*tPanel).visible = false
		pnlWebChrony.obj.(*tPanel).visible = true
		pnlWorkChrony.obj.(*tPanel).visible = false
	} else if node.obj.(*tTab).selected == 5 {
		pnlChrony.obj.(*tPanel).visible = false
		pnlSTVChrony.obj.(*tPanel).visible = false
		pnlSNMPChrony.obj.(*tPanel).visible = false
		pnlSystemChrony.obj.(*tPanel).visible = false
		pnlWebChrony.obj.(*tPanel).visible = false
		pnlWorkChrony.obj.(*tPanel).visible = true
	}
}


func btnPIDKillClick(node *Node){
}


func btnNameKillClick(node *Node){
}


func btnRunCmdClick(node *Node){
}


func btnMonitorClick(node *Node){
}


func btnNetstatClick(node *Node){
}


func btnNtpqClick(node *Node){
}


func cbxIburstClick(node *Node){
}


func cbxAllAllowClick(node *Node){
}


func cbxAllDenyClick(node *Node){
}


func btnStartChronyClick(node *Node){
}


func btnStopChronyClick(node *Node){
}


func btnRestartChronyClick(node *Node){
}


func btnActivityClick(node *Node){
}


func btnTrackingClick(node *Node){
}


func btnSourcesClick(node *Node){
}


func btnSourceStatsClick(node *Node){
}


func btnClientsClick(node *Node){
}


func btnConfigClick(node *Node){
}


func btnSaveConfigClick(node *Node){
}


func btnRestoreClick(node *Node){
}


func cbxRTCsyncClick(node *Node){
}

