package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/zxjsdp/specimen-go/config"
)

// IsFileExists returns true if file exists.
func IsFileExists(fileName string) bool {
	cleanName := strings.TrimSpace(fileName)
	if len(cleanName) == 0 {
		return false
	}
	_, err := os.Stat(cleanName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// CheckFileExists checks whether file exists.
func CheckFileExists(fileName, description, usage string) {
	cleanName := strings.TrimSpace(fileName)
	if len(cleanName) == 0 {
		log.Fatal(fmt.Sprintf("ERROR! No file name provided for [ %s ].%s",
			description, usage))
	}
	_, err := os.Stat(cleanName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(fmt.Sprintf("ERROR! File does not exist: [ %s ] for [ %s ].%s",
				description, fileName, usage))
		}
	}
}

// Get current working directory
func GetCurrentWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取路径失败！%s\n", err)
		return ""
	}
	return dir
}

// ReadLines reads a file and return all the lines.
func ReadLines(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// ReadContent reads a file and return file content as a string.
func ReadContent(filename string) string {
	lines := ReadLines(filename)

	return strings.Join(lines, "\n")
}

func GetDemoHTMLFilePath() string {
	demoHTMLFilePath := path.Join(GetCurrentWorkingDir(), config.DemoHTMLFileName)
	if !IsFileExists(demoHTMLFilePath) {
		WriteContent(demoHTMLFilePath, config.DemoHTMLContent)
	}

	return demoHTMLFilePath
}

// WriteContent write string content to a file.
func WriteContent(filename, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
