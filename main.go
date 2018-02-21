package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// MyMainWindow
type MyMainWindow struct {
	*walk.MainWindow
	searchFolder *walk.LineEdit
	searchText   *walk.LineEdit
	results      *walk.ListBox
	path         string
}

func main() {
	mw := &MyMainWindow{}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Grep Tool",
		Font:     Font{Family: "メイリオ", PointSize: 9},
		MinSize:  Size{800, 900},
		Layout:   VBox{},
		Children: []Widget{
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "検索対象",
					},
					LineEdit{
						Text:     "フォルダを指定してください",
						AssignTo: &mw.searchFolder,
					},
					PushButton{
						Text:      "開く",
						OnClicked: mw.openFolderClicked,
					},
				},
			},
			GroupBox{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "検索文字列",
					},
					LineEdit{
						AssignTo: &mw.searchText,
					},
					PushButton{
						Text:      "検索",
						OnClicked: mw.clicked,
					},
				},
			},
			ListBox{
				AssignTo: &mw.results,
				Row:      5,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

}

func (mw *MyMainWindow) clicked() {
	text := mw.searchText.Text()
	model := []string{}

	// grep
	file, err := os.Open(mw.path)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	hit := false
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
		res := strings.Index(sc.Text(), text)
		if res != -1 {
			model = append(model, fmt.Sprintf("%04d行目  %v", i, sc.Text()))
			hit = true
		}
	}
	if !hit {
		model = append(model, "0件でした")
	}
	mw.results.SetModel(model)
}

func (mw *MyMainWindow) openFolderClicked() {
	dlg := new(walk.FileDialog)
	dlg.FilePath = mw.path
	dlg.Title = "Select File"
	dlg.Filter = "Exe files All files (*.*)|*.*"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		//if ok, err := dlg.ShowBrowseFolder(mw); err != nil {
		return
	} else if !ok {
		return
	}
	mw.path = dlg.FilePath
	mw.searchFolder.SetText(mw.path)
	return
}