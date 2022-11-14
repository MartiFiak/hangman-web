package hangman

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"
)

func RoundPisteNoire(mode, wordToFind, word string, vowelsCount, attempts int, hangmanStep [][]string, useLettre []string) {

	tabsPos := 0
	modePos := 1
	psMode := "psMedium"

	var message string

	gameStatu := true

	var inputLetterOrWord string

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	WriteText("'ESC' to quit, '²' to restart, 'Left' or 'Right' to switch tabs, 'Up' or 'Down' to change difficulty.", 0, 0, termbox.ColorWhite)

	CreateBox(0, 1, 62, 3, termbox.ColorWhite)
	WriteText("Welcome", 20, 2, termbox.ColorWhite)
	WriteText("|", 29, 2, termbox.ColorWhite)
	WriteText("Game", 31, 2, termbox.ColorWhite)
	WriteText("|", 36, 2, termbox.ColorWhite)
	WriteText("Help", 38, 2, termbox.ColorWhite)

	WelcomePage(modePos, wordToFind)
	ShowHangman(hangmanStep, attempts)

	for {
		if attempts <= 0 && tabsPos == 1 && gameStatu {
			LosePage(hangmanStep, attempts, wordToFind)
			gameStatu = false
		} else if word == wordToFind && tabsPos == 1 && gameStatu {
			WinPage()
			gameStatu = false
		} else {
			ev := termbox.PollEvent()
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyEsc {
					break
				} else if ev.Key == termbox.KeyArrowLeft {
					if tabsPos > 0 {
						tabsPos--
					} else {
						tabsPos = 2
					}
					ChangePage(attempts, tabsPos, modePos, word, wordToFind, useLettre)
				} else if ev.Key == termbox.KeyArrowUp  && tabsPos == 0{
					if modePos > 0 {
						modePos--
					} else {
						modePos = 2
					}
					word, useLettre, psMode  = SelectMode(modePos, wordToFind)
				} else if ev.Key == termbox.KeyArrowDown && tabsPos == 0{
					if modePos < 2 {
						modePos++
					} else {
						modePos = 0
					}
					word, useLettre, psMode = SelectMode(modePos, wordToFind)
				} else if ev.Key == termbox.KeyArrowRight {
					if tabsPos < 2 {
						tabsPos++
					} else {
						tabsPos = 0
					}
					ChangePage(attempts, tabsPos, modePos, word, wordToFind, useLettre)
				} else if (ev.Key == termbox.KeyDelete || ev.Key == termbox.KeyBackspace) && tabsPos == 1 && Len(StringToSlice(inputLetterOrWord)) != 0 {
					inputLetterOrWord = inputLetterOrWord[:Len(StringToSlice(inputLetterOrWord))-1]
					ClearInput()
					WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
				} else if ev.Key == termbox.KeyEnter && tabsPos == 1 && Len(StringToSlice(inputLetterOrWord)) != 0 {
					word, vowelsCount, attempts, useLettre, message = InputProcessing(psMode, word, wordToFind, inputLetterOrWord, attempts, vowelsCount, useLettre, "")
					if vowelsCount == -100 {
						gameStatu = false
						ClearPage()
						WinPage()
					}
					if gameStatu {
						ClearInput()
						inputLetterOrWord = ""
						WriteText(message, 4, 29, termbox.ColorWhite)
						ShowHangman(hangmanStep, attempts)
						ReloadInfo(word, strconv.Itoa(attempts), SliceToString(useLettre))
					}
				}
				if ev.Ch == '²' {

					word = ""
					attempts = 10
					useLettre = []string{}

					rand.Seed(time.Now().UnixNano())

					wordToFind = Read(".\\assets\\words.txt")[rand.Intn(len(Read(".\\assets\\words.txt")))]

					for i := 0; i < Len(StringToSlice(wordToFind)); i++ {
						word += "_"
					}

					word, useLettre = RandomReveal(Len(StringToSlice(wordToFind))/2-1, word, wordToFind)

					RoundPisteNoire(mode, wordToFind, word, vowelsCount, attempts, hangmanStep, useLettre)
				}
				if tabsPos == 1 && gameStatu {
					switch ev.Ch {
					case 'a':
						inputLetterOrWord += "a"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'b':
						inputLetterOrWord += "b"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'c':
						inputLetterOrWord += "c"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'd':
						inputLetterOrWord += "d"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'e':
						inputLetterOrWord += "e"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'f':
						inputLetterOrWord += "f"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'g':
						inputLetterOrWord += "g"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'h':
						inputLetterOrWord += "h"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'i':
						inputLetterOrWord += "i"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'j':
						inputLetterOrWord += "j"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'k':
						inputLetterOrWord += "k"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'l':
						inputLetterOrWord += "l"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'm':
						inputLetterOrWord += "m"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'n':
						inputLetterOrWord += "n"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'o':
						inputLetterOrWord += "o"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'p':
						inputLetterOrWord += "p"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'q':
						inputLetterOrWord += "q"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'r':
						inputLetterOrWord += "r"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 's':
						inputLetterOrWord += "s"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 't':
						inputLetterOrWord += "t"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'u':
						inputLetterOrWord += "u"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'v':
						inputLetterOrWord += "v"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'w':
						inputLetterOrWord += "w"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'x':
						inputLetterOrWord += "x"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'y':
						inputLetterOrWord += "y"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					case 'z':
						inputLetterOrWord += "z"
						WriteText(inputLetterOrWord, 3, 13, termbox.ColorWhite)
						termbox.Flush()
					}

				}
			}

		}

	}

	termbox.Close()
}

func ReloadInfo(word, attempts, useLettre string) {
	WriteWord(word, 3, 19, termbox.ColorWhite)
	WriteWord(word, 3, 19, termbox.ColorWhite)
	WriteWord(useLettre, 3, 25, termbox.ColorWhite)
	WriteText("       ", 30, 7, termbox.ColorWhite)
	WriteText(attempts, 30, 7, termbox.ColorWhite)
}

func ClearInput() {
	for i := 12; i < 14; i++ {
		for j := 1; j < 41; j++ {
			termbox.SetCell(j, i, ' ', termbox.RGBToAttribute(0, 0, 0), termbox.RGBToAttribute(0, 0, 0))
			termbox.Flush()
		}
	}
	for i := 2; i < 60; i++ {
		termbox.SetCell(i, 29, ' ', termbox.RGBToAttribute(0, 0, 0), termbox.RGBToAttribute(0, 0, 0))
	}

}

func ChangePage(attempts, page, modePos int, word, wordToFind string, useLettre []string) {
	switch page {
	case 0:
		ClearPage()
		WelcomePage(modePos, wordToFind)
	case 1:
		ClearPage()
		GamePage(strconv.Itoa(attempts), word, SliceToString(useLettre))
	case 2:
		ClearPage()
		HelpPage()
	}
}
