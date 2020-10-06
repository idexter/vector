package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVector(t *testing.T) {
	v1 := NewVector(1, 2)
	assert.Equal(t, 2, v1.Dimension())
	assert.Equal(t, []float64{1, 2}, v1.ToFloat64Array())

	v2 := NewVector(-5, 3, 1, 7)
	assert.Equal(t, 4, v2.Dimension())
	assert.Equal(t, []float64{-5, 3, 1, 7}, v2.ToFloat64Array())

	assert.Panics(t, func() { NewVector(1) })
	assert.Panics(t, func() { NewVector() })
}

func TestNewVectorFromFloat64Array(t *testing.T) {
	v1 := NewVectorFromFloat64Array([]float64{1, 2})
	assert.Equal(t, 2, v1.Dimension())
	assert.Equal(t, []float64{1, 2}, v1.ToFloat64Array())

	v2 := NewVectorFromFloat64Array([]float64{-5, 3, 1, 7})
	assert.Equal(t, 4, v2.Dimension())
	assert.Equal(t, []float64{-5, 3, 1, 7}, v2.ToFloat64Array())

	assert.Panics(t, func() { NewVectorFromFloat64Array([]float64{1}) })
	assert.Panics(t, func() { NewVectorFromFloat64Array(nil) })
}

func TestVector_Dimension(t *testing.T) {
	v1 := NewVector(-1, 3)
	v2 := NewVector(-5, 4, 7)
	v3 := NewVector(-3, 2, 5, 10)

	assert.Equal(t, 2, v1.Dimension())
	assert.Equal(t, 3, v2.Dimension())
	assert.Equal(t, 4, v3.Dimension())
}

func TestVector_ToFloat64Array(t *testing.T) {
	v1 := NewVector(-1, 3)
	v2 := NewVector(-5, 4, 7)
	v3 := NewVector(-3, 2, 5, 10)

	assert.Equal(t, []float64{-1, 3}, v1.ToFloat64Array())
	assert.Equal(t, []float64{-5, 4, 7}, v2.ToFloat64Array())
	assert.Equal(t, []float64{-3, 2, 5, 10}, v3.ToFloat64Array())
}

func TestVector_Length(t *testing.T) {
	v1 := NewVector(2, 4)
	assert.Equal(t, 4.4721, math.Round(v1.Length()*10000)/10000)
}

func TestVector_Add(t *testing.T) {
	v1 := NewVector(1, -5)
	v2 := NewVector(8, 4)
	assert.Equal(t, NewVector(9, -1), v1.Add(v2))

	assert.Panics(t, func() {
		v1 := NewVector(1, 2)
		v2 := NewVector(3, 4, 5)
		v1.Add(v2)
	})
}

func TestVector_Subtract(t *testing.T) {
	v1 := NewVector(1, -3)
	v2 := NewVector(-6, 6)
	assert.Equal(t, NewVector(7, -9), v1.Subtract(v2))

	assert.Panics(t, func() {
		v1 := NewVector(1, 2)
		v2 := NewVector(3, 4, 5)
		v1.Subtract(v2)
	})
}

func TestVector_ScalarMultiply(t *testing.T) {
	v1 := NewVector(-1, -3)
	v2 := NewVector(-2, 3)
	v3 := NewVector(8, 4)

	assert.Equal(t, NewVector(-6, -18), v1.ScalarMultiply(6))
	assert.Equal(t, NewVector(10, -15), v2.ScalarMultiply(-5))
	assert.Equal(t, NewVector(2, 1), v3.ScalarMultiply(0.25))
}

func TestVector_DotProduct(t *testing.T) {
	v1 := NewVector(2, 5)
	v2 := NewVector(7, 1)
	assert.Equal(t, float64(19), v1.DotProduct(v2))

	v3 := NewVector(1, 2, 3)
	v4 := NewVector(-2, 0, 5)
	assert.Equal(t, float64(13), v3.DotProduct(v4))

	assert.Panics(t, func() {
		v1 := NewVector(1, 2)
		v2 := NewVector(3, 4, 5)
		v1.DotProduct(v2)
	})
}

func TestVector_CrossProduct(t *testing.T) {
	v1 := NewVector(1, -7, 1)
	v2 := NewVector(5, 2, 4)
	assert.Equal(t, NewVector(-30, 1, 37), v1.CrossProduct(v2))

	assert.Panics(t, func() {
		v1 := NewVector(1, 2)
		v2 := NewVector(3, 4)
		v1.CrossProduct(v2)
	})
}
