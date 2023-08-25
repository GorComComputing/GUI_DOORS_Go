package main

import (
    "fmt"
    //"math/rand"
    //"math"
    "syscall/js"
    //"time"
    "strconv"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    "encoding/json"
)


//var frmEvents *Node
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


func startEvents(frmMain *Node){
	//frmEvents = CreateForm(&layout, 100, 100, 500, 220, 0xD8DCC0, WIN, "Events", false, nil)
	frmMain.obj.(*tForm).x = 100
	frmMain.obj.(*tForm).y = 100
	frmMain.obj.(*tForm).sizeX = 600
	frmMain.obj.(*tForm).sizeY = 220
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	frmMain.obj.(*tForm).visible = false
	
					// 472, 202
	btnAddEvent = CreateBtn(frmMain, "btnAddEvent", 12, 22, 60, 20, 0xD8DCC0, 0x000000, "Add", nil)
	btnRefreshEvents = CreateBtn(frmMain, "btnRefreshEvents", 12 + 64, 22, 60, 20, 0xD8DCC0, 0x000000, "Refresh", btnRefreshEventsClick)
	
	lblCap1 = CreateLabel(frmMain, "lblCap1", 12, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Id", nil)
	lblCap2 = CreateLabel(frmMain, "lblCap2", 12 + 20, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Level", nil)
	lblCap3 = CreateLabel(frmMain, "lblCap3", 12 + 20 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Obj", nil)
	lblCap4 = CreateLabel(frmMain, "lblCap4", 12 + 20 + 50 + 30, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Source", nil)
	lblCap5 = CreateLabel(frmMain, "lblCap5", 12 + 20 + 50 + 50 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Event", nil)
	lblCap6 = CreateLabel(frmMain, "lblCap6", 12 + 20 + 50 + 50 + 50 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Body", nil)
	
	paginatorY := 0
	for i:=0; i < 5; i++ {
		eventsTable[i][0] = CreateLabel(frmMain, "", 12, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][1] = CreateLabel(frmMain, "", 12 + 20, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][2] = CreateLabel(frmMain, "", 12 + 20 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][3] = CreateLabel(frmMain, "", 12 + 20 + 50 + 30, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][4] = CreateLabel(frmMain, "", 12 + 20 + 50 + 50 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000,  "", nil)
		eventsTable[i][5] = CreateLabel(frmMain, "", 12 + 20 + 50 + 50 + 50 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		eventsTable[i][6] = CreateBtn(frmMain, "", 12 + 20 + 50 + 50 + 50 + 50 + 30, 22*(i+3), 40, 20, 0xD8DCC0, 0x000000, "Upd", nil)
		eventsTable[i][7] = CreateBtn(frmMain, "", 12 + 20 + 50 + 50 + 50 + 50 + 30 + 44, 22*(i+3), 40, 20, 0xD8DCC0, 0x000000, "Del", nil)
	}

	btnPrevEvents = CreateBtn(frmMain, "btnPrevEvents", 12 + 50, 22*(paginatorY+4), 40, 20, 0xD8DCC0, 0x000000, "Prev", btnPrevEventsClick)
	//btnPrevEvents.obj.(*tBtn).enabled = false
	lblCurEventsPage = CreateLabel(frmMain, "lblCurEventsPage", 12 + 50 + 50, 22*(paginatorY+4), 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurEventsPage), nil)
	btnNextEvents = CreateBtn(frmMain, "btnNextEvents", 12 + 50 + 50 + 15, 22*(paginatorY+4), 40, 20, 0xD8DCC0, 0x000000, "Next", btnNextEventsClick)
	
	memTest = CreateMemo(frmMain, "memTest", 400, 30, 100, 100, 0x000000, 0xF8FCF8, "", nil)
	
	refreshEventsTable()
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
	
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=get_evnt " + strconv.Itoa(CurEventsPage) + " 5", "").Get("response").String()
	fmt.Println("Responsed: ", result)
	
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
