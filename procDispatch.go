package main

import (
    "fmt"
    //"math/rand"
    //"math"
    //"syscall/js"
    //"time"
    "strconv"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    "encoding/json"
)


var btnAddEvent *Node
var btnRefreshEvents *Node

var btnPrevEvents *Node
var lblCurEventsPage *Node
var btnNextEvents *Node

var memTest *Node

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

var CurEventsPage = 20


var btnAddUser *Node
var btnRefreshUser *Node

var btnPrev *Node
var lblCurPage *Node
var btnNext *Node
var tblUsers *Node
var tblEvents *Node


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
	frmMain.obj.(*tForm).sizeX = 524
	frmMain.obj.(*tForm).sizeY = 252
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	frmMain.obj.(*tForm).visible = false
	
	
	listTabDispath := []string{"Events", "Users"} 
    pnlEventsDisp = CreatePanel(frmMain, "pnlEventsDisp", 2, 41, 520, 210, 0xd8dcc0, NONE, nil)
    pnlUsersDisp = CreatePanel(frmMain, "pnlUsersDisp", 2, 41, 520, 210, 0xd8dcc0, NONE, nil)
    pnlUsersDisp.obj.(*tPanel).visible = false
	tabDispatch = CreateTab(frmMain, "tabDispatch", 2, 20, 90, 20, 0xd8dcc0, 0x0, listTabDispath, tabDispathClick, nil)
	
	
	btnAddEvent = CreateBtn(pnlEventsDisp, "btnAddEvent", 12, 12, 70, 20, 0xD8DCC0, 0x000000, "Add", nil)
	btnRefreshEvents = CreateBtn(pnlEventsDisp, "btnRefreshEvents", 12 + 80, 12, 70, 20, 0xD8DCC0, 0x000000, "Refresh", btnRefreshEventsClick)
	

	listEvents := make([][]string, 5)
	for i := range listEvents {
    	listEvents[i] = make([]string, 6)
	}
    listEventsCols := []string{"Id", "Level", "Object", "Source", "Event", "Body"}
    tblEvents = CreateTable(pnlEventsDisp, " tblEvents", 12, 22 + 22, 362, 122, 0xf8fcf8, 0x0, listEventsCols, nil, listEvents, nil, nil)
	
	btnPrevEvents = CreateBtn(pnlEventsDisp, "btnPrevEvents", 90, 175, 60, 20, 0xD8DCC0, 0x000000, "Prev", btnPrevEventsClick)
	//btnPrevEvents.obj.(*tBtn).enabled = false
	lblCurEventsPage = CreateLabel(pnlEventsDisp, "lblCurEventsPage", 170, 175, 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurEventsPage), nil)
	btnNextEvents = CreateBtn(pnlEventsDisp, "btnNextEvents", 200, 175, 60, 20, 0xD8DCC0, 0x000000, "Next", btnNextEventsClick)
	
	memTest = CreateMemo(pnlEventsDisp, "memTest", 400, 44, 100, 100, 0x000000, 0xF8FCF8, "", nil)
	
	refreshEventsTable()
	
	
	btnAddUser = CreateBtn(pnlUsersDisp, "btnAddUser", 12, 12, 70, 20, 0xD8DCC0, 0x000000, "Add", nil)
	btnRefreshUser = CreateBtn(pnlUsersDisp, "btnRefreshUser", 12 + 80, 12, 70, 20, 0xD8DCC0, 0x000000, "Refresh", btnRefreshClick)
	

    listUsers := make([][]string, 5)
	for i := range listUsers {
    	listUsers[i] = make([]string, 5)
	}
    listUsersCols := []string{"Id", "Name", "Login", "Pswd", "Role"}
    tblUsers = CreateTable(pnlUsersDisp, "tblUsers", 12, 22 + 22, 302, 122, 0xf8fcf8, 0x0, listUsersCols, nil, listUsers, nil, nil)

	btnPrev = CreateBtn(pnlUsersDisp, "btnPrev", 90, 175, 60, 20, 0xD8DCC0, 0x000000, "Prev", btnPrevClick)
	btnPrev.obj.(*tBtn).enabled = false
	lblCurPage = CreateLabel(pnlUsersDisp, "lblCurPage", 170, 175, 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurUsersPage), nil)
	btnNext = CreateBtn(pnlUsersDisp, "btnNext", 200, 175, 60, 20, 0xD8DCC0, 0x000000, "Next", btnNextClick)
	
	refreshUsersTable()
}


func tabDispathClick(node *Node, x int, y int) {
	if node.obj.(*tTab).selected == 0 {
		pnlEventsDisp.obj.(*tPanel).visible = true
		pnlUsersDisp.obj.(*tPanel).visible = false
	} else {
		pnlUsersDisp.obj.(*tPanel).visible = true
		pnlEventsDisp.obj.(*tPanel).visible = false
	}
}


func btnRefreshEventsClick(node *Node){
	refreshEventsTable()	
}


func btnPrevEventsClick(node *Node){
	if CurEventsPage > 1 {
		CurEventsPage--
		if CurEventsPage == 1 {
			node.obj.(*tBtn).enabled = false
		}
	} 
	lblCurEventsPage.obj.(*tLabel).caption = strconv.Itoa(CurEventsPage)
	refreshEventsTable()
}


func btnNextEventsClick(node *Node){
	CurEventsPage++
	btnPrevEvents.obj.(*tBtn).enabled = true
	lblCurEventsPage.obj.(*tLabel).caption = strconv.Itoa(CurEventsPage)
	refreshEventsTable()
}


func refreshEventsTable(){
	for i:=0; i < 5; i++ {
		tblEvents.obj.(*tTable).list[i][0] = ""
		tblEvents.obj.(*tTable).list[i][1] = ""
		tblEvents.obj.(*tTable).list[i][2] = ""
		tblEvents.obj.(*tTable).list[i][3] = ""
		tblEvents.obj.(*tTable).list[i][4] = ""
		tblEvents.obj.(*tTable).list[i][5] = ""
		//eventsTable[i][6].obj.(*tBtn).visible = false
		//eventsTable[i][7].obj.(*tBtn).visible = false
	}
	
	result := Get("http://localhost:8085/api", "cmd=get_evnt " + strconv.Itoa(CurEventsPage) + " 5", "")
	
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


func btnRefreshClick(node *Node){
	refreshUsersTable()
}


func btnPrevClick(node *Node){
	if CurUsersPage > 1 {
		CurUsersPage--
		if CurUsersPage == 1 {
			node.obj.(*tBtn).enabled = false
		}
	} 
	lblCurPage.obj.(*tLabel).caption = strconv.Itoa(CurUsersPage)
	refreshUsersTable()
}


func btnNextClick(node *Node){
	CurUsersPage++
	btnPrev.obj.(*tBtn).enabled = true
	lblCurPage.obj.(*tLabel).caption = strconv.Itoa(CurUsersPage)
	refreshUsersTable()
}


func refreshUsersTable(){

	for i:=0; i < 5; i++ {
		tblUsers.obj.(*tTable).list[i][0] = ""
		tblUsers.obj.(*tTable).list[i][1] = ""
		tblUsers.obj.(*tTable).list[i][2] = ""
		tblUsers.obj.(*tTable).list[i][3] = ""
		tblUsers.obj.(*tTable).list[i][4] = ""
		//usersTable[i][5].obj.(*tBtn).visible = false
		//usersTable[i][6].obj.(*tBtn).visible = false
	}
	
	result := Get("http://localhost:8085/api", "cmd=get_usr " + strconv.Itoa(CurUsersPage) + " 5", "")
	
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
