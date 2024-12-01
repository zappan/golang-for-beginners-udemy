package util

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ClearScreen clears the screen
func ClearScreen() {
  if strings.Contains(runtime.GOOS, "windows") {
    // windows
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
  } else {
    // linux or mac
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
  }
}

