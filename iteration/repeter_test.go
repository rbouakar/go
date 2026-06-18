package iteration

import (
	"fmt"
	"testing"
)

func TestRepeter(t *testing.T) {
	repete := Repeter("a", 3)
	attendu := "aaa"

	if repete != attendu {
		t.Errorf("attendu %q mais recu %q", attendu, repete)
	}
}

func BenchmarkRepeter(b *testing.B) {
	for b.Loop() {
		Repeter("a", 3)
	}
}

func ExampleRepeter() {
	repete := Repeter("a", 3)
	fmt.Println(repete)
	//OUTPUT: aaa
}