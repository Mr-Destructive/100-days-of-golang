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
	log.Println(filepath.Match("paths*", filepath.Base(dir)))
	s, _ = filepath.Abs("operators/arithmetic/")
	log.Println(s)
	log.Println(dir)
	log.Println(filepath.Rel(s, dir))

	log.Println(filepath.Join("golang", "files"))
	log.Println(filepath.Join(s, "/files", "//read"))
}
