package problem

type Problem struct {
	Question string
	Answer   string
}

func New(record []string) Problem {
	return Problem{
		Question: record[0],
		Answer:   record[1],
	}
}
