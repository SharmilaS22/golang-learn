package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// [dy][dx]
	sl := [][]uint8{}
	for i := 0; i < dy; i++ {
		subsl := []uint8{}
		for j := 0; j < dx; j++ {
			subsl = append(subsl, uint8(j))
		}
		sl = append(sl, subsl)
	}
	return sl
}

func main1() {
	pic.Show(Pic)
}
