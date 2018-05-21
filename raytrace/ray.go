package raytrace

// Ray ...
type Ray struct {
	origin    Point
	direction Vector
}

// CreatePrime ...
func CreatePrime(x, y uint32, scene *Scene) *Ray {
	fovAdjustment := DegreeToRadius(scene.fov) / 2.0
	aspectRation := float32(scene.width / scene.height)

	sensorX := ((float32(x)+0.5)*2.0/float32(scene.width) - 1.0) * aspectRation * fovAdjustment
	sensorY := (1.0 - (float32(y)+0.5)*2.0/float32(scene.height)) * fovAdjustment

	return &Ray{
		origin:    *CreatePoint(0, 0, 0),
		direction: *CreateVector(sensorX, sensorY, -1.0).Normalize(),
	}
}