package main

import (
	"math"
	"strings"
	//"fmt"
)


var CC int 
var BC int 

var Black int = 0x000000
var White int = 0xFFFFFF
var Red int = 0xFF0000
var Green int = 0x00FF00
var Blue int = 0x0000FF
var Yellow int = 0xFFFF00
var Magenta int = 0xFF00FF
var Cyan int = 0x00FFFF


func SetColor(Color int){
    CC = Color
}


func SetBackColor(Color int){
    BC = Color
}


func ClearDevice(buffer []uint8){
    FillLB(buffer, 0, SIZE, BC)
}


func DrawPutPixel(buffer []uint8, x int, y int, Color int){
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


func FloodFillgl(buffer []uint8, x int, y int, bord int){
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


var YXbuf [GETMAX_Y]tXbuf


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


func PutPixel(buffer []uint8, x int, y int, C int){
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
var charA string = 
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pppppp " +
    "pp  pp " +
    "pp  pp " 
    ;//7x7
var charB string =
    "ppppp  " +
    "pp  pp " +
    "pp  pp " +
    "ppppp  " +
    "pp  pp " +
    "pp  pp " +
    "ppppp  "
    ;//7x7
var charC string =
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    "pp     " +
    "pp  pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var charD string =
    "ppppp  " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "ppppp  "
    ;//7x7
var charE string =
    "pppppp " +
    "pp     " +
    "pp     " +
    "ppppp  " +
    "pp     " +
    "pp     " +
    "pppppp "
    ;//7x7
var charF string =
    "pppppp " +
    "pp     " +
    "pp     " +
    "ppppp  " +
    "pp     " +
    "pp     " +
    "pp     "
    ;//7x7
var charG string =
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    "pp     " +
    "pp ppp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var charH string =
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pppppp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp "
    ;//7x7
var charI string =
    " pppp  " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    " pppp  "
    ;//7x7
var charJ string =
    "   ppp " +
    "    pp " +
    "    pp " +
    "    pp " +
    "pp  pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var charK string =
    "pp  pp " +
    "pp  pp " +
    "pp pp  " +
    "pppp   " +
    "pp pp  " +
    "pp  pp " +
    "pp  pp "
    ;//7x7
var charL string =
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pppppp "
    ;//7x7
var charM string =
    "p   p  " +
    "pp pp  " +
    "ppppp  " +
    "p p p  " +
    "p   p  " +
    "p   p  " +
    "p   p  "
    ;//7x7
var charN string =
    "pp  pp " +
    "pp  pp " +
    "ppp pp " +
    "pppppp " +
    "pp ppp " +
    "pp  pp " +
    "pp  pp "
    ;//7x7
var charO string =
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var charP string =
    "ppppp  " +
    "pp  pp " +
    "pp  pp " +
    "ppppp  " +
    "pp     " +
    "pp     " +
    "pp     "
    ;//7x7
var charQ string =
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " pppp  " +
    "   ppp "
    ;//7x7
var charR string =
    "ppppp  " +
    "pp  pp " +
    "pp  pp " +
    "ppppp  " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp "
    ;//7x7
var charS string =
    " pppp  " +
    "pp  pp " +
    "pp     " +
    " pppp  " +
    "    pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var charT string =
    "pppppp " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   "
    ;//7x7
var charU string =
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var charV string =
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " p  p  " +
    " pppp  " +
    "  pp   "
    ;//7x7
var charW string =
    "p    p " +
    "p    p " +
    "p    p " +
    "p pp p " +
    "p pp p " +
    " pppp  " +
    " p  p  "
    ;//7x7
var charX string =
    "pp  pp " +
    "pp  pp " +
    " pppp  " +
    "  pp   " +
    " pppp  " +
    "pp  pp " +
    "pp  pp "
    ;//7x7
var charY string =
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " pppp  " +
    "  pp   " +
    "  pp   " +
    "  pp   "
    ;//7x7
var charZ string =
    "pppppp " +
    "    pp " +
    "   pp  " +
    "  pp   " +
    " pp    " +
    "pp     " +
    "pppppp "
    ;//7x7
var charSpace string =
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       " +
    "       "
    ;//7x7
var charLine string =
    "       " +
    "       " +
    "       " +
    " ppppp " +
    "       " +
    "       " +
    "       "
    ;//7x7
var char0 string =
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    "pp ppp " +
    "ppp pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var char1 string =
    " pp    " +
    "ppp    " +
    " pp    " +
    " pp    " +
    " pp    " +
    " pp    " +
    "pppp   "
    ;//7x7
var char2 string =
    " pppp  " +
    "pp  pp " +
    "    pp " +
    "   ppp " +
    " ppp   " +
    "pp     " +
    "pppppp "
    ;//7x7
var char3 string =
    " pppp  " +
    "pp  pp " +
    "    pp " +
    "  ppp  " +
    "    pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var char4 string =
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    "pp  pp " +
    " ppppp " +
    "    pp " +
    "    pp "
    ;//7x7
var char5 string =
    "pppppp " +
    "pp     " +
    " ppp   " +
    "    pp " +
    "pp  pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var char6 string =
    " pppp  " +
    "pp  pp " +
    "pp     " +
    "ppppp  " +
    "pp  pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var char7 string =
    "pppppp " +
    "    pp " +
    "   pp  " +
    "  pp   " +
    "  pp   " +
    "  pp   " +
    "  pp   "
    ;//7x7
var char8 string =
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var char9 string =
    " pppp  " +
    "pp  pp " +
    "pp  pp " +
    " ppppp " +
    "    pp " +
    "pp  pp " +
    " pppp  "
    ;//7x7
var charUndefined string =
    "ppppppp" +
    "p     p" +
    "p     p" +
    "p     p" +
    "p     p" +
    "p     p" +
    "ppppppp"
    ;//7x7
var char2points string =
    "       " +
    "   pp  " +
    "   pp  " +
    "       " +
    "   pp  " +
    "   pp  " +
    "       "
    ;//7x7
var char_vert_line string =
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     " +
    "pp     "
    ;//7x7

	var interval int = 2
	var x = x_start
	str = strings.ToUpper(str)
	i := 0
    for j := 0; j < len(str); j++ {
        if str[j] == 'A' { DrawBitmapTransparent(buffer, charA, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'B' { DrawBitmapTransparent(buffer, charB, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'C' { DrawBitmapTransparent(buffer, charC, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'D' { DrawBitmapTransparent(buffer, charD, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'E' { DrawBitmapTransparent(buffer, charE, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'F' { DrawBitmapTransparent(buffer, charF, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'G' { DrawBitmapTransparent(buffer, charG, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'H' { DrawBitmapTransparent(buffer, charH, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'I' { DrawBitmapTransparent(buffer, charI, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'J' { DrawBitmapTransparent(buffer, charJ, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'K' { DrawBitmapTransparent(buffer, charK, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'L' { DrawBitmapTransparent(buffer, charL, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'M' { DrawBitmapTransparent(buffer, charM, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'N' { DrawBitmapTransparent(buffer, charN, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'O' { DrawBitmapTransparent(buffer, charO, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'P' { DrawBitmapTransparent(buffer, charP, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'Q' { DrawBitmapTransparent(buffer, charQ, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'R' { DrawBitmapTransparent(buffer, charR, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'S' { DrawBitmapTransparent(buffer, charS, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'T' { DrawBitmapTransparent(buffer, charT, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'U' { DrawBitmapTransparent(buffer, charU, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'V' { DrawBitmapTransparent(buffer, charV, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'W' { DrawBitmapTransparent(buffer, charW, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'X' { DrawBitmapTransparent(buffer, charX, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'Y' { DrawBitmapTransparent(buffer, charY, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 'Z' { DrawBitmapTransparent(buffer, charZ, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == ' ' { DrawBitmapTransparent(buffer, charSpace, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '-' { DrawBitmapTransparent(buffer, charLine, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '0' { DrawBitmapTransparent(buffer, char0, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '1' { DrawBitmapTransparent(buffer, char1, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '2' { DrawBitmapTransparent(buffer, char2, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '3' { DrawBitmapTransparent(buffer, char3, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '4' { DrawBitmapTransparent(buffer, char4, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '5' { DrawBitmapTransparent(buffer, char5, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '6' { DrawBitmapTransparent(buffer, char6, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '7' { DrawBitmapTransparent(buffer, char7, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '8' { DrawBitmapTransparent(buffer, char8, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '9' { DrawBitmapTransparent(buffer, char9, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == ':' { DrawBitmapTransparent(buffer, char2points, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == '|' { DrawBitmapTransparent(buffer, char_vert_line, x+7*i*scale, y, 7, 7, scale);
        } else if str[j] == 13 { i = 0; y += 7 + interval; continue;
        } else { DrawBitmapTransparent(buffer, charUndefined, x+7*i*scale, y, 7, 7, scale);}
        i++
    }
}






