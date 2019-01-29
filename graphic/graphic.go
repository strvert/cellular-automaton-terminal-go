package graphic

import (
    "github.com/nsf/termbox-go"
    "../chunk"
    "../chunkcontroller"
)

const CHUNK_SIZE = 64

const coldef = termbox.ColorDefault

type ScreenField struct {
    W int
    H int
    ChunkOffset [2]int
    CellOffset [2]int
}

func CalcDrawChunkNum(w, h int) (int, int) {
    cx := (w / CHUNK_SIZE) + 1
    cy := (h / CHUNK_SIZE) + 1
    return cx, cy
}

func DrawField(cc *chunkcontroller.Chunkcontroller, field ScreenField, cellStr string) (int, error) {
    chnumX, chnumY := CalcDrawChunkNum(field.W, field.H)
    choffX := field.ChunkOffset[0]
    choffY := field.ChunkOffset[1]
    ceoffX := field.CellOffset[0]
    ceoffY := field.CellOffset[1]

    for chx := choffX; chx < chnumX; chx++ {
        for chy := choffY; chy < chnumY; chy++ {
            currch, err := cc.GetChunk(chx, chy, true)
            if err != nil {
                return 0, err
            }
            err = DrawChunk(currch, [2]int{chx*CHUNK_SIZE+ceoffX, chy*CHUNK_SIZE+ceoffY}, cellStr)
            if err != nil {
                return 0, err
            }
        }
    }
    return 0, nil
}

func DrawChunk(c *chunk.Chunk, celloffset [2]int, cellStr string) error {
    for y := 0; y < CHUNK_SIZE; y++ {
        offsetX := 0
        for x := 0; x < CHUNK_SIZE; x++ {
            cell, err := c.GetCell(x, y)
            if err != nil {
                return err
            }

            cellColor := termbox.ColorBlack
            if cell == 1 {
                cellColor = termbox.ColorRed
            }

            for _, v := range cellStr {
                termbox.SetCell(x+offsetX+celloffset[0], y+celloffset[1], v, cellColor, coldef)
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
