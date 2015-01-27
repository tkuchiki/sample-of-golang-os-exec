package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-al")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println()
	}

	cmd.Wait()
}

