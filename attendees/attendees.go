package main

import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
  c "github.com/hstove/gender/classifier"
)

func main() {
  file, _ := os.Open("attendees.csv")
  defer file.Close()
  classifier := c.Classifier()

  reader := csv.NewReader(file)
  reader.TrailingComma = true
  i := 0
  genders := map[string]int {
    "F": 0,
    "M": 0,
  }

  for {
    i++
    record, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println(err)
    }
    name := record[0]
    gender, _ := c.Classify(classifier, name)
    // fmt.Println("\t\t", name, gender)
    genders[gender]++
  }

  total := float64(genders["M"] + genders["F"])
  fmt.Println("Total:", total)
  females := float64(genders["F"])
  diff := females / total
  percentage := fmt.Sprintf(" (%f%)", diff)
  fmt.Println("Female:", females, percentage)
  fmt.Println("Male:", genders["M"])
}