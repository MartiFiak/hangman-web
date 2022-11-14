package hangman

func SliceToString(s []string) string {
	/*
	Input : One Slice
	Convert Slice to String
	Output : A String
	*/
	str := ""
	for _, i := range s {
		str += string(i)
	}
	return str
}

func StringToSlice(s string) []string {
	/*
	Input: A String
	Convert String to Slice
	Output: One Slice
	*/
	str := []string{}
	for _, i := range s {
		str = append(str, string(i))
	}
	return str
}

func SliceSliceToSlice(s [][]string) []string{
	/*
	Input: A Slice of Slice
	Convert a Slice of Slice to Slice
	Output: One Slice
	*/
	var sToReturn []string
	for _, sIn := range s {
		sToReturn = append(sToReturn, SliceToString(sIn))
	}
	return sToReturn
}