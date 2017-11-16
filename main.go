package main

import (
	"flag"

	"fmt"
	"log"
	"strings"

	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/specimen"
	"github.com/zxjsdp/specimen-go/utils"
)

// 解析命令行，并输出 流水号文件名、物种记录及鉴定文件名、输出文件名
func parseArgument() (string, string, string) {
	snDataPtr := flag.String("s", "", "流水号文件.xlsx")
	offlineDataPtr := flag.String("d", "", "物种记录及鉴定文件.xlsx")
	outputDataPtr := flag.String("o", "", "输出文件.xlsx")

	flag.Parse()

	utils.CheckFileExists(*snDataPtr, "-s", config.USAGE)
	utils.CheckFileExists(*offlineDataPtr, "-d", config.USAGE)
	if len(strings.TrimSpace(*outputDataPtr)) == 0 {
		log.Fatal(fmt.Sprintf("ERROR! 参数不能为空：[ -o ]. %s", config.USAGE))
	}

	return *snDataPtr, *offlineDataPtr, *outputDataPtr
}

func main() {
	snDataFile, offlineDataFile, outputDataFile := parseArgument()
	specimen.RunSpecimenInfo(snDataFile, offlineDataFile, outputDataFile, true)
}
