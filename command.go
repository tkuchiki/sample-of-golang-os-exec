package main

import (
	"fmt"
	_ "log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")

	fmt.Printf("Path: %v\n", cmd.Path)
	fmt.Printf("Args: %v\n", cmd.Args)
	fmt.Printf("Env: %v\n", cmd.Env)
	fmt.Printf("Dir: %v\n", cmd.Dir)
	fmt.Printf("Stdin: %v\n", cmd.Stdin)
	fmt.Printf("Stdout: %v\n", cmd.Stdout)
	fmt.Printf("Stderr: %v\n", cmd.Stderr)
	fmt.Printf("ExtraFiles: %v\n", cmd.ExtraFiles)
	fmt.Printf("Process: %#v\n", cmd.Process)
	fmt.Printf("ProcessState: %#v\n", cmd.ProcessState)

	//if err != nil {
	//	log.Fatal(err)
	//}
}
