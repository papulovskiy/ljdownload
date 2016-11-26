package main

import (
    "flag"
    "fmt"
    // "net/http"
    // "strings"
    "github.com/PuerkitoBio/goquery"
)

var seq, start, stop int

var url_pattern string
var output_file string

type post struct {
    url   string
    title string
    text  string
}

func get_url(number int) string {
    return fmt.Sprintf(url_pattern, number)
}

func parse(number int) {
    doc, err := goquery.NewDocument(get_url(number))
    if err != nil {
        // log.Fatal(err)
    } else {
        url, url_err := doc.Find("a.link__control").Html()
        title, title_err := doc.Find("div.article__title").Html()
        text, text_err := doc.Find("div.article__main-text").Html()
        if url_err != nil || title_err != nil || text_err != nil {
            return
        }
        d := post{
            url:   url,
            title: title,
            text:  text,
        }
        fmt.Println(d)
    }

}

func init() {
    flag.IntVar(&start, "start", 1, "Start post number")
    flag.IntVar(&stop, "stop", 1, "Stop post number")

    flag.StringVar(&url_pattern, "url", "https://ljsear.ch/savedcopy?post=%v", "URL pattern, use %v as a post number placeholder")
    flag.StringVar(&output_file, "file", "download.json", "Output file")
}

func main() {
    flag.Parse()

    seq = start
    fmt.Printf("Starting from post %v\n", seq)

    parse(seq)
}
