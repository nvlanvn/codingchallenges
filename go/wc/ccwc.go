package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lang := os.Getenv("LANG")
	option := ""
	if len(os.Args) > 1 {
		option = os.Args[1]
	}
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	scanner := bufio.NewScanner(file)
	line_number := 0
	total_word := 0
	total_char := 0
	if option == "-m" && lang != "" {
		scanner.Split(bufio.ScanRunes)
	}
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		total_word += len(line)
		line_number += 1
		total_char += 1
	}
	switch option {
	case "-c":
		fmt.Println(fileInfo.Size(), fileInfo.Name())
	case "-l":
		fmt.Println(line_number, fileInfo.Name())
	case "-w":
		fmt.Println(total_word, fileInfo.Name())
	case "-m":
		fmt.Println(total_char, fileInfo.Name())
	default:
		fmt.Println(line_number, total_word, fileInfo.Size(), fileInfo.Name())
	}
}
