package calc

import (
	"math"
)

const Epsilon = math.SmallestNonzeroFloat64

type Func64 func(x float64) float64

func (f Func64) Limit(x, h float64) float64 {
	return f(x + h)
}

func (f Func64) Quotient(a, b float64) float64 {
	return (f(a) - f(b)) / (a - b)
}

func (f Func64) Derivative(x, h float64) float64 {
	return f.Quotient(x+h, x)
}

func (f Func64) Integral(a, b, h float64) float64 {
	if h < 0 {
		h = -h
	} else if h == 0 {
		h = Epsilon
	}
	if a > b {
		t := a
		a = b
		b = t
	}
	sum := float64(0)
	for i := a; i <= b; i += h {
		sum += f(i)
	}
	return sum * h
}
