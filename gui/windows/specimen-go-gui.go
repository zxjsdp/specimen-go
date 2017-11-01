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
}

type MyMainWindow struct {
	*walk.MainWindow
	combo1 *walk.ComboBox
	combo2 *walk.ComboBox
	combo3 *walk.ComboBox
	okButton *walk.PushButton
	cancelButton *walk.PushButton
	resultText *walk.TextEdit
	statusBar *walk.Label
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

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Specimen Go GUI",
		MinSize:  Size{500, 500},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout:Grid{Columns:2},
				Children:[]Widget{
					Label{
						Text: "Query file:",
					},
					ComboBox{
						Editable: true,
						AssignTo:&mw.combo1,
						Model: getXlsxFiles(),
						OnCurrentIndexChanged:mw.lb_ItemSelected_Combo1,
					},
					Label{
						Text: "Data file:",
					},
					ComboBox{
						Editable: true,
						AssignTo:&mw.combo2,
						Model: getXlsxFiles(),
					},
					Label{
						Text: "Output file:",
					},
					ComboBox{
						Editable: true,
						AssignTo:&mw.combo3,
						Model: getXlsxFiles(),
					},
				},
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
			Label{
				AssignTo: &mw.statusBar,
				Text: "",
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
