package app

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	c   = flag.Int("-c", 0755, "chmod number")
	dir = flag.String("-dir", "", "Use this flag if the golang directories are somewhere else.")
	// DIR = just assuming the dir for now
	DIR = "/usr/local/go"
	// GOPATH specifies the GOPATH in "go env"
	GOPATH = "$HOME/go"
)

// RemoveDir removes directories
func RemoveDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		// If there is an error here, specify the directories instead
		str := bufio.NewReader(os.Stdin)

		fmt.Println("Couldn't find directories. Specify directories to delete.")

		txt, _ := str.ReadString('\n')

		// Remove directory specified
		os.RemoveAll(txt)
	}
	fmt.Printf("Deleted Go directories %s and %s", DIR, GOPATH)
	return nil
}

// Uninstall golang
func Uninstall() error {
	// check directory exists
	dir, err := os.Stat(DIR)
	if os.IsNotExist(err) {
		return errors.New("Directory doesn't exist")
	}

	perm := dir.Mode().Perm()

	base8perm := fmt.Sprintf("%o\n", perm)

	if base8perm < string(433) {
		Chmod(DIR, os.FileMode(*c))
	}

	// Check if is dir and not a file, should return true
	if dir.IsDir() {
		// Remove directory
		RemoveDir(DIR)
		RemoveDir(GOPATH)
		fmt.Println("Deleted golang directory :)")
	}

	return nil
}
