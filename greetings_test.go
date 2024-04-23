package gotimegreetings

import (
	"testing"
)

func TestGreetGoodEveningEn(t *testing.T) {

	greeting := Greet("en")
	expected := "Good Evening"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}

func TestGreetGoodAfternoonEn(t *testing.T) {

	greeting := Greet("en")
	expected := "Good Afternoon"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}

func TestGreetGoodEveningPT_BR(t *testing.T) {

	greeting := Greet("pt_br")
	expected := "Boa noite"

	if greeting != expected {
		t.Errorf("Greet Test failed! Expected: %v, but got: %v", expected, greeting)
	}
}
