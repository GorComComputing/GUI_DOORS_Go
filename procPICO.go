package main

import (
	//"strings"
	//"strconv"
	//"fmt"
	//"time"
	"encoding/json"
	//"syscall/js"
)

var btnGPS *Node
var btnTemperatureIN *Node
var btnHumidity *Node
var btnLight *Node
var edtServo *Node
var btnServo *Node
var btnFRW *Node
var btnBCK *Node
var btnLFT *Node
var btnBLFT *Node
var btnRGH *Node
var btnBRGH *Node
var lblCar *Node
var pnlLeds *Node
var cbxLED *Node
var cbxRED *Node
var cbxYELLOW *Node
var cbxGREEN *Node
var cbxLASER *Node
var cbxPUMP *Node
var lblIPAddr *Node
var edtIPAddr *Node
var pnlData *Node
var lblTempOut *Node
var lblTempOutData *Node
var lblTempIn *Node
var lblTempInData *Node
var lblHum *Node
var lblHumData *Node
var lblLight *Node
var lblLightData *Node
var lblLon *Node
var lblLonData *Node
var lblLat *Node
var lblLatData *Node
var lblTimeGPS *Node
var lblTimeGPSData *Node
var lblDate *Node
var lblDateData *Node
var lblMoisture *Node
var lblMoistureData *Node
var btnRefreshData *Node

var PICO_IP string = "192.168.0.14"
var cnvCam bool = false


