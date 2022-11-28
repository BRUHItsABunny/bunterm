package bunterm

import (
	"os"
)

var DefaultTerminal = &BunTerminal{os.Stdout}

type BunTerminal struct {
	File *os.File
}

func NewBunTerminal(file *os.File) *BunTerminal {
	return &BunTerminal{File: file}
}

func (t *BunTerminal) ClearTerminal() error {
	return clearTerminal(t.File)
}

func (t *BunTerminal) ClearLine() error {
	return clearLine(t.File)
}

func (t *BunTerminal) MoveCursor(x, y int16) error {
	return moveCursor(t.File, x, y)
}

func (t *BunTerminal) GetTerminalSize() (int16, int16, error) {
	return getTerminalSize(t.File)
}
