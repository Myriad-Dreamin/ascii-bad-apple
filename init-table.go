package main


func initTable() {
	for i := 0; i < 64; i++ {
		bitTable[i] = uint64(1) << uint64(i)
	}

	frameTable[0] = FrameTime
	for i := 1; i < 5500; i++ {
		frameTable[i] = frameTable[i-1] + FrameTime
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 70; j++ {
			lineTable[i][j] = 2*j + i
		}
	}
}
