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

type Italian struct{}
func (Italian) LanguageName() string {return "Italian"}
func (Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}
func (Portuguese) LanguageName() string {return "Portuguese"}
func (Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

type German struct{}
func (German) LanguageName() string {return "German"}
func (German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}
