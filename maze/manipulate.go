package maze

import (
	"errors"
	"image"
	"image/color"
)

func PrepareMaze(img image.Image) (map[int][]int, image.Image, error) {
	oldImg := img.(*image.Paletted)

	newImg := image.NewRGBA(oldImg.Rect)
	w, h := newImg.Rect.Dx(), newImg.Rect.Dy()

	oldColor := oldImg.Pix
	newColor := make([]uint8, 4*len(oldColor))
	m := make(map[int][]int)
	var (
		indx, up, down, left, right int
		connected, direction        []int
	)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			indx = i*w + j
			newColor[indx*4+0] = oldColor[indx] * 255
			newColor[indx*4+1] = oldColor[indx] * 255
			newColor[indx*4+2] = oldColor[indx] * 255
			newColor[indx*4+3] = 255

			if oldColor[indx] != 0 {
				connected = []int{}
				up = (i-1)*w + j
				left = i*w + j - 1
				right = i*w + j + 1
				down = (i+1)*w + j
				direction = []int{up, left, right, down}
				for _, v := range direction {
					if v >= 0 && v < h*w {
						if oldColor[v] != 0 {
							connected = append(connected, v)
						}
					}
				}
				m[indx] = connected

			}
		}
	}
	newImg.Pix = newColor
	oldImg.Palette = append(oldImg.Palette, color.RGBA{255, 0, 0, 255})
	return m, newImg, nil
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

func SolvedColor(img image.Image, ent, ext int, sol []int) image.Image {
	nImg := img.(*image.RGBA)

	newColor := nImg.Pix

	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	for _, v := range sol {
		newColor[v*4+0] = blue.R
		newColor[v*4+1] = blue.G
		newColor[v*4+2] = blue.B
		newColor[v*4+3] = blue.A
	}

	newColor[ent*4+0] = green.R
	newColor[ent*4+1] = green.G
	newColor[ent*4+2] = green.B
	newColor[ent*4+3] = green.A

	newColor[ext*4+0] = red.R
	newColor[ext*4+1] = red.G
	newColor[ext*4+2] = red.B
	newColor[ext*4+3] = red.A

	nImg.Pix = newColor
	return nImg

}

func PathLength(sol []int) int {
	return len(sol) + 1
}
