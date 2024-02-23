// Package grades implements the type and constants to reference belt grades in ITF Taekwon-Do.
package grades

// A Grade is a belt grade in ITF Taekwon-Do.
type Grade int64

const (
	WhiteBelt  Grade = iota // 10th kup white belt.
	YellowTag               // 9th kup yellow tag.
	YellowBelt              // 8th kup yellow belt.
	GreenTag                // 7th kup green tag.
	GreenBelt               // 6th kup green belt.
	BlueTag                 // 5th kup blue tag.
	BlueBelt                // 4th kup blue belt.
	RedTag                  // 3rd kup red tag.
	RedBelt                 // 2nd kup red belt.
	BlackTag                // 1st kup black tag.
	BlackBelt1              // 1st dan black belt.
	BlackBelt2              // 2nd dan black belt.
	BlackBelt3              // 3rd dan black belt.
	BlackBelt4              // 4th dan black belt.
	BlackBelt5              // 5th dan black belt.
	BlackBelt6              // 6th dan black belt.
	BlackBelt7              // 7th dan black belt.
	BlackBelt8              // 8th dan black belt.
	BlackBelt9              // 9th dan black belt.

)
