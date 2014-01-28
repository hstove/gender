package main

import (
  "encoding/csv"
  "io"
  "os"
  "fmt"
  "log"
  "github.com/eaigner/shield"
)

var logger *log.Logger

func newClassifier () shield.Shield {
  return shield.New(
    shield.NewEnglishTokenizer(),
    shield.NewRedisStore("127.0.0.1:6379", "", logger, ""),
  )
}

func genderMap () map[string]int {
  res := map[string]int {
    "F": 0,
    "M": 0,
  }
  return res
}

func main() {
  file, _ := os.Open("sw.csv")
  defer file.Close()
  classifier := newClassifier()

  reader := csv.NewReader(file)
  reader.TrailingComma = true
  i := 0
  genders := make(map[string]map[string]int)
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
      fmt.Println(err)
    }
    role := record[0]
    if role == "Role" {
      continue
    }
    name := record[1]
    gender, _ := classifier.Classify(name)
    // fmt.Println(i, "\t", role, "\t\t", name, gender)
    if genders[role] == nil {
      fmt.Println("setup", role)
    } else {
      genders[role][gender]++
    }
  }
  for role, genderCounts := range genders {
    fmt.Println()
    fmt.Println(role)
    total := float64(genderCounts["M"] + genderCounts["F"])
    fmt.Println("Total:", total)
    females := float64(genderCounts["F"])
    diff := females / total
    percentage := fmt.Sprintf(" (%f%)", diff)
    fmt.Println("Female:", females, percentage)
    fmt.Println("Male:", genderCounts["M"])
  }
}