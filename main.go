package main

import (
    "os"
    "fmt"
    "github.com/gomarkdown/markdown"
    //"github.com/gomarkdown/markdown/html"
    //"github.com/gomarkdown/markdown/ast"    
    "github.com/gomarkdown/markdown/parser"
    "os/exec"
    "strings"
)

func printHelp() {
    fmt.Printf("Usage: \n  mdtogroff input.md output.pdf\n")
}

func mdToMS(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)
	// create HTML renderer with extensions
	renderer := newGroffRenderer()
    //fmt.Println(doc)
	return markdown.Render(doc, renderer)
}

func parseArgs(args []string) (string, string) {
    if len(args) != 2 {
        fmt.Println("Error, incorrect arguments")
        printHelp()
        Cleanup()
        os.Exit(1)
    }
    return args[0], args[1]
}

func runGroff(inputMSText string, outputFilename string) {
    cmd := exec.Command("groff", "-ms", "-tb", "-UT", "pdf")
    cmd.Stdin = strings.NewReader(inputMSText)
    outfile, err := os.Create(outputFilename)
    defer outfile.Close()
    cmd.Stdout = outfile
    
    err = cmd.Start(); if err != nil {
        panic(err)
    }
    cmd.Wait()
}

func main() {
    // Init and parse arguments
    Init()
    defer Cleanup()
    var inputFilename, outputFilename string = parseArgs(os.Args[1:]) 
    _, _ = inputFilename, outputFilename
    
    // Read markdown and perform parser conversion
    var inputMarkdown []byte = readFile(inputFilename)
    var outputManuscript []byte = mdToMS(inputMarkdown)
    
    // Perform additional "manual" conversion
    //htmlFinal = ".nr HM 2c\n" + htmlFinal
    // Images
    //replaceImages("")
    // Tables
    
    // Run commands to convert to PostScript then PDF
    _ = outputManuscript
    
    runGroff(string(outputManuscript), outputFilename)
}
