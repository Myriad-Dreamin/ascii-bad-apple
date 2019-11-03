package main

import (
	"encoding/binary"
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"log"
	"os"
	"sync"
	"time"
)

func write(l, r int) {
	sugar.WithFile(func(VideoSrc *os.File) {
		ret, err := VideoSrc.Seek(int64(l * FrameSize), 0)
		if err != nil || ret != int64(l * FrameSize) {
			log.Fatal(err)
		}
		for i := l; i < r; i++ {
			err = binary.Read(VideoSrc, binary.BigEndian, &bitFrame[i])
			if err != nil {
				log.Fatal(err)
			}
		}
	}, "binary_ascii_badapple.txt")
}


func Write() {
	initTable()
	var wg sync.WaitGroup

	for i := 0; i < 55; i++ {
		wg.Add(1)
		go func(i int) {
			write(i * 100, min(FrameCount, (i+1) * 100))
			wg.Done()
		}(i)
	}
	wg.Wait()
	sugar.WithWriteFile(func(logFile *os.File) {

		var start, end, frameStart = time.Now(), time.Now(), time.Now()
		for i := 0; i < FrameCount; i++ {
			frameStart = time.Now()
			outputInterlace(i)
			for {
				end = time.Now()
				if end.Sub(start) < frameTable[i] {
					continue
				}
				_, err := fmt.Fprintf(logFile, "frame:%d/5479  totaltime:%v  frame time:%v  delay_time:%v\n", i, end.Sub(start), end.Sub(frameStart), end.Sub(start) - frameTable[i])
				if err != nil {
					log.Fatalln("error!", err)
				}
				break
			}
		}
	}, "test.log")
}
