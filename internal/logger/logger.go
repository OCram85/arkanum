package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// getCaller walks up 2 frames (skipping getCaller itself and the Info/Warn/Error
// wrapper) and returns a short "package.Function" label for use as section.
func getCaller() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}
	// full name looks like: arkanum/internal/installer.DockerCLI
	// strip the module path prefix, keep "installer.DockerCLI"
	name := f.Name()
	if idx := strings.LastIndex(name, "/"); idx >= 0 {
		name = name[idx+1:]
	}
	return name
}

// Info prints a green info message with the caller's package.Function as section.
func Info(m string) {
	section := getCaller()
	fmt.Println(Green + "🧙 arkanum " + Reset + Blue + "[⚒️  " + section + "]" + Reset + ": " + m)
}

// Warn prints a yellow warning message to stderr with the caller as section.
func Warn(m string) {
	section := getCaller()
	fmt.Fprintln(os.Stderr, Yellow+"🧙 arkanum "+Reset+Blue+"[⚒️  "+section+"]"+Reset+": "+m)
}

// Error prints a red error message to stderr with the caller as section.
func Error(m string) {
	section := getCaller()
	msg := Red + "🧙 arkanum " + Reset + Blue + "[⚒️  " + section + "]" + Reset + ": " + m
	fmt.Fprintln(os.Stderr, msg)
}
