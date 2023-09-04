package main

import (
    "fmt"
    //"math/rand"
    //"math"
    //"syscall/js"
    //"time"
    "strconv"
    "strings"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    "encoding/json"
)


var btnAddDevice *Node
var btnRefreshDevices *Node 

var btnPrevDevices *Node
var lblCurDevicesPage *Node
var btnNextDevices *Node

type DevicesFromDB struct {
	Id int
	Name string
	IPaddr string
	Version string
	GNSS string
	PTP string
	PZG_VZG string
}

var CurDevicesPage = 1

var btnMenuDevice *Node
var btnSyncDevice *Node
var btnPTPDevice *Node
var btnGNSSDevice *Node
var btnDebugDevice *Node



var btnAddEvent *Node
var btnRefreshEvents *Node

var btnPrevEvents *Node
var lblCurEventsPage *Node
var btnNextEvents *Node

var memTest *Node

var pnlDevicesDisp *Node
var pnlEventsDisp *Node
var pnlUsersDisp *Node
var tabDispatch *Node

type EventsFromDB struct {
	Id int
	Level string
	Obj_id int
	Source string
	Event string
	Body string
	Is_checked bool
	Time string
}

var CurEventsPage = 1 //10


var btnAddUser *Node
var btnRefreshUser *Node

var btnPrev *Node
var lblCurPage *Node
var btnNext *Node
var tblUsers *Node
var tblEvents *Node
var tblDevices *Node


type UsersFromDB struct {
	Id int
	UserName string
	Login string
	Pswd string
	UserRole int
}

var CurUsersPage = 1


