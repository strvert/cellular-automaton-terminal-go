package main

import (
    "fmt"
    "os"
    "time"

    "github.com/nsf/termbox-go"

    "github.com/strvworks/cellular-automaton-terminal-go/chunk"
//    "github.com/strvworks/cellular-automaton-terminal-go/cellutil"
    "github.com/strvworks/cellular-automaton-terminal-go/graphic"
)

const CELL_STR = "██"

func main() {
    var ch chunk.Chunk

    ch.SetCell(0, 0, 1)
    ch.SetCell(0, 1, 1)
    ch.SetCell(3, 3, 1)

    if err := termbox.Init(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer termbox.Close()


    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for range ticker.C {
            graphic.DrawChunk(&ch, CELL_STR)

        }
    }()

MAINLOOP:
    for {
        neighbors, err := ch.GetNeighborhood(1, 1)
        if err != nil {
            fmt.Println(err)
            break MAINLOOP
        }
        graphic.DrawBottomMessage(fmt.Sprintf("%08b", neighbors), 0, 0)
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
