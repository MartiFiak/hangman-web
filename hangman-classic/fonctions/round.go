package hangman

import (
	"fmt"
)

func Round(mode, wordToFind, word string, vowelsCount, attempts int, ascii_letter, hangmanStep [][]string, useLettre []string, message string) bool {
	/*
	Input: Variable necessary for the game (Game mode, Word to find, the current state of the word, number of vowels already entered, the letters in ascci format, the steps of the hangman, the letters used, if the previous letter has already was used)
	Performs and calls the different actions of the game (This function is recursive)
	Output :
	*/
	if attempts <= 0 {
		EndScreen(attempts, wordToFind)
		return false
	} else if word == wordToFind {
		EndScreen(attempts, wordToFind)
		return false
	}else {
		Clear()
		fmt.Println(mode)
		fmt.Println(message)
		message = ""
		switch mode {
			case "psEasy":
				RoundPisteNoire(mode, wordToFind, word, vowelsCount, attempts, hangmanStep, useLettre)
				return false
			case "psMedium":
				RoundPisteNoire(mode, wordToFind, word, vowelsCount, attempts, hangmanStep, useLettre)
				return false
			case "psHard":
				RoundPisteNoire(mode, wordToFind, word, vowelsCount, attempts, hangmanStep, useLettre)
				return false
			case "ascii-art":
				AsciiArtScreen(ascii_letter, word, attempts)
				DisplayHangman(hangmanStep, attempts)
				fmt.Println()
			default:
				StandardScreen(word, attempts, hangmanStep, useLettre)
		}
	
		input := RequestLetterWord()
	
		word, vowelsCount, attempts, useLettre, message = InputProcessing(mode, word, wordToFind, input, attempts, vowelsCount, useLettre, message)
	
		if vowelsCount == -100 {
			Clear()
			return false
		}
	
		/*if mode == "hard" {
			vowelsCount = isVowel(input, vowelsCount)
			if vowelsCount >= 3 {
				attempts--
			}
		}*/
	
		fmt.Print("\n")
		
		return Round(mode, wordToFind, word, vowelsCount, attempts, ascii_letter, hangmanStep, useLettre, message)
	}
}