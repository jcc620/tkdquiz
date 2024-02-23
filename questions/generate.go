package questions

import "github.com/jcc620/tkdquiz/grades"

// GenerateQuestions returns qCount questions appropriate for those up to and including grade.
func GenerateQuestions(grade grades.Grade, qCount int) []Question {
	return questionBank
}
