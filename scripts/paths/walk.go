package main

import (
	"io/fs"
	"log"
	"path/filepath"
)

func main() {

	dir_path := "."
	walkfunc := iterate_files_walk
	walkdirfunc := iterate_files_walkdir
	err := filepath.WalkDir(dir_path, walkdirfunc)
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(dir_path, walkfunc)
	if err != nil {
		panic(err)
	}

}

func iterate_files_walkdir(path string, info fs.DirEntry, err error) error {
	var f string
	if info.IsDir() {
		f = "Directory"
	} else {
		f = "File"
	}
	log.Println(f + " : " + info.Name())

	if err != nil {
		return err
	} else {
		return nil
	}
}

func iterate_files_walk(path string, info fs.FileInfo, err error) error {
	var f string
	if info.IsDir() {
		f = "Directory"
	} else {
		f = "File"
	}
	log.Println(f + " : " + info.Name())
	if err != nil {
		return err
	} else {
		return nil
	}
}
