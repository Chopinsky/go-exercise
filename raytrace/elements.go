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

// Dot ...
func (v1 *Vector) Dot(v2 *Vector) float32 {
	return (v1.x*v2.x + v1.y*v2.y + v1.z*v2.z)
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

// CreateVector ...
func CreateVector(x, y, z float32) *Vector {
	return &Vector{
		x: x,
		y: y,
		z: z,
	}
}
