package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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

func GetCurrentWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取路径失败！%s\n", err)
		return ""
	}
	return dir
}
