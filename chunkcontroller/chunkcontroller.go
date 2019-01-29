package chunkcontroller

import (
    "errors"

    "../chunk"
)

const CHUNK_SIZE = 64

type Chunkcontroller struct {
    Chunkset map[[2]int]*chunk.Chunk
}

func NewChunkcontroller() (*Chunkcontroller) {
    var cc Chunkcontroller
    cc.Chunkset = map[[2]int]*chunk.Chunk{}
    return &cc
}

func (cc *Chunkcontroller) GetChunk(x, y int, gen bool) (*chunk.Chunk, error) {
    if val, ok := cc.Chunkset[[2]int{x, y}]; ok {
        return val, nil
    } else {
        if gen {
            cc.NewChunk(x, y)
            ch, err := cc.GetChunk(x, y, true)
            if err != nil {
                return nil, err
            }
            return ch, nil
        } else {
            return nil, errors.New("That chunk is not found")
        }
    }
}

func (cc *Chunkcontroller) SetCell(cx, cy, x, y, v int, aroundgen bool) (error) {
    if x >= CHUNK_SIZE || y >= CHUNK_SIZE || x < 0 || y < 0 {
        return errors.New("That cell is out of chunk")
    }
    if aroundgen {
        if (x == CHUNK_SIZE-1 || y == CHUNK_SIZE-1 || x == 0 || y == 0) && v == 1 {
            if y == 0 && x == CHUNK_SIZE-1 {
                if ! cc.CheckChunk(cx+1, cy-1) {
                    cc.NewChunk(cx+1, cy-1)
                }
            }
            if x == CHUNK_SIZE-1 {
                if ! cc.CheckChunk(cx+1, cy) {
                    cc.NewChunk(cx+1, cy)
                }
            }
            if x == CHUNK_SIZE-1 && y == CHUNK_SIZE-1 {
                if ! cc.CheckChunk(cx+1, cy+1) {
                    cc.NewChunk(cx+1, cy+1)
                }
            }
            if y == CHUNK_SIZE-1 {
                if ! cc.CheckChunk(cx, cy+1) {
                    cc.NewChunk(cx, cy+1)
                }
            }
            if y == CHUNK_SIZE-1 && x == 0 {
                if ! cc.CheckChunk(cx-1, cy+1) {
                    cc.NewChunk(cx-1, cy+1)
                }
            }
            if x == 0 {
                if ! cc.CheckChunk(cx-1, cy) {
                    cc.NewChunk(cx-1, cy)
                }
            }
            if x == 0 && y == 0 {
                if ! cc.CheckChunk(cx-1, cy-1) {
                    cc.NewChunk(cx-1, cy-1)
                }
            }
            if y == 0 {
                if ! cc.CheckChunk(cx, cy-1) {
                    cc.NewChunk(cx, cy-1)
                }
            }
        }
    }
    ch, err := cc.GetChunk(cx, cy, false)
    err = ch.SetCell(x, y, v)
    if err != nil {
        return err
    }
    return nil
}

func (cc *Chunkcontroller) GetCell(cx, cy, x, y int) (int, error) {
    ch, err := cc.GetChunk(cx, cy, false)
    if err != nil {
        return 0, err
    }
    cell, err := ch.GetCell(x, y)
    if err != nil {
        return 0, err
    }
    return cell, nil
}

func (cc *Chunkcontroller) GetNeighborhood(cx, cy, x, y int) (byte, error) {
    coords := [][]int{{x, y-1}, {x+1, y-1}, {x+1, y}, {x+1, y+1}, {x, y+1}, {x-1, y+1}, {x-1, y}, {x-1, y-1}}

    var neighbors byte = byte(0)
    var cell int = 0

    for i, coord := range coords {
        if coord[0] == CHUNK_SIZE || coord[1] == CHUNK_SIZE || coord[0] == -1 || coord[1] == -1 {
            if coord[1] == -1 && coord[0] == CHUNK_SIZE {
                aroundch, err := cc.GetChunk(cx+1, cy-1, false)
                if err != nil {
                    return byte(0), err
                }
                cell, err = aroundch.GetCell(0, CHUNK_SIZE-1)
                if err != nil {
                    return byte(0), err
                }
            }
            if coord[0] == CHUNK_SIZE {
                aroundch, err := cc.GetChunk(cx+1, cy, false)
                if err != nil {
                    return byte(0), err
                }
                cell, err = aroundch.GetCell(0, y)
                if err != nil {
                    return byte(0), err
                }
            }
            if coord[0] == CHUNK_SIZE && coord[1] == CHUNK_SIZE {
                aroundch, err := cc.GetChunk(cx+1, cy+1, false)
                if err != nil {
                    return byte(0), err
                }
                cell, err = aroundch.GetCell(0, 0)
                if err != nil {
                    return byte(0), err
                }
            }
            if coord[1] == CHUNK_SIZE {
                aroundch, err := cc.GetChunk(cx, cy+1, false)
                if err != nil {
                    return byte(0), err
                }
                cell, err = aroundch.GetCell(x, 0)
                if err != nil {
                    return byte(0), err
                }
            }
            if coord[1] == CHUNK_SIZE && coord[0] == -1 {
                aroundch, err := cc.GetChunk(cx-1, cy+1, false)
                if err != nil {
                    return byte(0), err
                }
                cell, err = aroundch.GetCell(CHUNK_SIZE-1, 0)
                if err != nil {
                    return byte(0), err
                }
            }
            if coord[0] == -1 {
                aroundch, err := cc.GetChunk(cx-1, cy, false)
                if err != nil {
                    return byte(0), err
                }
                cell, err = aroundch.GetCell(CHUNK_SIZE-1, y)
                if err != nil {
                    return byte(0), err
                }
            }
            if coord[0] == -1 && coord[1] == -1 {
                aroundch, err := cc.GetChunk(cx-1, cy-1, false)
                if err != nil {
                    return byte(0), err
                }
                cell, err = aroundch.GetCell(CHUNK_SIZE-1, CHUNK_SIZE-1)
                if err != nil {
                    return byte(0), err
                }
            }
            if coord[1] == -1 {
                aroundch, err := cc.GetChunk(cx, cy-1, false)
                if err != nil {
                    return byte(0), err
                }
                cell, err = aroundch.GetCell(x, CHUNK_SIZE-1)
                if err != nil {
                    return byte(0), err
                }
            }
            neighbors = neighbors | (byte(cell) << byte(7-i))
        } else {
            c, err := cc.GetChunk(cx, cy, false)
            if err != nil {
                return byte(0), err
            }
            cell, err := c.GetCell(coord[0], coord[1])
            if err != nil {
                return byte(0), err
            }
            neighbors = neighbors | (byte(cell) << byte(7-i))
        }
    }
    return neighbors, nil
}
func (cc *Chunkcontroller) CheckChunk(cx, cy int) (bool) {
    _, ok := cc.Chunkset[[2]int{cx, cy}]
    return ok
}

func (cc *Chunkcontroller) SetChunk(c *chunk.Chunk, x, y int) {
    cc.Chunkset[[2]int{x, y}] = c
}

func (cc *Chunkcontroller) NewChunk(x, y int) {
    var c chunk.Chunk
    cc.Chunkset[[2]int{x, y}] = &c
}

func (CC *Chunkcontroller) UpdateField() {

}
