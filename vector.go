package vector

import "math"

// Vector described mathematical Vector.
// It stores []float64 array under the hood.
type Vector struct {
	data []float64
}

// NewVector creates new vector.
// It panics when there is less than 2 arguments provided.
func NewVector(data ...float64) Vector {
	if len(data) < 2 {
		panic("vector should have at least 2 dimensions")
	}
	return Vector{data: data}
}

// NewVectorFromFloat64Array creates new vector from float64 slice.
// It panics when array length is less than 2.
func NewVectorFromFloat64Array(data []float64) Vector {
	if len(data) < 2 {
		panic("vector should have at least 2 dimensions")
	}
	return Vector{data: data}
}

// Dimension returns length of underlying array.
func (v Vector) Dimension() int {
	return len(v.data)
}

// Length returns vector length.
func (v Vector) Length() float64 {
	return math.Sqrt(v.DotProduct(v))
}

// ToFloat64Array Convert vector to []float64 array.
func (v Vector) ToFloat64Array() []float64 {
	return v.data
}

// Add adds v1 to current vector.
// It don't mutate current vector but return a new one.
// Also this function panics when vector dimensions are not the same.
func (v Vector) Add(v1 Vector) Vector {
	if v.Dimension() != v1.Dimension() {
		panic("vectors should have at least 2 points")
	}
	output := make([]float64, v.Dimension())
	for k := range v.data {
		output[k] = v.data[k] + v1.data[k]
	}
	return Vector{data: output}
}

// Subtract subtracts v1 from current vector.
// It don't mutate current vector but return a new one.
// Also this function panics when vector dimensions are not the same.
func (v Vector) Subtract(v1 Vector) Vector {
	if v.Dimension() != v1.Dimension() {
		panic("vectors should have same dimension")
	}
	output := make([]float64, v.Dimension())
	for k := range v.data {
		output[k] = v.data[k] - v1.data[k]
	}
	return Vector{data: output}
}

// ScalarMultiply multiply current vector with scalar.
// It don't mutate current vector but return a new one.
func (v Vector) ScalarMultiply(scalar float64) Vector {
	output := make([]float64, v.Dimension())
	for k := range v.data {
		output[k] = v.data[k] * scalar
	}
	return Vector{data: output}
}

// DotProduct do dot product of current vector with v1.
// It don't mutate current vector but return a new one.
// Also this function panics when vector dimensions are not the same.
func (v Vector) DotProduct(v1 Vector) float64 {
	if v.Dimension() != v1.Dimension() {
		panic("vectors should have same dimension")
	}
	var output float64
	for k := range v.data {
		output += v.data[k] * v1.data[k]
	}
	return output
}

// CrossProduct do cross product of current vector with v1.
// It don't mutate current vector but return a new one.
// Also this function panics when vector dimension are not equal to 3.
func (v Vector) CrossProduct(v1 Vector) Vector {
	if v.Dimension() == 3 && v1.Dimension() == 3 {
		output := make([]float64, 3)
		output[0] = v.data[1]*v1.data[2] - v.data[2]*v1.data[1]
		output[1] = v.data[2]*v1.data[0] - v.data[0]*v1.data[2]
		output[2] = v.data[0]*v1.data[1] - v.data[1]*v1.data[0]
		return Vector{data: output}
	}
	panic("operation supported only for 3 dimension vectors")
}
