package questions

import (
	"github.com/jcc620/tkdquiz/grades"
)

type Question struct {
	Query  string
	Answer string
	grade  grades.Grade
	qType  questionType
}

type questionType int64

const (
	angle questionType = iota
	weight
)