func startDispatch(frmMain *Node){
	frmMain.obj.(*tForm).x = 100
	frmMain.obj.(*tForm).y = 100
	frmMain.obj.(*tForm).sizeX = 904
	frmMain.obj.(*tForm).sizeY = 353
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	frmMain.obj.(*tForm).visible = false
	
	
	listTabDispath := []string{"Devices", "Events", "Users"} 
	pnlDevicesDisp = CreatePanel(frmMain, "pnlDevicesDisp", 2, 41, 900, 310, 0xd8dcc0, NONE, nil)
    pnlEventsDisp = CreatePanel(frmMain, "pnlEventsDisp", 2, 41, 900, 310, 0xd8dcc0, NONE, nil)
    pnlEventsDisp.obj.(*tPanel).visible = false
    pnlUsersDisp = CreatePanel(frmMain, "pnlUsersDisp", 2, 41, 900, 310, 0xd8dcc0, NONE, nil)
    pnlUsersDisp.obj.(*tPanel).visible = false
	tabDispatch = CreateTab(frmMain, "tabDispatch", 2, 20, 90, 20, 0xd8dcc0, 0x0, listTabDispath, tabDispathClick, nil)
	
	// Panel Devices
	btnAddDevice = CreateBitBtn(pnlDevicesDisp, "btnAddDevice", bmpAdd, 20, 7, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, nil)
	btnRefreshDevices = CreateBitBtn(pnlDevicesDisp, "btnRefreshDevices", bmpRefresh, 60, 7, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnRefreshDevicesClick)
	
	listDevices := make([][]string, 10)
	for i := range listDevices {
    	listDevices[i] = make([]string, 6)
	}
    listDevicesCols := []string{"Mode", "Name", "gnss_ref", "ptp_ref", "IP addr", "Version"}
    listDevicesSizeCols := []int{50, 170, 100, 100, 120, 80}
    tblDevices = CreateTable(pnlDevicesDisp, "tblDevices", 12, 22 + 22, 622, 222, 0xf8fcf8, 0x0, listDevicesCols, listDevicesSizeCols, nil, listDevices, 120, 20, nil, nil)
	
	btnPrevDevices = CreateBitBtn(pnlDevicesDisp, "btnPrevDevices", bmpBack, tblDevices.obj.(*tTable).x + tblDevices.obj.(*tTable).sizeX/2 - 50, 269, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnPrevDevicesClick)
	//btnPrevDevices.obj.(*tBitBtn).enabled = false
	lblCurDevicesPage = CreateLabel(pnlDevicesDisp, "lblCurDevicesPage", tblDevices.obj.(*tTable).x + tblDevices.obj.(*tTable).sizeX/2, 278, 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurDevicesPage), nil)
	btnNextDevices = CreateBitBtn(pnlDevicesDisp, "btnNextDevices", bmpNext, tblDevices.obj.(*tTable).x + tblDevices.obj.(*tTable).sizeX/2 + 30, 269, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnNextDevicesClick)
	
	btnMenuDevice = CreateBtn(pnlDevicesDisp, "btnMenuDevice", 648, 44, 70, 20, 0xD8DCC0, 0x000000, "Menu", btnMenuDeviceClick)
	btnSyncDevice = CreateBtn(pnlDevicesDisp, "btnSyncDevice", 648, 44+30, 70, 20, 0xD8DCC0, 0x000000, "Sync", btnSyncDeviceClick)
	btnPTPDevice = CreateBtn(pnlDevicesDisp, "btnPTPDevice", 648, 44+30+30, 70, 20, 0xD8DCC0, 0x000000, "PTP", btnPTPDeviceClick)
	btnGNSSDevice = CreateBtn(pnlDevicesDisp, "btnGNSSDevice", 648, 44+30+30+30, 70, 20, 0xD8DCC0, 0x000000, "GNSS", btnGNSSDeviceClick)
	btnDebugDevice = CreateBtn(pnlDevicesDisp, "btnDebugDevice", 648, 44+30+30+30+30, 70, 20, 0xD8DCC0, 0x000000, "Debug", btnDebugDeviceClick)
	
	
	
	// Panel Events
	btnAddEvent = CreateBitBtn(pnlEventsDisp, "btnAddEvent", bmpAdd, 20, 7, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, nil)
	btnRefreshEvents = CreateBitBtn(pnlEventsDisp, "btnRefreshEvents", bmpRefresh, 60, 7, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnRefreshEventsClick)
	

	listEvents := make([][]string, 10)
	for i := range listEvents {
    	listEvents[i] = make([]string, 6)
	}
    listEventsCols := []string{"Id", "Level", "Object", "Source", "Event", "Body"}
    listEventsSizeCols := []int{30, 100, 80, 80, 100, 80}
    tblEvents = CreateTable(pnlEventsDisp, " tblEvents", 12, 22 + 22, 472, 222, 0xf8fcf8, 0x0, listEventsCols, listEventsSizeCols, nil, listEvents, 120, 20, nil, nil)
	
	btnPrevEvents = CreateBitBtn(pnlEventsDisp, "btnPrevEvents", bmpBack, tblEvents.obj.(*tTable).x + tblEvents.obj.(*tTable).sizeX/2 - 50, 269, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnPrevEventsClick)
	//btnPrevEvents.obj.(*tBitBtn).enabled = false
	lblCurEventsPage = CreateLabel(pnlEventsDisp, "lblCurEventsPage", tblEvents.obj.(*tTable).x + tblEvents.obj.(*tTable).sizeX/2, 278, 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurEventsPage), nil)
	btnNextEvents = CreateBitBtn(pnlEventsDisp, "btnNextEvents", bmpNext, tblEvents.obj.(*tTable).x + tblEvents.obj.(*tTable).sizeX/2 + 30, 269, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnNextEventsClick)
	
	memTest = CreateMemo(pnlEventsDisp, "memTest", 760, 44, 100, 100, 0x000000, 0xF8FCF8, "", nil)
	

	
	// Panel Users
	btnAddUser = CreateBitBtn(pnlUsersDisp, "btnAddUser", bmpAdd, 20, 7, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, nil)
	btnRefreshUser = CreateBitBtn(pnlUsersDisp, "btnRefreshUser", bmpRefresh, 60, 7, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnRefreshClick)
	

    listUsers := make([][]string, 10)
	for i := range listUsers {
    	listUsers[i] = make([]string, 5)
	}
    listUsersCols := []string{"Id", "Name", "Login", "Pswd", "Role"}
    listUsersSizeCols := []int{30, 140, 100, 100, 40}
    tblUsers = CreateTable(pnlUsersDisp, "tblUsers", 12, 22 + 22, 412, 222, 0xf8fcf8, 0x0, listUsersCols, listUsersSizeCols, nil, listUsers, 120, 20, nil, nil)

	btnPrev = CreateBitBtn(pnlUsersDisp, "btnPrev", bmpBack, tblUsers.obj.(*tTable).x + tblUsers.obj.(*tTable).sizeX/2 - 50, 269, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnPrevClick)
	btnPrev.obj.(*tBitBtn).enabled = false
	lblCurPage = CreateLabel(pnlUsersDisp, "lblCurPage", tblUsers.obj.(*tTable).x + tblUsers.obj.(*tTable).sizeX/2, 278, 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurUsersPage), nil)
	btnNext = CreateBitBtn(pnlUsersDisp, "btnNext", bmpNext, tblUsers.obj.(*tTable).x + tblUsers.obj.(*tTable).sizeX/2 + 30, 269, 30, 30, 0xD8DCC0, 0x000000, "", FLAT, btnNextClick)
	
	refreshDevicesTable()
	refreshEventsTable()
	refreshUsersTable()
}


