package main

import (
	"fmt"
	"os"
	"time"
)

//Cette fonction permet de vérifier si le mot donner est complet, il affichera en cas de victore deux phrases.
func Victory() {
	var compt int
	var motcomplet string
	for i := 0; i < len(hidewords); i++ {
		motcomplet += hidewords[i]
	}
	var runemotcomplet = []rune(motcomplet)
	for i := 0; i < len(runemotcomplet); i++ {
		if runemotcomplet[i] <= 'z' && runemotcomplet[i] >= 'a' {
			compt += 1
		} else {
			compt += 0
		}
	}
	if compt == len(words) {
		fmt.Println("Pal mal, vous m'avez battu ^_^")
		time.Sleep(2 * time.Second)
		fmt.Println("On remet en jeu votre titre ?")
		os.Exit(0)
	}
}
