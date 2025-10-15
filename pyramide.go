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

func main() {
	Down()
	var t float64
	t = 2

	for j := 1; j <= 5; j++ {
		carre(t)
		Up()
		Right()
		Forward(1)
		Left()
		Down()
		t = t + 2
	}

	Say("pyramide finie")
}
