package hangmanweb

import (
	hc "hangmanweb/hangman-classic/fonctions"
	"math/rand"
	"strconv"
	"time"
)

var WordsFile = "./hangman-classic/assets/words.txt"

func InitGame(mode string) []string {

	switch mode {
	case "easy":
		WordsFile = "./hangman-classic/assets/words.txt"
	case "medium":
		WordsFile = "./hangman-classic/assets/words2.txt"
	case "hard":
		WordsFile = "./hangman-classic/assets/words3.txt"
	}

	rand.Seed(time.Now().UnixNano())
	wordToFind := hc.ReplaceAccentMaj(hc.Read(WordsFile)[rand.Intn(hc.Len(hc.Read(WordsFile)))])

	word := ""
	useLettre := []string{}

	for i := 0; i < hc.Len(hc.StringToSlice(wordToFind)); i++ {
		word += "_"
	}

	word, useLettre = hc.RandomReveal(hc.Len(hc.StringToSlice(wordToFind))/2-1, word, wordToFind)

	return []string{wordToFind, word, hc.SliceToString(useLettre)}
}

func InputTreatment(word, wordToFind, input, useLettre string, vowelsCount, attempts int) []string {
	if hc.Len(hc.StringToSlice(input)) != 0 {
		useLettreSlice := []string{}
		word, vowelsCount, attempts, useLettreSlice, _ = hc.InputProcessing("easy", word, wordToFind, input, attempts, 0, hc.StringToSlice(useLettre), "")
		if attempts <= 0 {
			if attempts < 0 {
				attempts = 0
			}
			return []string{"LoosePage", word, strconv.Itoa(vowelsCount), strconv.Itoa(attempts), hc.SliceToString(useLettreSlice)}
		}
		if input == wordToFind || word == wordToFind {
			return []string{"WinPage", word, strconv.Itoa(vowelsCount), strconv.Itoa(attempts), hc.SliceToString(useLettreSlice)}
		}
		return []string{"Okey", word, strconv.Itoa(vowelsCount), strconv.Itoa(attempts), hc.SliceToString(useLettreSlice)}
	} else {
		return []string{"Nope", word, strconv.Itoa(vowelsCount), strconv.Itoa(attempts), useLettre}
	}
}

func InputUsernameTreatment(input string) bool {
	return hc.Len(hc.StringToSlice(input)) != 0
}
