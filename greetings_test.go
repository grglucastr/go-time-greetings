package gotimegreetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetGoodEveningEn(t *testing.T) {

	greeting := Greet("en", 22)
	expected := "Good Evening"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}

func TestGreetGoodAfternoonEn(t *testing.T) {

	greeting := Greet("en", 17)
	expected := "Good Afternoon"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}

func TestGreetGoodMorningEn(t *testing.T) {

	greeting := Greet("en", 6)
	expected := "Good Morning"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}

func TestGreetGoodEveningPT_BR(t *testing.T) {

	greeting := Greet("pt_br", 18)
	expected := "Boa noite"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}

func TestGreetGoodAfternoonPT_BR(t *testing.T) {

	greeting := Greet("pt_br", 13)
	expected := "Boa tarde"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}

func TestGreetGoodMorningPT_BR(t *testing.T) {

	greeting := Greet("pt_br", 3)
	expected := "Bom dia"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}

func TestGreetingsFileNotFound(t *testing.T) {

	assert.Panics(t, func() { Greet("es", 2) }, "not panic")

}

func TestGreetingsFileBroken(t *testing.T) {

	assert.Panics(t, func() { Greet("err", 2) }, "not panic")

}
