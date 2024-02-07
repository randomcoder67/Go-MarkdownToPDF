package main

import (
    "os"
    "os/exec"
    "strings"
    "strconv"
    "fmt"
)

var _ = fmt.Println

func replaceTableHeaders(inputMS string) string {
    lenA := len(tableAlignmentLengths)
    for i:=0; i<lenA; i++ {
        var stringToReplace string = "ALIGNMENT_TO_REPLACE" + strconv.Itoa(lenA-1-i)
        var toReplaceWith string = ""
        for j:=0; j<tableAlignmentLengths[lenA-1-i]; j++ {
            toReplaceWith = toReplaceWith + "l "
        }
        toReplaceWith = toReplaceWith + "."
        toReplaceWith = strings.ReplaceAll(toReplaceWith, " .", ".")
        //fmt.Printf("%s, %s\n", stringToReplace, toReplaceWith)
        inputMS = strings.ReplaceAll(inputMS, stringToReplace, toReplaceWith)
    }
    return inputMS
}

func replaceImages(inputMS string) {
    for i, originalPath := range imageLocs {
        fmt.Println(originalPath, imageDests[i])
        cmd := exec.Command("convert", "-density", "200", "-units", "PixelsPerInch", originalPath, imageDests[i] + ".pdf")
        err := cmd.Start(); if err != nil {
	        panic(err)
	    }
	    cmd.Wait()
    }
}

func readFile(filename string) []byte {
    var dat []byte
    var err error
    dat, err = os.ReadFile(filename)
    
    if err != nil {
        panic(err)
    }
    
    return dat
}

func readFileToString(filename string) string {
    var dat []byte = readFile(filename)
    return string(dat)
}

func sliceContains(sliceA []string, stringA string) bool {
    for _, entry := range sliceA {
        if stringA == entry { return true }
    }
    return false
}
