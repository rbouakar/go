package main

import (
	"reflect"
	"testing"
)

func TestSomme(t *testing.T) {

	t.Run("collection de 5 nombres", func(t *testing.T) {

		nombres := []int{1, 2, 3, 4, 5}

		resultat := Somme(nombres)
		attendu := 15

		if resultat != attendu {
			t.Errorf("recu %d attendu %d donné %v", resultat, attendu, nombres)
		}

	})

}

func TestSommeTout(t *testing.T) {
	resultat := SommeTout([]int{1, 2}, []int{0, 9})
	attendu := []int{3, 9}

	if !reflect.DeepEqual(resultat, attendu) {
		t.Errorf("recu %v attendu %v", resultat, attendu)
	}
}

func TestSommeToutesQueues(t *testing.T) {

	verifierSommes := func(t testing.TB, resultat, attendu []int) {
		t.Helper()
		if !reflect.DeepEqual(resultat, attendu) {
			t.Errorf("recu %v attendu %v", resultat, attendu)
		}
	}

	t.Run("faire les sommes de quelques slices", func(t *testing.T) {
		resultat := SommeToutesQueues([]int{1, 2}, []int{0, 9})
		attendu := []int{2, 9}

		verifierSommes(t, resultat, attendu)
	})

	t.Run("sommer en sécurité des slices vides", func(t *testing.T) {
		resultat := SommeToutesQueues([]int{}, []int{3, 4, 5})
		attendu := []int{0, 9}
		
		verifierSommes(t, resultat, attendu)

	})
}
