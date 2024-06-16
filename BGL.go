package main

import (
	"math"
	//"strings"
	//"fmt"
	//"strconv"
)


var CC uint32 
var BC uint32 

var Black uint32 = 0x000000
var White uint32 = 0xFFFFFF
var Red uint32 = 0xFF0000
var Green uint32 = 0x00FF00
var Blue uint32 = 0x0000FF
var Yellow uint32 = 0xFFFF00
var Magenta uint32 = 0xFF00FF
var Cyan uint32 = 0x00FFFF


func SetColor(Color uint32){
    CC = Color
}


func SetBackColor(Color uint32){
    BC = Color
}


func ClearDevice(buffer []uint8){
    FillLB(buffer, 0, SIZE, BC)
}


func DrawPutPixel(buffer []uint8, x int, y int, Color uint32){
    FillLB(buffer, (BITMAP_WIDTH*y + x), 1, Color);
}


func DrawLinePP(buffer []uint8, x1 int, y1 int, x2 int, y2 int){
    var dx int = int(math.Abs(float64(x2-x1)));
    var dy int= int(math.Abs(float64(y2-y1)));
        var x int;
        var y int;
        var xend int;
        var yend int;
        var inc1 int;
        var inc2 int;
        var d int;
        var s int;
    if dx > dy {
        if x1 < x2 {
            x = x1; xend = x2; y = y1;
            if y2 >= y1 {
            	s = 1;
            } else {
            	s = -1;
            }
        } else {
            x = x2; xend = x1; y = y2;
            if y2 >= y1 {
            	s = -1;
            } else {
            	s = 1;
            }
        }
        inc1 = 2*(dy - dx);
        inc2 = 2*dy;
        d = 2*dy - dx;
        PutPixel(buffer, x, y, CC);
        for ;x < xend; {
        //while(x < xend){
            x++;
            if d > 0 {
                d += inc1;
                y += s;
            } else{
                d += inc2;
            }
            PutPixel(buffer, x, y, CC);
        }
    } else{
        if y1 < y2 {
            y = y1; yend = y2; x = x1;
            if x2 >= x1 {
            	s = 1;
            } else {
            	s = -1;
            }
        } else {
            y = y2; yend = y1; x = x2;
            if x2 >= x1 {
            	s = -1; 
            } else {
            	s = 1;
            }
        }
        inc1 = 2*(dx - dy);
        inc2 = 2*dx;
        d = 2*dx - dy;
        PutPixel(buffer, x, y, CC);
        for ;y < yend; {
        //while(y < yend){
            y++;
            if d > 0 {
                d += inc1;
                x += s;
            } else {
                d += inc2;
            }
            PutPixel(buffer, x, y, CC);
        }
    }
}


func LinePP(buffer []uint8, x1 int, y1 int, x2 int, y2 int){
	var code, x, y int
	x1 += xleft
	x2 += xleft
	y1 += ytop
	y2 += ytop
	
	code1 := Coding(x1, y1)
	code2 := Coding(x2, y2)
	inside := (code1 | code2) == 0
	for ; !(inside) && ((code1 & code2) == 0); {
		if code1 == 0 {
			x = x1
			x1 = x2
			x2 = x
			y = y1
			y1 = y2
			y2 = y
			code = code1
			code1 = code2
			code2 = code
		}
		// Здесь x1, y1 - снаружи
		if x1 < xleft {
			y1 = y1 + int(math.Round(float64((y2-y1)/(x2-x1)*(xleft-x1))))
			x1 = xleft
		} else if x1 > xright {
			y1 = y1 + int(math.Round(float64((y2-y1)/(x2-x1)*(xright-x1))))
			x1 = xright
		} else if y1 < ytop {
			x1 = x1 + int(math.Round(float64((x2-x1)/(y2-y1)*(ytop-y1))))
			y1 = ytop
		} else { // y1 > ybottom
			x1 = x1 + int(math.Round(float64((x2-x1)/(y2-y1)*(ybottom-y1))))
			y1 = ybottom
		}
		code1 = Coding(x1, y1)
		inside = (code1 | code2) == 0
	} 
	if inside {
		DrawLinePP(buffer, x1, y1, x2, y2)
	}
}


func Coding(x int, y int) int{
	code := 0
	if x < xleft {
		code += 8 
	} else if x > xright {
		code += 4
	}
	
	if y < ytop {
		code += 2
	} else if y > ybottom {
		code += 1
	}
	return code
}


func HLine(buffer []uint8, x1 int, y int, x2 int){
    if x1 < x2 {
        FillLB(buffer, BITMAP_WIDTH*y + x1, x2 - x1 + 1, CC);
    } else {
        FillLB(buffer, BITMAP_WIDTH*y + x2, x1 - x2 + 1, CC);
	}
}


func Circle(buffer []uint8, xc int, yc int, R int){
    var x int = 0;
    var y int = R;
    var d int = 3 - 2*R;
    Pixel8(buffer, xc, yc, 0, R);
    for ; x < y ; {
    //while(x < y){
        if d < 0 {
            d = d + 4*x + 6;
        } else {
            d = d + 4*(x - y) + 10;
            y--;
        }
        x++;
        Pixel8(buffer, xc, yc, x, y);
    }
}


func Pixel8(buffer []uint8, xc int, yc int, x int, y int){
    PutPixel(buffer, xc + x, yc + y, CC);
    PutPixel(buffer, xc - x, yc + y, CC);
    PutPixel(buffer, xc + x, yc - y, CC);
    PutPixel(buffer, xc - x, yc - y, CC);

    PutPixel(buffer, xc + y, yc + x, CC);
    PutPixel(buffer, xc - y, yc + x, CC);
    PutPixel(buffer, xc + y, yc - x, CC);
    PutPixel(buffer, xc - y, yc - x, CC);
}


