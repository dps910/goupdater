package app

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	extract = flag.String("-tarfile", "", "Extract given tar file to directory specified")
)

func init() {
	flag.Parse()
}

// Install golang to directory
func Install(dir string, r io.Reader) error {
	// Check that tar file exists

	info, err := os.Stat(dir)
	if err != nil {
		return errors.New("Couldn't get info of folder")
	}

	// Check if dir, if not return error
	if !info.IsDir() {
		return errors.New("This is not a directory")
	}

	// Check permissions, make sure that we have permission to read dir
	perm := info.Mode().Perm()

	// Make sure to print in octal representation (base 8) instead of base 10 https://stackoverflow.com/a/41259532
	base8perm := fmt.Sprintf("%o", perm)

	// 433 is r---wx-wx which I think would be the lowest chmod for reading.
	if base8perm != "755" {
		// Get user input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Do you want to change permissions to 755? [y/n] Current permissions: %s", perm)

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

		// Check if go directory exists, if it does then delete it (not going to yet, but when this is fully working)
		_, err := ioutil.ReadDir(dir)
		if err != nil {
			return errors.New("Directory doesn't exist")
		}
		fmt.Println("Directory does exist")

		// If go directory does exist, then we extract tar.gz file to /tmp temporarily and then move it
		readGzip, err := gzip.NewReader(r)
		if err != nil {
			return errors.New("Couldn't read file")
		}
		defer readGzip.Close()

		// Reader loops over tar file and creates the file structure in dir
		// Read files in archive
		tarfile := tar.NewReader(readGzip)
		for {
			header, err := tarfile.Next()

			switch {
			// If entries in the tar file are finished, call EOF()
			case err == io.EOF:
				fmt.Printf("Reached end of archive, size %d", header.Size)
				break
			case err != nil:
				return errors.New("tar.gz is nil")
			}

			switch header.Typeflag {
			// Directory
			case tar.TypeDir:
				if err := os.Mkdir(dir, 0755); err != nil {
					log.Fatalln(err)
				}

			// File
			case tar.TypeReg:
				// Open files in header
				f, err := os.OpenFile(header.Name, os.O_CREATE|os.O_RDWR, 0755)
				if err != nil {
					fmt.Println("Couldn't create file")
				}

				c, err := io.Copy(f, tarfile)
				if err != nil {
					fmt.Println("Couldn't copy read file to f")
				}
				fmt.Println(c)
			}
		}
	}

	return nil
}
