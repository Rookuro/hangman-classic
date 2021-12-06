package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var vie int = 10

func position() []string {
	readFile, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, "\n"+fileScanner.Text())
	}
	readFile.Close()
	for {
		if vie == 9 {
			fmt.Println(lines[0:7])
			return lines
		} else if vie == 8 {
			fmt.Println(lines[8:15])
			fmt.Println("GlaDoS : Le mot est pourtant simple, réfléchissez.")
			return lines
		} else if vie == 7 {
			fmt.Println(lines[16:23])
			return lines
		} else if vie == 6 {
			fmt.Println(lines[24:31])
			fmt.Println("GlaDoS : Y aurait-il une défaillance dans vôtre programme ?.")
			return lines
		} else if vie == 5 {
			fmt.Println(lines[32:39])
			return lines
		} else if vie == 4 {
			fmt.Println(lines[40:47])
			fmt.Println("GlaDoS : Je crains que vous vous soyez en position de faiblesse.")
			return lines
		} else if vie == 3 {
			fmt.Println(lines[48:55])
			return lines
		} else if vie == 2 {
			fmt.Println(lines[56:63])
			fmt.Println("GlaDoS : Je vais ajouter le mot échec à vôtre base de données.")
			return lines
		} else if vie == 1 {
			fmt.Println(lines[64:71])
			return lines
		} else if vie == 0 {
			fmt.Println(lines[72:79])
			fmt.Println("GlaDoS : On dirai que mon intelligence artificielle est supérieur à la vôtre.")
			os.Exit(0)
		}
	}
}
