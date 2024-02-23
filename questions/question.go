// Package questions implements functionality for generating a set of questions.
package questions

import (
	"github.com/jcc620/tkdquiz/grades"
)

// A Question is a quiz question.
type Question struct {
	Query  string // the actual question.
	Answer string // the answer to the question.
	grade  grades.Grade
	qType  questionType
}

type questionType int64

const (
	angle questionType = iota
	weight
)
