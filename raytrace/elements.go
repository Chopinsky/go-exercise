package raytrace

import "math"

// ORIGIN ...
var ORIGIN = &Point{
	x: 0,
	y: 0,
	z: 0,
}

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

// Add ...
func (p *Point) Add(other *Point) *Point {
	return &Point{
		p.x + other.x,
		p.y + other.y,
		p.z + other.z,
	}
}

// Dot ...
func (v1 *Vector) Dot(v2 *Vector) float32 {
	return (v1.x*v2.x + v1.y*v2.y + v1.z*v2.z)
}

// Multiply ...
func (v1 *Vector) Multiply(factor float32) *Vector {
	return &Vector{
		factor * v1.x,
		factor * v1.y,
		factor * v1.z,
	}
}

// ToPoint ...
func (v1 *Vector) ToPoint() *Point {
	return &Point{
		v1.x,
		v1.y,
		v1.z,
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

// CreatePoint ...
func CreatePoint(x, y, z float32) *Point {
	return &Point{
		x: x,
		y: y,
		z: z,
	}
}

// VectorFromPoints ...
func VectorFromPoints(tgt, src *Point) *Vector {
	return &Vector{
		x: tgt.x - src.x,
		y: tgt.y - src.y,
		z: tgt.z - src.z,
	}
}

// VectorFromPoint ...
func VectorFromPoint(p *Point) *Vector {
	return VectorFromPoints(p, ORIGIN)
}