func FloodFillgl(buffer []uint8, x int, y int, bord uint32){
    var xl int = x;
    var xr int = x;
    var yy int;
    for ;GetPixelgl(buffer, xl, y) != bord; {
        xl--;
    }
    xl++;
    for ;GetPixelgl(buffer, xr, y) != bord; {
        xr++;
    }
    xr--;
    if xl < xr {
        HLine(buffer, xl, y, xr);
    }
    yy = y - 1;
    
    
    for ok := true; ok; ok = (yy <= y + 1) {
	  x = xr;
      for ;x >= xl; {
      	for ;(x >= xl) && (GetPixelgl(buffer, x, yy) == bord) || (GetPixelgl(buffer, x, yy) == CC); {
            x--;
        }
        if x >= xl {
            FloodFillgl(buffer, x, yy, bord);
        }
        x--;
      }
      yy += 2;
	}
	
}


//var YXbuf [GETMAX_Y]tXbuf
var YXbuf []tXbuf = make([]tXbuf, GETMAX_Y)


type tPoint struct{
    x int
    y int
}

//typedef  struct tPoint tPoly[100];

const NMAX int = 100    

type tXbuf struct{
    m int          
    x [NMAX]int    
}


func DrawFillPoly(buffer []uint8, n int,  p []tPoint){
    //for i := 0; i < n; i++ {
        //p[i].y = BITMAP_HEIGHT - p[i].y;
    //}

    var ymin int = p[0].y;
    var ymax int = ymin;
    for i := 0; i <= n - 1; i++ {
        if p[i].y < ymin {
            ymin = p[i].y;
        } else if p[i].y > ymax {
            ymax = p[i].y;
        }
    }

    for y := ymin; y <= ymax; y++ {
        YXbuf[y].m = 0; 
	}

    var i1 int = n - 1;
    for i2 := 0; i2 <= n - 1; i2++ {
        if p[i1].y != p[i2].y {
            Edge(p[i1].x, p[i1].y, p[i2].x, p[i2].y); 
        }
        i1 = i2;
    }

    for y := ymin; y <= ymax; y++ {
        Sort(YXbuf[y]);
        for i := 0; i < YXbuf[y].m; i += 2 {
            HLine(buffer, YXbuf[y].x[i], y, YXbuf[y].x[i + 1]);
        }
    }
}


func FillPoly(buffer []uint8, n int,  p []tPoint){
	//var p1 []tPoint
	//var p2 []tPoint
	var m1, m2 int
	
	p1 := make([]tPoint, 2*n)
	p2 := make([]tPoint, 2*n)
	//fmt.Println(len(p1))
	
	/*for i := 0; i < 2*n-1; i++ {
		pnew := tPoint{x: 0, y: 0}
		p1 = append(p1, pnew)
		p2 = append(p2, pnew)
	}*/
	
	//SetLength(p1, 2*n)
	//SetLength(p2, 2*n)
	for i := 0; i < n; i++ {
		p1[i].x = p[i].x + xleft
		p1[i].y = p[i].y + ytop
		//pnew := tPoint{x: p[i].x + xleft, y: p[i].y + ytop}
		//p1 = append(p1, pnew)
	}
	//fmt.Println(p1)
	//fmt.Println(p2)
	ClipLeft(n, p1, &m2, p2)
	//fmt.Println(p2)
	//DrawFillPoly(buffer, m1, p1)
	if m2 > 0 {
		ClipTop(m2, p2, &m1, p1)
		if m1 > 0 {
			ClipRight(m1, p1, &m2, p2)
			if m2 > 0 {
				ClipBottom(m2, p2, &m1, p1)
				if m1 > 0 {
				//fmt.Println("m1 "+ string(m1))
					DrawFillPoly(buffer, m1, p1)
				}
			}
		}
	}
}


func ClipLeft(n int, p1 []tPoint, m *int, p2 []tPoint) {
	var x1, y1, x2, y2 int
	var inside1, inside2 bool
	*m = 0
	x1 = p1[n-1].x
	y1 = p1[n-1].y
	inside1 = x1 >= xleft
	for i := 0; i < n; i++ {
		x2 = p1[i].x
		y2 = p1[i].y
		inside2 = x2 >= xleft
		if inside1 != inside2 {
			p2[*m].y = y2 + int(math.Round(float64((y1-y2)/(x1-x2)*(xleft-x2))))
			p2[*m].x = xleft
			*m += 1
		}
		if inside2 {
			p2[*m] = p1[i]
			*m += 1
		}
		x1 = x2
		y1 = y2
		inside1 = inside2
	} 
}


func ClipRight(n int, p1 []tPoint, m *int, p2 []tPoint) {
	var x1, y1, x2, y2 int
	var inside1, inside2 bool
	*m = 0
	x1 = p1[n-1].x
	y1 = p1[n-1].y
	inside1 = x1 <= xright
	for i := 0; i < n; i++ {
		x2 = p1[i].x
		y2 = p1[i].y
		inside2 = x2 <= xright
		if inside1 != inside2 {
			p2[*m].y = y2 + int(math.Round(float64((y1-y2)/(x1-x2)*(xright-x2))))
			p2[*m].x = xright
			*m += 1
		}
		if inside2 {
			p2[*m] = p1[i]
			*m += 1
		}
		x1 = x2
		y1 = y2
		inside1 = inside2
	} 
}


func ClipTop(n int, p1 []tPoint, m *int, p2 []tPoint) {
	var x1, y1, x2, y2 int
	var inside1, inside2 bool
	*m = 0
	x1 = p1[n-1].x
	y1 = p1[n-1].y
	inside1 = y1 >= ytop
	for i := 0; i < n; i++ {
		x2 = p1[i].x
		y2 = p1[i].y
		inside2 = y2 >= ytop
		if inside1 != inside2 {
			p2[*m].x = x1 + int(math.Round(float64((x2-x1)/(y2-y1)*(ytop-y1))))
			p2[*m].y = ytop
			*m += 1
		}
		if inside2 {
			p2[*m] = p1[i]
			*m += 1
		}
		x1 = x2
		y1 = y2
		inside1 = inside2
	} 
}


