package main

import "fmt"

//Cette fonction est une fonction de transition qui permet d'utiliser deux fonction "Victory" qui affiche un texte si la lettre est
//présente et d'afficher un texte si le mot est complet.
//Et la fonction "position" qui permet d'afficher les positions du pendu, il affiche aussi le nombre de vie restante à l'utilisateur.
func Response(number rune) {
	if VerifLetter(string(number)) == true {
		fmt.Println("Bravo, vous avez trouvé une lettre !")
		Victory()
	} else {
		fmt.Println("Dommage cette lettre n'est pas dans le mot.")
		vie--
		position()
	}
	if vie > 1 {
		fmt.Println("Il vous reste :", vie, "vies")
	} else {
		fmt.Println("Il vous reste :", vie, "vie")
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println(hidewords)
	Letter()
}
