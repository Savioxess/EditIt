package main

import (
	"testing"
)

func TestInsertCharacter(t *testing.T) {
	content := [][]byte{{'a', 'c', 'b', '\n'}, {'x', 'h'}}
	want := [][]byte{{'a', 'x', 'c', 'b', '\n'}, {'x', 'h'}}
	cursor := Cursor{Y: 0, X: 1}
	inputChar := 'x'

	insertCharacter(&content, &cursor, byte(inputChar))
	var matched bool = true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test insertCharacter: Output: %v, Expected: %v`, content, want)
	}
}

func TestInsertCharacterEmptyContent(t *testing.T) {
	content := [][]byte{{}}
	want := [][]byte{{'a'}}
	cursor := Cursor{Y: 0, X: 1}
	inputChar := 'a'

	insertCharacter(&content, &cursor, byte(inputChar))
	var matched bool = true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test insertCharacter(Empty Content): Output: %v, Expected: %v`, content, want)
	}
}

func TestInsertCharacter2CharsInContent(t *testing.T) {
	content := [][]byte{{'t', 'h'}}
	want := [][]byte{{'t', 'h', 'i'}}
	cursor := Cursor{Y: 0, X: 2}
	inputChar := 'i'

	insertCharacter(&content, &cursor, byte(inputChar))
	var matched bool = true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test insertCharacter(2Chars Content): Output: %v, Expected: %v`, content, want)
	}
}
func TestInsertCharAtNewLine(t *testing.T) {
	content := [][]byte{{'a', '\n'}, {'c', 'b', '\n'}, {'x', 'h'}}
	want := [][]byte{{'a', 'x', '\n'}, {'c', 'b', '\n'}, {'x', 'h'}}
	cursor := Cursor{Y: 0, X: 1}
	inputChar := 'x'

	insertCharacter(&content, &cursor, byte(inputChar))
	var matched bool = true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test insertCharacter: Output: %v, Expected: %v`, content, want)
	}
}

func TestInsertCharBeforeNewLine(t *testing.T) {
	content := [][]byte{{'m', 'o', 'r', 'e', ' ', 'c', 'o', 'n', 't', 'e', 'n', 't'}}
	cursor := Cursor{Y: 0, X: 0}
	insertNewLine(&content, &cursor)
	want := [][]byte{{'a', '\n'}, {'m', 'o', 'r', 'e', ' ', 'c', 'o', 'n', 't', 'e', 'n', 't'}}
	cursor = Cursor{Y: 0, X: 0}
	inputChar := 'a'

	insertCharacter(&content, &cursor, byte(inputChar))
	var matched bool = true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test insertCharacter: Output: %v, Expected: %v`, content, want)
	}
}
func TestInsertNewLine(t *testing.T) {
	content := [][]byte{{'a', 'c', 'b', '\n'}, {'x', 'h'}}
	want := [][]byte{{'a', '\n'}, {'c', 'b', '\n'}, {'x', 'h'}}
	cursor := Cursor{Y: 0, X: 1}

	insertNewLine(&content, &cursor)
	var matched bool = true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test insertNewLine: Output: %v, Expected: %v`, content, want)
	}
}

func TestInsertNewLineAtMiddle(t *testing.T) {
	content := [][]byte{{'t', 'h', 'i', 's', ' ', 'a', 'n', 'd', ' ', 't', 'h', 'i', 's'}}
	want := [][]byte{{'t', 'h', 'i', 's', ' ', 'a', 'n', 'd', ' ', '\n'}, {'t', 'h', 'i', 's'}}
	cursor := Cursor{Y: 0, X: 9}

	insertNewLine(&content, &cursor)
	var matched bool = true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test insertNewLine: Output: %v, Expected: %v`, content, want)
	}
}

func TestLoadFileContent(t *testing.T) {
	filename := "test.txt"
	want := [][]byte{[]byte("more content")}

	content, file, err := loadFileContent(filename)
	defer file.Close()
	matched := true

	for i, _ := range *content {
		if string((*content)[i]) != string(want[i]) {
			matched = false
			break
		}
	}

	if err != nil || !matched {
		t.Fatalf("Test loadFileContent: Error: %v, Output: %s, Expected: %s", err, *content, want)
	}
}

func TestDeleteCharacterMiddle(t *testing.T) {
	content := [][]byte{{'a', 'c', 'b', '\n'}, {'x', 'h'}}
	cursor := Cursor{Y: 0, X: 2}
	want := [][]byte{{'a', 'b', '\n'}, {'x', 'h'}}
	deleteCharacter(&content, &cursor)
	matched := true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test DeleteCharacterMiddle: Output: %v, Expected: %v`, content, want)
	}
}

func TestDeleteCharacterStart(t *testing.T) {
	content := [][]byte{{'a', 'c', 'b', '\n'}, {'x', 'h', '\n'}, {'p', 'q', 'r'}}
	cursor := Cursor{Y: 1, X: 0}
	want := [][]byte{{'a', 'c', 'b', 'x', 'h', '\n'}, {'p', 'q', 'r'}}
	deleteCharacter(&content, &cursor)
	matched := true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test DeleteCharacterMiddle: Output: %v, Expected: %v`, content, want)
	}
}

func TestDeleteCharacterInLineWithNoNewline(t *testing.T) {
	content := [][]byte{{'a', 'c', 'b', '\n'}, {'x', 'h'}, {'p', 'q', 'r'}}
	cursor := Cursor{Y: 1, X: 2}
	want := [][]byte{{'a', 'c', 'b', '\n'}, {'x'}, {'p', 'q', 'r'}}
	deleteCharacter(&content, &cursor)
	matched := true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test DeleteCharacterInLineWithNoNewline: Output: %v, Expected: %v`, content, want)
	}
}

func TestDeleteCharacterAtBeginning(t *testing.T) {
	content := [][]byte{{'a', 'c', 'b'}}
	cursor := Cursor{Y: 0, X: 0}
	want := [][]byte{{'a', 'c', 'b'}}
	deleteCharacter(&content, &cursor)
	matched := true

	for i, _ := range content {
		if !matched {
			break
		}

		for j, _ := range content[i] {
			if content[i][j] != want[i][j] {
				matched = false
			}
		}
	}

	if !matched {
		t.Fatalf(`Test DeleteCharacterAtBeginning: Output: %v, Expected: %v`, content, want)
	}
}
