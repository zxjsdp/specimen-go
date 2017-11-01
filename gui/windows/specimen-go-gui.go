// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
)

import (
	"io/ioutil"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/utils"
	"github.com/zxjsdp/specimen-go/web"
)

const (
	Title    = "植物标本录入软件"
	Width    = 800
	Height   = 700
	IconPath = "icon.ico"
)

type MyMainWindow struct {
	*walk.MainWindow

	titleLabel  *walk.Label
	combo1      *walk.ComboBox
	openButton1 *walk.PushButton
	combo2      *walk.ComboBox
	openButton2 *walk.PushButton
	combo3      *walk.ComboBox
	openButton3 *walk.PushButton
	statusBar   *walk.Label
	startButton *walk.PushButton
	progressBar *walk.ProgressBar
	logView     *LogView

	previousFilePath string

	queryFile    string
	dataFile     string
	outputFile   string
	selectedFile string
}

func (mw *MyMainWindow) lb_ItemSelected_Combo1() {
	name := mw.combo1.Text()
	mw.queryFile = name
	mw.statusBar.SetText(fmt.Sprintf("已选择文件：%s", name))
}

func (mw *MyMainWindow) lb_ItemSelected_Combo2() {
	name := mw.combo2.Text()
	mw.dataFile = name
	mw.statusBar.SetText(fmt.Sprintf("已选择文件：%s", name))
}

func (mw *MyMainWindow) lb_ItemSelected_Combo3() {
	name := mw.combo3.Text()
	mw.outputFile = name
	mw.statusBar.SetText(fmt.Sprintf("已选择文件：%s", name))
}

func main() {
	RunMainWindow()
}

func getXlsxFiles() []string {
	files, err := ioutil.ReadDir("./")
	xlsxFiles := []string{}
	if err != nil {
		log.Fatal(err)
		return xlsxFiles
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), "xlsx") {
			xlsxFiles = append(xlsxFiles, f.Name())
		}
	}

	return xlsxFiles
}

func RunMainWindow() {
	mw := &MyMainWindow{}

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    Title,
		MinSize:  Size{Width: Width, Height: Height},
		Icon:     IconPath,
		Layout:   VBox{},

		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Separator{},
					Action{
						Text:        "退出",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "帮助",
						OnTriggered: mw.helpAction_Triggered,
					},
					Separator{},
					Action{
						Text:        "关于",
						OnTriggered: mw.aboutAction_Triggered,
					},
				},
			},
		},

		Children: []Widget{
			Label{
				AssignTo: &mw.titleLabel,
				Text:     Title,
				Font:     Font{Family: "Microsoft Yahei", PointSize: 15},
			},
			Composite{
				Layout: Grid{Columns: 3},
				Children: []Widget{
					Label{
						Text: "Query 文件：",
					},
					ComboBox{
						Editable: true,
						AssignTo: &mw.combo1,
						Model:    getXlsxFiles(),
						OnCurrentIndexChanged: mw.lb_ItemSelected_Combo1,
						ToolTipText:           "选取或者填写 query 文件名称",
					},
					PushButton{
						Text:     "...",
						AssignTo: &mw.openButton1,
						OnClicked: func() {
							mw.openButton_Triggered()
							mw.combo1.SetText(mw.selectedFile)
						},
					},

					Label{
						Text: "Data 文件：",
					},
					ComboBox{
						Editable: true,
						AssignTo: &mw.combo2,
						Model:    getXlsxFiles(),
						OnCurrentIndexChanged: mw.lb_ItemSelected_Combo2,
						ToolTipText:           "选取或者填写 data 文件名称",
					},
					PushButton{
						Text:     "...",
						AssignTo: &mw.openButton2,
						OnClicked: func() {
							mw.openButton_Triggered()
							mw.combo2.SetText(mw.selectedFile)
						},
					},

					Label{
						Text: "输出文件：",
					},
					ComboBox{
						Editable: true,
						AssignTo: &mw.combo3,
						Model:    getXlsxFiles(),
						OnCurrentIndexChanged: mw.lb_ItemSelected_Combo3,
						ToolTipText:           "选取或者填写输出文件名称",
					},
					PushButton{
						Text:     "...",
						AssignTo: &mw.openButton3,
						OnClicked: func() {
							mw.openButton_Triggered()
							mw.combo3.SetText(mw.selectedFile)
						},
					},
				},
			},

			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						AssignTo:    &mw.statusBar,
						Text:        "",
						ToolTipText: "提示信息",
					},
					HSpacer{},
					PushButton{
						Text:        "开始处理",
						AssignTo:    &mw.startButton,
						ToolTipText: "开始进行植物标本数据处理",
						OnClicked: func() {
							queryFile := mw.combo1.Text()
							dataFile := mw.combo2.Text()
							outputFile := mw.combo3.Text()

							if len(queryFile) == 0 || len(dataFile) == 0 || len(outputFile) == 0 {
								mw.statusBar.SetText("参数无效！")
								return
							}
							go mw.RunSpecimenInfoGoroutine(queryFile, dataFile, outputFile)
						},
					},
				},
			},
			ProgressBar{
				AssignTo:    &mw.progressBar,
				MinValue:    0,
				MaxValue:    100,
				Font:        Font{PointSize: 6},
				ToolTipText: "处理进度",
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	log.SetFlags(0)
	log.SetOutput(new(utils.LogWriter))
	lv, err := NewLogView(mw)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(lv)

	mw.Run()
}

