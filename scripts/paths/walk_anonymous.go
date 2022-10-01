package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var files []string
	var dirs []string

	root := "math/"
	// walk funciton with paramter as os.FileInfo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		dir_name := filepath.Base(root)

		// skip the path if it is a directory and also the nested directory
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
	log.Println("Files : ")
	for _, file := range files {
		log.Println(file)
	}

	err = filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		dir_name := filepath.Base(root)

		// add the path if the path is a directory
		if info.IsDir() == true && info.Name() != dir_name {
			dirs = append(dirs, path)
		} else {
			// add the path to the files slice
			files = append(files, path)
		}
		if err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		panic(err)
	}
	log.Println("Files : ")
	for _, file := range files {
		log.Println(file)
	}
	log.Println("Dirs: ")
	for _, dir := range dirs {
		log.Println(dir)
	}
}
