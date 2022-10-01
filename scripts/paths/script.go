package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	// get the working directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println(dir)

	// get the user home directory of the user
	dir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	log.Println(dir)

	// get the base name of the path
	s := filepath.Base("math/")
	log.Println(s)

	// Get the file info for a given path
	m, _ := os.Stat(s)
	log.Println(os.FileInfo(m).Name())

	// return true if the provided path is absolute
	log.Println(filepath.IsAbs(s))

	// get the paretn directory of the path
	log.Println(filepath.Dir(s))

	k := "bit.go"
	//os.Chdir("math/")
	p, _ := filepath.Abs(k)

	// teim the string before the suffix from the end
	log.Println(strings.TrimSuffix(k, path.Ext(k)))

	// get the extension of the path
	log.Println(path.Ext(k))

	// check if the path exists or not
	if t, err := os.Stat(p); os.IsNotExist(err) {
		log.Println("No, " + p + " does not exists")
	} else {
		log.Println(t.IsDir())
		log.Println("Yes, " + p + " exists")
	}

}
