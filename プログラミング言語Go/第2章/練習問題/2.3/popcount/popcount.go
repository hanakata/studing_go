package popcount

var pc [256]byte
var y int
var i uint64

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	for i = 0; i < 8; i++ {
		y += int(pc[byte(x>>(i*8))])
	}
	return y
}
