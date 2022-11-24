package hangmanweb

import (
	"net/http"
	"math/rand"
	"strconv"
)

func CookieSession(w http.ResponseWriter, r *http.Request, gameLaunch map[string]Hangman) string {

	cookies, err := r.Cookie("session_token")   // Recupère les cookies de session

	if err != nil {                           // Si il y en a pas
		token := rand.Intn(1000000000)              // On génére un token
		if !ContaintKey(strconv.Itoa(token), gameLaunch) {     // Si il existe pas encore
			http.SetCookie(w, &http.Cookie{                 // On ajoute le token au cookies de l'utilisateur
				Name: "session_token",
				Value: strconv.Itoa(token),
			})
			return strconv.Itoa(token)                       // On renvoie le token en string
		} else {
			return CookieSession(w, r, gameLaunch)         // Si il existe deja on le regenere
		}
	} else {
		return cookies.Value                        // Si le token est deja dans les cookies on le renvoie
	}

}