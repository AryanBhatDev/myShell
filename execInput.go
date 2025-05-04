package main

import (
	"os"
	"os/exec"
	"strings"
)



func execInput(input string) error {

    input = strings.TrimSuffix(input, "\n")

    args := strings.Split(input, " ")

    switch args[0] {
    case "cd":

        if len(args) < 2 {
            return ErrNoPath
        }

        return os.Chdir(args[1])
    case "exit":
        os.Exit(0)
    }

    cmd := exec.Command("bash", "-c", input)

    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout

    return cmd.Run()
}