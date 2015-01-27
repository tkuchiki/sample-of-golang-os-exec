package main

import (
        "fmt"
        "io"
        "os"
        "os/exec"
)

func main() {
        cmd := exec.Command("wc", "-w")
        stdin, err := cmd.StdinPipe()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        io.WriteString(stdin, "hoge foo bar")
        stdin.Close()
        out, err := cmd.Output()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        fmt.Println(string(out))
}

