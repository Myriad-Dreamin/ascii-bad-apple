package main

import(
	"fmt"
	"github.com/Azure/go-ansiterm/winterm"
	terminal "github.com/buger/goterm"
	"golang.org/x/sys/windows"
	"log"
	"os"
)

var handler windows.Handle

func gotoXY(x, y int16) {
	err := winterm.SetConsoleCursorPosition(uintptr(handler), winterm.COORD{
		X: x,
		Y: y,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func outputInterlaceT(i int) {
	var mode = i & 1
	for j := 0; j < 70; j++ {
		for k := 0; k < 5; k++ {
			for h := 0; h < 64; h++ {
				if (bitFrame[i][lineTable[mode][j]][k] & bitTable[h]) != 0 {
					frame[lineTable[mode][j]][k * 64 + h] = '*'
				} else {
					frame[lineTable[mode][j]][k * 64 + h] = ' '
				}
			}
		}
	}
	terminal.Clear()
	for j := 0; j < 70; j++ {
		terminal.MoveCursor(1, lineTable[mode][j] + 1)
		_, err := terminal.Println(string(frame[lineTable[mode][j]][:]))
		if err != nil {
			log.Fatal(err)
		}
	}
	terminal.Flush()
}


func outputInterlace(i int) {
	var mode = i & 1
	for j := 0; j < 70; j++ {
		for k := 0; k < 5; k++ {
			for h := 0; h < 64; h++ {
				if (bitFrame[i][lineTable[mode][j]][k] & bitTable[h]) != 0 {
					frame[lineTable[mode][j]][k * 64 + h] = '*'
				} else {
					frame[lineTable[mode][j]][k * 64 + h] = ' '
				}
			}
		}
	}
	for j := 0; j < 70; j++ {
		gotoXY(0, int16(lineTable[mode][j]))
		fmt.Println(string(frame[lineTable[mode][j]][:]))
	}
}

func init() {
	handler = windows.Handle(os.Stdout.Fd())
}
