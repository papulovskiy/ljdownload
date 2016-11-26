package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
    "sync"
    "time"

    "github.com/PuerkitoBio/goquery"
)

var seq, start, stop int

var url_pattern string

var wg sync.WaitGroup

var client *http.Client

type Post struct {
    Id    int    `json:"id"`
    Url   string `json:"url"`
    Title string `json:"title"`
    Text  string `json:"text"`
}

func get_url(number int) string {
    return fmt.Sprintf(url_pattern, number)
}

func parse(number int, results chan string) {
    response, err := client.Get(get_url(number))
    if err != nil {
        return
    }
    doc, err := goquery.NewDocumentFromResponse(response)
    if err == nil {
        url, url_err := doc.Find("a.link__control").Html()
        title, title_err := doc.Find("div.article__title").Html()
        text, text_err := doc.Find("div.article__main-text").Html()
        if url_err != nil || title_err != nil || text_err != nil {
            return
        }
        d := Post{
            Id:    number,
            Url:   url,
            Title: title,
            Text:  text,
        }
        json, err := json.Marshal(d)
        if err == nil {
            results <- fmt.Sprintf("%s", json)
        }
    }

}

func writer(results <-chan string) {
    for message := range results {
        fmt.Println(message)
    }
}

func init() {
    flag.IntVar(&start, "start", 1, "Start post number")
    flag.IntVar(&stop, "stop", 350000000, "Stop post number")

    flag.StringVar(&url_pattern, "url", "https://ljsear.ch/savedcopy?post=%v", "URL pattern, use %v as a post number placeholder")
}

func main() {
    flag.Parse()

    seq = start

    client = &http.Client{
        Transport: &http.Transport{
            MaxIdleConnsPerHost: 20,
        },
        Timeout: time.Duration(10) * time.Second,
    }

    results := make(chan string, 10)
    go writer(results)

    for seq = start; seq <= stop; seq++ {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            parse(num, results)
        }(seq)
        if seq%10 == 0 {
            wg.Wait()
            // time.Sleep(25 * time.Millisecond)
        }
    }
    wg.Wait()
}
