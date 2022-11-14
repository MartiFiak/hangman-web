package hangman

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear(){
	/*
	Input :
	Clear the terminal with a custom command depending on the os
	Output :
	*/
	if runtime.GOOS == "windows" {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
	}else {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}

func Len(input []string)int{
	/*
	Input :	
	Returns the length of the Slice returned based on the bone
	Output :
	*/
	if runtime.GOOS == "windows" {
        return len(input)
	}else {
        return len(input)+1
	}

}