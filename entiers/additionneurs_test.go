package entiers

import (
	"fmt"
	"testing"
)

func TestAdditonneurs( t*testing.T) {
	somme := Additionner(2,2)
	attendu := 4

	if somme != attendu {
		t.Errorf("attendu '%d' mais recu '%d'", attendu, somme)
	}
}

func ExampleAdditionner() {
	somme := Additionner(1, 5)
	fmt.Println(somme)
	//OUTPUT: 6
}