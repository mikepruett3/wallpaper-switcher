package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/reujab/wallpaper"
)

func main() {
	// Getting root variable from Command Line Flag
	// https://gobyexample.com/command-line-flags
	root := flag.String("path", "", "Path to Wallpapers, can be nested folders.")
	flag.Parse()

	// Define string Split variable
	var files []string

	// Find all files from the provided directory
	// https://zetcode.com/golang/find-file/
	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return nil
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Random Number Generator
	// https://golang.cafe/blog/golang-random-number-generator.html
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := len(files)
	randomIndex := rand.Intn(max-min+1) + min

	// Pick new wallpaper from randomIndex
	pick := files[randomIndex]

	// List all the wallpaper files found
	//for _, file := range files {
	//	fmt.Println(file)
	//}

	background, err := wallpaper.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current wallpaper:", background)
	fmt.Println("New wallpaper:", pick)

	// Change Current wallpaper to New wallpaper
	// https://github.com/reujab/wallpaper
	wallpaper.SetFromFile(pick)
}
