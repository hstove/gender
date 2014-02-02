package main

import (
  "fmt"
  "log"
  c "github.com/hstove/gender/classifier"
)

var logger *log.Logger

func main() {
  classifier := c.Classifier()

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
    gender, probability := c.Classify(classifier, name)
    fmt.Println(name, gender, probability)
  }
}