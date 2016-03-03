package calculus // import "github.com/TheDemx27/calculus"

import (
	"math"
	"fmt"
)

////////////////////////////////////Symbolic////////////////////////////////////

func (fnc Function) SymDiff() string {
	 fmt.Println(fnc.ParseToks())
	 return "compiles..."
	// exp := GroupTerms(fnc.GetToksAbstract, fnc.GetToksLit())
	//
	// switch (exp) {
	// 	case "C*V":
	// 	case "C*V^C":
	// 	case "C*V^V":
	// 	case "ln(V)":
	// 	case "sin(V)":
	// 	case "cos(V)":
	// 	case "-sin(V)":
	// 	case "-cos(V)":
	// 	case "1/V":
	// 	case "V*V":
	// 	case "V/V":
	// 	case "":
	// }
	// return exp
}

///////////////////////////////////Numerical////////////////////////////////////

/*
*	Determines how many samples to use
*/
func SetPrec(defaultPrec int, usrPrec []int) int {
	var prec int
	if len(usrPrec) > 0 {
		prec = usrPrec[0]
	} else {
		prec = defaultPrec
	}
	return prec
}

/*
*	Implements functions defined in code of the form: func someFunc(float64) float64
*/
type cFunc func(float64) float64

/*
*	Implementation of AntiDiff()
*/
func AntiDiff(f cFunc, a, b float64, i ...int) float64 {
		n := SetPrec(10, i)
    x, w := GlqNodes(n, f)
    var sum float64
    bma2 := (b - a) * .5
    bpa2 := (b + a) * .5
    for i, xi := range x {
        sum += w[i] * f(bma2*xi+bpa2)
    }
    return bma2 * sum
}

func GlqNodes(n int, f cFunc) (node []float64, weight []float64) {
    p := LegendrePoly(n)
    pn := p[n]
    n64 := float64(n)
    dn := func(x float64) float64 {
        return (x*pn(x) - p[n-1](x)) * n64 / (x*x - 1)
    }
    node = make([]float64, n)
    for i := range node {
        x0 := math.Cos(math.Pi * (float64(i+1) - .25) / (n64 + .5))
        node[i] = NewtonRaphson(pn, dn, x0)
    }
    weight = make([]float64, n)
    for i, x := range node {
        dnx := dn(x)
        weight[i] = 2 / ((1 - x*x) * dnx * dnx)
    }
    return
}

func LegendrePoly(n int) []cFunc {
    r := make([]cFunc, n+1)
    r[0] = func(float64) float64 { return 1 }
    r[1] = func(x float64) float64 { return x }
    for i := 2; i <= n; i++ {
        i2m1 := float64(i*2 - 1)
        im1 := float64(i - 1)
        rm1 := r[i-1]
        rm2 := r[i-2]
        invi := 1 / float64(i)
        r[i] = func(x float64) float64 {
            return (i2m1*x*rm1(x) - im1*rm2(x)) * invi
        }
    }
    return r
}

func NewtonRaphson(f, df cFunc, x0 float64) float64 {
    for i := 0; i < 30; i++ {
        x1 := x0 - f(x0)/df(x0)
        if math.Abs(x1-x0) <= math.Abs(x0*1e-15) {
            return x1
        }
        x0 = x1
    }
    panic("no convergence")
}

/*
*	Implementation of Diff()
*/
func Diff(fn cFunc, point float64, usrPrec ...int) float64 {
	prec := SetPrec(10000000, usrPrec)
	h := 1 / float64(prec)
	return (fn(point+h) - fn(point)) / h
}

/*
*	Wrappers for Diff and AntiDiff for evaluating the function inputted as a string with the Function struct
*/
func (fnc Function) AntiDiff(lower float64, upper float64, usrPrec ...int) float64 {
	if len(usrPrec) > 0 {
		return AntiDiff(fnc.Eval, lower, upper, usrPrec[0])
	} else {
		return AntiDiff(fnc.Eval, lower, upper)
	}
}

func (fnc Function) Diff(point float64, usrPrec ...int) float64 {
	if len(usrPrec) > 0 {
		return Diff(fnc.Eval, point, usrPrec[0])
	} else {
		return Diff(fnc.Eval, point)
	}
}
