package bunterm

import "fmt"

// SRC: https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
// SRC: https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
const esc = "\u001b"

func ansiClearTerminal() string {
	return fmt.Sprintf("%s[2J", esc)
}

func ansiClearLine() string {
	return fmt.Sprintf("%s[2K", esc) // may want to add \r
}

func ansiMoveCursor(x, y int16) string {
	return fmt.Sprintf("%s[%d;%dH", esc, y, x)
}
