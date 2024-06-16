package main

import (
    //"fmt"
    //"math/rand"
    //"math"
    //"syscall/js"
    //"time"
    //"strconv"
    //"net/http"
    //"io"
    //"math/rand"
)


var pnlCam1 *Node
var btnPhoto *Node
var imgCam *Node


func startCam1(frmMain *Node){
	setSize(frmMain, 644, 534)
	frmMain.obj.(*tForm).x = 220
	frmMain.obj.(*tForm).y = 110
	frmMain.obj.(*tForm).BC = 0xd8dcc0;
	
	frmMain.obj.(*tForm).visible = false
	
	pnlCam1 = CreatePanel(frmMain, "pnlCam1", 2, 18, 640, 480, 0x111111, NONE, nil)
	btnPhoto = CreateBtn(frmMain, "btnPhoto", 286, 504, 70, 24, 0xd8dcc0, 0x0, "Photo", btnPhotoClick)
	imgCam = CreateImage(pnlCam1, "imgCam", bmpWebCam1, 0, 0, 640, 480, nil)
}


func btnPhotoClick(node *Node){
	//getPhoto()
	//bmpWebCam = fetchFile("http://www.gorcom.online/img/pic.bmp")
	//bmpGoFile = fetchFile("http://www.gorcom.online/img/go.bmp?rnd=" + strconv.Itoa(rand.Intn(999999)))
}

/*
func getPhoto(){
	//var i int = 0
	//for ; i < len(process); i++ {
	//	if process[i].form.obj.(*tForm).visible && process[i].name == "Camera 1" {
			//result := js.Global().Call("getCamTime", "").Get("response").String()
			//js.Global().Call("HttpRequest", "http://"+ServerIP+":"+ServerPort+"/api?cmd=get_cam " + result, "").Get("response").String() 
	//		bmpWebCam = nil
			bmpWebCam = fetchFile("http://www.gorcom.online/img/pic.bmp")
			//js.Global().Call("camDraw", process[i].form.obj.(*tForm).x + 2, process[i].form.obj.(*tForm).y + 19)
			//showBMP(nil, bmpWebCam, process[i].form.obj.(*tForm).x + 2, process[i].form.obj.(*tForm).y + 19)
	//		break
	//	}
	//}
}*/
