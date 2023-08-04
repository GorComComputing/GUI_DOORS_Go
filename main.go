package main

import (
    "fmt"
    //"math/rand"
    //"math"
    "syscall/js"
    //"time"
    //"strconv"
)

func main() {
	message := "👋 Wasm TinyGo Landed! 🌍 Тестирование Gо в браузере"

  	fmt.Println(message)
  
    document := js.Global().Get("document")
    h2 := document.Call("createElement", "p")
    h2.Set("innerHTML", message)
    //document.Get("body").Call("appendChild", h2)
    
    body := document.Call("getElementById", "body_wasm")
    body.Call("appendChild", h2)
	//body.Set("innerHTML", "Dynamic Content")
	
	
	
	var (
		canvas, ctx, grd js.Value
	)
	//dombase = js.Global().Get("document")

	//button = document.Call("getElementById", "runButton")

	canvas = document.Call("createElement", "canvas")
	canvas.Set("width", 640)
	canvas.Set("height", 480)

	body.Call("appendChild", canvas)
	//body.Call("insertBefore", canvas)

	//body.Call("removeChild", button)

	ctx = canvas.Call("getContext", "2d")
	grd = ctx.Call("createLinearGradient", 0, 0, 640, 0)
	//grd.Call("addColorStop", 0, "#ffc107")
	//grd.Call("addColorStop", 0.33, "#ffc107")
	//grd.Call("addColorStop", 0.66, "#ffc107")
	grd.Call("addColorStop", 1, "#000000")
	ctx.Set("fillStyle", grd)
	ctx.Call("fillRect", 0, 0, 640, 480)

	ctx.Call("moveTo", 0, 0)
	/*for i := 0; i < 799; i++ {
		x := rand.Intn(800)
		y := rand.Intn(600)

		//ctx.Call("lineTo", i+x, (i*2/3)+y)
		//ctx.Call("stroke")
		//ctx.Set("strokeStyle", "#0d6efd")
		
		
		//ctx.Set("fillStyle", "#0d6efd")
		//ctx.Call("fillRect", x, y, 1, 1)

		time.Sleep(50 * time.Millisecond)
	}*/
	
	SetColor(0xFF0000)
	fmt.Println("CC: " + fmt.Sprintf("%x", CC))
	
	SetBackColor(0x0000FF)
	fmt.Println("BC: " + fmt.Sprintf("%x", BC))
	
	FillLB(0, 500, BC)
	fmt.Println("LB: " + fmt.Sprintf("%x", pBmp[10]))
	
	//ctx.Call("putImageData", pBmp,0,0)
	

imgData := canvas.Call("createImageData", 640, 480)
//const imgData = ctx.createImageData(100, 100);

/*for i := 0; i < 200; i += 4 {
  imgData[i+0] = 255
  imgData[i+1] = 0
  imgData[i+2] = 0
  imgData[i+3] = 255
}*/
ctx.Call("putImageData", imgData, 0, 0)
//ctx.putImageData(imgData, 10, 10);

	
	
	
/*	
    
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
    
    for  {
    
    
    
    var x int = 100
    var y int = 100
    
    var xstart int = x
    var ystart int = y
    var dy int = 0
    
    t += 20;
    if t > 1000  {t = -3.14*12}
    
    ctx.Set("fillStyle", grd)
    ctx.Call("fillRect", 0, 0, 640, 480)
	
    for i := 0; i < 66; i++ {
        for j := 0; j < 129; j++ {
                if(flag[i*129 + j] == byte('p')){
                    ctx.Set("fillStyle", "#ffc107")
                    ctx.Call("fillRect", (x+j + int(8*math.Cos(float64(j)/12.0+t))), (y+i + int(8*math.Sin(float64(j)/12.0+t))), 2, 2)
                    //PutPixel(x + j, (int)(y + i + 2*sin(j/12.+t)), 0xFFFF00);
                } else{
                    ctx.Set("fillStyle", "#dc3545")
                    ctx.Call("fillRect", (x+j + int(8*math.Cos(float64(j)/12.0+t))), (y+i + int(8*math.Sin(float64(j)/12.0+t))), 2, 2)
                    //PutPixel(x + j, (int)(y + i + 2*sin(j/12.+t)), 0xFF0000);
                }
                
		
                x += 2
                if x%4 == 0 {y++}
        }
        x = xstart
        dy += 2
        y = ystart + dy
    }
    
    
    time.Sleep(time.Millisecond)
	
}
*/	
 
    // Prevent the function from returning, which is required in a wasm module
    <-make(chan bool)
}





