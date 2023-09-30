package mock

type Pill int

//go:generate go run golang.org/x/tools/cmd/stringer -type=Pill

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
