package hangman

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nsf/termbox-go"
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

func ShowHangman(hangmanStep [][]string, attempts int) {
	/*
		Input : The breakdown of the hangman steps, the number of tries remaining
		Check the status of the trials
		Same that for DIsplayHangman but for Piste-Noire
		Output : Displays the hangman corresponding to the remaining number of tries
	*/
	if attempts < 0 {
		attempts = 0
	}
	if attempts != 10 {
		for i := 0; i < Len(hangmanStep[9-attempts])-1; i++ {
			WriteText(hangmanStep[9-attempts][i], 55, 15+i, termbox.ColorWhite)
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

func WelcomePage(modePos int, wordToFind string) {
	WriteText("Welcome", 20, 2, termbox.ColorRed)
	CreateBox(0, 5, 80, 21, termbox.ColorWhite)
	WriteText("[ Welcome ]", 5, 5, termbox.ColorWhite)

	SelectMode(modePos, wordToFind)

	WriteText("Game", 31, 2, termbox.ColorWhite)
	WriteText("Help", 38, 2, termbox.ColorWhite)

	WriteText("Welcome to termbox hangman version.", 3, 7, termbox.ColorWhite)
	WriteText("We hope that you are going to take some fun.", 3, 8, termbox.ColorWhite)
	WriteText("WARNING : Every key's input can't be release.", 3, 9, termbox.ColorWhite)
	WriteText("Press 'Enter' key's to valide your choose.", 3, 10, termbox.ColorWhite)

}

func GamePage(attempts, word, useLettre string) {

	WriteText("Game", 31, 2, termbox.ColorRed)
	CreateBox(0, 5, 20, 9, termbox.ColorWhite)
	WriteText("Welcome", 20, 2, termbox.ColorWhite)
	WriteText("Help", 38, 2, termbox.ColorWhite)

	CreateBox(22, 5, 42, 9, termbox.ColorRed)

	WriteText("Welcome", 2, 7, termbox.ColorWhite)
	WriteText(" Attempts ", 24, 5, termbox.ColorWhite)

	WriteText(attempts, 30, 7, termbox.ColorWhite)

	CreateBox(44, 5, 74, 27, termbox.ColorRed)
	WriteText(" Hangman ", 46, 5, termbox.ColorWhite)

	CreateBox(0, 11, 42, 15, termbox.ColorWhite)
	WriteText(" Words or Letters ", 3, 11, termbox.ColorWhite)

	CreateBox(0, 17, 42, 21, termbox.ColorBlue)
	WriteText(" Word ", 3, 17, termbox.ColorWhite)
	WriteWord(word, 3, 19, termbox.ColorWhite)

	CreateBox(0, 23, 42, 27, termbox.ColorBlue)
	WriteText(" Used Letters ", 3, 21, termbox.ColorWhite)
	WriteWord(useLettre, 3, 25, termbox.ColorWhite)

}

func HelpPage() {
	WriteText("Help", 38, 2, termbox.ColorRed)
	WriteText("Welcome", 20, 2, termbox.ColorWhite)
	WriteText("Game", 31, 2, termbox.ColorWhite)
	CreateBox(0, 5, 80, 21, termbox.ColorWhite)
	WriteText("[ Need Help ? ]", 5, 5, termbox.ColorWhite)

	WriteText("1. 'ESC' to quit.", 3, 7, termbox.ColorWhite)
	WriteText("2. 'ENTER' to confirm your choice.", 3, 8, termbox.ColorWhite)
	WriteText("3. 'BACKSPACE' or 'DEL' to delete the last letter.", 3, 9, termbox.ColorWhite)
	WriteText("4. '²' to restart.", 3, 10, termbox.ColorWhite)
}

func WinPage() {
	ClearPage()
	CreateBox(0, 5, 70, 25, termbox.ColorGreen)
	WriteText(" You Win ", 3, 5, termbox.ColorWhite)
	WriteText("Congratulation you WIN ! ", 2, 7, termbox.ColorWhite)
	WriteText("Press 'ESC' to quit. ", 2, 8, termbox.ColorWhite)
	WriteText("Press ' ² ' to restart. ", 2, 8, termbox.ColorWhite)
}

func LosePage(hangmanStep [][]string, attempts int, wordToFind string) {
	ClearPage()
	ShowHangman(hangmanStep, attempts)
	CreateBox(0, 5, 70, 25, termbox.ColorBlue)
	WriteText(" You Loose ", 3, 5, termbox.ColorWhite)
	WriteText("You lost, the word was "+wordToFind+". ", 2, 7, termbox.ColorWhite)
	WriteText("Press 'ESC' to quit. ", 2, 8, termbox.ColorWhite)
	WriteText("Press ' ² ' to restart. ", 2, 8, termbox.ColorWhite)
}

func ClearPage() {
	for i := 5; i < 50; i++ {
		for j := 0; j < 150; j++ {
			termbox.SetCell(j, i, ' ', termbox.RGBToAttribute(0, 0, 0), termbox.RGBToAttribute(0, 0, 0))
		}
	}
	termbox.Flush()
}

func SelectMode(modePos int, wordToFind string) (string, []string, string){
	psMode := ""
	md := []string{"Easy", "Medium", "Hard"}
	md1 := 0
	md2 := md[modePos]
	md3 := 0
	color := termbox.ColorWhite
	switch md[modePos] {
	case "Easy":
		md1 = 2
		md3 = 1
		color = termbox.ColorGreen
		psMode = "psEasy"
	case "Medium":
		md1 = 0
		md3 = 2
		color = termbox.ColorYellow
		psMode = "psMedium"
	case "Hard":
		md1 = 1
		md3 = 0
		color = termbox.ColorRed
		psMode = "psHard"
	}
	WriteText("Select difficulty :", 32, 14, termbox.ColorWhite)
	WriteText(md[md1]+"            ", 38, 16, termbox.ColorWhite)
	WriteText("<  "+md2+" >            ", 35, 17, color)
	WriteText(md[md3]+"            ", 38, 18, termbox.ColorWhite)
	word, useLettre := RandomRevealPS(modePos, wordToFind)
	return word, useLettre, psMode
}
