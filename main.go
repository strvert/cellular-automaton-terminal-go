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
    updateInterval := 500

    for i := 0; i < 10; i++ {
        ch.SetCell(i+3, 3, 1)
    }

    if err := termbox.Init(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer termbox.Close()


    ticker := time.NewTicker(time.Millisecond * time.Duration(updateInterval))
    updaterFunc := func() {
        for range ticker.C {
            graphic.DrawChunk(&ch, CELL_STR)
            err := ch.UpdateChunk()
            if err != nil {
                fmt.Println(err)
            }
        }
    }

    go updaterFunc()

MAINLOOP:
    for {
        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            switch ev.Key {
            case termbox.KeyEsc:
                ticker.Stop()
                break MAINLOOP

            case termbox.KeyArrowUp:
                if updateInterval-50 <= 0 {
                    continue
                }
                updateInterval -= 50
                ticker.Stop()
                ticker = time.NewTicker(time.Millisecond * time.Duration(updateInterval))
                go updaterFunc()

            case termbox.KeyArrowDown:
                if updateInterval+50 <= 0 {
                    continue
                }
                updateInterval += 50
                ticker.Stop()
                ticker = time.NewTicker(time.Millisecond * time.Duration(updateInterval))
                go updaterFunc()
            }
        }
    }

    fmt.Println("Stop")
}
