package hangman

import (
	"math/rand"
	"time"
)

func RandomReveal(nb int, word string, wordToFind string) (string, []string) {
	/*
	Input : A number, the state of a word, and the word to find
	Add nb letter of the word to find in the word state
	Output : Returns the new state of the word and a Slice containing the added letters
	*/
	unValide := []int{}

	_word := StringToSlice(word)
	var useLettre []string
	for i:= 0; i < nb; i++{
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(Len(StringToSlice(wordToFind)))
		isUse := false
		for _, _lettre := range useLettre{
			if _lettre == string(wordToFind[random]) {
				isUse = true
			}
		}
		_count := 0
		for _, _lettre := range wordToFind {
			if string(_lettre) == string(wordToFind[random]) {
				_count ++
			}
		}
		for _, _lettre := range unValide {
			if _lettre == random {
				isUse = true
			}
		}
		if !isUse && _count <= nb-i{
		 
			useLettre = append(useLettre, string(wordToFind[random]))
			for _index, _lettre := range wordToFind {
				if wordToFind[random] == wordToFind[_index]{
					_word[_index] = string(_lettre)
					if random == _index {
					} else {
						i++
					}
				}      
			}
			
		}else {
			if _count > nb-i {
				unValide = append(unValide, random)
			}
			i--
		}		
	}
	word = SliceToString(_word)
	return word, useLettre
}

func RandomRevealPS(mode int, wordToFind string) (string, []string){
	word := ""

	for i := 0; i < Len(StringToSlice(wordToFind)); i++ {
		word += "_"
	}

	nbToReveal := 0
	switch mode{
		case 0:
			nbToReveal = Len(StringToSlice(word))/2-1
		case 1:
			nbToReveal = Len(StringToSlice(word))/3-1
		case 2:
			nbToReveal = 0
	}

	return RandomReveal(nbToReveal, word, wordToFind)
}