package questions

import "github.com/jcc620/tkdquiz/grades"

var possibleAnswers map[questionType][]string = map[questionType][]string{
	angle: {"0°", "15°", "22.5°", "25°", "45°", "75°"},
}

var questionBank []Question = []Question{
	{
		Query:  "What is the angle of the front foot in Walking Stance?",
		Answer: "0°",
		grade:  grades.WhiteBelt,
		qType:  angle,
	},
	{
		Query:  "What is the angle of the back foot in Walking Stance?",
		Answer: "25°",
		grade:  grades.WhiteBelt,
		qType:  angle,
	},
	{
		Query:  "What is the angle of the feet in Parallel Stance?",
		Answer: "0°",
		grade:  grades.WhiteBelt,
		qType:  angle,
	},
	{
		Query:  "What is the angle of the feet in Attention Stance?",
		Answer: "22.5°",
		grade:  grades.WhiteBelt,
		qType:  angle,
	},
}
