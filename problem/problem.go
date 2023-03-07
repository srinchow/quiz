package problem

import "fmt"

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

func (p Problem) Print(idx int) {
	fmt.Printf("Problem %v, %v = ", idx+1, p.Question)
}
