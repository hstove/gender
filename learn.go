package main

import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
  "strconv"
  "strings"
  "runtime"
  c "github.com/hstove/gender/classifier"
  b "github.com/jbrukh/bayesian"
)

var start = 1950
var end = 2012

func worker(classifier *b.Classifier, year int, done chan bool) {
  filename := fmt.Sprintf("names/yob%d.txt", year)
  file, _ := os.Open(filename)
  defer file.Close()

  reader := csv.NewReader(file)
  for {
    record, err := reader.Read()
    if err == io.EOF {
        break
    }

    count, _ := strconv.ParseInt(record[2], 10, 8)
    name := strings.ToLower(record[0])
    idx := 0
    for idx <= int(count) {
      c.Learn(classifier, name, record[1])
      idx++
    }
  }
  fmt.Println("finished parsing", filename)
  done <- true
}

func main() {
  runtime.GOMAXPROCS(1)

  classifier := c.NewClassifier()

  done := make(chan bool, end - start)
  for i := start; i <= end; i++ {
    go worker(classifier, i, done)
  }
  for j := start; j <= end; j++ {
    <-done
  }
  classifier.WriteToFile("./classifier.serialized")
}