package main

import (
    "fmt"
    "math"
)

const (
    width, height = 600, 320
    cells = 100
    xyrange = 30.0
    xyscale = width / 2 / xyrange
    zscale = height * 0.4
    angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) 

func main() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j)
            bx, by := corner(i, j)
            cx, cy := corner(i, j+1)
            dx, dy := corner(i+1, j+1)
            if checkValidFloats([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
                fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", 
                    ax, ay, bx, by, cx, cy, dx, dy)
            } 
        }
    }
    fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
    // Find point (x,y) at corner of cell (i, j)
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z.
    z := f(x, y)
    
    // Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0, 0)
    return math.Sin(r) / r
}

// Exercise 3.1 is to verify that float values are not non-finite.
// checkValidFloats takes a slice of float values as inputs, checking for
// non-finite values. If any floats in the provided slice are non-finite
// or NaN, then returns false, otherwise returns true.
func checkValidFloats(f []float64) bool {
    for _, v := range f {
        // for any two NaN values a and b, a != b in go
        if math.IsInf(v, 1) || math.IsInf(v, -1) || math.IsNaN(v) {
            return false
        }
    }

    return true
}


