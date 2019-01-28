package graphic

import (
    "github.com/nsf/termbox-go"
    "../chunk"
)

const coldef = termbox.ColorDefault

func DrawChunk(c *chunk.Chunk, cellStr string) error {
    for y := 0; y < 64; y++ {
        offsetX := 0
        for x := 0; x < 64; x++ {
            cell, err := c.GetCell(x, y)
            if err != nil {
                return err
            }

            cellColor := termbox.ColorBlack
            if cell == 1 {
                cellColor = termbox.ColorRed
            }

            for _, v := range cellStr {
                termbox.SetCell(x+offsetX, y, v, cellColor, coldef)
                offsetX++
            }
            offsetX--
        }
    }
    termbox.Flush()
    return nil
}

func DrawBottomMessage(message string, offsetX, offsetY int) {
    _, wy := termbox.Size()
    wy -= 1
    for i, r := range message {
        termbox.SetCell(i+offsetX, wy+offsetY, r, termbox.ColorWhite, termbox.ColorDefault)
    }
    termbox.Flush()
}
