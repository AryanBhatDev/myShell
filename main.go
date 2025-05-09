package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("=> ")

        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        if err = execInput(input); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}
