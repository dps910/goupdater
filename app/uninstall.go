package app

import (
  "fmt"
  "os"
)

func Uninstall() error {
  // check directory exists
  dir, err := ("/usr/local/go"); os.isNotExist(err) {
    return errors.New("go directory does not exist")
  }

  perm := dir.Mode().Perm

  base8perm := os.Sprintf("%o\n", perm)

  if base8perm < string(433) {
    
  }

  // If go directory does not exist, install golang to that directory
  Install()

  
  return nil
}
