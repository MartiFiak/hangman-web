package main

import (
	"fmt"
	hangmanweb "hangmanweb/hangman-web"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"math/rand"
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

var gameLaunch = make(map[string]Hangman)

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

func ContaintKey(key string, _map map[string]Hangman) bool{
	if _, isPresent := _map[key]; isPresent {
		return true
	}
	return false
}

func CookieSession(w http.ResponseWriter, r *http.Request) string {

	cookies, err := r.Cookie("session_token")

	if err != nil {
		token := rand.Intn(1000000000)
		if !ContaintKey(strconv.Itoa(token), gameLaunch) {
			http.SetCookie(w, &http.Cookie{
				Name: "session_token",
				Value: strconv.Itoa(token),
			})
			return strconv.Itoa(token)
		} else {
			return CookieSession(w, r)
		}
	} else {
		return cookies.Value
	}

}

func ScoreHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/scoreboard.html"))
	tmpl.Execute(w, nil)
}

func StartGame(input, difficulty string,w http.ResponseWriter, r *http.Request) {
	dataList = hangmanweb.InitGame(difficulty)
	gameLaunch[CookieSession(w,r)] = Hangman{
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
				StartGame(gameLaunch[CookieSession(w,r)].PlayerName, gameLaunch[CookieSession(w,r)].Mode, w, r)
				http.Redirect(w, r, "/hangman", http.StatusFound)
			case "Leave":
				http.Redirect(w, r, "/home", http.StatusFound)
			}
			input := r.Form.Get("input")
			dataList = hangmanweb.InputTreatment(gameLaunch[CookieSession(w,r)].Word, gameLaunch[CookieSession(w,r)].WordToFind, input, gameLaunch[CookieSession(w,r)].LetterUsed, 0, gameLaunch[CookieSession(w,r)].Attempts)
			attempts, _ := strconv.Atoi(dataList[3])
			if dataList[0] == "Okey" {

				gameLaunch[CookieSession(w,r)] = Hangman{
					PlayerName: gameLaunch[CookieSession(w,r)].PlayerName,
					WordToFind: gameLaunch[CookieSession(w,r)].WordToFind,
					Attempts:   attempts,
					LetterUsed: dataList[4],
					Word:       dataList[1],
					Input:      input,
					Message:    gameLaunch[CookieSession(w,r)].Message,
					Mode:       gameLaunch[CookieSession(w,r)].Mode,
				}

				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else if dataList[0] == "Nope" {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else {
				attempts, _ := strconv.Atoi(dataList[3])

				gameLaunch[CookieSession(w, r)] = Hangman{
					PlayerName: gameLaunch[CookieSession(w,r)].PlayerName,
					WordToFind: gameLaunch[CookieSession(w,r)].WordToFind,
					Attempts:   attempts,
					LetterUsed: dataList[4],
					Word:       dataList[1],
					Input:      input,
					Message:    dataList[0],
					Mode:       gameLaunch[CookieSession(w,r)].Mode,
				}

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

	fmt.Println(gameLaunch)

	if gameLaunch[CookieSession(w,r)].Mode != "easy" && gameLaunch[CookieSession(w,r)].Mode != "medium" && gameLaunch[CookieSession(w,r)].Mode != "hard" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	} else {
		tmpl.Execute(w, gameLaunch[CookieSession(w,r)])
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	CookieSession(w,r)

	tmpl := template.Must(template.ParseFiles("./server/index.html"))

	gameLaunch[CookieSession(w,r)] = Hangman{
		PlayerName: "unknown",
		WordToFind: "",
		Attempts:   10,
		LetterUsed: "",
		Word:       "",
		Input:      "",
		Message:    "",
		Mode:       "",
	}

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		} else {
			difficulty := r.Form.Get("difficulty")
			input := r.FormValue("input")
			if hangmanweb.InputUsernameTreatment(input) {
				StartGame(input, difficulty, w, r)
				http.Redirect(w, r, "/hangman", http.StatusFound)
				return
			}
		}
	default:
	}
	tmpl.Execute(w, nil)

}
