package bunterm

import "fmt"

// SRC: // https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
const esc = "\u001b"

// TODO: Maybe ANSI codes don't always work, find edge cases and make workarounds

// ansiClearTerminalAlternate didn't work on Debian (5.16.18-1)
func ansiClearTerminalAlternate() string {
	return fmt.Sprintf("%s[K", esc)
}

func ansiClearTerminal() string {
	return fmt.Sprintf("%s[2J", esc)
}

func ansiClearLine() string {
	return fmt.Sprintf("%s[2K", esc) // may want to add \r
}

func ansiMoveCursor(x, y int16) string {
	return fmt.Sprintf("%s[%d;%dH", esc, y, x)
}

func ansiMoveCursorAlternate(x, y int16) string {
	return fmt.Sprintf("%s[%d;%df", esc, y, x)
}
