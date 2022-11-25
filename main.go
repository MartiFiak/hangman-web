package main

import (
	"encoding/csv"
	"fmt"
	hangmanweb "hangmanweb/hangman-web"
	"net/http"
	"os"
	"text/template"
)

var dataList []string
var sbUsersList hangmanweb.ScoreboardData
var globaldata hangmanweb.GlobalInfo
var gameLaunch = make(map[string]hangmanweb.Hangman)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("Game lauch in localhost:8080")
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./server"))
	http.Handle("/server/", http.StripPrefix("/server/", fs))

	http.HandleFunc("/logout", LogOutHundler)
	http.HandleFunc("/register", RegisterHundler)
	http.HandleFunc("/home", IndexHandler)
	http.HandleFunc("/", GameHandler)
	http.HandleFunc("/hangman", GameInputHandler)
	http.HandleFunc("/rules", RulesHandler)
	http.HandleFunc("/scoreboard", ScoreHandler)
	http.ListenAndServe(":"+port, nil)

}

func ScoreHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/scoreboard.html"))

	//Lecture de la base user
	usersDatabase, err := os.OpenFile("./server/database/users.csv", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	defer usersDatabase.Close()

	csvReaderUsersDB := csv.NewReader(usersDatabase)
	getDataUsersDB, err := csvReaderUsersDB.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	sbUsersList.UsersList = []hangmanweb.User{}

	if len(getDataUsersDB) != 0 {
		for _, userIngetData := range getDataUsersDB {
			sbUsersList.UsersList = append(sbUsersList.UsersList, hangmanweb.User{userIngetData[0], hangmanweb.AtoiWithoutErr(userIngetData[2]), hangmanweb.AtoiWithoutErr(userIngetData[3]), hangmanweb.AtoiWithoutErr(userIngetData[4])})
		}
	}
	// Tri des données        rang basé sur win*ratio

	for i := 0; i < len(sbUsersList.UsersList); i++ {
		for j := i; j < len(sbUsersList.UsersList); j++ {
			var iratio int
			var jratio int
			if sbUsersList.UsersList[i].GamePlay != 0 {
				iratio = sbUsersList.UsersList[i].Win * (sbUsersList.UsersList[i].Win * 100 / (sbUsersList.UsersList[i].GamePlay))
			} else {
				iratio = 1
			}
			if sbUsersList.UsersList[j].GamePlay != 0 {
				jratio = sbUsersList.UsersList[j].Win * (sbUsersList.UsersList[j].Win * 100 / (sbUsersList.UsersList[j].GamePlay))
			} else {
				jratio = 1
			}

			if jratio > iratio {
				sbUsersList.UsersList[i], sbUsersList.UsersList[j] = sbUsersList.UsersList[j], sbUsersList.UsersList[i]
			}
		}
	}

	// Penser a actualisé le scoarboard

	tmpl.Execute(w, sbUsersList)
}

func LogOutHundler(w http.ResponseWriter, r *http.Request) {
	hangmanweb.SetCookieAccount(w, "", "logout")
	http.Redirect(w, r, "/home", http.StatusFound)
}

func RegisterHundler(w http.ResponseWriter, r *http.Request){
	hangmanweb.SetCookieAccount(w, "", "register")
	http.Redirect(w, r, "/home", http.StatusFound)
}

func StartGame(input, difficulty string, w http.ResponseWriter, r *http.Request) {
	dataList = hangmanweb.InitGame(difficulty)
	hangmanweb.SetCookieAccount(w, input, "login")
	gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)] = hangmanweb.Hangman{
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
				StartGame(gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].PlayerName, gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Mode, w, r)
				http.Redirect(w, r, "/hangman", http.StatusFound)
			case "Leave":
				http.Redirect(w, r, "/home", http.StatusFound)
			}
			input := r.Form.Get("input")
			dataList = hangmanweb.InputTreatment(gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Word, gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].WordToFind, input, gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].LetterUsed, 0, gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Attempts)
			if dataList[0] == "Okey" {

				gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)] = hangmanweb.Hangman{
					PlayerName: gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].PlayerName,
					WordToFind: gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].WordToFind,
					Attempts:   hangmanweb.AtoiWithoutErr(dataList[3]),
					LetterUsed: dataList[4],
					Word:       dataList[1],
					Input:      input,
					Message:    gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Message,
					Mode:       gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Mode,
				}

				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else if dataList[0] == "Nope" {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else {
				gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)] = hangmanweb.Hangman{
					PlayerName: gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].PlayerName,
					WordToFind: gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].WordToFind,
					Attempts:   hangmanweb.AtoiWithoutErr(dataList[3]),
					LetterUsed: dataList[4],
					Word:       dataList[1],
					Input:      input,
					Message:    dataList[0],
					Mode:       gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Mode,
				}

				switch dataList[0] {
				case "WinPage":
					globaldata = hangmanweb.UpdateGlobalValue(r, true, globaldata)
					sbUsersList = hangmanweb.UpdateUserValue(true, w, r, sbUsersList, gameLaunch)
				case "LoosePage":
					globaldata = hangmanweb.UpdateGlobalValue(r, false, globaldata)
					sbUsersList = hangmanweb.UpdateUserValue(false, w, r, sbUsersList, gameLaunch)
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

	if gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Mode != "easy" && gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Mode != "medium" && gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Mode != "hard" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	} else {
		tmpl.Execute(w, gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)])
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	globaldata = hangmanweb.InitGlobalValue(r, globaldata)
	hangmanweb.CookieSession(w, r, gameLaunch)

	tmpl := template.Must(template.ParseFiles("./server/index.html"))

	gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)] = hangmanweb.Hangman{
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
			if hangmanweb.GetCookieAccount(r) != "" {
				StartGame(hangmanweb.GetCookieAccount(r), difficulty, w, r)
				http.Redirect(w, r, "/hangman", http.StatusFound)
			} else {
				input := r.FormValue("input")
				password := r.FormValue("password")
				if hangmanweb.InputUsernameTreatment(input, password) {
					StartGame(input, difficulty, w, r)
					http.Redirect(w, r, "/hangman", http.StatusFound)
					return
				}
			}
		}
	default:
	}
	tmpl.Execute(w, globaldata)

}
