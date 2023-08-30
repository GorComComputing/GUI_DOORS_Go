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

var lblCap1 *Node
var lblCap2 *Node
var lblCap3 *Node
var lblCap4 *Node
var lblCap5 *Node
var lblCap6 *Node

var eventsTable [5][8]*Node

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
var lblId *Node
var lblUserName *Node
var lblLogin *Node
var lblUserPswd *Node
var lblUserRole *Node

var usersTable [5][7]*Node

var btnPrev *Node
var lblCurPage *Node
var btnNext *Node


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
	frmMain.obj.(*tForm).sizeX = 600
	frmMain.obj.(*tForm).sizeY = 252
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	frmMain.obj.(*tForm).visible = false
	
	
	listTabDispath := []string{"Events", "Users"} 
    pnlEventsDisp = CreatePanel(frmMain, "pnlEventsDisp", 2, 41, 596, 210, 0xd8dcc0, NONE, nil)
    pnlUsersDisp = CreatePanel(frmMain, "pnlUsersDisp", 2, 41, 596, 210, 0xd8dcc0, NONE, nil)
    pnlUsersDisp.obj.(*tPanel).visible = false
	tabDispatch = CreateTab(frmMain, "tabDispatch", 2, 20, 90, 20, 0xd8dcc0, 0x0, listTabDispath, tabDispathClick, nil)
	
	
					// 472, 202
	btnAddEvent = CreateBtn(pnlEventsDisp, "btnAddEvent", 12, 22, 60, 20, 0xD8DCC0, 0x000000, "Add", nil)
	btnRefreshEvents = CreateBtn(pnlEventsDisp, "btnRefreshEvents", 12 + 64, 22, 60, 20, 0xD8DCC0, 0x000000, "Refresh", btnRefreshEventsClick)
	
	lblCap1 = CreateLabel(pnlEventsDisp, "lblCap1", 12, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Id", nil)
	lblCap2 = CreateLabel(pnlEventsDisp, "lblCap2", 12 + 20, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Level", nil)
	lblCap3 = CreateLabel(pnlEventsDisp, "lblCap3", 12 + 20 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Obj", nil)
	lblCap4 = CreateLabel(pnlEventsDisp, "lblCap4", 12 + 20 + 50 + 30, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Source", nil)
	lblCap5 = CreateLabel(pnlEventsDisp, "lblCap5", 12 + 20 + 50 + 50 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Event", nil)
	lblCap6 = CreateLabel(pnlEventsDisp, "lblCap6", 12 + 20 + 50 + 50 + 50 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Body", nil)
	
	paginatorY := 0
	for i:=0; i < 5; i++ {
		eventsTable[i][0] = CreateLabel(pnlEventsDisp, "", 12, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][1] = CreateLabel(pnlEventsDisp, "", 12 + 20, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][2] = CreateLabel(pnlEventsDisp, "", 12 + 20 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][3] = CreateLabel(pnlEventsDisp, "", 12 + 20 + 50 + 30, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][4] = CreateLabel(pnlEventsDisp, "", 12 + 20 + 50 + 50 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000,  "", nil)
		eventsTable[i][5] = CreateLabel(pnlEventsDisp, "", 12 + 20 + 50 + 50 + 50 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][6] = CreateBtn(pnlEventsDisp, "", 12 + 20 + 50 + 50 + 50 + 50 + 30, 22*(i+3), 40, 20, 0xD8DCC0, 0x000000, "Upd", nil)
		eventsTable[i][7] = CreateBtn(pnlEventsDisp, "", 12 + 20 + 50 + 50 + 50 + 50 + 30 + 44, 22*(i+3), 40, 20, 0xD8DCC0, 0x000000, "Del", nil)
	}

	btnPrevEvents = CreateBtn(pnlEventsDisp, "btnPrevEvents", 12 + 50, 22*(paginatorY+4), 40, 20, 0xD8DCC0, 0x000000, "Prev", btnPrevEventsClick)
	//btnPrevEvents.obj.(*tBtn).enabled = false
	lblCurEventsPage = CreateLabel(pnlEventsDisp, "lblCurEventsPage", 12 + 50 + 50, 22*(paginatorY+4), 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurEventsPage), nil)
	btnNextEvents = CreateBtn(pnlEventsDisp, "btnNextEvents", 12 + 50 + 50 + 15, 22*(paginatorY+4), 40, 20, 0xD8DCC0, 0x000000, "Next", btnNextEventsClick)
	
	memTest = CreateMemo(pnlEventsDisp, "memTest", 400, 30, 100, 100, 0x000000, 0xF8FCF8, "", nil)
	
	refreshEventsTable()
	
	
	btnAddUser = CreateBtn(pnlUsersDisp, "btnAddUser", 12, 22, 60, 20, 0xD8DCC0, 0x000000, "Add", nil)
	btnRefreshUser = CreateBtn(pnlUsersDisp, "btnRefreshUser", 12 + 64, 22, 60, 20, 0xD8DCC0, 0x000000, "Refresh", btnRefreshClick)
	
	lblId = CreateLabel(pnlUsersDisp, "lblId", 12, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Id", nil)
	lblUserName = CreateLabel(pnlUsersDisp, "lblUserName", 12 + 20, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Name", nil)
	lblLogin = CreateLabel(pnlUsersDisp, "lblLogin", 12 + 20 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Login", nil)
	lblUserPswd = CreateLabel(pnlUsersDisp, "lblUserPswd", 12 + 20 + 50 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Pswd", nil)
	lblUserRole = CreateLabel(pnlUsersDisp, "lblUserRole", 12 + 20 + 50 + 50 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Role", nil)
	
	paginatorY = 0
	for i:=0; i < 5; i++ {
		usersTable[i][0] = CreateLabel(pnlUsersDisp, "", 12, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		usersTable[i][1] = CreateLabel(pnlUsersDisp, "", 12 + 20, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		usersTable[i][2] = CreateLabel(pnlUsersDisp, "", 12 + 20 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		usersTable[i][3] = CreateLabel(pnlUsersDisp, "", 12 + 20 + 50 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		usersTable[i][4] = CreateLabel(pnlUsersDisp, "", 12 + 20 + 50 + 50 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000,  "", nil)
		usersTable[i][5] = CreateBtn(pnlUsersDisp, "", 12 + 20 + 50 + 50 + 50 + 30, 22*(i+3), 40, 20, 0xD8DCC0, 0x000000, "Upd", nil)
		usersTable[i][6] = CreateBtn(pnlUsersDisp, "", 12 + 20 + 50 + 50 + 50 + 30 + 44, 22*(i+3), 40, 20, 0xD8DCC0, 0x000000, "Del", nil)
	}

	btnPrev = CreateBtn(pnlUsersDisp, "btnPrev", 12 + 50, 22*(paginatorY+4), 40, 20, 0xD8DCC0, 0x000000, "Prev", btnPrevClick)
	btnPrev.obj.(*tBtn).enabled = false
	lblCurPage = CreateLabel(pnlUsersDisp, "lblCurPage", 12 + 50 + 50, 22*(paginatorY+4), 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurUsersPage), nil)
	btnNext = CreateBtn(pnlUsersDisp, "btnNext", 12 + 50 + 50 + 15, 22*(paginatorY+4), 40, 20, 0xD8DCC0, 0x000000, "Next", btnNextClick)
	
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
	refreshEventsTable()
}


func btnNextEventsClick(node *Node){
	CurEventsPage++
	btnPrevEvents.obj.(*tBtn).enabled = true
	refreshEventsTable()
}


func refreshEventsTable(){
	for i:=0; i < 5; i++ {
		eventsTable[i][0].obj.(*tLabel).caption = ""
		eventsTable[i][1].obj.(*tLabel).caption = ""
		eventsTable[i][2].obj.(*tLabel).caption = ""
		eventsTable[i][3].obj.(*tLabel).caption = ""
		eventsTable[i][4].obj.(*tLabel).caption = ""
		eventsTable[i][5].obj.(*tLabel).caption = ""
		eventsTable[i][6].obj.(*tBtn).visible = false
		eventsTable[i][7].obj.(*tBtn).visible = false
	}
	
	result := Get("http://localhost:8085/api", "cmd=get_evnt " + strconv.Itoa(CurEventsPage) + " 5", "")
	
	var events []EventsFromDB 
	err := json.Unmarshal([]byte(result), &events)
	if err != nil {
		fmt.Println(err)
	}
	
	paginatorY := 0
	for i, _ := range events {
		eventsTable[i][0].obj.(*tLabel).caption = strconv.Itoa(events[i].Id)
		eventsTable[i][1].obj.(*tLabel).caption = events[i].Level
		eventsTable[i][2].obj.(*tLabel).caption = strconv.Itoa(events[i].Obj_id)
		eventsTable[i][3].obj.(*tLabel).caption = events[i].Source
		eventsTable[i][4].obj.(*tLabel).caption = events[i].Event
		eventsTable[i][5].obj.(*tLabel).caption = events[i].Body
		eventsTable[i][6].obj.(*tBtn).visible = true
		eventsTable[i][7].obj.(*tBtn).visible = true
		
		paginatorY = i
	}
	
	btnPrevEvents.obj.(*tBtn).y = 22*(paginatorY+4)
	lblCurEventsPage.obj.(*tLabel).y = 22*(paginatorY+4)
	lblCurEventsPage.obj.(*tLabel).caption = strconv.Itoa(CurEventsPage)
	btnNextEvents.obj.(*tBtn).y = 22*(paginatorY+4)
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
	refreshUsersTable()
}


func btnNextClick(node *Node){
	CurUsersPage++
	btnPrev.obj.(*tBtn).enabled = true
	refreshUsersTable()
}


func refreshUsersTable(){

	for i:=0; i < 5; i++ {
		usersTable[i][0].obj.(*tLabel).caption = ""
		usersTable[i][1].obj.(*tLabel).caption = ""
		usersTable[i][2].obj.(*tLabel).caption = ""
		usersTable[i][3].obj.(*tLabel).caption = ""
		usersTable[i][4].obj.(*tLabel).caption = ""
		usersTable[i][5].obj.(*tBtn).visible = false
		usersTable[i][6].obj.(*tBtn).visible = false
	}
	
	result := Get("http://localhost:8085/api", "cmd=get_usr " + strconv.Itoa(CurUsersPage) + " 5", "")
	
	var users []UsersFromDB 
	err := json.Unmarshal([]byte(result), &users)
	if err != nil {
		fmt.Println(err)
	}
	
	paginatorY := 0
	for i, _ := range users {
		usersTable[i][0].obj.(*tLabel).caption = strconv.Itoa(users[i].Id)
		usersTable[i][1].obj.(*tLabel).caption = users[i].UserName
		usersTable[i][2].obj.(*tLabel).caption = users[i].Login
		usersTable[i][3].obj.(*tLabel).caption = users[i].Pswd
		usersTable[i][4].obj.(*tLabel).caption = strconv.Itoa(users[i].UserRole)
		usersTable[i][5].obj.(*tBtn).visible = true
		usersTable[i][6].obj.(*tBtn).visible = true
		
		paginatorY = i
		fmt.Println(users[i].Id, users[i].UserName, users[i].Login, users[i].UserRole)
	}
	
	btnPrev.obj.(*tBtn).y = 22*(paginatorY+4)
	lblCurPage.obj.(*tLabel).y = 22*(paginatorY+4)
	lblCurPage.obj.(*tLabel).caption = strconv.Itoa(CurUsersPage)
	btnNext.obj.(*tBtn).y = 22*(paginatorY+4)
}
