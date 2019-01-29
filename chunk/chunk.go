package chunk

import (
    "errors"
    "math/bits"
)

const CHUNK_SIZE = 64

type Chunk struct {
    cells [64][8]byte
}

func calcBitCoord(x, y int) (int, int, byte) {
    px := x / 8
    py := y
    shift := byte(x % 8)
    return px, py, shift
}

func (c *Chunk) GetCell(x, y int) (int, error) {
    if x >= CHUNK_SIZE || y >= CHUNK_SIZE || x < 0 || y < 0 {
        return 0, errors.New("That cell is out of chunk")
    }

    px, py, shift := calcBitCoord(x, y)
    return int((c.cells[py][px] >> byte(shift)) & 1), nil
}

func (c *Chunk) SetCell(x, y, v int) error {
    if x >= CHUNK_SIZE || y >= CHUNK_SIZE || x < 0 || y < 0 {
        return errors.New("That cell is out of chunk")
    }

    px, py, shift := calcBitCoord(x, y)
    if v == 1 {
        c.cells[py][px] = c.cells[py][px] | (byte(0x01) << byte(shift))
    } else {
        c.cells[py][px] = c.cells[py][px] & ^(byte(0x01) << byte(shift))
    }

    return nil
}

func (c *Chunk) GetNeighborhood(x, y int) (byte, error) {
    if x >= CHUNK_SIZE || y >= CHUNK_SIZE || x < 0 || y < 0 {
        return byte(0), errors.New("out of chunk")
    }

    coords := [][]int{{x, y-1}, {x+1, y-1}, {x+1, y}, {x+1, y+1}, {x, y+1}, {x-1, y+1}, {x-1, y}, {x-1, y-1}}

    var neighbors byte = byte(0)
    var cell int = 0
    for i, coord := range coords {
        if coord[0] >= CHUNK_SIZE || coord[1] >= CHUNK_SIZE || coord[0] < 0 || coord[1] < 0 {
            cell = 0
            neighbors = neighbors | (byte(cell) << byte(7-i))
        } else {
            cell, err := c.GetCell(coord[0], coord[1])
            if err != nil {
                return byte(0), err
            }
            neighbors = neighbors | (byte(cell) << byte(7-i))
        }
    }
    return neighbors, nil
}

func (c *Chunk) CalcNextCellState(x, y int) (int, error) {
    if x >= CHUNK_SIZE || y >= CHUNK_SIZE || x < 0 || y < 0 {
        return 0, errors.New("out of chunk")
    }
    neighbors, err := c.GetNeighborhood(x, y)
    if err != nil {
        return 0, err
    }

    ncount := bits.OnesCount8(neighbors)
    curr, err := c.GetCell(x, y)
    if err != nil {
        return 0, err
    }

    if curr == 1 {
        if ncount == 2 || ncount == 3 {
            return 1, nil
        } else {
            return 0, nil
        }
    } else {
        if ncount == 3 {
            return 1, nil
        } else {
            return 0, nil
        }
    }
    return 0, errors.New("Invalid cell state")
}

func (c *Chunk) UpdateChunk() error {
    var backChunk Chunk
    for y := 0; y < CHUNK_SIZE; y++ {
        for x := 0; x < CHUNK_SIZE; x++ {
            nextState, err := c.CalcNextCellState(x, y)
            if err != nil {
                return err
            }
            err = backChunk.SetCell(x, y, nextState)
            if err != nil {
                return err
            }
        }
    }
    c.cells = backChunk.cells
    return nil
}