func tabDispathClick(node *Node, x int, y int) {
	if node.obj.(*tTab).selected == 0 {
		pnlDevicesDisp.obj.(*tPanel).visible = true
		pnlEventsDisp.obj.(*tPanel).visible = false
		pnlUsersDisp.obj.(*tPanel).visible = false
	} else if node.obj.(*tTab).selected == 1 {
		pnlDevicesDisp.obj.(*tPanel).visible = false
		pnlEventsDisp.obj.(*tPanel).visible = true
		pnlUsersDisp.obj.(*tPanel).visible = false
	} else if node.obj.(*tTab).selected == 2 {
		pnlDevicesDisp.obj.(*tPanel).visible = false
		pnlEventsDisp.obj.(*tPanel).visible = false
		pnlUsersDisp.obj.(*tPanel).visible = true
	}
}


func btnMenuDeviceClick(node *Node){
	execProcess(2)  // Run Browser
	newPageBrowser(tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][4], "http://" + tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][4] + "/index.html")
}


func btnSyncDeviceClick(node *Node){
	execProcess(2)  // Run Browser
	newPageBrowser(tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][4], "http://" + tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][4] + "/sync.html")
}


func btnPTPDeviceClick(node *Node){
	execProcess(2)  // Run Browser
	newPageBrowser(tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][4], "http://" + tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][4] + "/ptp" + strings.ToLower(strings.TrimRight(tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][0], " ")) + ".html")
}


func btnGNSSDeviceClick(node *Node){
	execProcess(2)  // Run Browser
	newPageBrowser(tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][4], "http://" + tblDevices.obj.(*tTable).list[tblDevices.obj.(*tTable).selectedY][4] + "/gnss.html")
}


func btnDebugDeviceClick(node *Node){

}


// Panel Devices
func btnRefreshDevicesClick(node *Node){
	refreshDevicesTable()	
}


func btnPrevDevicesClick(node *Node){
	if CurDevicesPage > 1 {
		CurDevicesPage--
		if CurDevicesPage == 1 {
			node.obj.(*tBitBtn).enabled = false
		}
	} 
	lblCurDevicesPage.obj.(*tLabel).caption = strconv.Itoa(CurDevicesPage)
	refreshDevicesTable()
}


func btnNextDevicesClick(node *Node){
	CurDevicesPage++
	btnPrevDevices.obj.(*tBitBtn).enabled = true
	lblCurDevicesPage.obj.(*tLabel).caption = strconv.Itoa(CurDevicesPage)
	refreshDevicesTable()
}


func refreshDevicesTable(){
	for i:=0; i < 10; i++ {
		tblDevices.obj.(*tTable).list[i][0] = ""
		tblDevices.obj.(*tTable).list[i][1] = ""
		tblDevices.obj.(*tTable).list[i][2] = ""
		tblDevices.obj.(*tTable).list[i][3] = ""
		tblDevices.obj.(*tTable).list[i][4] = ""
		tblDevices.obj.(*tTable).list[i][5] = ""
		//eventsTable[i][6].obj.(*tBtn).visible = false
		//eventsTable[i][7].obj.(*tBtn).visible = false
	}
	
	result := Get("http://localhost:8085/api", "cmd=get_dev " + strconv.Itoa(CurDevicesPage) + " 10", "")
	
	var devices []DevicesFromDB 
	err := json.Unmarshal([]byte(result), &devices)
	if err != nil {
		fmt.Println(err)
	}
	
	for i, _ := range devices {
		tblDevices.obj.(*tTable).list[i][0] = devices[i].PZG_VZG
		tblDevices.obj.(*tTable).list[i][1] = devices[i].Name
		tblDevices.obj.(*tTable).list[i][2] = ""
		tblDevices.obj.(*tTable).list[i][3] = ""
		tblDevices.obj.(*tTable).list[i][4] = devices[i].IPaddr
		tblDevices.obj.(*tTable).list[i][5] = devices[i].Version
		//eventsTable[i][6].obj.(*tBtn).visible = true
		//eventsTable[i][7].obj.(*tBtn).visible = true
	}
}


