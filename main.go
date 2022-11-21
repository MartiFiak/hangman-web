package main

import (
	"fmt"
	hangmanweb "hangmanweb/hangman-web"
	hc "hangmanweb/hangman-classic/fonctions"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"net"
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
	tmpl.Execute(w, nil)
}

func StartGame(input, difficulty string,  r *http.Request) {
	dataList = hangmanweb.InitGame(difficulty)
	gameLaunch[r.Header.Get("X-Forwarded-For")] = Hangman{
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
				StartGame(gameLaunch[r.Header.Get("X-Forwarded-For")].PlayerName, gameLaunch[r.Header.Get("X-Forwarded-For")].Mode, r)
				http.Redirect(w, r, "/hangman", http.StatusFound)
			case "Leave":
				http.Redirect(w, r, "/home", http.StatusFound)
			}
			input := r.Form.Get("input")
			dataList = hangmanweb.InputTreatment(gameLaunch[r.Header.Get("X-Forwarded-For")].Word, gameLaunch[r.Header.Get("X-Forwarded-For")].WordToFind, input, gameLaunch[r.Header.Get("X-Forwarded-For")].LetterUsed, 0, gameLaunch[r.Header.Get("X-Forwarded-For")].Attempts)
			attempts, _ := strconv.Atoi(dataList[3])
			if dataList[0] == "Okey" {
				//gameLaunch[r.Header.Get("X-Forwarded-For")].Attempts = attempts
				//gameLaunch[r.Header.Get("X-Forwarded-For")].LetterUsed = dataList[4]
				//gameLaunch[r.Header.Get("X-Forwarded-For")].Word = dataList[1]
				//gameLaunch[r.Header.Get("X-Forwarded-For")].Input = input

				gameLaunch[r.Header.Get("X-Forwarded-For")] = Hangman{
					PlayerName: gameLaunch[r.Header.Get("X-Forwarded-For")].PlayerName,
					WordToFind: gameLaunch[r.Header.Get("X-Forwarded-For")].WordToFind,
					Attempts:   attempts,
					LetterUsed: dataList[4],
					Word:       dataList[1],
					Input:      input,
					Message:    gameLaunch[r.Header.Get("X-Forwarded-For")].Message,
					Mode:       gameLaunch[r.Header.Get("X-Forwarded-For")].Mode,
				}

				
				fmt.Println(gameLaunch)


				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else if dataList[0] == "Nope" {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			} else {
				attempts, _ := strconv.Atoi(dataList[3])

				gameLaunch[r.Header.Get("X-Forwarded-For")] = Hangman{
					PlayerName: gameLaunch[r.Header.Get("X-Forwarded-For")].PlayerName,
					WordToFind: gameLaunch[r.Header.Get("X-Forwarded-For")].WordToFind,
					Attempts:   attempts,
					LetterUsed: dataList[4],
					Word:       dataList[1],
					Input:      input,
					Message:    dataList[0],
					Mode:       gameLaunch[r.Header.Get("X-Forwarded-For")].Mode,
				}
				
				fmt.Println(gameLaunch)
				//gameLaunch[r.Header.Get("X-Forwarded-For")].Attempts = attempts
				//gameLaunch[r.Header.Get("X-Forwarded-For")].LetterUsed = dataList[4]
				//gameLaunch[r.Header.Get("X-Forwarded-For")].Word = dataList[1]
				//gameLaunch[r.Header.Get("X-Forwarded-For")].Input = input
				//gameLaunch[r.Header.Get("X-Forwarded-For")].Message = dataList[0]
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

	if gameLaunch[r.Header.Get("X-Forwarded-For")].Mode != "easy" && gameLaunch[r.Header.Get("X-Forwarded-For")].Mode != "medium" && gameLaunch[r.Header.Get("X-Forwarded-For")].Mode != "hard" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	} else {
		tmpl.Execute(w, gameLaunch[r.Header.Get("X-Forwarded-For")])
	}
}

func getMacAddr() ([]string, error) {
    ifas, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    var as []string
    for _, ifa := range ifas {
        a := ifa.HardwareAddr.String()
        if a != "" {
            as = append(as, a)
        }
    }
    return as, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./server/index.html"))

	//ips := r.Header.Get("X-Forwarded-For")
	//splitIps := strings.Split(ips, ",")

	macaddr, err := getMacAddr()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hc.SliceToString(macaddr))
	}

	gameLaunch[r.Header.Get("X-Forwarded-For")] = Hangman{
		PlayerName: "hc.SliceToString(macaddr)",
		WordToFind: "",
		Attempts:   10,
		LetterUsed: "",
		Word:       "",
		Input:      "",
		Message:    "",
		Mode:       "",
	}

	fmt.Println(gameLaunch)




	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		} else {
			difficulty := r.Form.Get("difficulty")
			input := r.FormValue("input")
			if hangmanweb.InputUsernameTreatment(input) {
				StartGame(input, difficulty, r)
				http.Redirect(w, r, "/hangman", http.StatusFound)
				return
			}
		}
	default:
	}
	tmpl.Execute(w, nil)

}
