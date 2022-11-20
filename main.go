package main

import (
	"fmt"
	hangmanweb "hangmanweb/hangman-web"
	hc "hangmanweb/hangman-classic/fonctions"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"strings"
)

var dataList []string

type Hangman struct {
	PlayerName string
	WordToFind string
	Attempts   int
	LetterUsed string
	Word       string
	Input      string
	Message    string
	Mode       string
}

var gameLaunch map[string]Hangman

var data Hangman

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("Game lauch in localhost:8080")
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./server"))
	http.Handle("/server/", http.StripPrefix("/server/", fs))

	http.HandleFunc("/home", IndexHandler)
	http.HandleFunc("/", GameHandler)
	http.HandleFunc("/hangman", GameInputHandler)
	http.HandleFunc("/rules", RulesHandler)
	http.HandleFunc("/scoreboard", ScoreHandler)
	http.ListenAndServe(":"+port, nil)
}

func ScoreHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/scoreboard.html"))
	tmpl.Execute(w, gameLaunch[data.PlayerName])
}

func StartGame(input, difficulty string) {
	dataList = hangmanweb.InitGame(difficulty)
	data = Hangman{
		PlayerName: input,
		WordToFind: dataList[0],
		Attempts:   10,
		LetterUsed: dataList[2],
		Word:       dataList[1],
		Input:      "",
		Message:    "Okey",
		Mode:       difficulty,
	}
}

func RulesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/rules.html"))
	tmpl.Execute(w, nil)
}

func GameInputHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		} else {
			endscreeninput := r.Form.Get("endscreeninput")
			switch endscreeninput {
			case "Restart":
				StartGame(data.PlayerName, data.Mode)
				http.Redirect(w, r, "/hangman", http.StatusFound)
			case "Leave":
				http.Redirect(w, r, "/home", http.StatusFound)
			}
			input := r.Form.Get("input")
			fmt.Println(input)
			dataList = hangmanweb.InputTreatment(data.Word, data.WordToFind, input, data.LetterUsed, 0, data.Attempts)
			attempts, _ := strconv.Atoi(dataList[3])
			if dataList[0] == "Okey" {
				data.Attempts = attempts
				data.LetterUsed = dataList[4]
				data.Word = dataList[1]
				data.Input = input
				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else if dataList[0] == "Nope" {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else {
				attempts, _ := strconv.Atoi(dataList[3])
				data.Attempts = attempts
				data.LetterUsed = dataList[4]
				data.Word = dataList[1]
				data.Input = input
				data.Message = dataList[0]
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/game.html"))
	fmt.Println(gameLaunch[data.PlayerName])
	if gameLaunch[data.PlayerName].Mode != "easy" && gameLaunch[data.PlayerName].Mode != "medium" && gameLaunch[data.PlayerName].Mode != "hard" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	} else {
		tmpl.Execute(w, gameLaunch[data.PlayerName])
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./server/index.html"))

	

	ips := r.Header.Get("X-Forwarded-For")
	//splitIps := strings.Split(ips, ",")
	gameLaunch[hc.SliceToString(strings.Split(r.Header.Get("X-Forwarded-For"), ","))] = Hangman{
		PlayerName: "unknown",
		WordToFind: "",
		Attempts:   10,
		LetterUsed: "",
		Word:       "",
		Input:      "",
		Message:    "",
		Mode:       "",
	}

	fmt.Println(gameLaunch[ips])




	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		} else {
			difficulty := r.Form.Get("difficulty")
			input := r.FormValue("input")
			if hangmanweb.InputUsernameTreatment(input) {
				StartGame(input, difficulty)
				http.Redirect(w, r, "/hangman", http.StatusFound)
				return
			}
		}
	default:
	}
	tmpl.Execute(w, nil)

}