// Panel Events
func btnRefreshEventsClick(node *Node){
	refreshEventsTable()	
}


func btnPrevEventsClick(node *Node){
	if CurEventsPage > 1 {
		CurEventsPage--
		if CurEventsPage == 1 {
			node.obj.(*tBitBtn).enabled = false
		}
	} 
	lblCurEventsPage.obj.(*tLabel).caption = strconv.Itoa(CurEventsPage)
	refreshEventsTable()
}


func btnNextEventsClick(node *Node){
	CurEventsPage++
	btnPrevEvents.obj.(*tBitBtn).enabled = true
	lblCurEventsPage.obj.(*tLabel).caption = strconv.Itoa(CurEventsPage)
	refreshEventsTable()
}


func refreshEventsTable(){
	for i:=0; i < 10; i++ {
		tblEvents.obj.(*tTable).list[i][0] = ""
		tblEvents.obj.(*tTable).list[i][1] = ""
		tblEvents.obj.(*tTable).list[i][2] = ""
		tblEvents.obj.(*tTable).list[i][3] = ""
		tblEvents.obj.(*tTable).list[i][4] = ""
		tblEvents.obj.(*tTable).list[i][5] = ""
		//eventsTable[i][6].obj.(*tBtn).visible = false
		//eventsTable[i][7].obj.(*tBtn).visible = false
	}
	
	result := Get("http://localhost:8085/api", "cmd=get_evnt " + strconv.Itoa(CurEventsPage) + " 10", "")
	
	var events []EventsFromDB 
	err := json.Unmarshal([]byte(result), &events)
	if err != nil {
		fmt.Println(err)
	}
	
	for i, _ := range events {
		tblEvents.obj.(*tTable).list[i][0] = strconv.Itoa(events[i].Id)
		tblEvents.obj.(*tTable).list[i][1] = events[i].Level
		tblEvents.obj.(*tTable).list[i][2] = strconv.Itoa(events[i].Obj_id)
		tblEvents.obj.(*tTable).list[i][3] = events[i].Source
		tblEvents.obj.(*tTable).list[i][4] = events[i].Event
		tblEvents.obj.(*tTable).list[i][5] = events[i].Body
		//eventsTable[i][6].obj.(*tBtn).visible = true
		//eventsTable[i][7].obj.(*tBtn).visible = true
	}
}


// Panel Users
func btnRefreshClick(node *Node){
	refreshUsersTable()
}


func btnPrevClick(node *Node){
	if CurUsersPage > 1 {
		CurUsersPage--
		if CurUsersPage == 1 {
			node.obj.(*tBitBtn).enabled = false
		}
	} 
	lblCurPage.obj.(*tLabel).caption = strconv.Itoa(CurUsersPage)
	refreshUsersTable()
}


func btnNextClick(node *Node){
	CurUsersPage++
	btnPrev.obj.(*tBitBtn).enabled = true
	lblCurPage.obj.(*tLabel).caption = strconv.Itoa(CurUsersPage)
	refreshUsersTable()
}


func refreshUsersTable(){

	for i:=0; i < 10; i++ {
		tblUsers.obj.(*tTable).list[i][0] = ""
		tblUsers.obj.(*tTable).list[i][1] = ""
		tblUsers.obj.(*tTable).list[i][2] = ""
		tblUsers.obj.(*tTable).list[i][3] = ""
		tblUsers.obj.(*tTable).list[i][4] = ""
		//usersTable[i][5].obj.(*tBtn).visible = false
		//usersTable[i][6].obj.(*tBtn).visible = false
	}
	
	result := Get("http://localhost:8085/api", "cmd=get_usr " + strconv.Itoa(CurUsersPage) + " 10", "")
	
	var users []UsersFromDB 
	err := json.Unmarshal([]byte(result), &users)
	if err != nil {
		fmt.Println(err)
	}
	
	for i, _ := range users {
		tblUsers.obj.(*tTable).list[i][0] = strconv.Itoa(users[i].Id)
		tblUsers.obj.(*tTable).list[i][1] = users[i].UserName
		tblUsers.obj.(*tTable).list[i][2] = users[i].Login
		tblUsers.obj.(*tTable).list[i][3] = users[i].Pswd
		tblUsers.obj.(*tTable).list[i][4] = strconv.Itoa(users[i].UserRole)
		//usersTable[i][5].obj.(*tBtn).visible = true
		//usersTable[i][6].obj.(*tBtn).visible = true
	}
}
