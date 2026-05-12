package school

import (
	"sort"
)

// Define the Grade and School types here.

type School struct {
	Grader []Grade
}

type Grade struct {
	grade_num int
	student   []string
}

func New() *School {
	s := new(School)

	return s
}

func (s *School) Add(student string, grade int) {

	s.Grader = append(s.Grader, Grade{grade, []string{student}})
}

func (s *School) Grade(level int) []string {

	var n []string

	for _, value := range s.Grader {
		if value.grade_num == level {
			n = append(n, value.student...)
		}
	}

	return n
}

func (s *School) Enrollment() []Grade {

	var unique []int

	var sorted []Grade

	for _, value := range s.Grader {
		skip := false
		for _, u := range unique {
			if value.grade_num == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, value.grade_num)
		}
	}

	sort.Ints(unique)

	for _, d := range unique {

		slc := s.Grade(d)

		sort.Strings(slc)

		sorted = append(sorted, Grade{d, slc})
	}

	return sorted
}
