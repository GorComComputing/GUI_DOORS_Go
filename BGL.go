package main

import (
	"math"
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


func ClearDevice(){
    FillLB(0, SIZE, BC)
}


func PutPixel(x int, y int, Color int){
    FillLB((BITMAP_WIDTH*y + x), 1, Color);
}


func LinePP(x1 int, y1 int, x2 int, y2 int){
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
        PutPixel(x, y, CC);
        for ;x < xend; {
        //while(x < xend){
            x++;
            if d > 0 {
                d += inc1;
                y += s;
            } else{
                d += inc2;
            }
            PutPixel(x, y, CC);
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
        PutPixel(x, y, CC);
        for ;y < yend; {
        //while(y < yend){
            y++;
            if d > 0 {
                d += inc1;
                x += s;
            } else {
                d += inc2;
            }
            PutPixel(x, y, CC);
        }
    }
}


func HLine(x1 int, y int, x2 int){
    if x1 < x2 {
        FillLB(BITMAP_WIDTH*y + x1, x2 - x1 + 1, CC);
    } else {
        FillLB(BITMAP_WIDTH*y + x2, x1 - x2 + 1, CC);
	}
}


func Circle(xc int, yc int, R int){
    var x int = 0;
    var y int = R;
    var d int = 3 - 2*R;
    Pixel8(xc, yc, 0, R);
    for ; x < y ; {
    //while(x < y){
        if d < 0 {
            d = d + 4*x + 6;
        } else {
            d = d + 4*(x - y) + 10;
            y--;
        }
        x++;
        Pixel8(xc, yc, x, y);
    }
}


func Pixel8(xc int, yc int, x int, y int){
    PutPixel(xc + x, yc + y, CC);
    PutPixel(xc - x, yc + y, CC);
    PutPixel(xc + x, yc - y, CC);
    PutPixel(xc - x, yc - y, CC);

    PutPixel(xc + y, yc + x, CC);
    PutPixel(xc - y, yc + x, CC);
    PutPixel(xc + y, yc - x, CC);
    PutPixel(xc - y, yc - x, CC);
}


func FloodFillgl(x int, y int, bord int){
    var xl int = x;
    var xr int = x;
    var yy int;
    for ;GetPixelgl(xl, y) != bord; {
        xl--;
    }
    xl++;
    for ;GetPixelgl(xr, y) != bord; {
        xr++;
    }
    xr--;
    if xl < xr {
        HLine(xl, y, xr);
    }
    yy = y - 1;
    
    
    for ok := true; ok; ok = (yy <= y + 1) {
	  x = xr;
      for ;x >= xl; {
      	for ;(x >= xl) && (GetPixelgl(x, yy) == bord) || (GetPixelgl(x, yy) == CC); {
            x--;
        }
        if x >= xl {
            FloodFillgl(x, yy, bord);
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


func FillPoly(n int,  p []tPoint){
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
            HLine(YXbuf[y].x[i], y, YXbuf[y].x[i + 1]);
        }
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


func DrawBitmapTransparent(monster string, xstart int, ystart int, sizeX int, sizeY int, scale int){
    var x int = xstart;
    var y int = ystart;

    for i := 0; i < sizeY; i++ {
        for scaleY := scale; scaleY > 0; scaleY-- {
        for j := 0; j < sizeX; j++ {
            if(monster[i*sizeX + j] == 'p'){
                for scaleX := scale; scaleX > 0; scaleX--{
                PutPixel(x + j, y + i, CC);
                if scaleX > 1 { x++;}
                }
            } else {
                for scaleX := scale; scaleX > 0; scaleX-- {
                //PutPixel(x + j, y + i, BC);
                if scaleX > 1 {x++;}
                }
            }
        }
        x = xstart;
        if scaleY > 1 {y++;}
        }
    }
}



func TextOutgl(str string, x int, y int, scale int){
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




    for i := 0; i < len(str); i++ {
        if str[i] == 'A' { DrawBitmapTransparent(charA, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'B' { DrawBitmapTransparent(charB, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'C' { DrawBitmapTransparent(charC, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'D' { DrawBitmapTransparent(charD, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'E' { DrawBitmapTransparent(charE, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'F' { DrawBitmapTransparent(charF, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'G' { DrawBitmapTransparent(charG, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'H' { DrawBitmapTransparent(charH, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'I' { DrawBitmapTransparent(charI, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'J' { DrawBitmapTransparent(charJ, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'K' { DrawBitmapTransparent(charK, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'L' { DrawBitmapTransparent(charL, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'M' { DrawBitmapTransparent(charM, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'N' { DrawBitmapTransparent(charN, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'O' { DrawBitmapTransparent(charO, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'P' { DrawBitmapTransparent(charP, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'Q' { DrawBitmapTransparent(charQ, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'R' { DrawBitmapTransparent(charR, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'S' { DrawBitmapTransparent(charS, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'T' { DrawBitmapTransparent(charT, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'U' { DrawBitmapTransparent(charU, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'V' { DrawBitmapTransparent(charV, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'W' { DrawBitmapTransparent(charW, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'X' { DrawBitmapTransparent(charX, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'Y' { DrawBitmapTransparent(charY, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == 'Z' { DrawBitmapTransparent(charZ, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == ' ' { DrawBitmapTransparent(charSpace, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '-' { DrawBitmapTransparent(charLine, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '0' { DrawBitmapTransparent(char0, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '1' { DrawBitmapTransparent(char1, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '2' { DrawBitmapTransparent(char2, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '3' { DrawBitmapTransparent(char3, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '4' { DrawBitmapTransparent(char4, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '5' { DrawBitmapTransparent(char5, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '6' { DrawBitmapTransparent(char6, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '7' { DrawBitmapTransparent(char7, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '8' { DrawBitmapTransparent(char8, x+7*i*scale, y, 7, 7, scale);}
        if str[i] == '9' { DrawBitmapTransparent(char9, x+7*i*scale, y, 7, 7, scale);}
    }
}






