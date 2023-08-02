package main

import (
    "fmt"
    "math/rand"
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
	for i := 0; i < 799; i++ {
		x := rand.Intn(100)
		y := rand.Intn(100)

		ctx.Call("lineTo", i+x, (i*2/3)+y)
		ctx.Call("stroke")
		ctx.Set("strokeStyle", "#0d6efd")
		time.Sleep(50 * time.Millisecond)
	}
	
	
	
 
    // Prevent the function from returning, which is required in a wasm module
    <-make(chan bool)
}





