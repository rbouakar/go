package main

func Somme(nombres []int) int {

	var somme int

	for _, nombre := range nombres {
		somme += nombre
	}
	return somme
}

func SommeTout(nombres ...[]int) []int {
	var sommes []int
	for _, nombre := range nombres {
		sommes = append(sommes, Somme(nombre))
	}
	return sommes
}

func SommeToutesQueues(nombres ...[]int) []int {
	var sommes []int
	for _, nombre := range nombres {

		if len(nombre) == 0 {
			sommes = append(sommes, 0)
		} else {
			queue := nombre[1:]
			sommes = append(sommes, Somme(queue))
		}
	}
	return sommes
}