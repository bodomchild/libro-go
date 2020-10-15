package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // primer color de la paleta
	blackIndex = 1 // segundo color de la paleta
)

func main() {
	Lissajous(os.Stdout)
}

func Lissajous(out io.Writer) {
	const (
		cycles  = 5     // numero de revoluciones del oscilador en x completas
		res     = 0.001 // resolucion angular
		size    = 100   // el lienzo de la imagen cubre [-size..+size]
		nframes = 64    // numero de cuadros de animacion
		delay   = 8     // demora entre cuadros en unidades de 10ms
	)
	freq := rand.Float64() * 3.0 // frecuencia relativa del oscilador en y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferencia de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim) // se ignoran errores de codificado
}

// TODO ejercicios 1.5 y 1.6
