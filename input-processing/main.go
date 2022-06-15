package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/software-engr-full-stack/backend-developer-tests/input-processing/pkg/inputprocessor"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	// TODO: Look for lines in the STDIN reader that contain "error" and output them.
	err := inputprocessor.Run(reader)
	if err != nil {
		panic(err)
	}
}
