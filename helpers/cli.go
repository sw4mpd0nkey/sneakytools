package helpers

/**
*	a lot of code stolen from https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
*
**/

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func CliInit() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func PrintBanner() {
	fmt.Println("####################################################################")
	fmt.Println("# S n e a k y  T o o l s")
	fmt.Println("####################################################################")
}

func PrintOptions() {
	fmt.Println("Please choose the type of operation you want to run: ")
	fmt.Println("1) TCP Scanner")
}
