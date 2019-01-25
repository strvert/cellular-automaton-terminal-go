package main

import (
    "fmt"
    "os"
    "math/bits"
    "./chunk"
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
            v, err := b.GetCell(x, y)
            if err != nil {
                fmt.Println(err)
            }
            if v == 1 {
                fmt.Print("█")
            } else {
                fmt.Print("▓")
            }
        }
        fmt.Println()
    }
    neighbor, err := b.GetNeigborhood(1, 1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("%b\n", neighbor.Data)
    fmt.Println(bits.OnesCount8(neighbor.Data))
}
