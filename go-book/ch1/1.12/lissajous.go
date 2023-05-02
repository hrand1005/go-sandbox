package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var (
	green   = color.RGBA{0, 255, 0, 255}
	red     = color.RGBA{255, 0, 0, 255}
	blue    = color.RGBA{0, 0, 255, 255}
	palette = []color.Color{
		color.Black,
		green,
		red,
		blue,
	}
)

type gifOpts struct {
	Cycles  int
	Res     float64
	Size    int
	NFrames int
	Delay   int
}

var defaultGifOpts = gifOpts{
	Cycles:  5,
	Res:     0.001,
	Size:    100,
	NFrames: 64,
	Delay:   8,
}

func NewGifOpts() gifOpts {
	return defaultGifOpts
}

func lissajous(out io.Writer) {
	lissajousWithGIFOpts(out, defaultGifOpts)
}

func lissajousWithGIFOpts(out io.Writer, gOpts gifOpts) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: gOpts.NFrames}
	phase := 0.0
	for i := 0; i < gOpts.NFrames; i++ {
		rect := image.Rect(0, 0, 2*gOpts.Size+1, 2*gOpts.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(gOpts.Cycles)*2*math.Pi; t += gOpts.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(gOpts.Size+int(x*float64(gOpts.Size)+0.5),
				gOpts.Size+int(y*float64(gOpts.Size)+0.5), uint8(i%3)+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, gOpts.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
