package main

import (
	"fmt"
	"github.com/BRUHItsABunny/bunterm"
	"os"
	"time"
)

func main() {
	var err error
	bTerm := bunterm.NewBunTerminal(os.Stdout)

	fmt.Println(bTerm.GetTerminalSize())

	time.Sleep(time.Second * time.Duration(5))

	for i := 0; i < 30; i++ {
		// Clear first
		err = bTerm.ClearTerminal()
		if err != nil {
			panic(err)
		}
		// Then move cursor
		err = bTerm.MoveCursor(0, 0)
		if err != nil {
			panic(err)
		}
		fmt.Println(time.Now().String() + "||") // pipe for error checking, if it doesn't work properly we expect more than just the pipe at the end
		time.Sleep(time.Second)
	}
}
