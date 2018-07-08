package raytrace

import (
	"errors"
	"image"
	"image/color"
	"math"
)

// Scene ...
type Scene struct {
	width   int
	height  int
	fov     float32
	spheres []Sphere
}

// CreateNewScene ...
func CreateNewScene(width, height int, fov float32, sphere Sphere) *Scene {
	return &Scene{
		width:   width,
		height:  height,
		fov:     fov,
		spheres: []Sphere{sphere},
	}
}

// Render ...
func (s *Scene) Render() *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, s.width, s.height))

	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			ray := CreatePrime(x, y, s)

			for _, sphere := range s.spheres {
				if sphere.Intersect(ray) >= 0 {
					m.Set(x, y, sphere.color)
				} else {
					m.Set(x, y, color.Black)
				}
			}
		}
	}

	//draw.Draw(m, m.Bounds(), image.Transparent, image.ZP, draw.Src)
	return m
}

// Trace ...
func (s *Scene) Trace(ray *Ray) (*Intersection, error) {
	var sphere Sphere
	var min float64 = -1

	for _, sphere := range s.spheres {
		distance := sphere.Intersect(ray)
		if (distance > 0) && (min < 0 || distance < min) {
			min = distance
		}
	}

	if min < 0 {
		return nil, errors.New("No intersection with this ray")
	}

	return NewIntersection(min, sphere), nil
}

// DegreeToRadius ...
func DegreeToRadius(deg float32) float32 {
	return deg * math.Pi / 180.0
}
