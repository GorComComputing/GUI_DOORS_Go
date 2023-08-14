package main

import (
    "fmt"
    //"math/rand"
    "math"
    //"syscall/js"
    //"time"
    //"strconv"
    //"net/http"
    //"io"
)

var frmFlag *Node
var frmDesktop *Node
var pnlTask *Node
var btnStart *Node
var btnFlag *Node
var btnWin1 *Node
var frmWin1 *Node
var btnEnter *Node
var btnCancel *Node
var btnOther *Node
var lblName *Node
var lblPswd *Node
var editName *Node
var editPswd *Node

func main() {
	message := "👋 Wasm started OK! 🌍"
  	fmt.Println(message)
  	
	//res, _ := http.DefaultClient.Get("http://localhost:8000")
	//if err != nil {
	//	fmt.Println("error making http request: \n")
	//}

	//fmt.Println("client: got response!\n")
	//fmt.Println("client: status code: " + strconv.Itoa(res.StatusCode))
	
/*js.FuncOf(func(this js.Value, args []js.Value) interface{} {

	go func(){

			res, _ := http.DefaultClient.Get("http://localhost:8000")
			defer res.Body.Close()
			

			b, _ := io.ReadAll(res.Body)

			fmt.Println("client: got response!")
			fmt.Println(string(b))
	}()
	
	return nil
})*/


	frmDesktop = CreateForm(&layout, 0, 0, BITMAP_WIDTH-1, BITMAP_HEIGHT-2, 0x0080C0, NONE, "", true, nil)
	pnlTask = CreatePanel(frmDesktop, 0, BITMAP_HEIGHT - 30, BITMAP_WIDTH - 1, 28, 0x30B410, nil)
	btnStart = CreateBtn(pnlTask, 0+2, BITMAP_HEIGHT - 30 + 2, 50, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick)
	btnFlag = CreateBtn(pnlTask, 0+2 + 52, BITMAP_HEIGHT - 30 + 2, 50, 28 - 4, 0xD8DCC0, 0x000000, "FLAG", btnFlagClick)
	btnWin1 = CreateBtn(pnlTask, 0+2 + 54 + 50, BITMAP_HEIGHT - 30 + 2, 50, 28 - 4, 0xD8DCC0, 0x000000, "WIN 1", btnWin1Click)
	
	frmWin1 = CreateForm(&layout, 100, 100, 300, 200, 0xD8DCC0, WIN, "WINDOW 1", true, nil)
	btnEnter = CreateBtn(frmWin1, 100 + 40, 100 + 20 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "ENTER", nil)
	btnCancel = CreateBtn(frmWin1, 100 + 40 + 70, 100 + 20 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "CANCEL", nil)
	btnOther = CreateBtn(frmWin1, 100 + 80 + 60, 100 + 20 + 30 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "OTHER", nil)
	
	lblName = CreateLabel(frmWin1, 100 + 12, 100 + 28, 80, 20, 0xD8DCC0, 0x000000, "USER NAME", nil)
	lblPswd = CreateLabel(frmWin1, 100 + 12, 100 + 28 + 30, 80, 20, 0xD8DCC0, 0x000000, "PASSWORD", nil)
	
	editName = CreateEdit(frmWin1, 100 + 80, 100 + 20, 80, 20, 0xF8FCF8, 0x000000, "MYUSERNAME", nil)
	editPswd = CreateEdit(frmWin1, 100 + 80, 100 + 20 + 30, 80, 20, 0xF8FCF8, 0x000000, "PSWD", nil)
	
	
	frmFlag = CreateForm(&layout, 50, 50, 300, 200, 0x000000, WIN, "FLAG", true, nil)


	
	/*btnEnter.obj.(*tBtn).x++
	btnCancel.obj.(*tBtn).x++
	btnOther.obj.(*tBtn).x++
	pnlTask.obj.(*tPanel).x++
  	editName.obj.(*tEdit).x++
  	editPswd.obj.(*tEdit).x++
  	lblName.obj.(*tLabel).x++
  	lblPswd.obj.(*tLabel).x++
  	btnStart.obj.(*tLabel).x++
  	//frmFlag.obj.(*tLabel).x++*/
 
    <-make(chan bool)
}


