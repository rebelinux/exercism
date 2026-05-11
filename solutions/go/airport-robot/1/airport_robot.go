package airportrobot

import (
	"fmt"
)

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

type Italian struct {
	lang string
}

func (i Italian) LanguageName() string {
	i.lang = "Italian"
	return i.lang
}

func (i Italian) Greet(n string) string {
	return fmt.Sprintf("Ciao %s!", n)
}

type Portuguese struct {
	lang string
}

func (i Portuguese) LanguageName() string {
	i.lang = "Portuguese"
	return i.lang
}

func (i Portuguese) Greet(n string) string {
	return fmt.Sprintf("Olá %s!", n)
}

func SayHello(n string, h Greeter) string {

	return fmt.Sprintf("I can speak %s: %s", h.LanguageName(), h.Greet(n))
}