func startPICO(frmMain *Node){ 
    frmMain.obj.(*tForm).x = BITMAP_WIDTH/2 - frmMain.obj.(*tForm).sizeX/2
	frmMain.obj.(*tForm).y = BITMAP_HEIGHT/2 - frmMain.obj.(*tForm).sizeY/2
    setSize(frmMain, 625, 250)
    
    frmMain.children[1].obj.(*tBitBtn).enabled = false
    
    frmMain.obj.(*tForm).BC = 0xd8dcc0

    btnGPS = CreateBtn(frmMain, "btnGPS", 10, 24, 130, 24, 0xd8dcc0, 0x0, "GPS", btnGPSClick)
    btnTemperatureIN = CreateBtn(frmMain, "btnTemperatureIN", 10, 56, 130, 24, 0xd8dcc0, 0x0, "Temperature IN", btnTemperatureINClick)
    btnHumidity = CreateBtn(frmMain, "btnHumidity", 10, 88, 130, 24, 0xd8dcc0, 0x0, "Humidity", btnHumidityClick)
    btnLight = CreateBtn(frmMain, "btnLight", 10, 120, 130, 24, 0xd8dcc0, 0x0, "Light", btnLightClick)
    
    edtServo = CreateEdit(frmMain, "edtServo", 108, 156, 30, 20, 0xf8fcf8, 0x0, "", nil, nil)
    btnServo = CreateBtn(frmMain, "btnServo", 10, 154, 90, 24, 0xd8dcc0, 0x0, "Servo", btnServoClick)
    
    btnFRW = CreateBtn(frmMain, "btnFRW", 318, 64, 50, 24, 0xd8dcc0, 0x0, "^", btnFRWClick)
    btnBCK = CreateBtn(frmMain, "btnBCK", 318, 132, 50, 24, 0xd8dcc0, 0x0, "v", btnBCKClick)
    btnLFT = CreateBtn(frmMain, "btnLFT", 270, 98, 50, 24, 0xd8dcc0, 0x0, "<", btnLFTClick)
    btnBLFT = CreateBtn(frmMain, "btnBLFT", 270, 166, 50, 24, 0xd8dcc0, 0x0, "<-bck", btnBLFTClick)
    btnRGH = CreateBtn(frmMain, "btnRGH", 368, 98, 50, 24, 0xd8dcc0, 0x0, ">", btnRGHClick)
    btnBRGH = CreateBtn(frmMain, "btnBRGH", 368, 166, 50, 24, 0xd8dcc0, 0x0, "bck->", btnBRGHClick)
    lblCar = CreateLabel(frmMain, "lblCar", 332, 98, 30, 20, 0xd8dcc0, 0x0, "CAR", nil)
    
    pnlLeds = CreatePanel(frmMain, "pnlLeds", 150, 24, 120, 155, 0xd8dcc0, NONE, nil)
    cbxLED = CreateCheckBox(pnlLeds, "cbxLED", 5, 5, 110, 20, 0xD8DCC0, 0x000000, "LED", false, cbxLEDClick)
    cbxRED = CreateCheckBox(pnlLeds, "cbxRED", 5, 30, 110, 20, 0xD8DCC0, 0x000000, "RED", false, cbxREDClick)
    cbxYELLOW = CreateCheckBox(pnlLeds, "cbxYELLOW", 5, 55, 110, 20, 0xD8DCC0, 0x000000, "YELLOW", false, cbxYELLOWClick)
    cbxGREEN = CreateCheckBox(pnlLeds, "cbxGREEN", 5, 80, 110, 20, 0xD8DCC0, 0x000000, "GREEN", false, cbxGREENClick)
    cbxLASER = CreateCheckBox(pnlLeds, "cbxLASER", 5, 105, 110, 20, 0xD8DCC0, 0x000000, "LASER", false, cbxLASERClick)
    cbxPUMP = CreateCheckBox(pnlLeds, "cbxPUMP", 5, 130, 110, 20, 0xD8DCC0, 0x000000, "PUMP", false, cbxPUMPClick)
    
  	lblIPAddr = CreateLabel(frmMain, "lblIPAddr", 56, 218, 90, 20, 0xd8dcc0, 0x0, "IP address:", nil)
  	edtIPAddr = CreateEdit(frmMain, "edtIPAddr", 150, 220, 120, 20, 0xf8fcf8, 0x0, PICO_IP, nil, nil)
  	
  	pnlData = CreatePanel(frmMain, "pnlData", 440, 24, 180, 220, 0xd8dcc0, NONE, nil)
  	lblTempOut = CreateLabel(pnlData, "lblTempOut", 5, 5, 120, 20, 0xd8dcc0, 0x0, "Temp Out:", nil)
  	lblTempOutData = CreateLabel(pnlData, "lblTempOutData", 80, 5, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	lblTempIn = CreateLabel(pnlData, "lblTempIn", 5, 25, 120, 20, 0xd8dcc0, 0x0, "Temp In:", nil)
  	lblTempInData = CreateLabel(pnlData, "lblTempInData", 80, 25, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	lblHum = CreateLabel(pnlData, "lblHum", 5, 45, 120, 20, 0xd8dcc0, 0x0, "Humidity:", nil)
  	lblHumData = CreateLabel(pnlData, "lblHumData", 80, 45, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	lblLight = CreateLabel(pnlData, "lblLight", 5, 65, 120, 20, 0xd8dcc0, 0x0, "Light:", nil)
  	lblLightData = CreateLabel(pnlData, "lblLightData", 80, 65, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	lblLon = CreateLabel(pnlData, "lblLon", 5, 105, 120, 20, 0xd8dcc0, 0x0, "Lon:", nil)
  	lblLonData = CreateLabel(pnlData, "lblLonData", 80, 105, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	lblLat = CreateLabel(pnlData, "lblLat", 5, 85, 120, 20, 0xd8dcc0, 0x0, "Lat:", nil)
  	lblLatData = CreateLabel(pnlData, "lblLatData", 80, 85, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	lblTimeGPS = CreateLabel(pnlData, "lblTime", 5, 125, 120, 20, 0xd8dcc0, 0x0, "Time:", nil)
  	lblTimeGPSData = CreateLabel(pnlData, "lblTimeData", 80, 125, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	lblDate = CreateLabel(pnlData, "lblDate", 5, 145, 120, 20, 0xd8dcc0, 0x0, "Date:", nil)
  	lblDateData = CreateLabel(pnlData, "lblDateData", 80, 145, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	lblMoisture = CreateLabel(pnlData, "lblMoisture", 5, 165, 120, 20, 0xd8dcc0, 0x0, "Moisture:", nil)
  	lblMoistureData = CreateLabel(pnlData, "lblMoistureData", 80, 165, 80, 20, 0xd8dcc0, 0x0, "", nil)
  	btnRefreshData = CreateBtn(pnlData, "btnRefreshData", 50, 190, 70, 24, 0xd8dcc0, 0x0, "Refresh", btnRefreshDataClick)
  	
  	//btnRefreshDataClick(btnRefreshData)
}


func btnGPSClick(node *Node){
	//result := Get("http://"+PICO_IP+"/post", "cmd=gps", "")
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=gps", "")
}


func btnTemperatureINClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=tempin", "")
}


func btnHumidityClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=hum", "")
}


func btnLightClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=light", "")
}


func btnServoClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=servo " + edtServo.obj.(*tEdit).text, "")
}


func btnFRWClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=car forward", "")
}


func btnBCKClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=car back", "")
}


func btnLFTClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=car left", "")
}


func btnBLFTClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=car bckleft", "")
}


func btnRGHClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=car right", "")
}


func btnBRGHClick(node *Node){
	Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=car bckright", "")
}


func cbxLEDClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=led off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=led on", "")
	}
}


func cbxREDClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=red off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=red on", "")
	}
}


func cbxYELLOWClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=yellow off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=yellow on", "")
	}
}


func cbxGREENClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=green off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=green on", "")
	}
}


func cbxLASERClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=laser off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=laser on", "")
	}
}


func cbxPUMPClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		//Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=laser off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		//Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=laser on", "")
	}
}


func btnRefreshDataClick(node *Node){
	result := Get("http://" + edtIPAddr.obj.(*tEdit).text + "/get", "cmd=state", "")
	//fmt.Printf(result)
	
	//fmt.Printf("Test")
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(result), &dat); err != nil {
    	//panic(err)
    }
    lblTempInData.obj.(*tLabel).caption = dat["TempIn"].(string) + " C"
    lblHumData.obj.(*tLabel).caption = dat["Hum"].(string) + " %"
	lblLightData.obj.(*tLabel).caption = dat["Light"].(string)
	lblLatData.obj.(*tLabel).caption = dat["Lat"].(string)
	lblLonData.obj.(*tLabel).caption = dat["Lon"].(string)
	lblTimeGPSData.obj.(*tLabel).caption = dat["Time"].(string)
	lblDateData.obj.(*tLabel).caption = dat["Date"].(string)
	lblTempOutData.obj.(*tLabel).caption = dat["TempOut"].(string) + " C"
	lblMoistureData.obj.(*tLabel).caption = dat["Moisture"].(string) + " %"
	edtServo.obj.(*tEdit).text = dat["Servo"].(string)

	if dat["Led"].(string) == "ON" {
		cbxLED.obj.(*tCheckBox).checked = true
	} else if dat["Led"].(string) == "OFF" {
		cbxLED.obj.(*tCheckBox).checked = false
	}
	
	if dat["Red"].(string) == "ON" {
		cbxRED.obj.(*tCheckBox).checked = true
	} else if dat["Red"].(string) == "OFF" {
		cbxRED.obj.(*tCheckBox).checked = false
	}
	
	if dat["Yellow"].(string) == "ON" {
		cbxYELLOW.obj.(*tCheckBox).checked = true
	} else if dat["Yellow"].(string) == "OFF" {
		cbxYELLOW.obj.(*tCheckBox).checked = false
	}
	
	if dat["Green"].(string) == "ON" {
		cbxGREEN.obj.(*tCheckBox).checked = true
	} else if dat["Green"].(string) == "OFF" {
		cbxGREEN.obj.(*tCheckBox).checked = false
	}
	
	if dat["Laser"].(string) == "ON" {
		cbxLASER.obj.(*tCheckBox).checked = true
	} else if dat["Laser"].(string) == "OFF" {
		cbxLASER.obj.(*tCheckBox).checked = false
	}
	
	if dat["Pump"].(string) == "ON" {
		cbxPUMP.obj.(*tCheckBox).checked = true
	} else if dat["Pump"].(string) == "OFF" {
		cbxPUMP.obj.(*tCheckBox).checked = false
	}
}






