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


func Line(x1 int, y1 int, x2 int, y2 int){
    var dx int = int(math.Abs(float64(x2-x1)))
    var dy int = int(math.Abs(float64(y2-y1)))
        var inc1 int
        var inc2 int
        var d int
        var s int
        var a int
        var aend int
    if dx > dy {
        if x1 < x2 {
            a = BITMAP_WIDTH*y1 + x1; aend = BITMAP_WIDTH*y2 + x2;
            if y2 >= y1 { 
            	s = BITMAP_WIDTH + 1;
            } else {
            	s = -BITMAP_WIDTH + 1;
            }
        } else {
            a = BITMAP_WIDTH*y2 + x2; aend = BITMAP_WIDTH*y1 + x1;
            if y1 >= y2 {
            	s = BITMAP_WIDTH + 1; 
            } else { 
            	s = -BITMAP_WIDTH + 1;
            }
        }
        inc1 = 2*(dy - dx);
        inc2 = 2*dy;
        d = 2*dy - dx;
        FillLB(a, 1, CC);
        for ; a != aend; {
        //while(a != aend){
            if d > 0 {
                d += inc1;
                a += s;
            } else {
                d += inc2;
                a++;
            }
            FillLB(a, 1, CC);
        }
    } else {
        if y1 < y2 {
            a = BITMAP_WIDTH*y1 + x1; aend = BITMAP_WIDTH*y2 + x2;
            if x1 <= x2 {
            	s = BITMAP_WIDTH + 1; 
            } else {
            	s = BITMAP_WIDTH - 1;
            }
        } else {
            a = BITMAP_WIDTH*y2 + x2; aend = BITMAP_WIDTH*y1 + x1;
            if x1 > x2 {
            	s = BITMAP_WIDTH + 1;
            } else {
            	s = BITMAP_WIDTH - 1;
            }
        }
        inc1 = 2*(dx - dy);
        inc2 = 2*dx;
        d = 2*dx - dy;
        FillLB(a, 1, CC);
        for ; a != aend; {
       // while(a != aend){
            if d > 0 {
                d += inc1;
                a += s;
            } else {
                d += inc2;
                a += BITMAP_WIDTH;
            }
            FillLB(a, 1, CC);
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






