package main 

import "fmt"
const anglais = "Anglais"
const francais = "Français"
const espagnol = "Espagnol"

const prefixeSalutAnglais = "Hello, "
const prefixeSalutFrancais = "Bonjour, "
const prefixeSalutEspagnol = "Hola, "


func Hello(name string, language string) string{

	if name == "" {
		name = "world"
	}

	return prefixeSalut(language) + name

}

func prefixeSalut(language string) (prefixe string) {
	switch language {
	
	case anglais:
		prefixe = prefixeSalutAnglais

	case espagnol:
		prefixe = prefixeSalutEspagnol

	default:
		prefixe = prefixeSalutFrancais

	}
	return
}

func main() {

	fmt.Println(Hello("world", "Espagnol"))
}