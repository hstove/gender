package main

import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
  c "github.com/hstove/gender/classifier"
)

func genderMap () map[string]int {
  res := map[string]int {
    "Female": 0,
    "Male": 0,
  }
  return res
}

func main() {
  file, _ := os.Open("sw.csv")
  defer file.Close()
  classifier := c.Classifier()

  reader := csv.NewReader(file)
  reader.TrailingComma = true
  i := 0
  genders := make(map[string]map[string]int)
  // genders["Organizer"] = genderMap()
  // genders["Mentor"] = genderMap()
  // genders["Speaker"] = genderMap()
  // genders["Judge"] = genderMap()
  genders["2012"] = genderMap()
  genders["2011"] = genderMap()
  genders["2013"] = genderMap()
  genders["2010"] = genderMap()
  genders["2014"] = genderMap()
  for {
    i++
    record, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println(err)
    }
    // role := record[0]
    year := record[2]
    if year == "Date" {
      continue
    }
    name := record[1]
    gender, _ := c.Classify(classifier, name)
    // fmt.Println(i, "\t", role, "\t\t", name, gender)
    if genders[year] == nil {
      fmt.Println("setup", year)
    } else {
      genders[year][gender]++
    }
  }
  for year, genderCounts := range genders {
    fmt.Println()
    fmt.Println(year)
    total := float64(genderCounts["M"] + genderCounts["F"])
    fmt.Println("Total:", total)
    females := float64(genderCounts["F"])
    diff := females / total
    percentage := fmt.Sprintf(" (%f%)", diff)
    fmt.Println("Female:", females, percentage)
    fmt.Println("Male:", genderCounts["M"])
  }
}