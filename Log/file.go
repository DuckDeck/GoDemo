package main

import "os"

type FileLogger struct {
	filePath    string
	fileObj     *os.File
	fileName    string
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// func NewFileLogger(l)
