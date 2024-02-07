package main

import (
    //"github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/html"
    "github.com/gomarkdown/markdown/ast"    
    //"github.com/gomarkdown/markdown/parser"
    "io"
    "strings"
)
    
func newGroffRenderer() *html.Renderer {
    opts := html.RendererOptions {
        Flags: html.CommonFlags,
        RenderNodeHook: myRenderHook,
    }
    return html.NewRenderer(opts)
}

func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
    //fmt.Printf("%T\n", node)
    if text, ok := node.(*ast.Text); ok{
        renderText(w, text)
        return ast.GoToNext, true
    }
    if lineBreak, ok := node.(*ast.Hardbreak); ok{
        renderLineBreak(w, lineBreak)
        return ast.GoToNext, true
    }
    if emph, ok := node.(*ast.Emph); ok{
        renderItalic(w, emph, entering)
        return ast.GoToNext, true
    }
    if strong, ok := node.(*ast.Strong); ok{
        renderBold(w, strong, entering)
        return ast.GoToNext, true
    }
    if monospace, ok := node.(*ast.Code); ok{
        renderMonospace(w, monospace, entering)
        return ast.GoToNext, true
    }
    if para, ok := node.(*ast.Paragraph); ok{
        renderParagraph(w, para, entering)
        return ast.GoToNext, true
    }
    if heading, ok := node.(*ast.Heading); ok{
        renderHeading(w, heading, entering)
        return ast.GoToNext, true
    }
    if quote, ok := node.(*ast.BlockQuote); ok{
        renderBlockQuote(w, quote, entering)
        return ast.GoToNext, true
    }
    if table, ok := node.(*ast.Table); ok{
        renderTable(w, table, entering)
        return ast.GoToNext, true
    }
    if tableRow, ok := node.(*ast.TableRow); ok{
        renderTableRow(w, tableRow, entering)
        return ast.GoToNext, true
    }
    if tableCell, ok := node.(*ast.TableCell); ok{
        renderTableCell(w, tableCell, entering)
        return ast.GoToNext, true
    }
    if _, ok := node.(*ast.TableHeader); ok{
        return ast.GoToNext, true
    }
    if _, ok := node.(*ast.TableBody); ok{
        return ast.GoToNext, true
    }
    if code, ok := node.(*ast.CodeBlock); ok{
        renderCodeBlock(w, code)
        return ast.GoToNext, true
    }
    if list, ok := node.(*ast.List); ok{
        renderList(w, list, entering)
        return ast.GoToNext, true
    }
    if listItem, ok := node.(*ast.ListItem); ok{
        renderListItem(w, listItem, entering)
        return ast.GoToNext, true
    }
    if image, ok := node.(*ast.Image); ok{
        if entering {
            inImage = true
            //fmt.Printf("IMAGE ALT TEXT: %s\n", image.Children[0].AsLeaf().Literal)
            //fmt.Printf("IMAGE PATH: %s\n", image.Destination)
            // Split image path to get only filename
            imagePathSplit := strings.Split(string(image.Destination), "/")
            // Create full final filename
            var startOfExtIndex int = strings.LastIndex(imagePathSplit[len(imagePathSplit)-1], ".")
            var finalPath string = TMP_DIR + "/" + IMAGE_DIR + "/" + imagePathSplit[len(imagePathSplit)-1][:startOfExtIndex]
            // Add original path to imageLocs
            imageLocs = append(imageLocs, string(image.Destination))
            
            // Check if finalPath is already in imageDests, if it is, modify it until it's not
            for sliceContains(imageDests, finalPath) {
                finalPath = finalPath + "A"
            }
            // Once done, add to imageDests
            imageDests = append(imageDests, finalPath)
            
            // Write image (left align) and alt text to ms text
            io.WriteString(w, ".PDFPIC -L " + finalPath + ".pdf" + "\n.br\n" + string(image.Children[0].AsLeaf().Literal) + "\n")
        } else { inImage = false }
        return ast.GoToNext, true
    }
    return ast.GoToNext, false
}
