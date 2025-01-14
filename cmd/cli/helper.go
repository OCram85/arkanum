package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"context"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

// Returns the current function block for use in log output.
func getCaller() string {
	// remove getCaller and say*** from stack
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	} else {
		return f.Name()
	}
}

func installSingleExtension(ctx context.Context, name string) error {
	say(fmt.Sprintf("Installing '%v'...", name))
	cmd := exec.CommandContext(ctx, "install-extension", name, "--force")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func execCommand(ctx context.Context, name string, arg ...string) error {
	cmd := exec.CommandContext(ctx, name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err:= cmd.Run(); err != nil {
		return err
	}
	return nil
}

func say(m string) {
	section := getCaller()
	fmt.Println(Green + "🧙 arkanum " + Reset + Blue + "[⚒️  " + section + "]" + Reset + ": " + m)
}

func sayWarning(m string) {
	section := getCaller()
	fmt.Println(Yellow + "🧙 arkanum " + Reset + Blue + "[⚒️  " + section + "]" + Reset + ": " + m)
}

func sayError(m string) error {
	section := getCaller()
	msg := (Red + "🧙 arkanum " + Reset + Blue + "[⚒️  " + section + "]" + Reset + ": " + m)
	fmt.Println(msg)
	return errors.New(msg)
}