func btnStartClick(node *Node){
	node.obj.(*tBtn).visible = false
	
}


func btnFlagClick(node *Node){
	frmFlag.obj.(*tForm).visible = !(frmFlag.obj.(*tForm).visible)
}


func btnWin1Click(node *Node){
	frmWin1.obj.(*tForm).visible = !(frmWin1.obj.(*tForm).visible)
}



var flag string = 
	"                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "             ppppppppppp      ppppppp      pppppppppp      pppppppppp    ppp     ppp    ppppppppppp    pppppppppp                " +
    "             ppppppppppp    ppppppppppp    ppppppppppp    ppppppppppp    ppp     ppp    ppppppppppp    ppppppppppp               " +
    "             ppp            ppp     ppp    ppp     ppp    ppp     ppp    ppp     ppp    ppp            ppp     ppp               " +
    "             ppp            ppp     ppp    ppp     ppp    ppp     ppp    ppp     ppp    ppp            ppp     ppp               "+
    "             ppp            ppp     ppp    ppp     ppp    ppp     ppp    ppp     ppp    ppppppppppp    pppppppppp                "+
    "             ppp            ppp     ppp    ppppppppppp    ppppppppppp    ppp     ppp    ppppppppppp    pppppppppp                "+
    "             ppp            ppp     ppp    pppppppppp      pppppppppp     pppppppppp    ppp            ppp     ppp               "+
    "             ppp            ppp     ppp    ppp                ppp ppp            ppp    ppp            ppp     ppp               "+
    "             ppp            ppppppppppp    ppp              ppp   ppp            ppp    ppppppppppp    ppppppppppp               "+
    "             ppp             ppppppppp     ppp            ppp     ppp            ppp    ppppppppppp    pppppppppp                "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "+
    "                                                                                                                                 "	
 
var t  float64 = 0


//export generateCheckerBoard
func generateCheckerBoard() {
	SetBackColor(0xFFFFFF) //0x111111
	ClearDevice()
	DrawNode(&layout)
	flagDraw(frmFlag.obj.(*tForm).x, frmFlag.obj.(*tForm).y)
}


func flagDraw(x int, y int) {
	//var x int = 100
    //var y int = 100
    
    var xstart int = x
    var ystart int = y
    var dy int = 0
    
    t += 20;
    if t > 1000  {t = -3.14*12}
    
    for y := 0; y < 400; y++ {
    	for x := 0; x < 640; x++ {
    			squareNumber := (y * BITMAP_WIDTH) + x;
      			squareRgbaIndex := squareNumber * 4;

      			graphicsBuffer[squareRgbaIndex + 0] = 17; 	// Red
      			graphicsBuffer[squareRgbaIndex + 1] = 17; 	// Green
      			graphicsBuffer[squareRgbaIndex + 2] = 17; 	// Blue
      			graphicsBuffer[squareRgbaIndex + 3] = 255; 	// Alpha
    	}
    }
    
    for i := 0; i < 66; i++ {
        for j := 0; j < 129; j++ {
        	xi := (x+j + int(8*math.Cos(float64(j)/12.0+t)))
            yi := (y+i + int(8*math.Sin(float64(j)/12.0+t)))
            squareNumber := (yi * BITMAP_WIDTH) + xi;
      		squareRgbaIndex := squareNumber * 4;
                
                if(flag[i*129 + j] == byte('p')){
                    graphicsBuffer[squareRgbaIndex + 0] = 0xFF; 	// Red
      				graphicsBuffer[squareRgbaIndex + 1] = 0xFF; 	// Green
      				graphicsBuffer[squareRgbaIndex + 2] = 0x00; 	// Blue
      				graphicsBuffer[squareRgbaIndex + 3] = 255; 		// Alpha
                } else{
                	graphicsBuffer[squareRgbaIndex + 0] = 0xFF; 	// Red
      				graphicsBuffer[squareRgbaIndex + 1] = 0x00; 	// Green
      				graphicsBuffer[squareRgbaIndex + 2] = 0x00; 	// Blue
      				graphicsBuffer[squareRgbaIndex + 3] = 255; 		// Alpha
                }

                x += 1  //2
                if x%2 == 0 {y++}
        }
        x = xstart
        dy += 1  //2
        y = ystart + dy
    }
}


