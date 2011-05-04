#!/usr/bin/env goscript

package main

import (
    "flag"
    "fmt"
)

func main() {
    fmt.Printf("goscript example\n")
    if flag.NArg() == 0 {
        fmt.Printf(" No args.\n")
    } else {
        fmt.Printf(" Args:\n")
        for i := 0; i < flag.NArg(); i++ {
            fmt.Printf("  %s\n", flag.Arg(i))
        }
    }
}

