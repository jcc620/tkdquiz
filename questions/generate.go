package questions

import (
	"math/rand"

	"github.com/jcc620/tkdquiz/grades"
)

// GenerateQuestions returns qCount questions appropriate for those up to and including grade.
func GenerateQuestions(grade grades.Grade, qCount int, seed int64) []Question {
	rand.Seed(seed)
	chosenQuestions := map[int]bool{}
	questions := []Question{}
	for len(questions) < qCount {
		nextQIdx := rand.Intn(len(questionBank))
		nextQ := questionBank[nextQIdx]
		if !chosenQuestions[nextQIdx] && nextQ.grade <= grade {
			questions = append(questions, nextQ)
		}
	}
	return questions
}