func ClipBottom(n int, p1 []tPoint, m *int, p2 []tPoint) {
	var x1, y1, x2, y2 int
	var inside1, inside2 bool
	*m = 0
	x1 = p1[n-1].x
	y1 = p1[n-1].y
	inside1 = y1 <= ybottom
	for i := 0; i < n; i++ {
		x2 = p1[i].x
		y2 = p1[i].y
		inside2 = y2 <= ybottom
		if inside1 != inside2 {
			p2[*m].x = x1 + int(math.Round(float64((x2-x1)/(y2-y1)*(ybottom-y1))))
			p2[*m].y = ybottom
			*m += 1
		}
		if inside2 {
			p2[*m] = p1[i]
			*m += 1
		}
		x1 = x2
		y1 = y2
		inside1 = inside2
	} 
}


func Edge(x1 int, y1 int, x2 int, y2 int){
    var y int;
    var yend int;
    var xf float32;
    var k float32 = float32(x2 - x1)/float32(y2 - y1);
    if y1 < y2 {
    	y = y1; 
    	yend = y2; 
    	xf = float32(x1);
    } else {
    	y = y2; 
    	yend = y1; 
    	xf = float32(x2);
    }
    for ;y < yend; {
       y++;
       xf += k;
       YXbuf[y].m++;
       YXbuf[y].x[YXbuf[y].m-1] = int(math.Round(float64(xf)));
    }
}


func Sort(a tXbuf){
    var y int;
    var j int;
    for i := 1; i < a.m; i++ {
        y = a.x[i];
        j = i - 1;
        for ;(j >= 0) && (y < a.x[j]); {
            a.x[j + 1] = a.x[j];
            j--;
        }
        a.x[j + 1] = y;
    }
}


func FillCircle(buffer []uint8, xc int, yc int, R int) {
	d := 3 - 2*R
	x := 0
	y := R
	Draw4(buffer, xc, yc, x, R)
	for ; x < y; {
		if d < 0 {
			d = d + 4*x + 6
		} else {
			d = d + 4*(x - y) + 10
			y = y - 1
		}
		x = x + 1
		Draw4(buffer, xc, yc, x, y)
	}
}


func Draw4(buffer []uint8, xc int, yc int, x int, y int) {
	HLine(buffer, xc + x, yc + y, xc - x)
	HLine(buffer, xc - y, yc + x, xc + y)
	HLine(buffer, xc + x, yc - y, xc - x)
	HLine(buffer, xc - y, yc - x, xc + y)
}


var xleft int
var ytop int
var xright int
var ybottom int
	
	
func SetViewPort(xl int, yt int, xr int, yb int){
	xleft = xl
	ytop = yt
	xright = xr
	ybottom = yb
}


var xleft_loc int = 0
var ytop_loc int = 0
var xright_loc int = GETMAX_X
var ybottom_loc int = GETMAX_Y
	
	
func SetLocalViewPort(xl int, yt int, xr int, yb int){
	xleft_loc = xl
	ytop_loc = yt
	xright_loc = xr
	ybottom_loc = yb
}


func PutPixel(buffer []uint8, x int, y int, C uint32){
	x += xleft	// Переход к абсолютным координатам
	y += ytop
	if x >= xleft && x <= xright && y >= ytop && y <= ybottom {
		DrawPutPixel(buffer, x, y, C)
	}
}


func DrawBitmapTransparent(buffer []uint8, monster string, xstart int, ystart int, sizeX int, sizeY int, scale int){
    var x int = xstart;
    var y int = ystart;

    for i := 0; i < sizeY; i++ {
        for scaleY := scale; scaleY > 0; scaleY-- {
        for j := 0; j < sizeX; j++ {
            if(monster[i*sizeX + j] == 'p'){
                for scaleX := scale; scaleX > 0; scaleX--{
                	if x + j >= xleft_loc && x + j <= xright_loc && y + i >= ytop_loc && y + i <= ybottom_loc {
                		PutPixel(buffer, x + j, y + i, CC);
                	}
                if scaleX > 1 { x++;}
                }
            } else {
                for scaleX := scale; scaleX > 0; scaleX-- {
                //PutPixel(buffer, x + j, y + i, BC);
                if scaleX > 1 {x++;}
                }
            }
        }
        x = xstart;
        if scaleY > 1 {y++;}
        }
    }
}


