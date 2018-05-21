package raytrace

import (
	"image"
	"image/draw"
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
func (s *Scene) Render() {
	m := image.NewRGBA(image.Rect(0, 0, s.width, s.height))
	draw.Draw(m, m.Bounds(), image.Transparent, image.ZP, draw.Src)
}

// DegreeToRadius ...
func DegreeToRadius(deg float32) float32 {
	return deg * math.Pi / 180.0
}
