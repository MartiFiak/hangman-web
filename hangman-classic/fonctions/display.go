package hangman

import (
	"fmt"
	"os"
	"strconv"
)

func EndScreen(attempts int, wordToFind string) {
	/*
		Input : Number of remaining tries & The word to find
		Checks the end conditions (victory or defeat) and deletes the save if it exists
		Output : Displays the exit screen corresponding to victory or defeat
	*/
	Clear()
	if attempts <= 0 { // Case of defeat
		for _, ligne := range Read("assets/snap.txt") {
			fmt.Println(ligne)
		}
	} else { // Win case
		for _, ligne := range Read("assets/bim.txt") {
			fmt.Println(ligne)
		}
	}
	if _, err := os.Stat("assets/save.txt"); err == nil { // Checking if a backup exists
		err = os.Remove("assets/save.txt")
		if err != nil {
			fmt.Println("Erreur lors de la suppression de la sauvegarde.")
		}
	}
	fmt.Println("Le mot à trouver était : " + wordToFind)
}

func AsciiArtScreen(ascii_letter [][]string, word string, attempts int) {
	/*
		Input : Ascii letter splitting, word state
		Iterates over the word state and displays the corresponding ascii letter
		Output : Displays the state of the word in ascii art
	*/
	fmt.Println(strconv.Itoa(attempts) + " attemps remaining.")
	for i := 0; i < Len(ascii_letter[0]); i++ {
		for _, _letter := range word {
			fmt.Print(ascii_letter[int(_letter)-32][i])
		}
		fmt.Println()
	}
}

func DisplayHangman(hangmanStep [][]string, attempts int) {
	/*
		Input : The breakdown of the hangman steps, the number of tries remaining
		Check the status of the trials
		Output : Displays the hangman corresponding to the remaining number of tries
	*/
	if attempts != 10 {
		for i := 0; i < Len(hangmanStep[10-attempts]); i++ {
			fmt.Println(hangmanStep[9-attempts][i])
		}
	}
}

func StandardScreen(word string, attempts int, hangmanStep [][]string, useLettre []string) {
	/*
		Input : Current word status, remaining tries, hangman steps, letters already used
		Uses inputs to achieve basic interface display
		Output : Displays the basic interface
	*/
	fmt.Println(word + "                                            " + strconv.Itoa(attempts) + "/10\n")
	DisplayHangman(hangmanStep, attempts)
	fmt.Println("____________________Already Use____________________")
	fmt.Print("\n")
	for _, _alphabetLettre := range "abcdefghijklmnopqrstuvwxyz" {
		isAdd := false
		for _, _lettre := range useLettre {
			if _lettre == string(_alphabetLettre) {
				fmt.Print(string(_alphabetLettre))
				isAdd = true
				break
			}
		}
		if !isAdd {
			fmt.Print(" ")
		}
		fmt.Print(" ")
	}
	fmt.Print("\n")
	fmt.Println("___________________________________________________")
	fmt.Print("\n")
}
