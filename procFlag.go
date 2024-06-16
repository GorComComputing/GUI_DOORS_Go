package main

import (
    //"fmt"
    //"math/rand"
    "math"
    //"syscall/js"
    //"time"
    //"strconv"
    //"net/http"
    //"io"
)





func startFlag(frmMain *Node){
/*	setSize(frmMain, BITMAP_WIDTH, BITMAP_HEIGHT)
	//setSize(frmMain, 380, 340)
	frmMain.obj.(*tForm).x = 0
	frmMain.obj.(*tForm).y = 0
	frmMain.obj.(*tForm).mode = FLAT;
	frmMain.obj.(*tForm).BC = 0x111111;
	
	//frmMain.obj.(*tForm).visible = true
	frmMain.obj.(*tForm).visible = false
	
	//cnvFlag = CreateCanvas(frmDesktop, "cnvFlag", 0, 0, frmDesktop.obj.(*tForm).sizeX-10, frmDesktop.obj.(*tForm).sizeY-10, nil) //376, 321
	for y := 0; y < cnvFlag.obj.(*tCanvas).sizeY; y++ {
    	for x := 0; x < cnvFlag.obj.(*tCanvas).sizeX; x++ {
    			squareNumber := (y * cnvFlag.obj.(*tCanvas).sizeX) + x;
      			squareRgbaIndex := squareNumber * 4;

      			cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 0] = 0; 	// Red
      			cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 1] = 0; 	// Green
      			cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 2] = 0; 	// Blue
      			cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 3] = 255; 	// Alpha
    	}
    }*/
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


func flagDraw(x int, y int) {
    
    var xstart int = x
    var ystart int = y
    var dy int = 0
    
    t += 20;
    if t > 1000  {t = -3.14*12}
    
    for y := 0; y < cnvFlag.obj.(*tCanvas).sizeY; y++ {
    	for x := 0; x < cnvFlag.obj.(*tCanvas).sizeX; x++ {
    			squareNumber := (y * cnvFlag.obj.(*tCanvas).sizeX) + x;
      			squareRgbaIndex := squareNumber * 4;

      			cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 0] = 0x00; 	// Red
      			cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 1] = 0x80; 	// Green
      			cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 2] = 0xC0; 	// Blue
      			cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 3] = 255; 	// Alpha
    	}
    }
    
    for i := 0; i < 66; i++ {
        for j := 0; j < 129; j++ {
        	xi := (x+j + int(8*math.Cos(float64(j)/12.0+t)))
            yi := (y+i + int(8*math.Sin(float64(j)/12.0+t)))
            squareNumber := (yi * cnvFlag.obj.(*tCanvas).sizeX) + xi;
      		squareRgbaIndex := squareNumber * 4;
      		//squareNumber2 := (yi * cnvFlag.obj.(*tCanvas).sizeX + 1) + xi;
      		//squareRgbaIndex2 := squareNumber2 * 4;
                
                if(flag[i*129 + j] == byte('p')){
                    cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 0] = 0xFF; 	// Red
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 1] = 0xFF; 	// Green
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 2] = 0x00; 	// Blue
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 3] = 255; 		// Alpha
      				
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 4] = 0xFF; 	// Red
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 5] = 0xFF; 	// Green
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 6] = 0x00; 	// Blue
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 7] = 255; 		// Alpha
      				
      				/*cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 0] = 0xFF; 	// Red
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 1] = 0xFF; 	// Green
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 2] = 0x00; 	// Blue
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 3] = 255; 		// Alpha
      				
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 4] = 0xFF; 	// Red
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 5] = 0xFF; 	// Green
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 6] = 0x00; 	// Blue
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 7] = 255; 		// Alpha*/
                } else{
                	cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 0] = 0xFF; 	// Red
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 1] = 0x00; 	// Green
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 2] = 0x00; 	// Blue
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 3] = 255; 		// Alpha
      				
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 4] = 0xFF; 	// Red
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 5] = 0x00; 	// Green
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 6] = 0x00; 	// Blue
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex + 7] = 255; 		// Alpha
      				
      				/*cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 0] = 0xFF; 	// Red
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 1] = 0x00; 	// Green
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 2] = 0x00; 	// Blue
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 3] = 255; 		// Alpha
      				
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 4] = 0xFF; 	// Red
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 5] = 0x00; 	// Green
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 6] = 0x00; 	// Blue
      				cnvFlag.obj.(*tCanvas).buffer[squareRgbaIndex2 + 7] = 255; 		// Alpha*/
                }

                x += 1  //2
                if x%2 == 0 {y++}
        }
        x = xstart
        dy += 1  //2
        y = ystart + dy
    }
}

