package main

import "time"

var bitTable [64]uint64
var frameTable [5500]time.Duration
var lineTable [2][70]int
var bitFrame [5479][140][5]uint64
var frame[140][320]uint8

const (
	FrameTime = 40 * time.Millisecond
	FrameSize = 140 * 5 * 8
	FrameCount = 5479
)