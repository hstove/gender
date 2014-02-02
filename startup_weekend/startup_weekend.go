package main

import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
  c "github.com/hstove/gender/classifier"
  "strings"
)

func genderMap () map[string]map[string]int {
  result := map[string]map[string]int {}
  years := []string{
    "2010",
    "2011",
    "2012",
    "2013",
    "2014",
  }
  for _, year := range years {
    result[year] = map[string]int {
      "Female": 0,
      "Male": 0,
    }
  }
  return result
}

func main() {
  file, _ := os.Open("startup_weekend/sw.csv")
  defer file.Close()
  classifier := c.Classifier()

  reader := csv.NewReader(file)
  reader.TrailingComma = true
  i := 0
  genders := make(map[string]map[string]map[string]int)
  // fmt.Println(genders)
  genders["Organizer"] = genderMap()
  genders["Mentor"] = genderMap()
  genders["Speaker"] = genderMap()
  genders["Judge"] = genderMap()
  for {
    i++
    record, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      // fmt.Println(err)
    }
    role := record[0]
    // fmt.Println(record)
    year := record[2]
    if year == "Date" {
      continue
    }
    name := strings.Split(record[1], " ")[0]
    gender, _ := c.Classify(classifier, name)
    // fmt.Println(name, gender)
    if genders[role] == nil {
      fmt.Println("setup", year)
    } else {
      genders[role][year][gender]++
    }
  }
  // fmt.Println(genders)
  for role, years := range genders {
    fmt.Println()
    fmt.Println(role)
    for year, genderCounts := range years {
      // these years don't have enough data
      if year == "2010" {
        continue
      }
      if year == "2014" {
        continue
      }
      fmt.Println(year)
      total := float64(genderCounts["Male"] + genderCounts["Female"])
      fmt.Println("Total:", total)
      females := float64(genderCounts["Female"])
      diff := females / total
      fmt.Println("Female:", females, "(", diff * 100, "%)")
      fmt.Println("Male:", genderCounts["Male"])
    }
  }
}