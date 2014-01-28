package main

import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
  "log"
  "github.com/eaigner/shield"
  "strconv"
  "strings"
  "runtime"
)

var logger *log.Logger
var start = 1960
var end = 2004

func newClassifier () shield.Shield {
  return shield.New(
    shield.NewEnglishTokenizer(),
    shield.NewRedisStore("127.0.0.1:6379", "", logger, ""),
  )
}

func worker(year int, done chan bool) {
  classifier := newClassifier()
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
    count = count / 10 + 1
    idx := 0
    for idx <= int(count) {
      err := classifier.Learn(record[1], name)
      if err != nil {
        fmt.Println("Error: ", err)
      }
      idx++
    }
  }
  fmt.Println("finished parsing", filename)
  done <- true
}

func main() {
  nCPU := runtime.NumCPU()
  runtime.GOMAXPROCS(nCPU)
  fmt.Println("Number of CPUs: ", nCPU)

  classifier := newClassifier()
  classifier.Reset()

  done := make(chan bool, end - start)
  for i := start; i <= end; i++ {
    go worker(i, done)
  }
  for j := start; j <= end; j++ {
    <-done
  }
}