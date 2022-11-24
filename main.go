package main

import (
	"encoding/csv"
	"fmt"
	hangmanweb "hangmanweb/hangman-web"
	"net/http"
	"os"
	"strconv"
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

	http.HandleFunc("/home", IndexHandler)
	http.HandleFunc("/", GameHandler)
	http.HandleFunc("/hangman", GameInputHandler)
	http.HandleFunc("/rules", RulesHandler)
	http.HandleFunc("/scoreboard", ScoreHandler)
	http.ListenAndServe(":"+port, nil)

}

func UpdateUserValue(win bool, w http.ResponseWriter, r *http.Request) {
	userDatabase, err := os.OpenFile("./server/database/users.csv", os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	csvReaderUsersDB := csv.NewReader(userDatabase)
	getDataUsersDB, err := csvReaderUsersDB.ReadAll()

	sbUsersList.UsersList = []hangmanweb.User{}

	if len(getDataUsersDB) != 0 {
		for ligne, userIngetData := range getDataUsersDB {
			if userIngetData[0] == gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].PlayerName {
				if win {
					sbUsersList.UsersList = append(sbUsersList.UsersList, hangmanweb.User{userIngetData[0], AtoiWithoutErr(userIngetData[2]) + 1, AtoiWithoutErr(userIngetData[3]), AtoiWithoutErr(userIngetData[4]) + 1})
					getDataUsersDB[ligne][2] = strconv.Itoa(AtoiWithoutErr(getDataUsersDB[ligne][2]) + 1)
					getDataUsersDB[ligne][4] = strconv.Itoa(AtoiWithoutErr(getDataUsersDB[ligne][4]) + 1)
				} else {
					sbUsersList.UsersList = append(sbUsersList.UsersList, hangmanweb.User{userIngetData[0], AtoiWithoutErr(userIngetData[2]), AtoiWithoutErr(userIngetData[3]) + 1, AtoiWithoutErr(userIngetData[4]) + 1})
					getDataUsersDB[ligne][3] = strconv.Itoa(AtoiWithoutErr(getDataUsersDB[ligne][3]) + 1)
					getDataUsersDB[ligne][4] = strconv.Itoa(AtoiWithoutErr(getDataUsersDB[ligne][4]) + 1)
				}
			} else {
				sbUsersList.UsersList = append(sbUsersList.UsersList, hangmanweb.User{userIngetData[0], AtoiWithoutErr(userIngetData[2]), AtoiWithoutErr(userIngetData[3]), AtoiWithoutErr(userIngetData[4])})
			}
		}
	}
	userDatabase.Close()

	userDatabase, err = os.OpenFile("./server/database/users.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	csvWriterUsersDB := csv.NewWriter(userDatabase)

	err = csvWriterUsersDB.WriteAll(getDataUsersDB)
	if err != nil {
		fmt.Println(err)
	}
	defer csvWriterUsersDB.Flush()
}

func UpdateGlobalValue(save bool) {

	globalDatabase, err := os.OpenFile("./server/database/global.csv", os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	csvReaderGlobalDB := csv.NewReader(globalDatabase)
	getDataGlobalDB, err := csvReaderGlobalDB.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	if len(getDataGlobalDB) != 0 {
		globaldata.DeadSanta, _ = strconv.Atoi(getDataGlobalDB[0][0])
		globaldata.SaveSanta, _ = strconv.Atoi(getDataGlobalDB[0][1])
	} else {
		globaldata.DeadSanta = 0
		globaldata.SaveSanta = 0
	}
	globalDatabase.Close()

	globalDatabase, err = os.OpenFile("./server/database/global.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	csvWriterGlobalDB := csv.NewWriter(globalDatabase)

	if save {
		globaldata.SaveSanta++
	} else {
		globaldata.DeadSanta++
	}

	newData := []string{strconv.Itoa(globaldata.DeadSanta), strconv.Itoa(globaldata.SaveSanta)}
	err = csvWriterGlobalDB.Write(newData)
	if err != nil {
		fmt.Println(err)
	}
	defer csvWriterGlobalDB.Flush()

	if globaldata.SaveSanta+globaldata.DeadSanta != 0 {
		globaldata.Ratio = globaldata.SaveSanta * 100 / (globaldata.SaveSanta + globaldata.DeadSanta)
	} else {
		globaldata.Ratio = 50
	}

	globaldata.Total = globaldata.DeadSanta + globaldata.SaveSanta

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
			sbUsersList.UsersList = append(sbUsersList.UsersList, hangmanweb.User{userIngetData[0], AtoiWithoutErr(userIngetData[2]), AtoiWithoutErr(userIngetData[3]), AtoiWithoutErr(userIngetData[4])})
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

func AtoiWithoutErr(str string) int {
	inte, _ := strconv.Atoi(str)
	return inte
}

func StartGame(input, difficulty string, w http.ResponseWriter, r *http.Request) {
	dataList = hangmanweb.InitGame(difficulty)
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
			attempts, _ := strconv.Atoi(dataList[3])
			if dataList[0] == "Okey" {

				gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)] = hangmanweb.Hangman{
					PlayerName: gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].PlayerName,
					WordToFind: gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].WordToFind,
					Attempts:   attempts,
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
				attempts, _ := strconv.Atoi(dataList[3])

				gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)] = hangmanweb.Hangman{
					PlayerName: gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].PlayerName,
					WordToFind: gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].WordToFind,
					Attempts:   attempts,
					LetterUsed: dataList[4],
					Word:       dataList[1],
					Input:      input,
					Message:    dataList[0],
					Mode:       gameLaunch[hangmanweb.CookieSession(w, r, gameLaunch)].Mode,
				}

				switch dataList[0] {
				case "WinPage":
					UpdateGlobalValue(true)
					UpdateUserValue(true, w, r)
				case "LoosePage":
					UpdateGlobalValue(false)
					UpdateUserValue(false, w, r)
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

	globaldata = hangmanweb.InitGlobalValue(globaldata)
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
			input := r.FormValue("input")
			password := r.FormValue("password")
			if hangmanweb.InputUsernameTreatment(input, password) {
				StartGame(input, difficulty, w, r)
				http.Redirect(w, r, "/hangman", http.StatusFound)
				return
			}
		}
	default:
	}
	tmpl.Execute(w, globaldata)

}
