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
)

var logger *log.Logger

func main() {
  classifier := shield.New(
    shield.NewEnglishTokenizer(),
    shield.NewRedisStore("127.0.0.1:6379", "", logger, ""),
  )
  classifier.Reset()

  for i := 1950; i <= 2012; i++ {
    filename := fmt.Sprintf("names/yob%d.txt", i)
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
        classifier.Learn(record[1], name)
        idx++
      }
    }
    fmt.Println("finished parsing", filename)
  }
}