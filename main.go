package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

func main() {

    markdownFilePath := os.Args[1]
    htmlFilePath := os.Args[2]

    content, err := ioutil.ReadFile(markdownFilePath)
    if err != nil {
        log.Fatal(err)
    }

    // Get only the latest release, given that they are separated by ##
    lenght := strings.Index(string(content)[2:], "##")
    content = content[0:lenght]
	html := markdown.ToHTML(content, nil, nil)

	f, err := os.OpenFile(htmlFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        log.Fatal(err)
    }

	f.WriteAt(html, 0)

    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}
