package main

import (
  "fmt"
  "log"
  "github.com/eaigner/shield"
)

var logger *log.Logger

func main() {
  classifier := shield.New(
    shield.NewEnglishTokenizer(),
    shield.NewRedisStore("127.0.0.1:6379", "", logger, ""),
  )

  names := []string {
    "hank",
    "mark",
    "hannah",
    "rachael",
    "edward",
    "norah",
    "henry",
    "charlie",
    "ben",
    "claire",
    "matt",
    "lauren",
    "marsha",
  }

  for i, _ := range names {
    name := names[i]
    result, _ := classifier.Classify(name)
    fmt.Println(name, result)
  }
}