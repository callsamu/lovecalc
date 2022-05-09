package core

// Simple PODs for containing input and results
// used for serialization

type Couple struct {
	FirstName  string
	SecondName string
}

type Match struct {
	Couple
	CoupleName  string
	Probability float64
}
