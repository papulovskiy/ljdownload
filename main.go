package main

import (
    "flag"
    "fmt"
    // "github.com/PuerkitoBio/goquery"
)

var seq int
var output string

func init() {
    flag.IntVar(&seq, "start", 1, "Start post number")

}
func main() {
    flag.Parse()
    fmt.Printf("Starting from post %v\n", seq)
}
