package hangmanweb

/*
Contient les informations de la partie courante.
*/
type Hangman struct {
	PlayerName string
	UserLevel  int
	UserXpAv   float64
	WordToFind string
	Attempts   int
	LetterUsed string
	Word       string
	Input      string
	Message    string
	Mode       string
}

/*
Contient les informations global du jeu.
*/
type GlobalInfo struct {
	Username   string
	UserLevel  int
	UserXpAv   float64
	Status     string
	DeadSanta  int
	SaveSanta  int
	Ratio      int
	Total      int
	ErrMessage string
}

/*
Contient les informations d'un utilisateur.
*/
type User struct {
	Username string
	Win      int
	Loose    int
	GamePlay int
	Level    int
	Exp      int
}

/*
Contient tous les utilisateurs et leurs informations.
*/
type ScoreboardData struct {
	UsersList []User
}
