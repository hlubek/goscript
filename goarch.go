package main

import (
    "flag"
    "fmt"
    "os"
    "runtime"
    )

func main() {
    if flag.NArg() == 0 || flag.Arg(0) == "-h" {
        fmt.Printf("usage: goarch <compiler|linker|ext>\n")
        os.Exit(1)
    }
    letter_map := map[string]string {
        "386": "8",
        "amd64": "6",
        "arm": "5",
    }
    letter := letter_map[runtime.GOARCH]
    switch flag.Arg(0) {
        case "compiler":
            fmt.Printf("%sg", letter)
        case "linker":
            fmt.Printf("%sl", letter)
        case "ext":
            fmt.Printf("%s", letter)
    }
}

