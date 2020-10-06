# Vector

![test](https://github.com/idexter/vector/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/idexter/vector)](https://goreportcard.com/report/github.com/idexter/vector)
[![codecov](https://codecov.io/gh/idexter/vector/branch/master/graph/badge.svg?token=LHA4hN5boC)](undefined)

This library provides basic implementation of linear algebra "vector" data structure with basic operations.
It was made as small test assignment. So only basic things has been implemented.

Supported operations:
- Length
- Add
- Subtract
- ScalarMultiply
- DotProduct
- CrossProduct

Almost all functions within this package panics when vectors used in a wrong way.
For example: It panics if 2 vectors has different dimensions. 
For more information see package documentation.

## Test

Just run `go test -v -race -cover ./...`
