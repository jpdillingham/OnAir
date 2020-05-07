package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-ps"
)

func main() {
	// tasklist /FI "WINDOWTITLE eq Zoom Meeting" /FO CSV /NH
	// exit code
	rc := 0

	ps, _ := ps.Processes()

	for i, _ := range ps {
		fmt.Println(ps[i].Executable())
	}

	fmt.Println("")
	os.Exit(rc)
}
