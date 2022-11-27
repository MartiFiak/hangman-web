package hangmanweb

func ContaintKey(key string, _map map[string]Hangman) bool{
	/*
	Verify if key is in map
	*/
	if _, isPresent := _map[key]; isPresent {
		return true
	}
	return false
}