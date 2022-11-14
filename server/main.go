package main

import (
	"fmt"
	"hangmanweb"
	"net/http"
	"strconv"
	"text/template"
)

var dataList = hangmanweb.InitGame()

type Hangman struct {
	WordToFind string
	MaVariable string
	Attempts   int
	LetterUsed string
	Word       string
	Input      string
    Message    string
}

var data = Hangman{
	WordToFind: dataList[0],
	Attempts:   10,
	LetterUsed: dataList[2],
	Word:       dataList[1],
	Input:      "",
	Message:      "",
}

func main() {
	http.HandleFunc("/", Handler)

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hangman", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("index.html"))

	switch r.Method {
	case "POST": // Gestion d'erreur
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
                data.Message = dataList[0]
                tmpl.Execute(w, data)
				return
            }
		}
	default:
		tmpl.Execute(w, data)
	}

}
