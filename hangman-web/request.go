package hangmanweb

import (
	"encoding/csv"
	"fmt"
	hc "hangmanweb/hangman-classic/fonctions"
	"math/rand"
	"os"
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
		word, vowelsCount, attempts, useLettreSlice, _ = hc.InputProcessing("easy", word, wordToFind, hc.ReplaceAccentMaj(input), attempts, 0, hc.StringToSlice(useLettre), "")
		if attempts <= 0 {
			if attempts < 0 {
				attempts = 0
			}
			return []string{"LoosePage", word, strconv.Itoa(vowelsCount), strconv.Itoa(attempts), hc.SliceToString(useLettreSlice)}
		}
		if hc.ReplaceAccentMaj(input) == wordToFind || word == wordToFind {
			return []string{"WinPage", word, strconv.Itoa(vowelsCount), strconv.Itoa(attempts), hc.SliceToString(useLettreSlice)}
		}
		return []string{"Okey", word, strconv.Itoa(vowelsCount), strconv.Itoa(attempts), hc.SliceToString(useLettreSlice)}
	} else {
		return []string{"Nope", word, strconv.Itoa(vowelsCount), strconv.Itoa(attempts), useLettre}
	}
}

func RegisterUser(input, password, confirmpassword string)bool{
	if hc.Len(hc.StringToSlice(input)) != 0 && hc.Len(hc.StringToSlice(password)) != 0 && password == confirmpassword {
		if !UserExist(input) {
			usersDatabase, err := os.OpenFile("./server/database/users.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
			if err != nil {
				fmt.Println(err)
			}
			defer usersDatabase.Close()

			csvWriterGlobalDB := csv.NewWriter(usersDatabase)

			newData := []string{input, password, "0", "0", "0", "1", "0"} // Username, Password, Win, Lose, GamePlay, level, xp
			err = csvWriterGlobalDB.Write(newData)
			if err != nil {
				fmt.Println(err)
			}
			defer csvWriterGlobalDB.Flush()
			return true
		}
	}
	return false
}

func InputUsernameTreatment(input, password string) string {
	if hc.Len(hc.StringToSlice(input)) != 0 && hc.Len(hc.StringToSlice(password)) != 0 {

		usersDatabase, err := os.OpenFile("./server/database/users.csv", os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
		}

		defer usersDatabase.Close()

		csvReaderUserslDB := csv.NewReader(usersDatabase)
		getDataUsersDB, err := csvReaderUserslDB.ReadAll()
		if err != nil {
			fmt.Println(err)
		}

		for _, user := range getDataUsersDB {
			if user[0] == input {
				if user[1] == password {
					return "true"
				} else {
					return "WrongPassWord"

				}
			}
		}

		return "false"

	} else {
		return "false"
	}
}
