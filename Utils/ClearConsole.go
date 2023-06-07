package Utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen(){
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
  	} else {
		return
	}
}