func (mw *MyMainWindow) RunSpecimenInfoGoroutine(queryFile, dataFile, outputFile string) {
	mw.startButton.SetEnabled(false)
	mw.startButton.SetText("处理中...")
	defer mw.startButton.SetEnabled(true)
	defer mw.startButton.SetText("开始处理")

	log.Printf("开始读取 entry 数据文件 ...\n")
	mw.progressBar.SetValue(1)
	entryDataMatrix := files.GetDataMatrix(dataFile)
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDataSlice)
	log.Printf("读取 entry 数据结束！\n")
	mw.progressBar.SetValue(10)

	log.Printf("开始读取 marker 数据文件 ...\n")
	mw.progressBar.SetValue(20)
	markerDataMatrix := files.GetDataMatrix(queryFile)
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)
	log.Printf("读取 marker 数据结束！\n")
	mw.progressBar.SetValue(30)

	log.Printf("开始提取网络信息，这可能会花费一些时间，请耐心等待 ...\n")
	mw.progressBar.SetValue(40)
	speciesNames := converters.ExtractSpeciesNames(entryDataSlice)
	webInfoMap := web.GenerateWebInfoMap(speciesNames)
	log.Printf("提取网络信息结束！\n")
	mw.progressBar.SetValue(60)

	log.Printf("开始整合本地数据及网络信息 ...\n")
	mw.progressBar.SetValue(70)
	resultDataSlice := make([]entities.ResultData, len(markerDataSlice))
	for i, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, entryDataMap, webInfoMap)
		resultDataSlice[i] = resultData
	}
	log.Printf("整合本地数据及网络信息结束！\n")
	mw.progressBar.SetValue(80)

	log.Printf("开始将结果信息写入 xlsx 输出文件...\n")
	mw.progressBar.SetValue(90)
	files.SaveDataMatrix(outputFile, resultDataSlice)

	log.Printf("任务完成！\n")
	mw.progressBar.SetValue(100)
}

func (mw *MyMainWindow) openButton_Triggered() {
	if filePath, err := mw.openFile(); err != nil {
		log.Print(err)
		return
	} else {
		mw.selectedFile = filePath
	}
}

func (mw *MyMainWindow) openFile() (string, error) {
	dlg := new(walk.FileDialog)

	dlg.FilePath = mw.previousFilePath
	dlg.Filter = "Xlsx Files (*.xlsx)"
	dlg.Title = "Select an xlsx file"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.statusBar.SetText("打开文件失败！" + err.Error())
		return "", err
	} else if !ok {
		mw.statusBar.SetText("打开文件失败！")
		return "", err
	}

	mw.previousFilePath = dlg.FilePath

	mw.statusBar.SetText(fmt.Sprintf("已选择文件：%s", dlg.FilePath))

	return dlg.FilePath, nil
}

func (mw *MyMainWindow) helpAction_Triggered() {
	walk.MsgBox(mw, "帮助", "使用帮助", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) aboutAction_Triggered() {
	walk.MsgBox(mw, "关于", "Specimen GUI v1.2.0 by zxjsdp", walk.MsgBoxIconInformation)
}
