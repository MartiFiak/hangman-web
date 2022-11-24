package hangmanweb

import (
	"strconv"
	"fmt"
)

/*
Utilise la fonction Atoi et gère les erreurs pour renvoyer le résultat de Atoi sans l'erreur.
*/
func AtoiWithoutErr(str string) int {
	intstr, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return intstr
}