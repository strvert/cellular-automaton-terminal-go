package main

import (
    "fmt"
    "os"
    "time"
//    "math/bits"

    "github.com/nsf/termbox-go"

    "github.com/strvworks/cellular-automaton-terminal-go/chunk"
//    "github.com/strvworks/cellular-automaton-terminal-go/cellutil"
    "github.com/strvworks/cellular-automaton-terminal-go/graphic"
)

const CELL_STR = "██"

func main() {
    var b chunk.Chunk

    for x := 0; x < 30; x++ {
        b.SetCell(x, x, 1)
    }

    if err := termbox.Init(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer termbox.Close()


    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for range ticker.C {
            graphic.DrawChunk(&b, CELL_STR)
        }
    }()

MAINLOOP:
    for {
        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            switch ev.Key {
            case termbox.KeyEsc:
                break MAINLOOP
            }
        }
    }

    ticker.Stop()
    fmt.Println("Stop")
}
