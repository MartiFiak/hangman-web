package hangmanweb

import (
	"math/rand"
	"net/http"
	"strconv"
)

func CookieSession(w http.ResponseWriter, r *http.Request, gameLaunch map[string]Hangman) string {

	cookies, err := r.Cookie("session_token") // Recupère les cookies de session

	if err != nil { // Si il y en a pas
		token := rand.Intn(1000000000)                     // On génére un token
		if !ContaintKey(strconv.Itoa(token), gameLaunch) { // Si il existe pas encore
			http.SetCookie(w, &http.Cookie{ // On ajoute le token au cookies de l'utilisateur
				Name:  "session_token",
				Value: strconv.Itoa(token),
			})
			return strconv.Itoa(token) // On renvoie le token en string
		} else {
			return CookieSession(w, r, gameLaunch) // Si il existe deja on le regenere
		}
	} else {
		return cookies.Value // Si le token est deja dans les cookies on le renvoie
	}

}

func GetCookieAccount(r *http.Request) string {
	cookies, err := r.Cookie("username")
	if err != nil {
		return ""
	} else {
		return cookies.Value
	}
}

func GetCookieStatus(r *http.Request) string {
	cookies, err := r.Cookie("status")
	if err != nil {
		return ""
	} else {
		return cookies.Value
	}
}

func SetCookieAccount(w http.ResponseWriter, input, status string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: input,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "status",
		Value: status,
	})
}
