package raytrace

import (
	"image"
	"image/color"
	"math"
)

// Scene ...
type Scene struct {
	width  int
	height int
	fov    float32
	sphere Sphere
}

// CreateNewScene ...
func CreateNewScene(width, height int, fov float32, sphere Sphere) *Scene {
	return &Scene{
		width:  width,
		height: height,
		fov:    fov,
		sphere: sphere,
	}
}

// Render ...
func (s *Scene) Render() *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, s.width, s.height))

	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			ray := CreatePrime(x, y, s)

			if s.sphere.Intersect(ray) {
				m.Set(x, y, s.sphere.color)
			} else {
				m.Set(x, y, color.Black)
			}
		}
	}

	//draw.Draw(m, m.Bounds(), image.Transparent, image.ZP, draw.Src)
	return m
}

// DegreeToRadius ...
func DegreeToRadius(deg float32) float32 {
	return deg * math.Pi / 180.0
}
