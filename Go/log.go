package main

import (
	"log"
	"os"
	"time"
)

var f *os.File

func ensureDir(path string) {
	mode := os.ModePerm
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
}

func getWorkingDir() (dir string) {
	var err error
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getTimeForLogName() (strTime string) {
	currentTime := time.Now()
	strTime = currentTime.Format("[2006-01-02 15-04-05]")
	return
}

func initLog(name string) {
	var err error
	dir := getWorkingDir()
	dir = dir + "/log"
	timeStr := getTimeForLogName()
	path := dir + "/" + name + timeStr + ".log"
	ensureDir(dir)
	f, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
