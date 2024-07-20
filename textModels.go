package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Cursor struct {
	Y int
	X int
}

func insertCharacter(content *[][]byte, cursor *Cursor, inputChar byte) {
	var row int = (*cursor).Y

	if cursor.X > len((*content)[row])-1 {
		(*content)[row] = append((*content)[row], inputChar)
		cursor.X += 1
		return
	}

	(*content)[row] = append((*content)[row], 0)

	for i := len((*content)[row]) - 2; i >= (*cursor).X; i-- {
		(*content)[row][i+1] = (*content)[row][i]

		if i == (*cursor).X {
			(*content)[row][i] = inputChar
			cursor.X += 1
			return
		}
	}
}

func insertNewLine(content *[][]byte, cursor *Cursor) {
	row := cursor.Y
	col := cursor.X

	*content = append(*content, []byte{})

	for i := len(*content) - 2; i > row; i-- {
		(*content)[i+1] = (*content)[i]
	}

	newRow := make([]byte, len((*content)[row][col:]))
	copy(newRow, (*content)[row][col:])
	(*content)[row] = (*content)[row][:col]
	(*content)[row+1] = newRow

	insertCharacter(content, cursor, byte('\n'))

	cursor.Y++
	cursor.X = 0
}

func loadFileContent(filename string) (*[][]byte, *os.File, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return nil, nil, err
	}

	content := [][]byte{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = append(content, []byte(scanner.Text()+"\n"))
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	if len(content) == 0 {
		return &[][]byte{{}}, file, nil
	}

	content[len(content)-1] = content[len(content)-1][:len(content[len(content)-1])-1]
	return &content, file, nil
}

func deleteCharacter(content *[][]byte, cursor *Cursor) {
	if cursor.Y == 0 && cursor.X == 0 {
		return
	}

	if cursor.Y > 0 && cursor.X == 0 {
		sizeLineBefore := len((*content)[cursor.Y-1]) - 1
		lineBeforeWithoutNewLine := (*content)[cursor.Y-1][:len((*content)[cursor.Y-1])-1]
		tempContent := *content
		changedLineBefore := []byte(string(lineBeforeWithoutNewLine) + string((*content)[cursor.Y]))

		*content = tempContent[:cursor.Y]
		*content = append(*content, tempContent[cursor.Y+1:]...)
		(*content)[cursor.Y-1] = changedLineBefore
		cursor.Y -= 1
		cursor.X = sizeLineBefore
		return
	}

	(*content)[cursor.Y] = []byte(string((*content)[cursor.Y][:cursor.X-1]) + string((*content)[cursor.Y][cursor.X:]))
	cursor.X -= 1
}

func saveContentToFile(file *os.File, content *[][]byte) error {
	_, err := file.Seek(0, io.SeekStart)

	if err != nil {
		log.Fatalf("Failed to seek: %s", err)
	}

	err = file.Truncate(0)

	if err != nil {
		log.Fatalf("Failed to truncate: %s", err)
	}
	writer := bufio.NewWriter(file)

	for _, val := range *content {
		_, err := writer.WriteString(string(val))

		if err != nil {
			return err
		}
	}

	err = writer.Flush()

	if err != nil {
		return err
	}

	return nil
}
