package raytrace

import (
	"errors"
	"image/color"
	"math"
)

// Element ...
type Element struct {
	element Intersectable
}

// Sphere ...
type Sphere struct {
	center Point
	radius float64
	color  color.Color
}

// Plane ...
type Plane struct {
	p0     Point
	normal Vector
	color  color.Color
}

// Intersectable ...
type Intersectable interface {
	Intersect(ray *Ray) float64
}

// Color ...
func (e *Element) Color() (*color.Color, error) {
	switch el := e.element.(type) {
	case *Sphere:
		return &el.color, nil
	case *Plane:
		return &el.color, nil
	default:
		return nil, errors.New("Element type is not supported")
	}
}

// Intersect ...
func (e *Element) Intersect(ray *Ray) (float64, error) {
	switch el := e.element.(type) {
	case *Sphere:
		return el.Intersect(ray), nil
	case *Plane:
		return el.Intersect(ray), nil
	default:
		return -1, errors.New("Element type is not supported")
	}
}

// Intersect ...
func (s *Sphere) Intersect(ray *Ray) float64 {
	line := MakeVector(s.center, ray.origin)

	adjacent := float64(line.Dot(&ray.direction))
	distance := float64(line.Dot(line)) - adjacent*adjacent
	scope := s.radius * s.radius

	if distance > scope {
		return -1
	}

	thc := math.Sqrt(scope - distance)
	t0 := adjacent - thc
	t1 := adjacent + thc

	if t0 < 0.0 && t1 < 0.0 {
		return -1
	} else if t0 < t1 {
		return t0
	} else {
		return t1
	}
}

// CreateSphere ...
func CreateSphere(center Point, radius float64, color color.Color) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
		color:  color,
	}
}

// Intersect ...
func (p *Plane) Intersect(ray *Ray) float64 {
	// TODO ...
	return -1.0
}

// CreateColor ...
func CreateColor(red, green, blue uint8) *color.RGBA {
	return &color.RGBA{red, green, blue, 1}
}
