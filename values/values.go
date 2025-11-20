package values

type Value int

const (
	Rock Value = iota
	Paper
	Scissors
)

// String liefert den Wert als String.
func (v Value) String() string {
	switch v {
	case Rock:
		return "Stein"
	case Paper:
		return "Papier"
	case Scissors:
		return "Schere"
	default:
		return "Gib was gescheites an"
	}
}

// Beats gibt an, ob der Wert v den Wert w schlägt.
func (v Value) Beats(w Value) bool {
	if v == Rock && w == Scissors {
		return true
	}
	if v == Paper && w == Rock {
		return true
	}
	if v == Scissors && w == Paper {
		return true
	}
	return false
}
