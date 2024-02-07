package main

import (
    "os"
    "fmt"
)

const (
    TEMP_MS_FILENAME string = "tempMS.ms"
    TEMP_PS_FILENAME string = "tempPS.ps"
    IMAGE_DIR string = "images"
)

var HOME_DIR string
var TMP_DIR string

func doMkdir(dirname string) {
    var err error
    err = os.Mkdir(dirname, 0755)
    if err != nil {
        panic(err)
    }
}

func Init() {
    var err error
    HOME_DIR, err = os.UserHomeDir()
    if err != nil {
        panic(err)
    }
    
    TMP_DIR = fmt.Sprintf("/tmp/mdToGroff%d", os.Getpid())
    
    doMkdir(TMP_DIR)
    doMkdir(TMP_DIR + "/" + IMAGE_DIR)
    
}

func Cleanup() {
    var err error
    err = os.RemoveAll(TMP_DIR)
    if err != nil {
        panic(err)
    }
}
