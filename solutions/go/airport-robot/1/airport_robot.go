package airportrobot

import (
	"fmt"
)

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Portuguese struct{}
type Italian struct{}
type German struct{}

func (ig Italian) LanguageName() string {
	return "Italian"
}
func (ig Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (pg Portuguese) LanguageName() string {
	return "Portuguese"
}
func (pg Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func (gg German) LanguageName() string {
	return "German"
}

func (gg German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}
