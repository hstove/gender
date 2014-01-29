package classifier

import (
  "github.com/jbrukh/bayesian"
  "strings"
)

const (
  Boy bayesian.Class = "Male"
  Girl bayesian.Class = "Female"
)

func NewClassifier() *bayesian.Classifier {
  return bayesian.NewClassifier(Boy, Girl)
}

func Classifier() *bayesian.Classifier {
  classifier, _ := bayesian.NewClassifierFromFile("classifier.serialized")
  return classifier
}

func Gender(gender string) bayesian.Class {
  if gender == "M" {
    return Boy
  }
  return Girl
}

func Learn(classifier *bayesian.Classifier, name string, gender string) {
  classifier.Learn([]string{name}, Gender(gender))
}

func Classify (classifier *bayesian.Classifier, name string) (string, float64) {
  name = strings.ToLower(name)
  scores, likely, _ := classifier.ProbScores([]string{name})
  gender := GenderFromInt(likely)
  return string(gender), scores[likely]
}

func GenderFromInt(gender int) bayesian.Class {
  if gender == 0 {
    return Boy
  }
  return Girl
}