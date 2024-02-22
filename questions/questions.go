package questions

import "github.com/jcc620/tkdquiz/grades"

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
}
