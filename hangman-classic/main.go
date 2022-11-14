package hangman

import (
	"math/rand"
	"time"

	hangman "hangmanweb/hangman-classic/fonctions"
)

func Main() {
	mode := "classique"

	hangmanStep := [][]string{}

	for i := 0; i < hangman.Len(hangman.Read("assets/hangman.txt")); i += 8 {
		hangmanStep = append(hangmanStep, hangman.Read("assets/hangman.txt")[i:i+8])
	}

	word := ""
	attempts := 10
	useLettre := []string{}
	message := "Start"

	rand.Seed(time.Now().UnixNano())

	dictRoad := "./assets/words.txt"

	wordToFind := hangman.ReplaceAccentMaj(hangman.Read(dictRoad)[rand.Intn(hangman.Len(hangman.Read(dictRoad)))])

	for i := 0; i < hangman.Len(hangman.StringToSlice(wordToFind)); i++ {
		word += "_"
	}

	word, useLettre = hangman.RandomReveal(hangman.Len(hangman.StringToSlice(wordToFind))/2-1, word, wordToFind)

	hangman.Round(mode, wordToFind, word, 0, attempts, hangmanStep, useLettre, message)
}
