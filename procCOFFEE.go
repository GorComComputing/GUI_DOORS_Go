package main

import (
	//"strings"
	//"strconv"
	//"fmt"
	//"time"
	"encoding/json"
)


var edtServoX *Node
var btnServoX *Node
var edtServoY *Node
var btnServoY *Node
var pnlLeds_COFFEE  *Node
var cbxRED_COFFEE *Node
var cbxPUMP_COFFEE *Node
var cbxTRACK *Node
var cbxFUN *Node
var cbxRELAY *Node
var lblIPAddr_COFFEE *Node
var edtIPAddr_COFFEE *Node
var pnlData_COFFEE  *Node
/*var lblTempOut *Node
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
var lblMoistureData *Node*/
var btnRefreshData_COFFEE *Node
var trkServoX *Node

var COFFEE_IP string = "192.168.0.12"


func startCOFFEE(frmMain *Node){ 
    frmMain.obj.(*tForm).x = BITMAP_WIDTH/2 - frmMain.obj.(*tForm).sizeX/2
	frmMain.obj.(*tForm).y = BITMAP_HEIGHT/2 - frmMain.obj.(*tForm).sizeY/2
    setSize(frmMain, 464, 250)
    
    frmMain.children[1].obj.(*tBitBtn).enabled = false
    
    frmMain.obj.(*tForm).BC = 0xd8dcc0
    
    edtServoX = CreateEdit(frmMain, "edtServoX", 108, 26, 30, 20, 0xf8fcf8, 0x0, "90", nil, nil)
    btnServoX = CreateBtn(frmMain, "btnServoX", 10, 24, 90, 24, 0xd8dcc0, 0x0, "Servo X", btnServoXClick)
    edtServoY = CreateEdit(frmMain, "edtServoY", 108, 58, 30, 20, 0xf8fcf8, 0x0, "90", nil, nil)
    btnServoY = CreateBtn(frmMain, "btnServoY", 10, 56, 90, 24, 0xd8dcc0, 0x0, "Servo Y", btnServoYClick)
    
    //trkServoX = CreateTrackBar(frmMain, "trkServoX", 10, 96, 120, 20, 0xFFFFFF, 0x0, "90", nil)
    
    pnlLeds_COFFEE = CreatePanel(frmMain, "pnlLeds_COFFEE", 150, 24, 120, 155, 0xd8dcc0, NONE, nil)
    cbxRED_COFFEE = CreateCheckBox(pnlLeds_COFFEE, "cbxRED_COFFEE", 5, 5, 110, 20, 0xD8DCC0, 0x000000, "RED", false, cbxRED_COFFEE_Click)
    cbxPUMP_COFFEE = CreateCheckBox(pnlLeds_COFFEE, "cbxPUMP_COFFEE", 5, 30, 110, 20, 0xD8DCC0, 0x000000, "PUMP", false, cbxPUMP_COFFEE_Click)
    cbxTRACK = CreateCheckBox(pnlLeds_COFFEE, "cbxTRACK", 5, 55, 110, 20, 0xD8DCC0, 0x000000, "TRACK", false, cbxTRACKClick)
    cbxFUN = CreateCheckBox(pnlLeds_COFFEE, "cbxFUN", 5, 80, 110, 20, 0xD8DCC0, 0x000000, "FUN", false, cbxFUNClick)
    cbxRELAY = CreateCheckBox(pnlLeds_COFFEE, "cbxRELAY", 5, 105, 110, 20, 0xD8DCC0, 0x000000, "RELAY", false, cbxRELAYClick)
    
  	lblIPAddr_COFFEE = CreateLabel(frmMain, "lblIPAddr_COFFEE", 56, 218, 90, 20, 0xd8dcc0, 0x0, "IP address:", nil)
  	edtIPAddr_COFFEE = CreateEdit(frmMain, "edtIPAddr_COFFEE", 150, 220, 120, 20, 0xf8fcf8, 0x0, COFFEE_IP, nil, nil)
  	
  	pnlData_COFFEE = CreatePanel(frmMain, "pnlData_COFFEE", 280, 24, 180, 220, 0xd8dcc0, NONE, nil)
  	/*lblTempOut = CreateLabel(pnlData, "lblTempOut", 5, 5, 120, 20, 0xd8dcc0, 0x0, "Temp Out:", nil)
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
  	lblMoistureData = CreateLabel(pnlData, "lblMoistureData", 80, 165, 80, 20, 0xd8dcc0, 0x0, "", nil)*/
  	btnRefreshData_COFFEE = CreateBtn(pnlData_COFFEE, "btnRefreshData_COFFEE", 50, 190, 70, 24, 0xd8dcc0, 0x0, "Refresh", btnRefreshData_COFFEE_Click)
  	
  	//btnRefreshData_COFFEE_Click(btnRefreshData_COFFEE)
}



func btnServoXClick(node *Node){
	Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=servox " + edtServoX.obj.(*tEdit).text, "")
}

func btnServoYClick(node *Node){
	Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=servoy " + edtServoY.obj.(*tEdit).text, "")
}


func cbxRED_COFFEE_Click(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=red off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=red on", "")
	}
}


func cbxPUMP_COFFEE_Click(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=pump off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=pump on", "")
	}
}


func cbxTRACKClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=track off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=track on", "")
	}
}


func cbxFUNClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=fun off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=fun on", "")
	}
}


func cbxRELAYClick(node *Node){
	if(node.obj.(*tCheckBox).checked){
		node.obj.(*tCheckBox).checked = false
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=relay off", "")
	} else {
		node.obj.(*tCheckBox).checked = true
		Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=relay on", "")
	}
}


func btnRefreshData_COFFEE_Click(node *Node){
	result := Get("http://" + edtIPAddr_COFFEE.obj.(*tEdit).text + "/get", "cmd=state", "")
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(result), &dat); err != nil {
    	panic(err)
    }
	edtServoX.obj.(*tEdit).text = dat["ServoX"].(string)
	edtServoY.obj.(*tEdit).text = dat["ServoY"].(string)
	
	if dat["Red"].(string) == "ON" {
		cbxRED_COFFEE.obj.(*tCheckBox).checked = true
	} else if dat["Red"].(string) == "OFF" {
		cbxRED_COFFEE.obj.(*tCheckBox).checked = false
	}
	
	if dat["Pump"].(string) == "ON" {
		cbxPUMP_COFFEE.obj.(*tCheckBox).checked = true
	} else if dat["Pump"].(string) == "OFF" {
		cbxPUMP_COFFEE.obj.(*tCheckBox).checked = false
	}
	
	if dat["Track"].(string) == "ON" {
		cbxTRACK.obj.(*tCheckBox).checked = true
	} else if dat["Track"].(string) == "OFF" {
		cbxTRACK.obj.(*tCheckBox).checked = false
	}
	
	if dat["Fun"].(string) == "ON" {
		cbxFUN.obj.(*tCheckBox).checked = true
	} else if dat["Fun"].(string) == "OFF" {
		cbxFUN.obj.(*tCheckBox).checked = false
	}
	
	if dat["Relay"].(string) == "ON" {
		cbxRELAY.obj.(*tCheckBox).checked = true
	} else if dat["Relay"].(string) == "OFF" {
		cbxRELAY.obj.(*tCheckBox).checked = false
	}	
}






