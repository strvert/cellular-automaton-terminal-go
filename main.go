package main

import (
    "fmt"
    "github.com/celluar-automaton-terminal-go/chunk"
)

const BUFSIZE = 1024

func main() {
    var b chunk.Chunk

    for x := 0; x < 30; x++ {
        b.SetCell(x, 0, 1)
        b.SetCell(0, x, 1)
    }

    for y := 0; y < 64; y++ {
        for x := 0; x < 64; x++ {
            fmt.Print(b.GetCell(x, y))
        }
        fmt.Println()
    }
}