func TextOutgl(buffer []uint8, str string, x_start int, y int, scale int){
var char41 string = 
    "   p   " +
    "  ppp  " +
    " pp pp " +
    "pp   pp" +
    "pp   pp" +
    "ppppppp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "       " 
    ;//10x7
var char42 string =
    "pppppp " +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " ppppp " +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    "pppppp " +
    "       " 
    ;//10x7
var char43 string =
    "  pppp " +
    " pp  pp" +
    "pp    p" +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp    p" +
    " pp  pp" +
    "  pppp " +
    "       " 
    ;//10x7
var char44 string =
    "ppppp  " +
    " pp pp " +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " pp pp " +
    "ppppp  " +
    "       " 
    ;//10x7
var char45 string =
    "ppppppp" +
    " pp  pp" +
    " pp   p" +
    " pp p  " +
    " pppp  " +
    " pp p  " +
    " pp   p" +
    " pp  pp" +
    "ppppppp" +
    "       " 
    ;//10x7
var char46 string =
    "ppppppp" +
    " pp  pp" +
    " pp   p" +
    " pp p  " +
    " pppp  " +
    " pp p  " +
    " pp    " +
    " pp    " +
    "pppp   " +
    "       " 
    ;//10x7
var char47 string =
    "  pppp " +
    " pp  pp" +
    "pp    p" +
    "pp     " +
    "pp     " +
    "pp pppp" +
    "pp   pp" +
    " pp  pp" +
    "  ppp p" +
    "       " 
    ;//10x7
var char48 string =
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "ppppppp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "       " 
    ;//10x7
var char49 string =
    "  pppp " +
    "   pp  " +
    "   pp  " +
    "   pp  " +
    "   pp  " +
    "   pp  " +
    "   pp  " +
    "   pp  " +
    "  pppp " +
    "       " 
    ;//10x7
var char4A string =
    "   pppp" +
    "    pp " +
    "    pp " +
    "    pp " +
    "    pp " +
    "    pp " +
    "pp  pp " +
    "pp  pp " +
    " pppp  " +
    "       " 
    ;//10x7
var char4B string =
    "pppp pp" +
    " pp  pp" +
    " pp pp " +
    " pp pp " +
    " pppp  " +
    " pp pp " +
    " pp pp " +
    " pp  pp" +
    "pppp pp" +
    "       " 
    ;//10x7
var char4C string =
    "pppp   " +
    " pp    " +
    " pp    " +
    " pp    " +
    " pp    " +
    " pp    " +
    " pp   p" +
    " pp  pp" +
    "ppppppp" +
    "       " 
    ;//10x7
var char4D string =
    "pp    pp" +
    "ppp  ppp" +
    "pppppppp" +
    "pp pp pp" +
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    "        " 
    ;//8x10
var char4E string =
    "pp   pp" +
    "ppp  pp" +
    "pppp pp" +
    "pp pppp" +
    "pp  ppp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "       " 
    ;//7x10
var char4F string =
    "  ppp  " +
    " pp pp " +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    " pp pp " +
    "  ppp  " +
    "       " 
    ;//7x10
var char50 string =
    "pppppp " +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " ppppp " +
    " pp    " +
    " pp    " +
    " pp    " +
    "pppp   " +
    "       " 
    ;//7x10
var char51 string =
    " ppppp " +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp p pp" +
    "pp pppp" +
    " ppppp " +
    "    pp " +
    "    ppp" 
    ;//7x10
var char52 string =
    "pppppp " +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " ppppp " +
    " pp pp " +
    " pp  pp" +
    " pp  pp" +
    "ppp  pp" +
    "       " 
    ;//7x10
var char53 string =
    " ppppp " +
    "pp   pp" +
    "pp   pp" +
    " pp    " +
    "  ppp  " +
    "    pp " +
    "pp   pp" +
    "pp   pp" +
    " ppppp " +
    "       " 
    ;//7x10
var char54 string =
    "pppppppp" +
    "pp pp pp" +
    "p  pp  p" +
    "   pp   " +
    "   pp   " +
    "   pp   " +
    "   pp   " +
    "   pp   " +
    "  pppp  " +
    "        " 
    ;//8x10
var char55 string =
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    " ppppp " +
    "       " 
    ;//7x10
var char56 string =
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    " pp  pp " +
    "  pppp  " +
    "   pp   " +
    "        " 
    ;//8x10
var char57 string =
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    "pp pp pp" +
    "pp pp pp" +
    "pppppppp" +
    " pp  pp " +
    " pp  pp " +
    "        " 
    ;//8x10
var char58 string =
    "pp    pp" +
    "pp    pp" +
    " pp  pp " +
    "  pppp  " +
    "   pp   " +
    "  pppp  " +
    " pp  pp " +
    "pp    pp" +
    "pp    pp" +
    "        " 
    ;//8x10
var char59 string =
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    " pp  pp " +
    "  pppp  " +
    "   pp   " +
    "   pp   " +
    "   pp   " +
    "  pppp  " +
    "        " 
    ;//8x10
var char5A string =
    "pppppppp" +
    "pp    pp" +
    "p    pp " +
    "    pp  " +
    "  pppp  " +
    "  pp    " +
    " pp    p" +
    "pp    pp" +
    "pppppppp" +
    "        " 
    ;//8x10
var char20 string =
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       "
    ;//7x11
var charLine string =
    "       " +
    "       " +
    "       " +
    " ppppp " +
    "       " +
    "       " +
    "       "
    ;//7x7
var char30 string =
    " ppppp " +
    "pp   pp" +
    "pp  ppp" +
    "pp pppp" +
    "pppp pp" +
    "ppp  pp" +
    "pp   pp" +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       " 
    ;//7x11
var char31 string =
    "  pp   " +
    " ppp   " +
    "pppp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "pppppp " +
    "       " +
    "       " 
    ;//7x11
var char32 string =
    " ppppp " +
    "pp   pp" +
    "     pp" +
    "    pp " +
    "   pp  " +
    "  pp   " +
    " pp    " +
    "pp   pp" +
    "ppppppp" +
    "       " +
    "       " 
    ;//7x11
var char33 string =
    " ppppp " +
    "pp   pp" +
    "     pp" +
    "     pp" +
    "  pppp " +
    "     pp" +
    "     pp" +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       " 
    ;//7x11
var char34 string =
    "    pp " +
    "   ppp " +
    "  pppp " +
    " pp pp " +
    "pp  pp " +
    "ppppppp" +
    "    pp " +
    "    pp " +
    "   pppp" +
    "       " +
    "       " 
    ;//7x11
var char35 string =
    "ppppppp" +
    "pp     " +
    "pp     " +
    "pp     " +
    "pppppp " +
    "     pp" +
    "     pp" +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       " 
    ;//7x11
var char36 string =
    "  ppp  " +
    " pp    " +
    "pp     " +
    "pp     " +
    "pppppp " +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       " 
    ;//7x11
var char37 string =
    "ppppppp" +
    "pp   pp" +
    "     pp" +
    "    pp " +
    "   pp  " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "       " +
    "       " 
    ;//7x11
var char38 string =
    " ppppp " +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    " ppppp " +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       " 
    ;//7x11
var char39 string =
    " ppppp " +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    " pppppp" +
    "     pp" +
    "     pp" +
    "     pp" +
    " ppppp " +
    "       " +
    "       " 
    ;//7x11
var charUndefined string =
    "ppppppp" +
    "p     p" +
    "p     p" +
    "p     p" +
    "p     p" +
    "p     p" +
    "ppppppp"
    ;//7x7
var charMinimized string =
    "       " +
    "  ppppp" +
    "  p   p" +
    "ppppp p" +
    "p   ppp" +
    "p   p  " +
    "ppppp  "
    ;//7x7
var char3A string =
    "    " +
    "    " +
    " pp " +
    " pp " +
    "    " +
    "    " +
    " pp " +
    " pp " +
    "    " +
    "    " +
    "    "
    ;//4x11 
var char_vert_line string =
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "       " +
    "       "
    ;//7x11
var char61 string =
    "       " +
    "       " +
    "       " +
    " ppp   " +
    "   pp  " +
    " pppp  " +
    "pp pp  " +
    "pp pp  " +
    " ppp pp" +
    "       " +
    "       "
    ;//7x11
var char62 string =
    "ppp    " +
    " pp    " +
    " pp    " +
    " pppp  " +
    " pp pp " +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    "pp ppp " +
    "       " +
    "       "
    ;//7x11
var char63 string =
    "       " +
    "       " +
    "       " +
    " ppppp " +
    "pp   pp" +
    "pp     " +
    "pp     " +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       "
    ;//7x11
var char64 string =
    "   ppp " +
    "    pp " +
    "    pp " +
    "  pppp " +
    " pp pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " ppp pp" +
    "       " +
    "       "
    ;//7x11
var char65 string =
    "       " +
    "       " +
    "       " +
    " ppppp " +
    "pp   pp" +
    "ppppppp" +
    "pp     " +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       "
    ;//7x11
var char66 string =
    "  ppp  " +
    " pp pp " +
    " pp  p " +
    " pp    " +
    "ppppp  " +
    " pp    " +
    " pp    " +
    " pp    " +
    "pppp   " +
    "       " +
    "       "
    ;//7x11
var char67 string =
    "       " +
    "       " +
    "       " +
    " ppp pp" +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " ppppp " +
    "    pp " +
    "pp  pp " +
    " pppp  "
    ;//7x11
var char68 string =
    "ppp    " +
    " pp    " +
    " pp    " +
    " pp pp " +
    " ppp pp" +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    "ppp  pp" +
    "       " +
    "       "
    ;//7x11
var char69 string =
    "  pp   " +
    "  pp   " +
    "       " +
    " ppp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    " pppp  " +
    "       " +
    "       "
    ;//7x11
var char6A string =
    "    pp " +
    "    pp " +
    "       " +
    "   ppp " +
    "    pp " +
    "    pp " +
    "    pp " +
    "    pp " +
    "p   pp " +
    " p  pp " +
    " pppp  "
    ;//7x11
var char6B string =
    "p p    " +
    " pp    " +
    " pp    " +
    " pp  pp" +
    " pp pp " +
    " pppp  " +
    " pp pp " +
    " pp  pp" +
    "ppp  pp" +
    "       " +
    "       "
    ;//7x11
var char6C string =
    " ppp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    " pppp  " +
    "       " +
    "       "
    ;//7x11
var char6D string =
    "        " +
    "        " +
    "        " +
    "ppp  pp " +
    "pppppppp" +
    "pp pp pp" +
    "pp pp pp" +
    "pp pp pp" +
    "pp pp pp" +
    "        " +
    "        "
    ;//8x11
var char6E string =
    "       " +
    "       " +
    "       " +
    "pp ppp " +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    "       " +
    "       "
    ;//7x11
var char6F string =
    "       " +
    "       " +
    "       " +
    " ppppp " +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       "
    ;//7x11
var char70 string =
    "       " +
    "       " +
    "       " +
    "pp ppp " +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " pp  pp" +
    " ppppp " +
    " pp    " +
    "pppp   "
    ;//7x11
var char71 string =
    "       " +
    "       " +
    "       " +
    " ppp pp" +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " ppppp " +
    "    pp " +
    "    pp " +
    "   pppp"
    ;//7x11
var char72 string =
    "       " +
    "       " +
    "       " +
    "pp ppp " +
    " ppp pp" +
    " pp  pp" +
    " pp    " +
    " pp    " +
    "pppp   " +
    "       " +
    "       "
    ;//7x11
var char73 string =
    "       " +
    "       " +
    "       " +
    " ppppp " +
    "pp   pp" +
    " ppp   " +
    "   ppp " +
    "pp   pp" +
    " ppppp " +
    "       " +
    "       "
    ;//7x11
var char74 string =
    "   p   " +
    "  pp   " +
    "  pp   " +
    "pppppp " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp pp" +
    "   ppp " +
    "       " +
    "       "
    ;//7x11
var char75 string =
    "       " +
    "       " +
    "       " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " ppp pp" +
    "       " +
    "       "
    ;//7x11
var char76 string =
    "        " +
    "        " +
    "        " +
    "pp    pp" +
    "pp    pp" +
    "pp    pp" +
    " pp  pp " +
    "  pppp  " +
    "   pp   " +
    "        " +
    "        "
    ;//8x11
var char77 string =
    "        " +
    "        " +
    "        " +
    "pp    pp" +
    "pp    pp" +
    "pp pp pp" +
    "pp pp pp" +
    "pppppppp" +
    " pp  pp " +
    "        " +
    "        "
    ;//8x11
var char78 string =
    "       " +
    "       " +
    "       " +
    "pp   pp" +
    " pp pp " +
    "  ppp  " +
    "  ppp  " +
    " pp pp " +
    "pp   pp" +
    "       " +
    "       "
    ;//7x11
var char79 string =
    "       " +
    "       " +
    "       " +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    "pp   pp" +
    " pppppp" +
    "     pp" +
    "    pp " +
    " pppp  "
    ;//7x11
var char7A string =
    "       " +
    "       " +
    "       " +
    "ppppppp" +
    "pp  pp " +
    "   pp  " +
    "  pp   " +
    " pp  pp" +
    "ppppppp" +
    "       " +
    "       "
    ;//7x11
var char7B string =
    "       " +
    "  pppp " +
    " pp    " +
    " pp    " +
    " pp    " +
    "pp     " +
    " pp    " +
    " pp    " +
    " pp    " +
    "  pppp " +
    "       " 
    ;//7x11 
    
var char7D string =
    "       " +
    "pppp   " +
    "   pp  " +
    "   pp  " +
    "   pp  " +
    "    pp " +
    "   pp  " +
    "   pp  " +
    "   pp  " +
    "pppp   " +
    "       " 
    ;//7x11 
var char7E string =
    "       " +
    "       " +
    "       " +
    "       " +
    " pp  p " +
    "p  pp  " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       "
    ;//7x11 
    
    
    
var char21 string =
    "    " +
    " pp " +
    "pppp" +
    "pppp" +
    " pp " +
    " pp " +
    "    " +
    " pp " +
    " pp " +
    "    " +
    "    "
    ;//4x11
var char22 string =
    "pp   pp" +
    "pp   pp" +
    " p   p " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       "
    ;//7x11
var char23 string =
    " pp pp " +
    " pp pp " +
    "ppppppp" +
    " pp pp " +
    " pp pp " +
    " pp pp " +
    "ppppppp" +
    " pp pp " +
    " pp pp " +
    "       " +
    "       "
    ;//7x11
var char24 string =
    "   pp  " +
    "   pp  " +
    " ppppp " +
    "pp   pp" +
    "pp    p" +
    "pp     " +
    " ppppp " +
    "     pp" +
    "p    pp" +
    "pp   pp" +
    " ppppp " +
    "   pp  " +
    "   pp  " 
    ;//8x13
var char25 string =
    "       " +
    "       " +
    "pp    p" +
    "pp   pp" +
    "    pp " +
    "   pp  " +
    "  pp   " +
    " pp  pp" +
    "pp   pp" +
    "       " +
    "       "
    ;//7x11   
var char26 string =
    "  ppp  " +
    " pp pp " +
    " pp pp " +
    "  ppp  " +
    " ppp pp" +
    "pp ppp " +
    "pp  pp " +
    "pp  pp " +
    " ppp pp" +
    "       " +
    "       "
    ;//7x11 
var char27 string =
    " pp " +
    " pp " +
    " pp " +
    "pp  " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    "
    ;//4x11    
var char28 string =
    "  pp" +
    " pp " +
    "pp  " +
    "pp  " +
    "pp  " +
    "pp  " +
    "pp  " +
    " pp " +
    "  pp" +
    "    " +
    "    "
    ;//7x11 
var char29 string =
    "pp  " +
    " pp " +
    "  pp" +
    "  pp" +
    "  pp" +
    "  pp" +
    "  pp" +
    " pp " +
    "pp  " +
    "    " +
    "    "
    ;//4x11 
var char2A string =
    "        " +
    "        " +
    " pp  pp " +
    "  pppp  " +
    "pppppppp" +
    "  pppp  " +
    " pp  pp " +
    "        " +
    "        " +
    "        " +
    "        "
    ;//8x11  
var char2B string =
    "        " +
    "   pp   " +
    "   pp   " +
    "   pp   " +
    "pppppppp" +
    "   pp   " +
    "   pp   " +
    "   pp   " +
    "        " +
    "        " +
    "        "
    ;//8x11  
    
var char2C string =
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    " pp " +
    " pp " +
    " pp " +
    "pp  " +
    "    "
    ;//4x11 
    
var char2E string =
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    " pp " +
    " pp " +
    "    " +
    "    "
    ;//4x11 
var char2F string =
    "      p" +
    "     pp" +
    "    pp " +
    "   pp  " +
    "  pp   " +
    " pp    " +
    "pp     " +
    "p      " +
    "       " +
    "       " +
    "       "
    ;//7x11

var char3B string =
    "    " +
    "    " +
    " pp " +
    " pp " +
    "    " +
    "    " +
    " pp " +
    " pp " +
    "pp  " +
    "    " +
    "    "
    ;//4x11    
var char3C string =
    "    pp " +
    "   pp  " +
    "  pp   " +
    " pp    " +
    "pp     " +
    " pp    " +
    "  pp   " +
    "   pp  " +
    "    pp " +
    "       " +
    "       "
    ;//7x11 
var char3D string =
    "       " +
    "       " +
    "       " +
    "ppppppp" +
    "       " +
    "       " +
    "ppppppp" +
    "       " +
    "       " +
    "       " +
    "       "
    ;//7x11     
var char3E string =
    "pp     " +
    " pp    " +
    "  pp   " +
    "   pp  " +
    "    pp " +
    "   pp  " +
    "  pp   " +
    " pp    " +
    "pp     " +
    "       " +
    "       "
    ;//7x11 
var char3F string =
    " ppppp " +
    "pp   pp" +
    "pp   pp" +
    "    pp " +
    "   pp  " +
    "   pp  " +
    "       " +
    "   pp  " +
    "   pp  " +
    "       " +
    "       "
    ;//7x11   
var char40 string =
    " ppppp " +
    "pp   pp" +
    "pp   pp" +
    "pp pppp" +
    "pp pppp" +
    "pp pppp" +
    "pp ppp " +
    "pp     " +
    " ppppp " +
    "       " +
    "       "
    ;//7x11 
    
var char5B string =
    "    " +
    "pppp" +
    "pp  " +
    "pp  " +
    "pp  " +
    "pp  " +
    "pp  " +
    "pp  " +
    "pp  " +
    "pppp" +
    "    " 
    ;//4x11 
var char5C string =
    "p      " +
    "pp     " +
    " pp    " +
    "  pp   " +
    "   pp  " +
    "    pp " +
    "     pp" +
    "      p" +
    "       " +
    "       " +
    "       "
    ;//7x11
var char5D string =
    "    " +
    "pppp" +
    "  pp" +
    "  pp" +
    "  pp" +
    "  pp" +
    "  pp" +
    "  pp" +
    "  pp" +
    "pppp" +
    "    " 
    ;//4x11 
var char5E string =
    "   p   " +
    "  ppp  " +
    " pp pp " +
    "pp   pp" +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " 
    ;//7x11     
var char5F string =
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "ppppppp" +
    "       " 
    ;//7x11 
var char60 string =
    " pp " +
    " pp " +
    " pp " +
    "pp  " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    " +
    "    "
    ;//4x11 
    

	var interval int = 2
	var x = x_start
	//str = strings.ToUpper(str)
    for j := 0; j < len(str); j++ {
    //if (str[j] >=0x20 && str[j] <= 0x7E) || str[j] == 13 {
        if str[j] == 'A' { DrawBitmapTransparent(buffer, char41, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'B' { DrawBitmapTransparent(buffer, char42, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'C' { DrawBitmapTransparent(buffer, char43, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'D' { DrawBitmapTransparent(buffer, char44, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'E' { DrawBitmapTransparent(buffer, char45, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'F' { DrawBitmapTransparent(buffer, char46, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'G' { DrawBitmapTransparent(buffer, char47, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'H' { DrawBitmapTransparent(buffer, char48, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'I' { DrawBitmapTransparent(buffer, char49, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'J' { DrawBitmapTransparent(buffer, char4A, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'K' { DrawBitmapTransparent(buffer, char4B, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'L' { DrawBitmapTransparent(buffer, char4C, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'M' { DrawBitmapTransparent(buffer, char4D, x, y, 8, 10, scale); x += 9;
        } else if str[j] == 'N' { DrawBitmapTransparent(buffer, char4E, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'O' { DrawBitmapTransparent(buffer, char4F, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'P' { DrawBitmapTransparent(buffer, char50, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'Q' { DrawBitmapTransparent(buffer, char51, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'R' { DrawBitmapTransparent(buffer, char52, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'S' { DrawBitmapTransparent(buffer, char53, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'T' { DrawBitmapTransparent(buffer, char54, x, y, 8, 10, scale); x += 9;
        } else if str[j] == 'U' { DrawBitmapTransparent(buffer, char55, x, y, 7, 10, scale); x += 8;
        } else if str[j] == 'V' { DrawBitmapTransparent(buffer, char56, x, y, 8, 10, scale); x += 9;
        } else if str[j] == 'W' { DrawBitmapTransparent(buffer, char57, x, y, 8, 10, scale); x += 9;
        } else if str[j] == 'X' { DrawBitmapTransparent(buffer, char58, x, y, 8, 10, scale); x += 9;
        } else if str[j] == 'Y' { DrawBitmapTransparent(buffer, char59, x, y, 8, 10, scale); x += 9;
        } else if str[j] == 'Z' { DrawBitmapTransparent(buffer, char5A, x, y, 8, 10, scale); x += 9;
        } else if str[j] == ' ' { DrawBitmapTransparent(buffer, char20, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '-' { DrawBitmapTransparent(buffer, charLine, x, y, 7, 7, scale); x += 7;
        } else if str[j] == '0' { DrawBitmapTransparent(buffer, char30, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '1' { DrawBitmapTransparent(buffer, char31, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '2' { DrawBitmapTransparent(buffer, char32, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '3' { DrawBitmapTransparent(buffer, char33, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '4' { DrawBitmapTransparent(buffer, char34, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '5' { DrawBitmapTransparent(buffer, char35, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '6' { DrawBitmapTransparent(buffer, char36, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '7' { DrawBitmapTransparent(buffer, char37, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '8' { DrawBitmapTransparent(buffer, char38, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '9' { DrawBitmapTransparent(buffer, char39, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '|' { DrawBitmapTransparent(buffer, char_vert_line, x, y, 7, 11, scale); x += 7;
        
        } else if str[j] == 'a' { DrawBitmapTransparent(buffer, char61, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'b' { DrawBitmapTransparent(buffer, char62, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'c' { DrawBitmapTransparent(buffer, char63, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'd' { DrawBitmapTransparent(buffer, char64, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'e' { DrawBitmapTransparent(buffer, char65, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'f' { DrawBitmapTransparent(buffer, char66, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'g' { DrawBitmapTransparent(buffer, char67, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'h' { DrawBitmapTransparent(buffer, char68, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'i' { DrawBitmapTransparent(buffer, char69, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'j' { DrawBitmapTransparent(buffer, char6A, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'k' { DrawBitmapTransparent(buffer, char6B, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'l' { DrawBitmapTransparent(buffer, char6C, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'm' { DrawBitmapTransparent(buffer, char6D, x, y, 8, 11, scale); x += 9;
        } else if str[j] == 'n' { DrawBitmapTransparent(buffer, char6E, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'o' { DrawBitmapTransparent(buffer, char6F, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'p' { DrawBitmapTransparent(buffer, char70, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'q' { DrawBitmapTransparent(buffer, char71, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'r' { DrawBitmapTransparent(buffer, char72, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 's' { DrawBitmapTransparent(buffer, char73, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 't' { DrawBitmapTransparent(buffer, char74, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'u' { DrawBitmapTransparent(buffer, char75, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'v' { DrawBitmapTransparent(buffer, char76, x, y, 8, 11, scale); x += 9;
        } else if str[j] == 'w' { DrawBitmapTransparent(buffer, char77, x, y, 8, 11, scale); x += 9;
        } else if str[j] == 'x' { DrawBitmapTransparent(buffer, char78, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'y' { DrawBitmapTransparent(buffer, char79, x, y, 7, 11, scale); x += 8;
        } else if str[j] == 'z' { DrawBitmapTransparent(buffer, char7A, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '{' { DrawBitmapTransparent(buffer, char7B, x, y, 7, 11, scale); x += 8;
        
        } else if str[j] == '}' { DrawBitmapTransparent(buffer, char7D, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '~' { DrawBitmapTransparent(buffer, char7E, x, y, 7, 11, scale); x += 8;
        
        } else if str[j] == '!' { DrawBitmapTransparent(buffer, char21, x, y, 4, 11, scale); x += 5;
        } else if str[j] == 0x22 { DrawBitmapTransparent(buffer, char22, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '#' { DrawBitmapTransparent(buffer, char23, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '$' { DrawBitmapTransparent(buffer, char24, x, y, 7, 13, scale); x += 8;
        } else if str[j] == '%' { DrawBitmapTransparent(buffer, char25, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '&' { DrawBitmapTransparent(buffer, char26, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '\'' { DrawBitmapTransparent(buffer, char27, x, y, 4, 11, scale); x += 5;
        } else if str[j] == '(' { DrawBitmapTransparent(buffer, char28, x, y, 4, 11, scale); x += 5;
        } else if str[j] == ')' { DrawBitmapTransparent(buffer, char29, x, y, 4, 11, scale); x += 5;
        } else if str[j] == '*' { DrawBitmapTransparent(buffer, char2A, x, y, 8, 11, scale); x += 9;
        } else if str[j] == '+' { DrawBitmapTransparent(buffer, char2B, x, y, 8, 11, scale); x += 9;
        } else if str[j] == ',' { DrawBitmapTransparent(buffer, char2C, x, y, 4, 11, scale); x += 5;
        
        } else if str[j] == '.' { DrawBitmapTransparent(buffer, char2E, x, y, 4, 11, scale); x += 5;
        } else if str[j] == '/' { DrawBitmapTransparent(buffer, char2F, x, y, 7, 11, scale); x += 8;
        
        } else if str[j] == ':' { DrawBitmapTransparent(buffer, char3A, x, y, 4, 11, scale); x += 5;
        } else if str[j] == ';' { DrawBitmapTransparent(buffer, char3B, x, y, 4, 11, scale); x += 5;
        } else if str[j] == '<' { DrawBitmapTransparent(buffer, char3C, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '=' { DrawBitmapTransparent(buffer, char3D, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '>' { DrawBitmapTransparent(buffer, char3E, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '?' { DrawBitmapTransparent(buffer, char3F, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '@' { DrawBitmapTransparent(buffer, char40, x, y, 7, 11, scale); x += 8;
        
        } else if str[j] == '[' { DrawBitmapTransparent(buffer, char5B, x, y, 4, 11, scale); x += 5;
        } else if str[j] == '\\' { DrawBitmapTransparent(buffer, char5C, x, y, 7, 11, scale); x += 8;
        } else if str[j] == ']' { DrawBitmapTransparent(buffer, char5D, x, y, 4, 11, scale); x += 5;
        } else if str[j] == '^' { DrawBitmapTransparent(buffer, char5E, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '_' { DrawBitmapTransparent(buffer, char5F, x, y, 7, 11, scale); x += 8;
        } else if str[j] == '`' { DrawBitmapTransparent(buffer, char60, x, y, 4, 11, scale); x += 5;
        } else if str[j] == 0x0A { x = x_start; y += 12 + interval; continue;
        //} else if str[j] == 0x0D { x = x_start; y += 12 + interval; continue;
        } else if str[j] == 0x09 { x += 32; continue;
        } else if str[j] == 0x01 { DrawBitmapTransparent(buffer, charMinimized, x, y, 7, 7, scale); x += 7;
        } else { DrawBitmapTransparent(buffer, charUndefined, x, y, 7, 7, scale); x += 7;}
        // Перенос
        /*if x + 8 > xright_loc {
        x = x_start; y += 12 + interval; continue;
        }*/
    //}
    }
}


func showBMP(buffer []uint8, Wall []byte, x int, y int){
	if Wall != nil {
		pos := 0x1000000*int(Wall[0x0D]) + 0x10000*int(Wall[0x0C]) + 0x100*int(Wall[0x0B]) + int(Wall[0x0A]) //0x36
		width := 0x1000000*int(Wall[0x15]) + 0x10000*int(Wall[0x14]) + 0x100*int(Wall[0x13]) + int(Wall[0x12])
		height := 0x1000000*int(Wall[0x19]) + 0x10000*int(Wall[0x18]) + 0x100*int(Wall[0x17]) + int(Wall[0x16])
    	for i := 0; i < height; i++ {
        	for j := 0; j < width; j++ {
        		if 0x10000*uint32(Wall[pos+2]) + 0x100*uint32(Wall[pos+1]) + uint32(Wall[pos]) != 0xFEFEFE {
        			if x + j >= xleft_loc && x + j <= xright_loc && y + i >= ytop_loc && y + i <= ybottom_loc {
						PutPixel(nil, x + j, y + height - i, 0x10000*uint32(Wall[pos+2]) + 0x100*uint32(Wall[pos+1]) + uint32(Wall[pos])) 
					}	
				}
				pos += 3
			}
			if width%4 != 0 {
				pos += (((width/4 + 1)*4) - width)*3
			}
		}
	}
}




