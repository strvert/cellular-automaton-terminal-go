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

func (cc *Chunkcontroller) SetCell(cx, cy, x, y, v int, gen bool) (error) {
    if gen {
        newChunk := [2]int{0, 0}
        if x == CHUNK_SIZE-1 {
            newChunk[0] += 1
        }
        if y == CHUNK_SIZE-1 {
            newChunk[1] += 1
        }
        if x == 0 {
            newChunk[0] -= 1
        }
        if y == 0 {
            newChunk[1] -= 1
        }
        cc.NewChunk(cx+newChunk[0], cy+newChunk[1])
    }
    ch, err := cc.GetChunk(cx, cy, false)
    err = ch.SetCell(x, y, v)
    if err != nil {
        return err
    }
}

func (cc *Chunkcontroller) GetCell(cx, cy, x, y int, gen bool) (int, error) {
    ch, err := cc.GetChunk(cx, cy, gen)
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
        if coord[0] >= CHUNK_SIZE || coord[1] >= CHUNK_SIZE || coord[0] < 0 || coord[1] < 0 {
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
