package main

import (
    "fmt"
)

const BUFSIZE = 1024

type chunk struct {
    cells [64][8]byte
}

func (c *chunk) GetCell(x, y int) int {
    px := x / 8
    py := y
    shift := x % 8
    return int((c.cells[py][px] >> byte(shift)) & 1)
}

func (c *chunk) SetCell(x, y, n int) {
    px := x / 8
    py := y
    shift := x % 8
    if n == 1 {
        c.cells[py][px] = c.cells[py][px] | (byte(0x80) >> byte(shift))
    } else {
        c.cells[py][px] = c.cells[py][px] & ^(byte(0x80) >> byte(shift))
    }
}

func main() {
    var b chunk

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
