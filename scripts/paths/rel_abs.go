package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println(dir)
	s := filepath.Base("math/")
	k := "/"
	log.Println(filepath.IsAbs(k))

	log.Println(filepath.Base(dir))

	// return true if the pattern matches with the path
	log.Println(filepath.Match("paths*", filepath.Base(dir)))

	// get the absolute path of the provided path
	s, _ = filepath.Abs("operators/arithmetic/")
	log.Println(s)
	log.Println(dir)

	// get the relative path between the base and target directory
	log.Println(filepath.Rel(s, dir))

	// join the two path strings as one path
	log.Println(filepath.Join("golang", "files"))
	log.Println(filepath.Join(s, "/files", "//read"))
}
