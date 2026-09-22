//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
)

const FILE = "TEST.INP"

var (
	inp = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func runFile() {
	exec.Command("./solution").Run()
	exec.Command("./main").Run()
}

func check() bool {
	cmd := exec.Command("diff", "-w", "TEST.OUT", "TEST.ANS")
	output, err := cmd.CombinedOutput()

	if err == nil {
		return true
	} else {
		fmt.Fprintln(out, string(output))
		return false
	}
}

func gen() {
	fIn, _ := os.Create(FILE)
	Iout := bufio.NewWriter(fIn)

	fmt.Fprintln(Iout, rand.IntN(100), rand.IntN(100))

	Iout.Flush()
	fIn.Close()
}

func main() {
	defer out.Flush()
	exec.Command("go", "build", "-o", "solution", "solution.go").Run()
	exec.Command("go", "build", "-o", "main", "main.go").Run()

	for test := 1; test <= 100; test++ {
		gen()
		runFile()
		if check() {
			fmt.Fprintln(out, "Test #", test, ": AC")
		} else {
			fmt.Fprintln(out, "Test #", test, ": WA")
			return
		}
	}
}
