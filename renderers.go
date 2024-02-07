package main

import (
    //"github.com/gomarkdown/markdown"
    //"github.com/gomarkdown/markdown/html"
    "github.com/gomarkdown/markdown/ast"    
    //"github.com/gomarkdown/markdown/parser"
    "io"
    "strconv"
    "regexp"
)

var inBlockQuote bool = false
var tableAlignmentLengths []int
var currentTableLength int = 0
var pastTableLength int = 0
var inStrong bool = false
var orderedList bool = false
var atFirstListEntry bool = false
var inList bool = false
var inImage bool = false

var imageLocs = []string{}
var imageDests = []string{}

func renderList(w io.Writer, p *ast.List, entering bool) {
    if entering {
        inList = true
        atFirstListEntry = true
        if p.ListFlags == 17 {
            orderedList = true
            var startNumber string = strconv.Itoa(p.Start + 1) // Glitch here, not working
            //fmt.Println(startNumber)
            io.WriteString(w, ".nr step " + startNumber + " 1\n")
        } else {
            orderedList = false
        }
    } else {
        inList = false
    }
}

func renderListItem(w io.Writer, p *ast.ListItem, entering bool) {
    if entering {
        if orderedList {
            if atFirstListEntry {
                io.WriteString(w, ".IP \\n[step] 2\n")
            } else {
                io.WriteString(w, ".IP \\n+[step]\n")
            }
        } else {
            if atFirstListEntry {
                io.WriteString(w, ".IP \\[bu] 2\n")
            } else {
                io.WriteString(w, ".IP \\[bu]\n")
            }
        }
        atFirstListEntry = false
    } else {
        io.WriteString(w, "\n")
    }
}

func renderText(w io.Writer, p *ast.Text) {
    var textString string = string(p.Literal)
    if inStrong {
        io.WriteString(w, "\n.B \"" + textString + "\"\\c\n")
    } else if !inImage {
        io.WriteString(w, textString)
    }
}

func renderLineBreak(w io.Writer, p *ast.Hardbreak) {
    io.WriteString(w, "\n.br\n")
}

func renderItalic(w io.Writer, p *ast.Emph, entering bool) {
    if entering {
        if inStrong {
            inStrong = false
            io.WriteString(w, "\n.BI \"")
        } else {
            io.WriteString(w, "\n.I \"")
        }
    } else {
        io.WriteString(w, "\"\\c\n")
    }
}

func renderBold(w io.Writer, p *ast.Strong, entering bool) {
    if entering {
        inStrong = true
    } else {
        inStrong = false
    }
}

func renderMonospace(w io.Writer, p *ast.Code, entering bool) {
    var codeString string = string(p.AsLeaf().Literal)
    codeString = "\n.CW \"" + codeString + "\"\\c\n"
    io.WriteString(w, codeString)
}

func renderTable(w io.Writer, p *ast.Table, entering bool) {
    numTableString := strconv.Itoa(len(tableAlignmentLengths))
    if entering {
        io.WriteString(w, ".TS\ntab(;) allbox;\nALIGNMENT_TO_REPLACE" + numTableString + "\n")
    } else {
        tableAlignmentLengths = append(tableAlignmentLengths, pastTableLength)
        io.WriteString(w, ".TE\n")
    }
}

func renderTableRow(w io.Writer, p *ast.TableRow, entering bool) {
    if !entering {
        pastTableLength = currentTableLength
        currentTableLength = 0
        io.WriteString(w, "\n")
    }
}

func renderTableCell(w io.Writer, p *ast.TableCell, entering bool) {
    if !entering {
        currentTableLength++
        io.WriteString(w, ";")
    }
}

func renderHeading(w io.Writer, p *ast.Heading, entering bool) {
    if entering {
        var fontSize string = strconv.Itoa(18-(p.Level*5))
        io.WriteString(w, "\n.ps +" + fontSize + "\n.B\n")
    } else {
        io.WriteString(w, "\n.br\n.ps\n")
    }
}

func renderBlockQuote(w io.Writer, p *ast.BlockQuote, entering bool) {
    if entering {
        inBlockQuote = true
        io.WriteString(w, ".LP\n> ")
    } else {
        inBlockQuote = false
        io.WriteString(w, "\n")
    }
}

func renderParagraph(w io.Writer, p *ast.Paragraph, entering bool) {
    if !inBlockQuote && !inList {
        if entering {
            io.WriteString(w, ".LP\n")
        } else {
            io.WriteString(w, "\n")
        }
    }
}

func renderCodeBlock(w io.Writer, p *ast.CodeBlock) {
    var codeString string = string(p.AsLeaf().Literal)
    //var language string = string(p.Info)
    codeString = ".CW\n" + codeString
    var re = regexp.MustCompile(`\n([^\s])`)
    codeString = re.ReplaceAllString(codeString, "\n.br\n$1")
    io.WriteString(w, codeString)
}
