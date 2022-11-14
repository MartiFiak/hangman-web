package hangman

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func RequestLetterWord() string {
	/*
	Input :
	Request a String from the user
	Output: The String entered by the user
	*/
	var input string
	fmt.Print("Enter letter or word : ")
	fmt.Scanln(&input)
	input = ReplaceAccentMaj(input)
	return input
}

func InputProcessing(mode, word, wordToFind, input string, attempts, vowelsCount int, useLettre []string, message string) (string, int, int, []string, string) {
	/*
	Input: Important variable in the course of the game
	Get user input, convert it the right way (no accents or capitals) and process it
	Output: Returns modified values ​​after handling user input
	*/
	if Len(StringToSlice(input)) > 1 {
		if input == "stop" {
			if _, err := os.Stat("assets/save.txt"); err == nil && (mode != "psEasy" && mode != "psMedium" && mode != "psHard"){
				fmt.Println("WARNING : Save was found !\nCreate a new save delete the oldone.\nContinue ? [y/n]")
				reponse := RequestLetterWord()
				if reponse == "y" {
					err = os.Remove("assets/save.txt")
					if err != nil {
						fmt.Println("Erreur lors de la suppression de la sauvegarde.")
					}
				} else {
					return word, -100, attempts, useLettre, message
				}
			}
			file, err := os.OpenFile("assets/save.txt", os.O_CREATE|os.O_WRONLY, 0600)
			defer file.Close()
			if err != nil {
				panic(err)
			}
			saveData, err := json.Marshal([]string{mode, word, wordToFind, strconv.Itoa(attempts), SliceToString(useLettre)})
			if err != nil {
				panic(err)
			}
			_, err = file.Write(saveData)
			if err != nil {
				panic(err)
			}
			return word, -100, attempts, useLettre, message
		} else if input == wordToFind{
			if  mode != "psEasy" && mode != "psMedium" && mode != "psHard" {
				EndScreen(attempts, wordToFind)
			}
			return word, -100, attempts, useLettre, message
		} else {
			return word, vowelsCount, attempts - 2, useLettre, message
		}
	} else if len(input) == 1 {
		if mode == "hard" || mode == "psHard" {
			vowelsCount = isVowel(input, vowelsCount)
			if  vowelsCount > 3 {
				attempts--
			}
		}
		for _, le := range useLettre {
			if le == input {
				if mode == "hard" || mode == "psHard" || mode == "psMedium" {
					return word, vowelsCount, attempts - 1, useLettre, input + " was already used."
				}
				return word, vowelsCount, attempts, useLettre, input + " was already used."
			}
		}
		useLettre = append(useLettre, string(input))

		if Find(wordToFind, input) {

			var _word string

			for _index, _lettre := range wordToFind {
				if input == string(_lettre) {
					_word += input
				} else {
					_word += string(word[_index])
				}
			}

			word = _word

			return word, vowelsCount, attempts, useLettre, message
		} else {
			return word, vowelsCount, attempts - 1, useLettre, input + " is not in the word."
		}
	} else {
		return word, vowelsCount, attempts, useLettre, message
	}
}

func isVowel(input string, vowelsCount int) int {
	/*
	Input: A letter and the number of vowels
	Increase the number of vowels by 1 if the letter is a vowel
	Output: The new number of vowels
	*/
	switch input {
	case "a":
		vowelsCount++
	case "e":
		vowelsCount++
	case "i":
		vowelsCount++
	case "o":
		vowelsCount++
	case "u":
		vowelsCount++
	case "y":
		vowelsCount++
	}
	return vowelsCount
}

func Find(mot string, lettre string) bool {
	/*	
	Input: One word and one letter
	Scan the word to see if the letter belongs to the word
	Output: Returns if the letter is in the word
	*/
	for _, _lettre := range mot {
		if string(_lettre) == lettre {
			return true
		}
	}
	return false
}

func ReplaceAccentMaj(entry string) string {
	/*
	Input: One word
	Goes through the word and removes accents and capitals from words
	Output: The lowercase word without accents
	*/
	entrySlice := StringToSlice(entry)
	var toReturn string
	for _, _lettre := range entrySlice {
		if _lettre == "é" || _lettre == "è" || _lettre == "ê" || _lettre == "ë" || _lettre == "É" || _lettre == "È" || _lettre == "Ê" || _lettre == "Ë" || _lettre == "E" {
			toReturn += "e"
		} else if _lettre == "à" || _lettre == "â" || _lettre == "À" || _lettre == "Â" || _lettre == "A" {
			toReturn += "a"
		} else if _lettre == "î" || _lettre == "ï" || _lettre == "Î" || _lettre == "Ï" || _lettre == "I" {
			toReturn += "i"
		} else if _lettre == "ô" || _lettre == "Ô" || _lettre == "O" {
			toReturn += "o"
		} else if _lettre == "û" || _lettre == "ü" || _lettre == "ù" || _lettre == "Û" || _lettre == "Ü" || _lettre == "Ù" || _lettre == "U" {
			toReturn += "u"
		} else if _lettre == "ÿ" || _lettre == "Ÿ" || _lettre == "Y" {
			toReturn += "y"
		} else {
			if _lettre[0] <= 'Z' && _lettre[0] >= 'A' {
				toReturn += string(_lettre[0] + 32)
			} else {
				toReturn += _lettre
			}
		}
	}
	return toReturn
}
