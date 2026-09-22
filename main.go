//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	inp *bufio.Scanner
	out *bufio.Writer
)

func nextStr() string {
	inp.Scan()
	return inp.Text()
}

func nextInt() int {
	x, _ := strconv.Atoi(nextStr())
	return x
}

func openFile() {
	inpFile := "TEST.INP"
	outFile := "TEST.OUT"

	if fIn, err := os.Open(inpFile); err == nil {
		inp = bufio.NewScanner(fIn)
		if fOut, err := os.Create(outFile); err == nil {
			out = bufio.NewWriter(fOut)
		} else {
			out = bufio.NewWriter(os.Stdout)
		}
	} else {
		inp = bufio.NewScanner(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
	}

	inp.Split(bufio.ScanWords)
	inp.Buffer(make([]byte, 1024), 1e7)
}

func main() {
	openFile()
	defer out.Flush()

	a := nextInt()
	b := nextInt()
	fmt.Fprintln(out, a+b+67)
}
