package main

import (
	"log"
	"os"

	"seehuhn.de/go/ncurses"
)

func main() {
	mainWindow := ncurses.Init()
	height, width := mainWindow.GetMaxYX()
	commandWindow := ncurses.NewWin(1, width, height-1, 0)
	cursor := Cursor{X: 0, Y: 0}
	content := &[][]byte{{}}
	defer ncurses.EndWin()

	var commandColorPair ncurses.ColorPair = 1

	if err := commandColorPair.Init(ncurses.ColorBlue, ncurses.ColorBlack); err != nil {
		log.Println(err)
		return
	}

	fileName := os.Args[1]

	content, file, err := loadFileContent(fileName)

	defer file.Close()

	if err != nil {
		log.Println(err)
		return
	}

	setContentOnWindow(mainWindow, content)
	commandWindow.SetBackground(" ", ncurses.AttrBold, commandColorPair)
	setViewToNormal(mainWindow, commandWindow, &cursor)
	handleInteraction(mainWindow, commandWindow, &cursor, content, file)
}
