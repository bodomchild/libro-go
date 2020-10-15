// Surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyRange       = 30.0
	xyScale       = width / 2 / xyRange
	zScale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Draw() string {
	var svg string
	zMin, zMax := minmax() // Ej 3.3, copiado de internet
	svg += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if !check(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			svg += fmt.Sprintf("<polygon style='stroke: %s; fill: white;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j, zMin, zMax), ax, ay, bx, by, cx, cy, dx, dy) // La funcion color es parte del ej 3.3
		}
	}
	svg += fmt.Sprintln("</svg>")
	return svg
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y) // Ejemplo del libro
	//z := eggbox(x, y) // Forma de caja de huevos
	//z := saddle(x, y) // Forma de montura

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 { // Ej 3.2
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 { // Ej 3.2
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return y*y/a2 - x*x/b2
}

func check(points ...float64) bool { // Ej 3.1
	for _, p := range points {
		if math.IsNaN(p) || math.IsInf(p, 0) {
			return false
		}
	}
	return true
}

// minmax returns the min and max values for z given the min/max values of x
// and y and assuming a square domain.
func minmax() (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xOff := 0; xOff <= 1; xOff++ {
				for yOff := 0; yOff <= 1; yOff++ {
					x := xyRange * (float64(i+xOff)/cells - 0.5)
					y := xyRange * (float64(j+yOff)/cells - 0.5)
					z := f(x, y) // Cambiar esto si se cambia la forma que se quiere dibujar
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return
}

func color(i, j int, zMin, zMax float64) string {
	min := math.NaN()
	max := math.NaN()
	for xOff := 0; xOff <= 1; xOff++ {
		for yOff := 0; yOff <= 1; yOff++ {
			x := xyRange * (float64(i+xOff)/cells - 0.5)
			y := xyRange * (float64(j+yOff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zMax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zMin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}
