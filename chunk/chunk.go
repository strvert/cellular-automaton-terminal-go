package chunk

type Chunk struct {
    cells [64][8]byte
}

func (c *Chunk) GetCell(x, y int) int {
    px := x / 8
    py := y
    shift := x % 8
    return int((c.cells[py][px] >> byte(shift)) & 1)
}

func (c *Chunk) SetCell(x, y, n int) {
    px := x / 8
    py := y
    shift := x % 8
    if n == 1 {
        c.cells[py][px] = c.cells[py][px] | (byte(0x80) >> byte(shift))
    } else {
        c.cells[py][px] = c.cells[py][px] & ^(byte(0x80) >> byte(shift))
    }
}
