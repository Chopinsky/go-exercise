package raytrace

import (
	"image/color"
	"math"
)

// Point ...
type Point struct {
	x float32
	y float32
	z float32
}

// Vector ...
type Vector struct {
	x float32
	y float32
	z float32
}

// Sphere ...
type Sphere struct {
	center Point
	radius float64
	color  color.Color
}

// Intersectable ...
type Intersectable interface {
	Intersect(ray *Ray) bool
}

// Dot ...
func (v1 *Vector) Dot(v2 *Vector) float32 {
	return (v1.x*v2.x + v1.y*v2.y + v1.z*v2.z)
}

// Intersect ...
func (s *Sphere) Intersect(ray *Ray) bool {
	line := MakeVector(s.center, ray.origin)

	adjacent := line.Dot(&ray.direction)
	distance := float64(line.Dot(line) - adjacent*adjacent)
	scope := s.radius * s.radius

	return (distance < scope)
}

// CreatePoint ...
func CreatePoint(x, y, z float32) *Point {
	return &Point{
		x: x,
		y: y,
		z: z,
	}
}

// CreateVector ...
func CreateVector(x, y, z float32) *Vector {
	return &Vector{
		x: x,
		y: y,
		z: z,
	}
}

// Normalize ...
func (v1 *Vector) Normalize() *Vector {
	length := float32(math.Sqrt(float64(v1.x*v1.x + v1.y*v1.y + v1.z*v1.z)))

	if length > 0.0 {
		v1.x = v1.x / length
		v1.y = v1.y / length
		v1.z = v1.z / length
	}

	return v1
}

// CreateSphere ...
func CreateSphere(center Point, radius float64, color color.Color) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
		color:  color,
	}
}

// CreateColor ...
func CreateColor(red, green, blue uint8) *color.RGBA {
	return &color.RGBA{red, green, blue, 1}
}
