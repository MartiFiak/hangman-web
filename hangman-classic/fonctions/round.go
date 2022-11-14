package hangman

import (
	"fmt"
)

func Round(mode, wordToFind, word string, vowelsCount, attempts int, hangmanStep [][]string, useLettre []string, message string) bool {
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

		StandardScreen(word, attempts, hangmanStep, useLettre)
	
		input := RequestLetterWord()
	
		word, vowelsCount, attempts, useLettre, message = InputProcessing(mode, word, wordToFind, input, attempts, vowelsCount, useLettre, message)
	
		if vowelsCount == -100 {
			Clear()
			return false
		}
	
		fmt.Print("\n")
		
		return Round(mode, wordToFind, word, vowelsCount, attempts, hangmanStep, useLettre, message)
	}
}