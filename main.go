package main

import (
	//"fmt"
	"syscall/js"
	//"strings"
	//"unsafe"
	//"strconv"
)


// Объявление глобальных переменных
var (
	ServerIP       string
	ServerPort     string
	ServerProtocol string

	RootDir    string = "/DOORS/"
	DesktopDir string = "/DOORS/"

	mainUser string = ""
)


// Список программ
var programs = []*tProgram{
	{"Files", startExplorer, &bmpFolder_small},
	{"Notepad", startNotepad, &bmpNotepad},
	{"Internet", startBrowser, &bmpBrowser},
	//{"Flag", startFlag, &bmpProgram}, 
	//{"SNMP", startSNMP, &bmpProgram},
	//{"Dispatcher", startDispatch, &bmpProgram},
	{"Terminal", startTerminal, &bmpProgram},
	{"Virtual Machine", startVM, &bmpProgram},
	//{"Chrony", startChrony, &bmpProgram},
	{"COFFEE", startCOFFEE, &bmpProgram},
	{"PICO", startPICO, &bmpProgram},
	{"Camera 1", startCam1, &bmpCamera},
	{"Camera 2", startCam2, &bmpCamera},
	
	//{"ARM", startARM, &bmpProgram},
}
	

func main() {
	// Получение параметров сервера из глобального контекста JavaScript
	ServerIP = js.Global().Get("location").Get("host").String()
	ServerPort = js.Global().Get("location").Get("port").String()
	ServerProtocol = js.Global().Get("location").Get("protocol").String()
	
	// Получение размеров окна браузера
  	BITMAP_WIDTH = js.Global().Get("innerWidth").Int()
    BITMAP_HEIGHT = js.Global().Get("innerHeight").Int()
    	
    // Настройка графических параметров
    SIZE = BITMAP_WIDTH*BITMAP_HEIGHT 
	GETMAX_X = BITMAP_WIDTH - 1 
	GETMAX_Y = BITMAP_HEIGHT - 1
	BUFFER_SIZE = BITMAP_WIDTH*BITMAP_HEIGHT * 4
	
	// Инициализация системы
	initDOORS(programs)
	
	// Отрисовка событий
	eventDraw()
  	
  	// Блокировка завершения программы
	select {}
    //<-make(chan bool)
}


















