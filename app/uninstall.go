package app

import (
  "fmt"
  "os"
)

func Uninstall() error {
  // check directory exists
  _, err := ("/usr/local/go"); os.isNotExist(err) {
    return errors.New("go directory does not exist")
  }

  // If go directory does not exist, install golang to that directory
  Install()

  
  return nil
}
