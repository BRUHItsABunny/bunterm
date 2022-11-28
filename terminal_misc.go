//go:build !windows

package bunterm

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func clearTerminal(file *os.File) error {
	_, err := file.WriteString(ansiClearTerminal())
	return err
}

func clearLine(file *os.File) error {
	_, err := file.WriteString(ansiClearLine() + "\r")
	return err
}

func moveCursor(file *os.File, x, y int16) error {
	_, err := file.WriteString(ansiMoveCursor(x, y))
	return err
}

func getTerminalSize(file *os.File) (int16, int16, error) {
	x, y, err := term.GetSize(int(file.Fd()))
	if err != nil {
		return 0, 0, fmt.Errorf("term.GetSize: %w", err)
	}
	return int16(x), int16(y), nil
}
