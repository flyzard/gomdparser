package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
)

func main() {

    markdownFilePath, latestReleaseHtmlFilePath, htmlFilePath, nReleases := getArguments(os.Args)

    content, err := ioutil.ReadFile(markdownFilePath)
    if err != nil {
        log.Fatal(err)
    }

    // Convert the full release notes file to MD
    convertContentToMd(content, htmlFilePath)
    
    // Get only the latest release, given that they are separated by ##
    lenght := strings.Index(string(content)[nReleases + 1:], "##")
    content = content[0:lenght]

    // Convert the nReleases to html
	convertContentToMd(content, latestReleaseHtmlFilePath)
}

func getArguments(givenArgs []string) (string, string, string, int) {

    // default names when no arguments passed
    args := [3]string{"release_notes.md", "www/latest_release.html", "www/releases.html" }

    size := len(givenArgs)

   for i := 1; i < 4; i++ {
       if size >= i+1 {
           args[i-1] = givenArgs[i]
       }

       if i == 1 {
            if _, err := os.Stat(args[i-1]); err != nil {
                log.Fatal(err)
            }
        }
   }

    // for i := range args {
    //     if size >= i+2 {
    //         args[i] = givenArgs[i+1]
    //     }

    //     if i == 1 {
    //         if _, err := os.Stat(args[i]); err != nil {
    //             log.Fatal(err)
    //         }
    //     }
    // }

    nReleases := 1
    if size > 4 {
        nReleases, _ = strconv.Atoi(givenArgs[4])
    }

    return args[0], args[1], args[2], nReleases
}


func convertContentToMd(content []byte, filPath string) {
    html := markdown.ToHTML(content, nil, nil)
    f, err := os.OpenFile(filPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        log.Fatal(err)
    }

    f.WriteAt(html, 0)

    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}
