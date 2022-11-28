//go:build windows

package bunterm

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	kFillConsoleOutputCharacter = kernel32.NewProc("FillConsoleOutputCharacterW")
	kFillConsoleOutputAttribute = kernel32.NewProc("FillConsoleOutputAttribute")
)

func clearTerminal(file *os.File) error {
	var written uint32
	info, err := getInfo(file)

	length := uint32(info.Size.X * info.Size.Y)
	_, _, err = kFillConsoleOutputCharacter.Call(
		file.Fd(),
		uintptr(' '),
		uintptr(length),
		*(*uintptr)(unsafe.Pointer(&windows.Coord{})),
		uintptr(unsafe.Pointer(&written)),
	)
	if err != nil && validateError(err) {
		return fmt.Errorf("kFillConsoleOutputCharacter.Call: %w", err)
	}

	_, _, err = kFillConsoleOutputAttribute.Call(
		file.Fd(),
		uintptr(info.Attributes),
		uintptr(length),
		*(*uintptr)(unsafe.Pointer(&windows.Coord{})),
		uintptr(unsafe.Pointer(&written)),
	)

	if err != nil && validateError(err) {
		return fmt.Errorf("kFillConsoleOutputAttribute.Call: %w", err)
	}
	return nil
}

func clearLine(file *os.File) error {
	var written uint32

	info, err := getInfo(file)

	coords := windows.Coord{
		X: info.Window.Left,
		Y: info.CursorPosition.Y,
	}
	length := uint32(info.Size.X)
	_, _, err = kFillConsoleOutputAttribute.Call(file.Fd(), uintptr(info.Attributes), uintptr(length), *(*uintptr)(unsafe.Pointer(&coords)), uintptr(unsafe.Pointer(&written)))
	if err != nil && validateError(err) {
		return fmt.Errorf("kFillConsoleOutputAttribute.Call: %w", err)
	}

	_, _, err = kFillConsoleOutputCharacter.Call(file.Fd(), uintptr(' '), uintptr(length), *(*uintptr)(unsafe.Pointer(&coords)), uintptr(unsafe.Pointer(&written)))
	if err != nil && validateError(err) {
		return fmt.Errorf("kFillConsoleOutputCharacter.Call: %w", err)
	}
	return nil
}

func moveCursor(file *os.File, x, y int16) error {
	err := windows.SetConsoleCursorPosition(windows.Handle(file.Fd()), windows.Coord{
		X: x,
		Y: y,
	})
	if err != nil {
		return fmt.Errorf("windows.SetConsoleCursorPosition: %w", err)
	}
	return nil
}

func getInfo(file *os.File) (*windows.ConsoleScreenBufferInfo, error) {
	var info windows.ConsoleScreenBufferInfo
	err := windows.GetConsoleScreenBufferInfo(windows.Handle(file.Fd()), &info)
	if err != nil {
		return nil, fmt.Errorf("windows.GetConsoleScreenBufferInfo: %w", err)
	}
	return &info, nil
}

func getTerminalSize(file *os.File) (int16, int16, error) {
	info, err := getInfo(file)
	if err != nil {
		return 0, 0, err
	}
	return info.Size.X, info.Size.Y, nil
}

func validateError(err error) bool {
	return err.Error() != "The operation completed successfully."
}
