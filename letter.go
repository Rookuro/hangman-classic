package main

import "fmt"

//Cette fonction permet de prendre seulement une lettre ou un mot, il vérifit aussi les caractères entrées par l'utilisateur.
//Transforme les majuscules en minuscules.
func Letter() {
	a := 0
	for a <= 0 {
		var number []rune
		fmt.Println("Choisissez une lettre :")
		fmt.Scanf("%s\n", &interaction)
		for i := 0; i < len(interaction); i++ {
			number = append(number, rune(interaction[i]))
		}
		switch interaction {
		case interaction:
			if len(number) == 1 {
				if number[0] >= 97 && number[0] <= 122 {
					a++
					Word(string(number[0]))
					Response(number[0])
				} else if number[0] >= 65 && number[0] <= 90 {
					a++
					number[0] += 32
					Word(string(number[0]))
					Response(number[0])
				} else {
					fmt.Println("Seul les lettres sont acceptées.")
				}
			} else if len(interaction) > 1 {
				for i := 0; i < len(number); i++ {
					if number[i] >= 97 && number[i] <= 122 {
						a++
					} else if number[i] >= 65 && number[i] <= 90 {
						a++
						number[i] += 32
					}
				}
				verifword(string(number))
			}
		}
	}
}

//Cette fonction permet de vérifier la lettre dans le mot, si elle est présente il retourne "true" au sinon il retourne "faux".
func VerifLetter(number string) bool {
	var index int
	number = string(number)
	for i := 0; i < len(words); i++ {
		if tab[i] == number {
			hidewords[i] = number
			index++
		}
	}
	if index >= 1 {
		return true
	} else {
		return false
	}
}

func PrintLetter() {
	fmt.Print("\n")
	for i := 0; i < len(hidewords); i++ {
		fmt.Print(hidewords[i], " ")
	}
	fmt.Print("\n", "\n")
}
