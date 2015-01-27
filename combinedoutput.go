package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, _ := exec.Command("ls", "-X").CombinedOutput()
	fmt.Println(string(out))

	out, _ = exec.Command("ls", "-al").CombinedOutput()
	fmt.Println(string(out))
}
