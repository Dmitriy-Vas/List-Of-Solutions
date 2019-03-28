package Hough_Transform

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func Hough(input string, output string, ntx int, mry int) bool {
	f, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return false
	}
	im, err := png.Decode(f)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if err = f.Close(); err != nil {
		fmt.Println(err)
		return false
	}
	nimx := im.Bounds().Max.X
	mimy := im.Bounds().Max.Y
	mry = int(mry/2) * 2
	him := image.NewGray(image.Rect(0, 0, ntx, mry))
	draw.Draw(him, him.Bounds(), image.NewUniform(color.White),
		image.ZP, draw.Src)

	rmax := math.Hypot(float64(nimx), float64(mimy))
	dr := rmax / float64(mry/2)
	dth := math.Pi / float64(ntx)

	for jx := 0; jx < nimx; jx++ {
		for iy := 0; iy < mimy; iy++ {
			col := color.GrayModel.Convert(im.At(jx, iy)).(color.Gray)
			if col.Y == 255 {
				continue
			}
			for jtx := 0; jtx < ntx; jtx++ {
				th := dth * float64(jtx)
				r := float64(jx)*math.Cos(th) + float64(iy)*math.Sin(th)
				iry := mry/2 - int(math.Floor(r/dr+.5))
				col = him.At(jtx, iry).(color.Gray)
				if col.Y > 0 {
					col.Y--
					him.SetGray(jtx, iry, col)
				}
			}
		}
	}
	if f, err = os.Create(output); err != nil {
		fmt.Println(err)
		return false
	}
	if err := png.Encode(f, him); err != nil {
		fmt.Println(err)
		return false
	}
	if cErr := f.Close(); cErr != nil && err == nil {
		fmt.Println(err)
		return false
	}
	return true
}
