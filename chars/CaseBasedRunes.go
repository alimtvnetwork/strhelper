package chars

// CaseBasedRunes
//
// if case-sensitive then returns as is or else lower both and return
type CaseBasedRunes struct {
	ToRunes       []rune
	ComparingRune rune
}
