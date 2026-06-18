package iteration

import "strings"

func Repeter(a string, nbr int) string {
	var repete strings.Builder
	for i := 0; i < nbr; i++ {
		repete.WriteString(a)
	}
	return repete.String()

}