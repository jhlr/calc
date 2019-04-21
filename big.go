package calc

import (
	"math/big"
)

type Float = *big.Float

type FuncBig func(r, x Float)

func (f FuncBig) Quotient(r, a, b Float) {
	temp1 := reuse()
	temp2 := reuse()
	defer release(temp1, temp2)

	// (f(a) - f(b)) / (a - b)
	f(temp1, a)
	f(temp2, b)
	temp1.Sub(temp1, temp2)
	temp2.Sub(a, b)
	r.Quo(temp1, temp2)
}

func (f FuncBig) Derivative(r, x, h Float) {
	temp := reuse()
	defer release(temp)

	// (f(x+h) - f(x))/h
	temp.Add(x, h)
	f.Quotient(r, temp, x)
}

func (f FuncBig) Limit(r, x, h Float) {
	temp := reuse().Add(x, h)
	defer release(temp)
	f(r, temp)
}

func (f FuncBig) Integral(r, a, b, h Float) {
	absh := reuse().Abs(h)
	i := reuse()
	temp := reuse()
	sum := reuse().SetFloat64(0)
	defer release(i, temp, sum, absh)

	if absh.Sign() == 0 {
		absh.SetFloat64(Epsilon)
	}
	if a.Cmp(b) > 0 {
		t := a
		a = b
		b = t
	}
	for i.Set(a); i.Cmp(b) <= 0; i.Add(i, absh) {
		f(temp, i)
		sum.Add(sum, temp)
	}
	r.Mul(sum, absh)
}
