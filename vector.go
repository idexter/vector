package vector

import "math"

type Vector struct {
	data []float64
}

func NewVector(data ...float64) Vector {
	if len(data) < 2 {
		panic("vector should have at least 2 dimensions")
	}
	return Vector{data: data}
}

func NewVectorFromFloat64Array(data []float64) Vector {
	if len(data) < 2 {
		panic("vector should have at least 2 dimensions")
	}
	return Vector{data: data}
}

func (v Vector) Dimension() int {
	return len(v.data)
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.DotProduct(v))
}

func (v Vector) ToFloat64Array() []float64 {
	return v.data
}

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

func (v Vector) ScalarMultiply(scalar float64) Vector {
	output := make([]float64, v.Dimension())
	for k := range v.data {
		output[k] = v.data[k] * scalar
	}
	return Vector{data: output}
}

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
