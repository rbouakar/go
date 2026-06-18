package main 

import "fmt"
const anglais = "Anglais"
const francais = "Français"
const espagnol = "Espagnol"

const prefixeSalutAnglais = "Hello, "
const prefixeSalutFrancais = "Bonjour, "
const prefixeSalutEspagnol = "Hola, "


func Hello(name string) string{
	return prefixeSalutAnglais + name
}

func main() {

	fmt.Print(Hello("world"))
}