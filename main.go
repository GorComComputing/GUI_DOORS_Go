package main

import (
    "fmt"
    //"math/rand"
    "math"
    //"syscall/js"
    //"time"
    "strconv"
    //"net/http"
    //"io"
)



//const CHECKERBOARD_SIZE int = 640;
const BUFFER_SIZE int = SIZE * 4;
var graphicsBuffer [BUFFER_SIZE]uint8;

var layout = Node{parent: nil, previous: nil, children: nil}
var list []*Node

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
	
	//layout = Node{parent: nil, previous: nil, children: nil}
	btnStart := CreateBtn(&layout, 0+2, BITMAP_HEIGHT - 30 + 2, 50, 28 - 4, 0x50A0F8, 0xF8FCF8, "START", btnStartClick) 
	btnEnter := CreateBtn(btnStart, 100 + 80, 100 + 20 + 30 + 30, 300, 240, 0xD8DCC0, 0x000000, "ENTER", nil)
	btnCancel := CreateBtn(btnStart, 100 + 80 + 60, 100 + 20 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "CANCEL", nil)
	btnOther := CreateBtn(btnEnter, 100 + 80 + 60, 100 + 20 + 30 + 30 + 30, 60, 24, 0xD8DCC0, 0x000000, "OTHER", nil)
	btnEnter.obj.x++
	btnCancel.obj.x++
	btnOther.obj.x++
	
	
	/*SetBackColor(0x111111)
	ClearDevice()
 
	DrawDesktop(0x0080C0)
	DrawTaskbar(0, BITMAP_HEIGHT - 30, BITMAP_WIDTH - 1, 28, 0x30B410);
	//btnStart.obj.Draw()
	
	DrawWindow(100, 100, 300, 200, 0xD8DCC0, "WINDOW 1")
	DrawLabel(100 + 12, 100 + 28, 80, 20, 0xD8DCC0, 0x000000, "USER NAME")
	DrawEdit(100 + 80, 100 + 20, 80, 20, 0xF8FCF8, 0x000000, "MYUSERNAME")
	DrawLabel(100 + 12, 100 + 28 + 30, 80, 20, 0xD8DCC0, 0x000000, "PASSWORD")
	DrawEdit(100 + 80, 100 + 20 + 30, 80, 20, 0xF8FCF8, 0x000000, "PSWD")
	//btnEnter.obj.Draw()

	DrawLayout(&layout)
	fmt.Println("Draw: END")
	
	PutPixel(100, 100, 0xFF0000)
	SetColor(0xFFFF00)
	//LinePP(200, 300, 100, 20)
	//HLine(50, 400, 120)
	Circle(80, 80, 15)
	
	SetColor(0xFF0000)
	fmt.Println("Pixel: " + strconv.FormatUint(0xFFFFFFFF & uint64(GetPixelgl(100, 100)),16))
	if (uint32(GetPixelgl(100, 100))) == 0xFF0000 {
		fmt.Println("YES") 
	}
	FloodFillgl(80, 80, 0xFFFF00)
  */
  
  
    <-make(chan bool)
}


// Function to return a pointer (Index) to our buffer in wasm memory
//export getGraphicsBufferPointer
func getGraphicsBufferPointer() *[BUFFER_SIZE]uint8 {
  return &graphicsBuffer
}


// Function to return the size of our buffer in wasm memory
//export getGraphicsBufferSize
func getGraphicsBufferSize() int {
  return BUFFER_SIZE;
}


//export eventClick
func eventClick(x int, y int)  {
	fmt.Println("Event: " + strconv.Itoa(x) + " " + strconv.Itoa(y))
	list = nil
	Click(&layout, x, y)
	//fmt.Println(list)
	fmt.Println("CLICKED: " + list[len(list)-1].obj.caption)
	if list[len(list)-1].obj.onClick != nil {
		list[len(list)-1].obj.onClick(list[len(list)-1])
	}
}

