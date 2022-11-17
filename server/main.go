package main

import (
	"fmt"
	"hangmanweb"
	"net/http"
	"strconv"
	"text/template"
)

var dataList []string

type Hangman struct {
	PlayerName string
	WordToFind string
	MaVariable string
	Attempts   int
	LetterUsed string
	Word       string
	Input      string
	Message    string
}

var data Hangman

func main() {
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/hangman", GameHandler)
	http.HandleFunc("/rules", RulesHandler)
	http.ListenAndServe(":8080", nil)
}

func RulesHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("rules.html"))
	tmpl.Execute(w, nil)
}

func GameHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("game.html"))

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		} else {
			input := r.Form.Get("input")
			dataList = hangmanweb.InputTreatment(data.Word, data.WordToFind, input, data.LetterUsed, 0, data.Attempts)
			attempts, _ := strconv.Atoi(dataList[3])
			if dataList[0] == "Okey" {
				data.Attempts = attempts
				data.LetterUsed = dataList[4]
				data.Word = dataList[1]
				data.Input = input
				tmpl.Execute(w, data)
				return
			} else if dataList[0] == "Nop" {
				tmpl.Execute(w, data)
				return
			} else {
				data.Attempts = attempts
				data.LetterUsed = dataList[4]
				data.Word = dataList[1]
				data.Input = input
				data.Message = dataList[0]
				tmpl.Execute(w, data)
				return
			}
		}
	default:
		tmpl.Execute(w, data)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("index.html"))

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		} else {
			input := r.FormValue("input")
			if hangmanweb.InputUsernameTreatment(input) {
				dataList = hangmanweb.InitGame()
				data = Hangman{
					PlayerName: input,
					WordToFind: dataList[0],
					Attempts:   10,
					LetterUsed: dataList[2],
					Word:       dataList[1],
					Input:      "",
					Message:    "",
				}
				http.Redirect(w, r, "/hangman", http.StatusFound)
				return
			}
		}
	default:
	}
	tmpl.Execute(w, nil)

}
