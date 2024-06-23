package main

import (
    //"fmt"
    //"syscall/js"
    "math/rand"
    //"strconv"

)


// Глобальные переменные для размеров экрана и буфера графики
var (
	BITMAP_WIDTH  int = 1920 // 1920 //1600 //1024 //640 
	BITMAP_HEIGHT int = 1080 // 1080 //900 //768 //480
	SIZE          int = BITMAP_WIDTH * BITMAP_HEIGHT
	GETMAX_X      int = BITMAP_WIDTH - 1
	GETMAX_Y      int = BITMAP_HEIGHT - 1
	BUFFER_SIZE   int = SIZE * 4 * 2 // Размер буфера в байтах (RGBA * 2)
)

// Глобальный буфер для графики
var graphicsBuffer []uint8 = make([]uint8, BUFFER_SIZE, BUFFER_SIZE)


// Функция для заполнения части буфера определенным цветом
func FillLB(buffer []uint8, start int, count int, value uint32){
	var buf []uint8
    if buffer == nil {
        buf = graphicsBuffer
    } else {
        buf = buffer
    }

    for i := start * 4; i < (start+count)*4; i += 4 {
        buf[i+0] = uint8(value >> 16)       // Красный
        buf[i+1] = uint8(value >> 8)        // Зеленый
        buf[i+2] = uint8(value)             // Синий
        buf[i+3] = 255                      // Альфа
    }
}


// Функция для заполнения буфера случайными цветами
func FillLBrnd(){
    for i := 0; i < BUFFER_SIZE; i++ {
        if graphicsBuffer[i] != 1 {
        	graphicsBuffer[i] = uint8(rand.Intn(255))
        }
    }
}


// Функция для получения цвета пикселя по координатам
func GetPixelgl(buffer []uint8, x int, y int) uint32 { 
    var val uint32 = 0
    squareNumber := (y * BITMAP_WIDTH) + x
    squareRgbaIndex := squareNumber * 4

    val += uint32(graphicsBuffer[squareRgbaIndex+0]) << 16 // Красный
    val += uint32(graphicsBuffer[squareRgbaIndex+1]) << 8  // Зеленый
    val += uint32(graphicsBuffer[squareRgbaIndex+2])       // Синий
    val += uint32(graphicsBuffer[squareRgbaIndex+3]) << 24 // Альфа

    return val
}


// Функция, возвращающая указатель на буфер графики для WebAssembly
//export getGraphicsBufferPointer
func getGraphicsBufferPointer() *uint8 {		//*[BUFFER_SIZE]uint8
	return &graphicsBuffer[0]
}


// Функция, возвращающая размер буфера графики для WebAssembly
//export getGraphicsBufferSize
func getGraphicsBufferSize() int {
	return BUFFER_SIZE
}