func Click(node *Node, x int, y int) {
	
	if node.obj != nil {
		if node.obj.x < x && (node.obj.x + node.obj.sizeX) > x && node.obj.y < y && (node.obj.y + node.obj.sizeY) > y {
			list = append(list, node)
			//fmt.Println("Click: " + list[len(list)-1].obj.caption)
		}
	} else {
		//fmt.Println("Click: none")
	}
	if node.children != nil {
		for i := 0; i < len(node.children); i++ { 
			Click(node.children[i], x, y)
		}
	}
	//fmt.Println(list)
	return
}


func btnStartClick(node *Node){
	node.obj.visible = false
	fmt.Println("IT WORKED !!! Click: START")
}



var flag string = "                                                                                                                                 " +
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
	SetBackColor(0x111111)
	ClearDevice()
 
	DrawDesktop(0x0080C0)
	DrawTaskbar(0, BITMAP_HEIGHT - 30, BITMAP_WIDTH - 1, 28, 0x30B410);
	//btnStart.obj.Draw()
	
	DrawWindow(100, 100, 300, 200, 0xD8DCC0, "WINDOW 1")
	DrawLabel(100 + 12, 100 + 28, 80, 20, 0xD8DCC0, 0x000000, "USER NAME")
	DrawEdit(100 + 80, 100 + 20, 80, 20, 0xF8FCF8, 0x000000, "MYUSERNAME")
	DrawLabel(100 + 12, 100 + 28 + 30, 80, 20, 0xD8DCC0, 0x000000, "PASSWORD")
	DrawEdit(100 + 80, 100 + 20 + 30, 80, 20, 0xF8FCF8, 0x000000, "PSWD")
	//btnEnter.obj.Draw()

	DrawLayout(&layout)
	fmt.Println("Draw: END")
	
	PutPixel(100, 100, 0xFF0000)
	SetColor(0xFFFF00)
	//LinePP(200, 300, 100, 20)
	//HLine(50, 400, 120)
	Circle(80, 80, 15)
	
	SetColor(0xFF0000)
	fmt.Println("Pixel: " + strconv.FormatUint(0xFFFFFFFF & uint64(GetPixelgl(100, 100)),16))
	if (uint32(GetPixelgl(100, 100))) == 0xFF0000 {
		fmt.Println("YES") 
	}
	FloodFillgl(80, 80, 0xFFFF00)	
}


func flagDraw() {
	var x int = 100
    var y int = 100
    
    var xstart int = x
    var ystart int = y
    var dy int = 0
    
    t += 20;
    if t > 1000  {t = -3.14*12}
    
    for y := 0; y < BITMAP_WIDTH; y++ {
    	for x := 0; x < BITMAP_HEIGHT; x++ {
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

                x += 2
                if x%4 == 0 {y++}
        }
        x = xstart
        dy += 2
        y = ystart + dy
    }
}

//export generateCheckerBoard
/*func generateCheckerBoard(
  darkValueRed uint8,
  darkValueGreen uint8,
  darkValueBlue uint8,
  lightValueRed uint8,
  lightValueGreen uint8,
  lightValueBlue uint8,
) {
  for y := 0; y < CHECKERBOARD_SIZE; y++ {
    for x := 0; x < CHECKERBOARD_SIZE; x++ {
      isDarkSquare := true;

      if y % 2 == 0 {
        isDarkSquare = false;
      }

      if x % 2 == 0 {
        isDarkSquare = !isDarkSquare;
      }

      squareValueRed := darkValueRed;
      squareValueGreen := darkValueGreen;
      squareValueBlue := darkValueBlue;
      if !isDarkSquare {
      squareValueRed = lightValueRed;
      squareValueGreen = lightValueGreen;
      squareValueBlue = lightValueBlue;
      }

      squareNumber := (y * CHECKERBOARD_SIZE) + x;
      squareRgbaIndex := squareNumber * 4;

      graphicsBuffer[squareRgbaIndex + 0] = squareValueRed; 	// Red
      graphicsBuffer[squareRgbaIndex + 1] = squareValueGreen; 	// Green
      graphicsBuffer[squareRgbaIndex + 2] = squareValueBlue; 	// Blue
      graphicsBuffer[squareRgbaIndex + 3] = 255; 				// Alpha 
    }
  }
}*/
