package main 

import "testing"

func TestHello(t *testing.T) {

	verifierMessageCorrect := func(t testing.TB, result, waited string) {
		t.Helper()
		if result != waited {
			t.Errorf("result ùq waited %q", result, waited)
		}

	}
	t.Run("dire bonjour aux personnes", func(t *testing.T) {
		result := Hello("Chris")
		waited := "Hello, Chris"
		verifierMessageCorrect(t, result, waited)
	})

	t.Run("dire 'Hello, world' quand une chaine vide est fournie", func(t *testing.T) {
		result := Hello("")
		waited := "Hello, world"
		verifierMessageCorrect(t, result, waited)
	})
	 
}
