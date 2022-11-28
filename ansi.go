package bunterm

import "fmt"

// SRC: // https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html

func ansiClearTerminal() string {
	return "\033[K"
}

func ansiClearTerminalAlternate() string {
	return "\033[2J"
}

func ansiClearLine() string {
	return "\033[2K" // may want to add \r
}

// TODO: which one when?

func ansiMoveCursor(x, y int16) string {
	return fmt.Sprintf("\033[<%d>;<%d>H", y, x)
}

func ansiMoveCursorAlternate(x, y int16) string {
	return fmt.Sprintf("\033[<%d>;<%d>f", y, x)
}
