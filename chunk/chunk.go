package chunk

import (
    "errors"
)


type Chunk struct {
    cells [64][8]byte
}

type Neighborhood struct {
    Data byte
}

func calcBitCoord(x, y int) (int, int, byte) {
    px := x / 8
    py := y
    shift := byte(x % 8)
    return px, py, shift
}

func (c *Chunk) GetCell(x, y int) (int, error) {
    if x >= 64 && y >= 64 {
        return 0, errors.New("out of chunk")
    }

    px, py, shift := calcBitCoord(x, y)
    return int((c.cells[py][px] >> byte(shift)) & 1), nil
}

func (c *Chunk) SetCell(x, y, n int) error {
    if x >= 64 && y >= 64 {
        return errors.New("out of chunk")
    }

    px, py, shift := calcBitCoord(x, y)
    if n == 1 {
        c.cells[py][px] = c.cells[py][px] | (byte(0x01) << byte(shift))
    } else {
        c.cells[py][px] = c.cells[py][px] & ^(byte(0x01) << byte(shift))
    }

    return nil
}

func (c *Chunk) GetNeigborhood(x, y int) (Neighborhood, error) {
    coords := [][]int{{x, y-1}, {x+1, y-1}, {x+1, y}, {x+1, y+1}, {x, y+1}, {x-1, y+1}, {x-1, y}, {x-1, y-1}}

    var neighbors Neighborhood = Neighborhood{byte(0)}
    for i, coord := range coords {
        cell, err := c.GetCell(coord[0], coord[1])
        if err != nil {
            return Neighborhood{0}, err
        }
        neighbors.Data = neighbors.Data | byte(cell) << byte(7-i)
    }
    return neighbors, nil
}
