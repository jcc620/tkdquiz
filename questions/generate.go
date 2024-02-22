package questions

import "github.com/jcc620/tkdquiz/grades"

func GenerateQuestions(_ grades.Grade, _, _ int) []Question {
	return questionBank
}
