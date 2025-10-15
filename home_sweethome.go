package main
import (. "github.com/univ-tlse2/traceur")

func carre(taille float64) {
	var i int
	i = 1
	for i <= 4 {
		Forward(taille)
		Right()
		i = i + 1
	}
}

func triangle(taille float64) {
	var i int
	i = 1
	for i <= 3 {
		Forward(taille)
		Right()
		Right()
		Right()
		Right()
		Right()
		Right()
		i = i + 1
	}
}

func main() {
	Down()
	Color("brown")
	carre(6)
	Up()
	Forward(6)
	Down()
	Color("red")
	triangle(6)
	Say("maison terminee")
}
