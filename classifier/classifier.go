package classifier

import (
	"bytes"
	"strings"

	_ "embed"

	"github.com/jbrukh/bayesian"
)

const (
	Boy  bayesian.Class = "Male"
	Girl bayesian.Class = "Female"
	None bayesian.Class = "N/A"
)

func NewClassifier() *bayesian.Classifier {
	return bayesian.NewClassifier(Boy, Girl, None)
}

//go:embed classifier.serialized
var classifier []byte

func Classifier() *bayesian.Classifier {
	classifier, _ := bayesian.NewClassifierFromReader(bytes.NewReader(classifier))
	return classifier
}

func Gender(gender string) bayesian.Class {
	if gender == "M" {
		return Boy
	}
	return Girl
}

func Learn(classifier *bayesian.Classifier, name string, gender string) {
	strings.ToLower(name)
	classifier.Learn([]string{name}, Gender(gender))
}

func Classify(classifier *bayesian.Classifier, name string) (string, float64) {
	name = strings.ToLower(name)
	scores, likely, _ := classifier.ProbScores([]string{name})
	gender := GenderFromInt(likely)
	probability := scores[likely]
	if probability == 0.5936993620304059 {
		gender = None
		probability = 0
	}
	return string(gender), probability
}

func GenderFromInt(gender int) bayesian.Class {
	if gender == 0 {
		return Boy
	}
	return Girl
}
