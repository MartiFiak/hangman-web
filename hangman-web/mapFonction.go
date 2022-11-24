package hangmanweb

func ContaintKey(key string, _map map[string]Hangman) bool{
	if _, isPresent := _map[key]; isPresent {
		return true
	}
	return false
}