package points

import "math/rand"

type Point struct {
	X int
	Y int
}

type DataGenerator interface {
	GeneratePoint() Point
	GeneratePoints(n int) []Point
}

type PointGenerator struct {
	xlimit int
	ylimit int
}

func NewPointsGenerator() DataGenerator {
	return &PointGenerator{
		xlimit: 10,
		ylimit: 10,
	}
}

func (p *PointGenerator) GeneratePoints(n int) []Point {
	points := make([]Point, 0, n)

	for range n {
		points = append(points, p.GeneratePoint())
	}

	return points
}

func (p *PointGenerator) GeneratePoint() Point {
	return Point{
		X: rand.Intn(p.xlimit),
		Y: rand.Intn(p.ylimit),
	}
}
