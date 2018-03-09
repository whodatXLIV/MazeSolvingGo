package maze

import (
	"errors"
	"image"
	"image/color"
)

func PrepareMaze(img image.Image) (map[int][]int, error) {
	oldImg := img.(*image.Paletted)

	w, h := oldImg.Rect.Dx(), oldImg.Rect.Dy()

	m := make(map[int][]int)
	var (
		indx, up, down, left, right int
		connected, direction        []int
	)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			indx = i*w + j
			if oldImg.Pix[indx] != 0 {
				connected = []int{}
				up = (i-1)*w + j
				left = i*w + j - 1
				right = i*w + j + 1
				down = (i+1)*w + j
				direction = []int{up, left, right, down}
				for _, v := range direction {
					if v >= 0 && v < h*w {
						if oldImg.Pix[v] != 0 {
							connected = append(connected, v)
						}
					}
				}
				m[indx] = connected

			}
		}
	}
	oldImg.Palette = append(oldImg.Palette, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255})
	return m, nil
}

func FindEntrance(img image.Image) (int, error) {
	b := img.Bounds()
	w := b.Max.X
	var p int
	var rp uint8

	for i := 0; i < w; i++ {
		rp = img.(*image.Paletted).Pix[i]
		if rp > 0 {
			p = i
			return p, nil
		}
	}

	err := errors.New("No Entrance Found")
	return p, err
}

func FindExit(img image.Image) (int, error) {
	b := img.Bounds()
	w, h := b.Max.X, b.Max.Y
	var p, indx int
	var rp uint8
	for i := 0; i < w; i++ {
		indx = h*(w-1) - 1 + i
		rp = img.(*image.Paletted).Pix[indx]
		if rp > 0 {
			p = indx
			return p, nil
		}
	}

	err := errors.New("No Exit Found")
	return p, err
}

func SolvedColor(img image.Image, ent, ext int, sol []int) {
	nImg := img.(*image.Paletted)

	for _, v := range sol {
		nImg.Pix[v] = 4

	}

	nImg.Pix[ent] = 3

	nImg.Pix[ext] = 2

}

func PathLength(sol []int) int {
	return len(sol) + 1
}
