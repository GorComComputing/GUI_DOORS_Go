package main

import (
    "fmt"
    "math/rand"
    //"math"
    //"syscall/js"
    //"time"
    "strconv"
    //"strings"
    //"net/http"
    //"io"
    //"bytes"
    //"io/ioutil"
    //"encoding/json"
    "time"
    //"encoding/base64"
    //"encoding/hex"
    //_ "github.com/stealthrocket/net/http"
    //"github.com/Nigel2392/requester"
    //"encoding/ascii85"
    //"unicode/utf8"
)


type tProgram struct {
	name string 
    start func(*Node)
    picture *[]byte
}


var process []*tProc

type tProc struct {
	name string 
    form *Node 
    btn  *Node
    isRun bool
    picture []byte
}


// Инициализация DOORS 
func initDOORS(programs []*tProgram) {
	
	fmt.Println("👋 DOORS started OK! 🌍")
	
	//Декодируем base64 в bmp
  	bmpFolder = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Folder.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Folder)
  	bmpFile = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/File.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64File)
  	bmpFolder_small = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Folder_small.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Folder_small)
  	bmpFile_small = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/File_small.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64File_small)
  	bmpProgram = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Program.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Program)
	bmpNotepad = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Notepad.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Notepad)
  	bmpForm_close = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Form_close.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Form_close)
  	bmpComboBox = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/ComboBox.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64ComboBox)
  	bmpLogo_menu = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Logo_menu.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Logo_menu)
  	bmpNew_file = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/New_file.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64New_file)
  	bmpOpen_file = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Open_file.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Open_file)
  	bmpSave_file = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Save_file.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Save_file)
  	bmpUp = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Up.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Up)
  	bmpHelp = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Help.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Help)
  	bmpBrowser = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Browser.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Browser)
  	bmpPrograms = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Programs.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Programs)
  	bmpSettings = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Settings.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Settings)
  	bmpBack = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Back.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Back)
  	bmpNext = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Next.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Next)
  	bmpAdd = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Add.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Add)
  	bmpAdd2 = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Add2.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Add2)
  	bmpRefresh = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Refresh.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64Refresh)
  	bmpDorFile = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/dor.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64DorFile)
  	bmpGoFile = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/go.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))
  	bmpCFile = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/c.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64CFile)
  	bmpHTMLFile = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/html.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64HTMLFile)
  	bmpASMFile = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/asm.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))//base64.StdEncoding.DecodeString(b64ASMFile)
	bmpWebCam1 = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/pic2.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))
	bmpWebCam2 = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/poll.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))
	bmpFieldDesktop = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/windows.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))
	//bmpSkyFlag = fetchFile("http://www.gorcom.online/img/Sky3.bmp")
	bmpCamera = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Camera.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))
	bmpKey = fetchFile("http://" + ServerIP + ":" + ServerPort + "/img/Key.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))
	//bmpKey = fetchFileJPEG("http://www.gorcom.online/img/Key.jpg?rnd=" + strconv.Itoa(rand.Intn(999999)))
	

	
	
	/*fmt.Println("HERE 1")
	testJPEG("http://www.gorcom.online/img/Key.jpg?rnd=" + strconv.Itoa(rand.Intn(999999)))
	fmt.Println("HERE 2")*/
	
	startDesktop()
	startRAD()
	// Запуск программ
	for i := 0; i < len(programs); i++ {
		startProcess(programs[i].name, programs[i].start, programs[i].picture) 
	} 	
}


var xTask int = 2 + 71
func startProcess(name string, onStart func(*Node), picture *[]byte){
	/*obj := tBtn{name: "btnTask"+name, x: xTask, y: 2, sizeX: 80, sizeY: 28 - 4, BC: 0xD8DCC0, TC: 0x000000, caption: name, visible: true, pressed: false, enabled: true, onClick: btnTaskClick}
	node := Node{typ: BUTTON, parent: pnlTask, previous: nil, children: nil, obj: &obj}
	pnlTask.children = append(pnlTask.children, &node)*/
	var pic []byte
	if picture != nil {
		pic = *picture
	}
	frm := CreateForm(&layout, "frm" + name, pic, 400, 400, 200, 130, 0xD8DCC0, WIN, name, true, nil)
	frm.obj.(*tForm).visible = false
	
	proc := tProc{name: name, form: frm, btn: nil, isRun: false, picture: pic} //, btn: &node
	process = append(process, &proc)
	//xTask += 81
	//layout.children[len(layout.children)-2].obj.(*tForm).focused = false
	//layout.children[len(layout.children)-1].obj.(*tForm).focused = true
	onStart(frm)
}


func btnTaskClick(node *Node){
	var i int = 0
	for ; i < len(process); i++ {
		if node == process[i].btn {
			process[i].form.obj.(*tForm).visible = true //!(process[i].form.obj.(*tForm).visible)
			i := findNode(process[i].form)
			fmt.Println(i)
			if i > 0 {
				sortChildren(i)
			}
			break
		}
	}
}


func onTimer() {
	if cnvFlag != nil {
		flagDraw(cnvFlag.obj.(*tCanvas).x+50, cnvFlag.obj.(*tCanvas).y+50)
	}
	
	cursor = !(cursor)
	
	t := time.Now()
	lblTime.obj.(*tLabel).caption = strconv.Itoa(t.Hour()) + ":" + fmt.Sprintf("%02d", t.Minute())
	
	//getPhoto()

	
	
	//SetColor(0xFFFF00)
	//LinePP(cnvFlag.obj, 10, 10, 100, 100)
	//Circle(cnvFlag.obj, 50, 50, 30)
	
	
    
/*
	canvas := js.Global().Get("document").Call("getElementById", "cnvs1")

	context := canvas.Call("getContext", "2d")
	
	context.Set("fillStyle", "#ffc107")
	//context.Call("clearRect", 50, 50, 100, 100)
	context.Call("fillRect", 50, 50, 50, 50)

	ctx.Set("fillStyle", grd)
	ctx.Set("strokeStyle", "#FF0000")
	ctx.Call("fillRect", 0, 0, 40, 80)
	ctx.Call("fillText", "Fill text", 20, 50)
*/
}


func execProcess(num int) {
	if !(process[num].isRun) {
		process[num].isRun = true
		process[num].form.obj.(*tForm).visible = true
	
		obj := tBitBtn{name: "btnTask"+process[num].name, x: xTask, y: 2, sizeX: 100, sizeY: 28 - 4, BC: 0xD8DCC0, TC: 0x000000, caption: process[num].name, visible: true, pressed: false, enabled: true, picture: process[num].picture, mode: BORDER, onClick: btnTaskClick}
		obj.inalign = LEFT
		node_new := Node{typ: BIT_BUTTON, parent: pnlTask, previous: nil, children: nil, obj: &obj}
		pnlTask.children = append(pnlTask.children, &node_new)
		//obj.pressed = true
	
		process[num].btn = &node_new
		xTask += 101
		layout.children[len(layout.children)-2].obj.(*tForm).focused = false
		process[num].form.obj.(*tForm).focused = true
	
		i := findNode(process[num].form)
		fmt.Println(i)
		if i > 0 {
			sortChildren(i)
		}	
	}
}
