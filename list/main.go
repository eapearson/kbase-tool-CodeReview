package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	var dir string

	if len(os.Args) > 1 {
		fmt.Printf("hmm %s\n", os.Args[1])
		dir = os.Args[1]
	} else {
		dir = "."
	}

	ScourDir(dir)
}

func DirList(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(strings.ToUpper(file.Name()))
			DirList(file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}
}

func Scour(dir string, parentPath string) {
	fmt.Printf("**scouring** %s, %s\n", dir, parentPath)
	dirPath := ""
	if parentPath == "" {
		dirPath = dir
	} else {
		dirPath = parentPath + "/" + dir
	}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("dir : " + strings.ToUpper(file.Name()))

			Scour(file.Name(), dirPath)
		} else {
			fmt.Println("file: " + file.Name())
		}
	}
}

func ScourDirs(dir string) {
	parentPath := ""
	Scour(dir, parentPath)
}

func Scourer(filePath string, info os.FileInfo, err error, stats map[string]int) error {
	if err != nil {
		return err
	}
	if info.Mode().IsRegular() {
		// fmt.Printf("%s\n", path)
		ext := path.Ext(filePath)
		stats[ext]++
	}
	return nil
}

func ScourDir(dir string) {
	stats := make(map[string]int)
	walker := func(path string, info os.FileInfo, err error) error {
		return Scourer(path, info, err, stats)
	}
	filepath.Walk(dir, walker)
	for k, v := range stats {
		fmt.Printf("%s = %d\n", k, v)
	}
}
