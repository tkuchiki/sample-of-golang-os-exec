package main

import (
        "bytes"
        "fmt"
        "os"
        "os/exec"
)

func main() {
        cmd := exec.Command("ls", "-al")
        var stdout bytes.Buffer
        cmd.Stdout = &stdout

        err := cmd.Run()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        fmt.Println(stdout.String())
}

