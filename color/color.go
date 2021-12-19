package color

import (
	"os"
	"runtime"
)

type Writer struct {
	w *os.File
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func Init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

func NewWriter(w *os.File) Writer {
	Init()
	return Writer{w}
}

func (c Writer) Write(b []byte) (int, error) {
	n, err := c.w.Write([]byte(Cyan))

	if err != nil {
		return n, err
	}

	n2, err2 := c.w.Write(b)

	if err2 != nil {
		return n + n2, err2
	}

	n3, err3 := c.w.Write([]byte(Reset))

	if err3 != nil {
		return n + n2 + n3, err3
	}

	return n + n2 + n3, nil
}
