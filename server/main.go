package main

import (
	"net/http"
	"text/template"
    "hangmanweb"
)

const (
	link= "<link rel='stylesheet' href='style.css'>"
)

// Je crée ma structure
type Test struct {
    MaVariable string
    LetterUsed string
    Word string
}

func main() {
    http.HandleFunc("/", Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler
    //
    fs := http.FileServer(http.Dir("./"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    //

    http.HandleFunc("/hangman", Handler) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
    http.ListenAndServe(":8080", nil)    // On lance le serveur local sur le port 8080

    hangmanweb.Hangmanweb()
}

func Handler(w http.ResponseWriter, r *http.Request) {
    // J'utilise la librairie tmpl pour créer un template qui va chercher mon fichier index.html
    tmpl := template.Must(template.ParseFiles("index.html"))

    // Je crée une variable qui définit ma structure
    data := Test{
        MaVariable: "Oueba",
        LetterUsed: "hasdw",
        Word: "Hello",
    }

    // J'execute le template avec les données
    tmpl.Execute(w, data)
}