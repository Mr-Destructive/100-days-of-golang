package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	//"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println(dir)
	s := filepath.Base("math/")
	log.Println(s)
	m, _ := os.Stat(s)
	log.Println(os.FileInfo(m).Name())
	//log.Println(filepath.IsAbs(s))
	//log.Println(filepath.Dir(s))

	k := "bit.go"
	//os.Chdir("math/")
	p, _ := filepath.Abs(k)
	log.Println(strings.TrimSuffix(k, path.Ext(k)))
	log.Println(path.Ext(k))

	if t, err := os.Stat(p); os.IsNotExist(err) {
		log.Println("No, " + p + " does not exists")
	} else {
		log.Println(t.IsDir())
		log.Println("Yes, " + p + " exists")
	}

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	iterate(currentDirectory)

	// Read Text from a file
	//txt, _ := os.ReadFile(p)
	//log.Println(string(txt))
	var files []string

	root := "../"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		dir_name := filepath.Base(root)
		if info.IsDir() == true && info.Name() != dir_name {
			return filepath.SkipDir
		} else {
			files = append(files, path)
			return nil
		}
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		log.Println(file)
	}
}

func iterate(path string) {
	count := 0
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		count += 1
		var f string
		if info.IsDir() {
			f = "Directory"
		} else {
			f = "File"
		}
		log.Printf("%s Name: %s\n", f, info.Name())
		return nil
	})
	log.Println(count)
}
