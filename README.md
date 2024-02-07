# Go Markdown To PDF

![Static Badge](https://img.shields.io/badge/Linux-grey?logo=linux)
![Static Badge](https://img.shields.io/badge/Golang-007D9C)
![Static Badge](https://img.shields.io/badge/Usage-PDF_Compilation-blue)
![GitHub Release](https://img.shields.io/github/v/tag/randomcoder67/Go-MarkdownToPDF)

Golang program to convert Markdown files to PDF format, without using LaTeX or Pandoc. compilation is done using Groff, with Ghostscript used to convert PostScript file to PDF.

## Dependancies

* ImageMagick
* Groff
* Ghostscript

### Arch Linux

`sudo pacman -S imagemagick groff ghostscript`

### Debain/Ubuntu

`sudo apt imagemagick groff ghostscript`

### Fedora

`sudo dnf install ImageMagick groff ghostscript`

## Installation

`make`  
`make install`  

## Usage

`mdtopdf input.md output.pdf`
