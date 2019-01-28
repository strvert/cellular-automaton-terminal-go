package chunkcontroller

import (
    "errors"

    "github.com/strvworks/cellular-automaton-terminal-go/chunk"
)

type Chunkcontroller struct {
    Chunkset map[[2]int]*chunk.Chunk
}

func NewChunkcontroller (*Chunkcontroller) {
    var cc Chunkcontroller
    cc.Chunkset = {}
    return &cc
}

func (cc *Chunkcontroller) GetChunk(x, y int) (*chunk.Chunk, error) {
    if val, ok := cc.Chunkset[[2]int{x, y}]; ok {
        return val, nil
    } else {
        return nil, errors.New("That chunk is not found")
    }
}

func (cc *Chunkcontroller) SetChunk(c *chunk.Chunk, x, y int) (*chunk.Chunk) {
    cc.Chunkset[[2]int{x, y}] = c
}

func (cc *Chunkcontroller) NewChunk(x, y int) (*chunk.Chunk) {
    cc.Chunkset[[2]int{x, y}] = chunk.Chunk
}
