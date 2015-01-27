package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("echo", "hoge").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	_, err = exec.Command("echoechoecho", "hoge").Output()

	fmt.Printf("%T\n", err)
	fmt.Printf("%v\n", err.Error())
}
