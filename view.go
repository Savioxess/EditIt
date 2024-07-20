package main

import (
	"fmt"
	"os"

	"seehuhn.de/go/ncurses"
)

type Command int

const (
	EXIT Command = iota
	SAVE
	SAVEANDEXIT
	NONE
)

func setViewToNormal(mainWindow *ncurses.Window, commandWindow *ncurses.Window, cursor *Cursor) {
	commandWindow.Erase()
	commandWindow.AddStr("---- NORMAL ----")
	commandWindow.Refresh()
	mainWindow.Move(cursor.Y, cursor.X)
}

func handleInteraction(mainWindow *ncurses.Window, commandWindow *ncurses.Window, cursor *Cursor, content *[][]byte, file *os.File) {
	for {
		inputChar := mainWindow.GetCh()
		//fmt.Println(inputChar)

		switch inputChar {
		case ':':
			command := commandInput(commandWindow)

			if command == EXIT {
				return
			}

			if command == SAVEANDEXIT {
				saveContentToFile(file, content)
				return
			}

			if command == SAVE {
				saveContentToFile(file, content)
			}

			setViewToNormal(mainWindow, commandWindow, cursor)
		case 'i':
			insertMode(mainWindow, commandWindow, cursor, content)
			setViewToNormal(mainWindow, commandWindow, cursor)
		case 'a':
			if (*content)[cursor.Y][cursor.X] != 10 {
				cursor.X += 1
				insertMode(mainWindow, commandWindow, cursor, content)
				setViewToNormal(mainWindow, commandWindow, cursor)
			}
		case 'h':
			if cursor.X <= 0 {
				continue
			}

			cursor.X -= 1
			mainWindow.Move(cursor.Y, cursor.X)
		case 'l':
			if cursor.X > len((*content)[cursor.Y]) && (*content)[cursor.Y][len((*content)[cursor.Y])-1] != 10 {
				continue
			}

			if cursor.X >= len((*content)[cursor.Y])-1 {
				continue
			}

			cursor.X += 1
			mainWindow.Move(cursor.Y, cursor.X)
		case 'j':
			if cursor.Y >= len(*content)-1 {
				continue
			}

			cursor.Y += 1
			cursor.X = len((*content)[cursor.Y]) - 1
			mainWindow.Move(cursor.Y, cursor.X)
		case 'k':
			if cursor.Y <= 0 {
				continue
			}

			cursor.Y -= 1
			cursor.X = len((*content)[cursor.Y]) - 1
			mainWindow.Move(cursor.Y, cursor.X)
		}
	}
}

func insertMode(mainWindow *ncurses.Window, commandWindow *ncurses.Window, cursor *Cursor, content *[][]byte) {
	commandWindow.Erase()
	commandWindow.AddStr(fmt.Sprintf("---- INSERT ---- col: %d, row: %d", cursor.X, cursor.Y))
	commandWindow.Refresh()
	mainWindow.Move(cursor.Y, cursor.X)

	for {
		inputChar := mainWindow.GetCh()

		if inputChar == 27 {
			return
		}

		if inputChar == ncurses.KeyBackspace {
			//fmt.Println(content)
			deleteCharacter(content, cursor)
		}

		if inputChar == 10 {
			insertNewLine(content, cursor)
		} else if inputChar >= 32 && inputChar <= 126 {
			insertCharacter(content, cursor, byte(inputChar))
		}

		if inputChar == 9 {
			insertCharacter(content, cursor, byte(' '))
			insertCharacter(content, cursor, byte(' '))
		}

		setContentOnWindow(mainWindow, content)
		mainWindow.Move(cursor.Y, cursor.X)
	}
}

func commandInput(commandWindow *ncurses.Window) Command {
	commandWindow.Erase()
	commandWindow.AddStr(":")
	commandWindow.Refresh()
	inputCommand := commandWindow.Readline(10)

	if inputCommand == "q" {
		return EXIT
	}

	if inputCommand == "w" {
		return SAVE
	}

	if inputCommand == "wq" {
		return SAVEANDEXIT
	}

	return NONE
}

func setContentOnWindow(mainWindow *ncurses.Window, content *[][]byte) {
	mainWindow.Erase()
	for _, val := range *content {
		mainWindow.AddStr(string(val))
	}

	mainWindow.Refresh()
}
