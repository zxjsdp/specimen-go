// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
	"strings"
)

type SpecimenArgument struct {
	QueryFile string
	DataFile string
	OutputFile string

	SelectedFile string
}

type MyMainWindow struct {
	*walk.MainWindow

	titleLabel *walk.Label
	combo1 *walk.ComboBox
	openButton1 *walk.PushButton
	combo2 *walk.ComboBox
	openButton2 *walk.PushButton
	combo3 *walk.ComboBox
	openButton3 *walk.PushButton
	statusBar *walk.Label
	okButton *walk.PushButton
	cancelButton *walk.PushButton
	resultText *walk.TextEdit

	previousFilePath string
}

func (mw *MyMainWindow) lb_ItemSelected_Combo1() {
	name := mw.combo1.Text()
	mw.statusBar.SetText(fmt.Sprintf("已选择文件：%s", name))
}

func (mw *MyMainWindow) lb_ItemSelected_Combo2() {
	name := mw.combo2.Text()
	mw.statusBar.SetText(fmt.Sprintf("已选择文件：%s", name))
}

func (mw *MyMainWindow) lb_ItemSelected_Combo3() {
	name := mw.combo3.Text()
	mw.statusBar.SetText(fmt.Sprintf("已选择文件：%s", name))
}

func main() {
	var argument = new(SpecimenArgument)
	RunMainWindow(argument)
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

func RunMainWindow(argument *SpecimenArgument) {
	mw := &MyMainWindow{}
	var openAction *walk.Action

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Specimen Go GUI",
		MinSize:  Size{700, 700},
		Icon:"icon.ico",
		Layout:   VBox{},

		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						//OnTriggered: mw.openAction_Triggered,
					},
					Separator{},
					Action{
						Text:        "Exit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "About",
						OnTriggered: mw.aboutAction_Triggered,
					},
				},
			},
		},

		Children: []Widget{
			Label{
				AssignTo: &mw.titleLabel,
				Text: "植物标本数据处理软件",
				Font:Font{Family:"Microsoft Yahei",PointSize:15},
			},
			Composite{
				Layout:Grid{Columns:3},
				Children:[]Widget{
					Label{
						Text: "Query 文件：",
					},
					ComboBox{
						Editable: true,
						AssignTo:&mw.combo1,
						Model: getXlsxFiles(),
						OnCurrentIndexChanged:mw.lb_ItemSelected_Combo1,
					},
					PushButton{
						Text:"...",
						AssignTo:&mw.openButton1,
						OnClicked:func() {
							mw.openButton_Triggered(argument)
							mw.combo1.SetText(argument.SelectedFile)
						},
					},

					Label{
						Text: "Data 文件：",
					},
					ComboBox{
						Editable: true,
						AssignTo:&mw.combo2,
						Model: getXlsxFiles(),
						OnCurrentIndexChanged:mw.lb_ItemSelected_Combo2,
					},
					PushButton{
						Text:"...",
						AssignTo:&mw.openButton2,
						OnClicked:func() {
							mw.openButton_Triggered(argument)
							mw.combo2.SetText(argument.SelectedFile)
						},
					},

					Label{
						Text: "输出文件：",
					},
					ComboBox{
						Editable: true,
						AssignTo:&mw.combo3,
						Model: getXlsxFiles(),
						OnCurrentIndexChanged:mw.lb_ItemSelected_Combo3,
					},
					PushButton{
						Text:"...",
						AssignTo:&mw.openButton3,
						OnClicked:func() {
							mw.openButton_Triggered(argument)
							mw.combo3.SetText(argument.SelectedFile)
						},
					},
				},
			},
			Label{
				AssignTo: &mw.statusBar,
				Text: "",
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:     "Clear",
						AssignTo:&mw.cancelButton,
						OnClicked: func() {
							mw.resultText.SetText(fmt.Sprintf("query: %s, data: %s, output: %s\n",
								argument.QueryFile, argument.DataFile, argument.OutputFile))
						},
					},
					PushButton{
						Text:      "Start",
						AssignTo:&mw.okButton,
						OnClicked: func() {
							mw.resultText.SetText(argument.QueryFile + argument.DataFile + argument.OutputFile)
						},
					},
				},
			},
			TextEdit{
				AssignTo: &mw.resultText,
				ReadOnly: true,
				Text:     fmt.Sprintf("%+v", "The quick fox jumped over the lazy dog!"),
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}

func (mw *MyMainWindow) openButton_Triggered(argument *SpecimenArgument) {
	if filePath, err := mw.openFile(); err != nil {
		log.Print(err)
		return
	} else {
		argument.SelectedFile = filePath
	}
}

func (mw *MyMainWindow) openFile() (string, error) {
	dlg := new(walk.FileDialog)

	dlg.FilePath = mw.previousFilePath
	dlg.Filter = "Xlsx Files (*.xlsx)"
	dlg.Title = "Select an xlsx file"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.statusBar.SetText("打开文件失败！"+ err.Error())
		return "", err
	} else if !ok {
		mw.statusBar.SetText("打开文件失败！")
		return "", err
	}

	mw.previousFilePath = dlg.FilePath

	mw.statusBar.SetText(fmt.Sprintf("已选择文件：%s", dlg.FilePath))

	return dlg.FilePath, nil
}

func (mw *MyMainWindow) aboutAction_Triggered() {
	walk.MsgBox(mw, "About", "Specimen GUI", walk.MsgBoxIconInformation)
}
