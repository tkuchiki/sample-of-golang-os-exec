package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ls")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path)

	_, err = exec.LookPath("hogehogehoge")

	fmt.Printf("%T\n", err)
	fmt.Printf("%v\n", err.Error())
}
