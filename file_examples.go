package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func IterateOverADir() {
	cwd := GetCurrentWorkingDirectory()
	filepath.Walk(cwd, func(cwd string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(info.Name())
		return nil
	})
}

func CheckIfExists(filename string) bool {
	fileinfo, err := os.Stat(filename)

	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
		return false
	}

	fmt.Println(fileinfo)
	return true
}

func GetCurrentWorkingDirectory() string {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return cwd
}

func WriteToAFile(filename string, content string) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	bytesWritten, err := writer.WriteString(content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytesWritten)
	writer.Flush()
}

func WordByWordLine(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func WordByWordScan(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func IsADirectory(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
	}

	if info.IsDir() {
		return true
	}

	return false
}

func main() {
	filename := "dsa"
	fmt.Println(CheckIfExists(filename))
}
