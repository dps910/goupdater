package app

import (
  "fmt"
  "os"
  "flag"
)

var (
  c = flag.Int("-c", 0755, "chmod number")
  // just assuming the dir for now
  DIR = "/usr/local/go"
)

func Uninstall() error {
  // check directory exists
  dir, err := (DIR; os.isNotExist(err) {
    return errors.New("go directory does not exist")
  }

  perm := dir.Mode().Perm

  base8perm := os.Sprintf("%o\n", perm)

  if base8perm < string(433) {
    Chmod(DIR, os.FileMode(c))
  }

  // If go directory does not exist, install golang to that directory
  Install()

  
  return nil
}
