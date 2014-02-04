package main

import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
  c "github.com/hstove/gender/classifier"
  "strings"
)

func main() {
  file, _ := os.Open("attendees/attendees.csv")
  defer file.Close()
  classifier := c.Classifier()

  reader := csv.NewReader(file)
  reader.TrailingComma = true
  i := 0
  genders := map[string]map[string]int {}
  years := []string{
    "2009",
    "2010",
    "2011",
    "2012",
    "2013",
    "2014",
  }
  for _, year := range years {
    genders[year] = map[string]int {
      "Female": 0,
      "Male": 0,
      "N/A": 0,
    }
  }

  for {
    i++
    record, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println("ERROR: ", err)
    }
    if len(record) == 0 {
      continue
    }
    name := strings.Split(record[0], " ")[0]
    year := strings.Split(record[1], "-")[0]
    gender, _ := c.Classify(classifier, name)
    genders[year][gender]++
  }
  for year, genderCounts := range genders {
    fmt.Println(year)
    total := float64(genderCounts["Male"] + genderCounts["Female"])
    fmt.Println("Total:", total)
    fmt.Println("N/A:", genderCounts["N/A"])
    females := float64(genderCounts["Female"])
    diff := females / total
    fmt.Println("Female:", females, "(", diff * 100, "%)")
    fmt.Println("Male:", genderCounts["Male"])
    fmt.Println()
  }
}