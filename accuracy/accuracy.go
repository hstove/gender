package main

import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
  "strings"
  c "github.com/hstove/gender/classifier"
)

func main() {
  classifier := c.Classifier()
  year := 2000
  filename := fmt.Sprintf("names/yob%d.txt", year)
  file, _ := os.Open(filename)
  defer file.Close()

  reader := csv.NewReader(file)
  wrong := 0.0
  rows := 0.0
  for {
    record, err := reader.Read()
    if err == io.EOF {
        break
    }

    name := strings.ToLower(record[0])
    gender := c.Gender(record[1])
    guess, _ := c.Classify(classifier, name)
    if guess != string(gender) {
      // fmt.Println(name, gender, guess)
      wrong++
    }
    rows++
  }
  // Calculate the accuracy of the classifier
  // based a sample of the Census data.
  accuracy := 1. - wrong / rows
  fmt.Println("Rows:", rows)
  fmt.Println("Wrong:", wrong)
  fmt.Println("Accuracy:", accuracy)
}