package hangman

import (
	"bufio"
	"log"
	"os"
)

func Read(txt string) []string {
	/*
	Input: Location of a .txt file
	Open the text file and save the different lines in a Slice
	Output: Slice containing the lines of the .txt file
	*/
	lst_mot := []string{}
	f, err := os.Open(txt)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lst_mot = append(lst_mot, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lst_mot
}
