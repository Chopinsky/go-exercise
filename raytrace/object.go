package raytrace

import "math"

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

// Color ...
type Color struct {
	red   float32
	green float32
	blue  float32
}

// Sphere ...
type Sphere struct {
	center Point
	radius float64
	color  Color
}

// Intersectable ...
type Intersectable interface {
	Intersect(ray *Ray) bool
}

// Intersect ...
func (s *Sphere) Intersect(ray *Ray) bool {
	return false
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
func (v *Vector) Normalize() *Vector {
	length := float32(math.Sqrt(float64(v.x*v.x + v.y*v.y + v.z*v.z)))

	if length > 0.0 {
		v.x = v.x / length
		v.y = v.y / length
		v.z = v.z / length
	}

	return v
}

// CreateSphere ...
func CreateSphere(center Point, radius float64, color Color) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
		color:  color,
	}
}

// CreateColor ...
func CreateColor(red, green, blue float32) *Color {
	return &Color{
		red:   red,
		green: green,
		blue:  blue,
	}
}
