package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Chmod(stringordir string, num os.FileMode) error {
	// Return new reader with buffer of default size (4096 bytes)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to change permissions to 755? [y/n]")

	// Read string
	text, _ := reader.ReadString('\n')
	if text == "y" {
		err := os.Chmod(dir, num)
		if err != nil {
			return errors.New("Couldn't chmod dir")
		}
	} else if text == "n" {
		fmt.Print("Program will not be able to run unless it has permissions to read and write.")
		os.Exit(0)
	} else {
		os.Exit(0)
	}
	return nil
}
