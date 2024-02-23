package questions

import (
	"math/rand"

	"github.com/jcc620/tkdquiz/grades"
)

// GenerateQuestions returns n random questions appropriate for those up to and including grade.
// seed is used to seed the radnom number generator.
func GenerateQuestions(grade grades.Grade, n int, seed int64) []Question {
	rand.Seed(seed)
	chosenQuestions := map[int]bool{}
	questions := []Question{}
	for len(questions) < n {
		nextQIdx := rand.Intn(len(questionBank))
		nextQ := questionBank[nextQIdx]
		if !chosenQuestions[nextQIdx] && nextQ.grade <= grade {
			questions = append(questions, nextQ)
		}
	}
	return questions
}

// GenerateWrongAnswers returns a slice of incorrect but plausible answers for q.
func (q *Question) GenerateWrongAnswers(n int) []string {
	wrongAnswerBank := possibleAnswers[q.qType]
	chosenWrongAnswers := map[int]bool{}
	wrongAnswers := []string{}
	for len(wrongAnswers) < n {
		nextAIdx := rand.Intn(len(wrongAnswerBank))
		nextA := wrongAnswerBank[nextAIdx]
		if !chosenWrongAnswers[nextAIdx] && nextA != q.Answer {
			wrongAnswers = append(wrongAnswers, nextA)
		}
	}
	return wrongAnswers
}
