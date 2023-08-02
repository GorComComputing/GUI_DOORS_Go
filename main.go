package main

import (
    "fmt"
    //"math/rand"
    "math"
    "syscall/js"
    "time"
)

func main() {
	message := "👋 Wasm TinyGo Landed! 🌍"

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
	canvas.Set("width", 800)
	canvas.Set("height", 600)

	body.Call("appendChild", canvas)
	//body.Call("insertBefore", canvas)

	//body.Call("removeChild", button)

	ctx = canvas.Call("getContext", "2d")
	grd = ctx.Call("createLinearGradient", 0, 0, 800, 0)
	//grd.Call("addColorStop", 0, "#ffc107")
	//grd.Call("addColorStop", 0.33, "#ffc107")
	//grd.Call("addColorStop", 0.66, "#ffc107")
	grd.Call("addColorStop", 1, "#ffc107")
	ctx.Set("fillStyle", grd)
	ctx.Call("fillRect", 0, 0, 800, 600)

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
    ctx.Call("fillRect", 0, 0, 800, 600)
	
    for i := 0; i < 66; i++ {
        for j := 0; j < 129; j++ {
                /*if(flag[i*129 + j] == 'p'){
                    PutPixel(x + j, (int)(y + i + 2*sin(j/12.+t)), 0xFFFF00);
                } else{
                    PutPixel(x + j, (int)(y + i + 2*sin(j/12.+t)), 0xFF0000);
                }*/
                ctx.Set("fillStyle", "#dc3545")
		ctx.Call("fillRect", (x+j + int(8*math.Cos(float64(j)/12.0+t))), (y+i + int(8*math.Sin(float64(j)/12.0+t))), 1, 1)
                //x += 2
                if x%4 == 0 {y++}
        }
        x = xstart
        dy += 0 //2
        y = ystart + dy
    }
    
    
    time.Sleep(1 * time.Millisecond)
	
}
	
 
    // Prevent the function from returning, which is required in a wasm module
    <-make(chan bool)
}





