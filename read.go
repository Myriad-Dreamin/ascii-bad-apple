package main

import (
	"encoding/binary"
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"log"
	"os"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}


func Read() {
	sugar.WithFile(func(VideoSrc *os.File) {
		sugar.WithWriteFile(func(VideoDst *os.File) {
			for i := 0; i < 5479; i++ {
				for j := 0; j < 140; j++ {
					_, err := fmt.Fscanln(VideoSrc, &bitFrame[i][j][0], &bitFrame[i][j][1], &bitFrame[i][j][2], &bitFrame[i][j][3], &bitFrame[i][j][4])
					if err != nil {
						log.Fatal(err)
					}
					for k := 0; k < 5; k++ {
						err = binary.Write(VideoDst, binary.BigEndian, bitFrame[i][j][k])
						if err != nil {
							log.Fatal(err)
						}
					}
				}
				if i%100 == 0 {
					fmt.Printf("Frame %d/5479 is loded!\n", i)
				}
			}
		}, "binary_ascii_badapple.txt")
	}, "ascii_badapple.txt")
}
