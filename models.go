package main

// CubicRequest represents the request body for the POST request
type CubicRequest struct {
	First  Cubic `json:"first"`
	Second Cubic `json:"second"`
}

// CubicResponse represents the response body
type CubicResponse struct {
	Success bool    `json:"success"`
	Volume  float64 `json:"volume,omitempty"`
}

// Vector3 represents a 3D vector or point
type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Cubic represents a cubic object with a center point and dimensions
type Cubic struct {
	Center     Vector3 `json:"center"`
	Dimensions Vector3 `json:"dimensions"`
}
