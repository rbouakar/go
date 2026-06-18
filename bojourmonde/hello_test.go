package main 

import "testing"

func TestHello(t *testing.T) {

	verifierMessageCorrect := func(t testing.TB, result, waited string) {
		t.Helper()
		if result != waited {
			t.Errorf("result %q waited %q", result, waited)
		}

	}
	t.Run("dire bonjour aux personnes", func(t *testing.T) {
		result := Hello("Chris", "")
		waited := "Bonjour, Chris"
		verifierMessageCorrect(t, result, waited)
	})

	t.Run("dire 'Hello, world' quand une chaine vide est fournie", func(t *testing.T) {
		result := Hello("", "")
		waited := "Bonjour, world"
		verifierMessageCorrect(t, result, waited)
	})

	t.Run("en espagnol", func(t *testing.T) {
		result := Hello("Elodie", "Espagnol")
		waited := "Hola, Elodie"
		verifierMessageCorrect(t, result, waited)
	})
	 
	t.Run("en anglais", func(t *testing.T) {
		result := Hello("", "Anglais")
		waited := "Hello, world"
		verifierMessageCorrect(t, result, waited)
	})

}
