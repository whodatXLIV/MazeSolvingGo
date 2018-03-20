package maze

import (
	"errors"
	"image"
	"image/color"
)

func PrepareMaze(img image.Image) (map[int][]int, error) {
	maze := img.(*image.Paletted)

	w, h := maze.Rect.Dx(), maze.Rect.Dy()
	c := maze.Pix
	m := make(map[int][]int)
	var (
		indx, up, down, left, right int
		connected, direction        []int
	)
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			indx = i*w + j
			if c[indx] != 0 {
				connected = []int{}
				up = (i-1)*w + j
				left = i*w + j - 1
				right = i*w + j + 1
				down = (i+1)*w + j
				direction = []int{up, left, right, down}

				for _, v := range direction {
					if v >= 0 && v < h*w {
						if c[v] != 0 {
							connected = append(connected, v)
						}
					}
				}
				m[indx] = connected

			}
		}
	}
	maze.Palette = append(maze.Palette, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255})
	return m, nil
}

func PrepareMazeExtra(img image.Image) (map[int][]int, error) {
	maze := img.(*image.Paletted)

	w, h := maze.Rect.Dx(), maze.Rect.Dy()
	c := maze.Pix
	m := make(map[int][]int)
	mt := make(map[int]int)
	var (
		indx, up, down, left, right int
		connected                   int
	)
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			indx = i*w + j
			if c[indx] != 0 {
				connected = 0
				up = int(c[(i-1)*w+j])
				left = int(c[i*w+j-1])
				right = int(c[i*w+j+1])
				down = int(c[(i+1)*w+j])

				if left == 1 {
					if right == 1 {
						if up == 1 || down == 1 {
							connected = indx

						}
					} else {
						connected = indx
					}

				}

				m[indx] = connected

			}
		}
	}
	maze.Palette = append(maze.Palette, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255})
	return m, nil
	// Think about using struct to hold begining end of maze
}
func FindEntrance(img image.Image) (int, error) {
	b := img.Bounds()
	w := b.Max.X
	var p int
	c := img.(*image.Paletted).Pix

	for i := 0; i < w; i++ {
		if c[i] > 0 {
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
	c := img.(*image.Paletted).Pix

	for i := 0; i < w; i++ {
		indx = h*(w-1) - 1 + i
		if c[indx] > 0 {
			p = indx
			return p, nil
		}
	}

	err := errors.New("No Exit Found")
	return p, err
}

func SolvedColor(img image.Image, ent, ext int, sol []int) {
	// Palette colors [Black white red green blue]
	maze := img.(*image.Paletted)

	for _, v := range sol {
		maze.Pix[v] = 4

	}

	maze.Pix[ent] = 3

	maze.Pix[ext] = 2

}

func PathLength(sol []int) int {
	return len(sol) + 1
}
