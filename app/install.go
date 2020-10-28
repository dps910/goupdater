package app

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	dir = "$HOME/test"
)

func Install() error {
	info, err := os.Stat("dir")
	if err != nil {
		return errors.New("Couldn't get info of folder")
	}

	// Check if dir, if not return error
	if !info.IsDir() {
		return errors.New("This is not a directory.")
	}

	// Check permissions, make sure that we have permission to read dir
	perm := info.Mode().Perm()

	// Make sure to print in octal representation (base 8) instead of base 10 https://stackoverflow.com/a/41259532
	base8perm := fmt.Sprintf("%o\n", perm)

	// 433 is r---wx-wx which I think would be the lowest chmod for reading.
	if base8perm < string(433) {
		// Get user input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to change permissions to 755? [y/n]")

		// Read string
		text, _ := reader.ReadString('\n')
		if text == "y" {
			err = os.Chmod(dir, 0755)
			if err != nil {
				return errors.New("Couldn't change dir permissions")
			}
		} else if text == "n" {
			// Make program exit
			fmt.Print("This program can't work then :S")
			os.Exit(0)
		} else {
			os.Exit(0)
		}

		// Chcek if /usr/local/go exists, if it does then delete it (not going to yet, but when this is fully working)
		godir := "~/test"
		_, err := ioutil.ReadDir(godir)
		if err != nil {
			return errors.New("Directory doesn't exist")
		}

		// If directory does exist, delete directory
		_ = os.Remove(godir)
	}

	return nil
}

func main() {
	Install()
}
