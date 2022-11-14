package main

import (
	"fmt"
	hangman "hangman/fonctions"
	"math/rand"
	"os"
	"io/ioutil"
	"strconv"
	"time"
	"encoding/json"
)

func main() {
	mode := "classique"

	ascii_letter := [][]string{}

	if hangman.Len(hangman.SliceSliceToSlice(ascii_letter)) == 0 {
		for i := 0; i < hangman.Len(hangman.Read(".\\assets\\standard.txt")); i += 9 {
			ascii_letter = append(ascii_letter, hangman.Read(".\\assets\\standard.txt")[i:i+9])
		}
	}

	hangmanStep := [][]string{}

	for i := 0; i < hangman.Len(hangman.Read("assets/hangman.txt")); i += 8 {
		hangmanStep = append(hangmanStep, hangman.Read("assets/hangman.txt")[i:i+8])
	}
	
	word := ""
	attempts := 10
	useLettre := []string{}
	message := ""

	rand.Seed(time.Now().UnixNano())

	dictRoad := ""

	switch hangman.Len(os.Args) {
	case 1:
		mode = "psMedium"
		dictRoad = "./assets/words.txt"
	case 2:
		dictRoad = os.Args[1]
		if _, err := os.Stat(os.Args[1]); err != nil {
			fmt.Println("Please enter a valide file.")
			return
		}
	case 3:
		if os.Args[1] == "--startWith" {
			if os.Args[2] == "save.txt" {
				saveData, err := ioutil.ReadFile(".\\assets\\save.txt")
				if err != nil {
					panic(err)
				}
				var saveDataSlice []string
				err = json.Unmarshal(saveData, &saveDataSlice)
				if err != nil{
					panic(err)
				}
				message = "-- Your save is start --"
				attempts, _ := strconv.Atoi(saveDataSlice[3])
				hangman.Round(saveDataSlice[0], saveDataSlice[2], saveDataSlice[1], 0, attempts, ascii_letter, hangmanStep, hangman.StringToSlice(saveDataSlice[4]), message)
			}
		}
		if os.Args[2] == "--letterFile" {
			fmt.Println("Please enter ascii letter file.")
			return
		}
		if os.Args[1] == "--hard" {
			dictRoad = os.Args[2]
			if _, err := os.Stat(os.Args[2]); err != nil {
				fmt.Println("Please enter a valide file.")
				return
			}
			mode = "hard"
		} else if os.Args[2] == "--hard" {
			dictRoad = os.Args[1]
			if _, err := os.Stat(os.Args[1]); err != nil {
				fmt.Println("Please enter a valide file.")
				return
			}
			mode = "hard"
		} else {
			fmt.Println("To many arguments")
			return
		}
	case 4:
		if os.Args[2] == "--letterFile"{
			mode = "ascii-art"
			dictRoad = os.Args[1]
			if _, err := os.Stat(os.Args[3]); err != nil {
				fmt.Println("Please enter a valide file.")
				return
			}
			letterFileRoad := os.Args[3]
			if _, err := os.Stat(os.Args[3]); err != nil {
				fmt.Println("Please enter a valide file.")
				return
			}
			ascii_letter = [][]string{}
			for i := 0; i < hangman.Len(hangman.Read(letterFileRoad)); i += 9 {
				ascii_letter = append(ascii_letter, hangman.Read(letterFileRoad)[i:i+9])
			}

		}else {
			fmt.Println("To many arguments")
			return
		}
	default:
		fmt.Println("To many arguments")
		return
	}

	wordToFind := hangman.ReplaceAccentMaj(hangman.Read(dictRoad)[rand.Intn(hangman.Len(hangman.Read(dictRoad)))])

	for i := 0; i < hangman.Len(hangman.StringToSlice(wordToFind)); i++ {
		word += "_"
	}

	switch mode {
	case "hard":
		word, useLettre = hangman.RandomReveal(hangman.Len(hangman.StringToSlice(wordToFind))/3-1, word, wordToFind)
	default:
		word, useLettre = hangman.RandomReveal(hangman.Len(hangman.StringToSlice(wordToFind))/2-1, word, wordToFind)
	}
	if _, err := os.Stat("assets/save.txt"); err == nil {
		var reponse string
		fmt.Println("We found a save.\nWant you to launch them ? [y/n]")
		fmt.Scanln(&reponse)

		switch reponse {
		case "y":
			saveData, err := ioutil.ReadFile(".\\assets\\save.txt")
			if err != nil {
				panic(err)
			}
			var saveDataSlice []string
			err = json.Unmarshal(saveData, &saveDataSlice)
			if err != nil {
				panic(err)
			}
			attempts, _ := strconv.Atoi(saveDataSlice[3])
			message = "-- Your save is start --"
			hangman.Round(saveDataSlice[0], saveDataSlice[2], saveDataSlice[1], 0, attempts, ascii_letter, hangmanStep, hangman.StringToSlice(saveDataSlice[4]), message)
			return
		default:
			hangman.Round(mode, wordToFind, word, 0, attempts, ascii_letter, hangmanStep, useLettre, message)
			return
		}

	} else {
		hangman.Round(mode, wordToFind, word, 0, attempts, ascii_letter, hangmanStep, useLettre, message)
	}
}
