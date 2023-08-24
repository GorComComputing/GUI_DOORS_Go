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


//var frmUsers *Node
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


func startUsers(frmMain *Node){
	//frmUsers = CreateForm(&layout, 100, 100, 400, 220, 0xD8DCC0, WIN, "Users", false, nil)
	frmMain.obj.(*tForm).x = 150
	frmMain.obj.(*tForm).y = 150
	frmMain.obj.(*tForm).sizeX = 400
	frmMain.obj.(*tForm).sizeY = 220
	frmMain.children[0].obj.(*tBitBtn).x = frmMain.obj.(*tForm).sizeX - 17
	frmMain.obj.(*tForm).visible = false
	
	btnAddUser = CreateBtn(frmMain, "btnAddUser", 12, 22, 60, 20, 0xD8DCC0, 0x000000, "Add", nil)
	btnRefreshUser = CreateBtn(frmMain, "btnRefreshUser", 12 + 64, 22, 60, 20, 0xD8DCC0, 0x000000, "Refresh", btnRefreshClick)
	
	lblId = CreateLabel(frmMain, "lblId", 12, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Id", nil)
	lblUserName = CreateLabel(frmMain, "lblUserName", 12 + 20, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Name", nil)
	lblLogin = CreateLabel(frmMain, "lblLogin", 12 + 20 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Login", nil)
	lblUserPswd = CreateLabel(frmMain, "lblUserPswd", 12 + 20 + 50 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Pswd", nil)
	lblUserRole = CreateLabel(frmMain, "lblUserRole", 12 + 20 + 50 + 50 + 50, 22 + 22, 50, 20, 0xD8DCC0, 0x0000FF, "Role", nil)
	
	paginatorY := 0
	for i:=0; i < 5; i++ {
		usersTable[i][0] = CreateLabel(frmMain, "", 12, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		usersTable[i][1] = CreateLabel(frmMain, "", 12 + 20, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		usersTable[i][2] = CreateLabel(frmMain, "", 12 + 20 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		usersTable[i][3] = CreateLabel(frmMain, "", 12 + 20 + 50 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000, "", nil)
		usersTable[i][4] = CreateLabel(frmMain, "", 12 + 20 + 50 + 50 + 50, 22*(i+3), 50, 20, 0xD8DCC0, 0x000000,  "", nil)
		usersTable[i][5] = CreateBtn(frmMain, "", 12 + 20 + 50 + 50 + 50 + 30, 22*(i+3), 40, 20, 0xD8DCC0, 0x000000, "Upd", nil)
		usersTable[i][6] = CreateBtn(frmMain, "", 12 + 20 + 50 + 50 + 50 + 30 + 44, 22*(i+3), 40, 20, 0xD8DCC0, 0x000000, "Del", nil)
	}

	btnPrev = CreateBtn(frmMain, "btnPrev", 12 + 50, 22*(paginatorY+4), 40, 20, 0xD8DCC0, 0x000000, "Prev", btnPrevClick)
	btnPrev.obj.(*tBtn).enabled = false
	lblCurPage = CreateLabel(frmMain, "lblCurPage", 12 + 50 + 50, 22*(paginatorY+4), 20, 20, 0xD8DCC0, 0x0000FF, strconv.Itoa(CurUsersPage), nil)
	btnNext = CreateBtn(frmMain, "btnNext", 12 + 50 + 50 + 15, 22*(paginatorY+4), 40, 20, 0xD8DCC0, 0x000000, "Next", btnNextClick)
	
	refreshUsersTable()
}


func btnRefreshClick(node *Node){
	refreshUsersTable()
	
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=get_usr " + strconv.Itoa(CurUsersPage) + " 5").Get("response").String()
	memTerminal.obj.(*tMemo).text = result	
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
	
	result := js.Global().Call("HttpRequest", "http://localhost:8085/api?cmd=get_usr " + strconv.Itoa(CurUsersPage) + " 5").Get("response").String()
	fmt.Println("Responsed: ", result)
	
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
