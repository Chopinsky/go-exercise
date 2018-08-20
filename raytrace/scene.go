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

				m.Set(x, y, clr)
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

// GetColor ...
func (s *Scene) GetColor(ray *Ray, intersect *Intersection) (color.Color, error) {
	hitPoint := ray.origin.Add(ray.direction.Multiply(float32(intersect.distance)).ToPoint())

	surfaceNormal, err := intersect.object.SurfaceNormal(hitPoint)
	if err != nil {
		return nil, errors.New("Unable to render scene: " + err.Error())
	}

	if _, err := s.Trace(ray); err != nil {
		return color.Black, nil
	}

	lightDirection := s.light.direction.Multiply(-1).Normalize()
	lightPower := surfaceNormal.Dot(lightDirection)

	if lightPower <= 0.0 {
		lightPower = 0.0
		return color.Black, nil
	}

	lightPower = lightPower * s.light.intensity
	lightReflected := intersect.object.albedo / math.Pi

	if clr, err := intersect.object.Color(); err == nil {
		pr, pg, pb, _ := clr.RGBA()
		lr, lg, lb, _ := s.light.color.RGBA()
		factor := lightPower * lightReflected

		dr := factor * float32(pr-(pr-lr)/2.0)
		dg := factor * float32(pg-(pg-lg)/2.0)
		db := factor * float32(pb-(pb-lb)/2.0)

		c := color.RGBA{
			uint8(dr),
			uint8(dg),
			uint8(db),
			1.0,
		}

		return c, nil
	}

	return nil, errors.New("Unable to render scene")
}

// DegreeToRadius ...
func DegreeToRadius(deg float32) float32 {
	return deg * math.Pi / 180.0
}
