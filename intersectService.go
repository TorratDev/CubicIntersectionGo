package main

import (
	"math"
)

// IntersectService defines the interface for the intersect logic
type IntersectService interface {
	Intersects(first, second Cubic) bool
	IntersectedVolume(first, second Cubic) float64
}

// Example implementation of IntersectService
type intersectServiceImpl struct{}

// Intersects checks if two cubes intersect (simple example)
func (i intersectServiceImpl) Intersects(first, second Cubic) bool {
	// Check X-axis overlap
	if math.Abs(first.Center.X-second.Center.X) > (first.Dimensions.X/2 + second.Dimensions.X/2) {
		return false
	}
	// Check Y-axis overlap
	if math.Abs(first.Center.Y-second.Center.Y) > (first.Dimensions.Y/2 + second.Dimensions.Y/2) {
		return false
	}
	// Check Z-axis overlap
	if math.Abs(first.Center.Z-second.Center.Z) > (first.Dimensions.Z/2 + second.Dimensions.Z/2) {
		return false
	}
	return true
}

// IntersectedVolume returns the intersected volume between two cubes
func (i intersectServiceImpl) IntersectedVolume(first, second Cubic) float64 {
	// Calculate overlap on the X-axis
	xOverlap := math.Max(0, first.Dimensions.X/2+second.Dimensions.X/2-math.Abs(first.Center.X-second.Center.X))
	// Calculate overlap on the Y-axis
	yOverlap := math.Max(0, first.Dimensions.Y/2+second.Dimensions.Y/2-math.Abs(first.Center.Y-second.Center.Y))
	// Calculate overlap on the Z-axis
	zOverlap := math.Max(0, first.Dimensions.Z/2+second.Dimensions.Z/2-math.Abs(first.Center.Z-second.Center.Z))

	// Return the volume of the overlap
	return xOverlap * yOverlap * zOverlap
}
