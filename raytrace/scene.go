package raytrace

import (
	"errors"
	"image"
	"image/color"
	"math"
)

// Light ...
type Light struct {
	direction Vector
	color     color.Color
	intensity float32
}

// Scene ...
type Scene struct {
	width    int
	height   int
	fov      float32
	elements []Element
	light    Light
}

// CreateNewScene ...
func CreateNewScene(width, height int, fov float32, element Element, light Light) *Scene {
	return &Scene{
		width:    width,
		height:   height,
		fov:      fov,
		elements: []Element{element},
		light:    light,
	}
}

// Render ...
func (s *Scene) Render() *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, s.width, s.height))

	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			ray := CreatePrime(x, y, s)

			for _, element := range s.elements {
				distance, err := element.Intersect(ray)
				if err != nil || distance < 0.0 {
					m.Set(x, y, color.Black)
					continue
				}

				clr, err := element.Color()
				if err != nil {
					m.Set(x, y, color.Black)
					continue
				}

				m.Set(x, y, *clr)
			}
		}
	}

	//draw.Draw(m, m.Bounds(), image.Transparent, image.ZP, draw.Src)
	return m
}

// Trace ...
func (s *Scene) Trace(ray *Ray) (*Intersection, error) {
	var el Element
	var min float64 = -1

	for _, element := range s.elements {
		distance, err := element.Intersect(ray)
		if err != nil {
			continue
		}

		if (distance > 0) && (min < 0 || distance < min) {
			min = distance
			el = element
		}
	}

	if min < 0 {
		return nil, errors.New("No intersection with this ray")
	}

	return NewIntersection(min, el), nil
}

// DegreeToRadius ...
func DegreeToRadius(deg float32) float32 {
	return deg * math.Pi / 180.0
}
