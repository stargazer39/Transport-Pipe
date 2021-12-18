package main

import (
	"os"
	"runtime"
)

type ColorWriter struct {
	w *os.File
}

var ColorReset = "\033[0m"
var ColorRed = "\033[31m"
var ColorGreen = "\033[32m"
var ColorYellow = "\033[33m"
var ColorBlue = "\033[34m"
var ColorPurple = "\033[35m"
var ColorCyan = "\033[36m"
var ColorGray = "\033[37m"
var ColorWhite = "\033[97m"

func ColorInit() {
	if runtime.GOOS == "windows" {
		ColorReset = ""
		ColorRed = ""
		ColorGreen = ""
		ColorYellow = ""
		ColorBlue = ""
		ColorPurple = ""
		ColorCyan = ""
		ColorGray = ""
		ColorWhite = ""
	}
}

func NewColorWriter(w *os.File) ColorWriter {
	ColorInit()
	return ColorWriter{w}
}

func (c ColorWriter) Write(b []byte) (int, error) {
	n, err := c.w.Write([]byte(ColorCyan))

	if err != nil {
		return n, err
	}

	n2, err2 := c.w.Write(b)

	if err2 != nil {
		return n + n2, err2
	}

	n3, err3 := c.w.Write([]byte(ColorReset))

	if err3 != nil {
		return n + n2 + n3, err3
	}

	return n + n2 + n3, nil
}
