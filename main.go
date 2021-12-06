package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var words string
var hidewords []string
var tab []string
var interaction string
var UseWords []string
var vie int = 10

func main() {
	mot()
	Hide(words)
	random(tab)
	Letter()
}

func mot() {
	fmt.Println("Bonjour, mon nom est GlaDoS, je suis une intelligence artificielle.")
	time.Sleep(2 * time.Second)
	fmt.Println("Aujourd'hui nous allons tester vos compétences, vous êtes prêt ? C'est partie !")
	file, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	answers := string(file)
	table := strings.Split(answers, "\n")
	rand.Seed(time.Now().UnixNano())
	words = (table[rand.Intn(len(table))])
	for i := 0; i < len(words); i++ {
		tab = append(tab, string(words[i]))
	}
	// fmt.Println(words)
}

func Hide(words string) {
	for i := 0; i < len(words); i++ {
		hidewords = append(hidewords, "_ ")
	}
}

func random(tab []string) {
	n := len(words)/2 - 1
	for i := 0; i < n; i++ {
		alea := rand.Intn(len(words))
		hidewords[alea] = tab[alea]
	}
	fmt.Println(hidewords)
}

func Letter() {
	a := 0
	for a <= 0 {
		var number rune
		fmt.Println("Choisissez une lettre :")
		fmt.Scanf("%s\n", &interaction)
		for i := 0; i < len(interaction); i++ {
			number = rune(interaction[i])
		}
		switch interaction {
		case interaction:
			if len(interaction) > 1 {
				fmt.Println("Choisissez seulement UNE lettre.")
			} else if number >= 97 && number <= 122 {
				a++
				Word(string(number))
				Response(number)
			} else if number >= 65 && number <= 90 {
				a++
				number += 32
				Word(string(number))
				Response(number)
			} else {
				fmt.Println("Seul les lettres sont acceptées.")
			}
		}
	}
}

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

func Word(number string) {
	var a int
	for i := 0; i < len(UseWords); i++ {
		if number == UseWords[i] {
			a++
		}
	}
	if a == 0 {
		UseWords = append(UseWords, number)
	} else {
		fmt.Println("Vous avez déjà écris cette lettre.")
		fmt.Println("Lettres déjà testées :", UseWords)
		fmt.Println("Il vous reste :", vie, "vies")
		fmt.Println("-----------------------------------------------")
		Letter()
	}
	fmt.Println("Lettres déjà testées :", UseWords)
}

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
