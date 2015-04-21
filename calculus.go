package calculus

import (
	"math"
	"strconv"
	"strings"
)

import (
	"github.com/TheDemx27/calculus/compute"
)

func Diff(f string, point float64, usrprec ...int) float64 {
	var prec float64 = 100000.0
	if len(usrprec) > 0 {
		prec = float64(usrprec[0])
	}
	h := 1 / prec
	return (fnc(point+h, f) - fnc(point, f)) / h
}

func AntiDiff(f string, lower float64, upper float64, usrprec ...int) float64 {
	prec := 1000
	if len(usrprec) > 0 {
		prec = usrprec[0]
	}
	parts := make([]float64, 2*prec+1)
	r := upper - lower
	nf := float64(prec)
	dx0 := r / nf
	parts[0] = fnc(lower, f) * dx0
	parts[1] = fnc(lower+dx0*.5, f) * dx0 * 4
	x0 := lower + dx0
	for i := 1; i < prec; i++ {
		x1 := lower + float64(i+1)*r/nf
		xmid := (x0 + x1) * .5
		dx := x1 - x0
		parts[2*i] = fnc(x0, f) * dx * 2
		parts[2*i+1] = fnc(xmid, f) * dx * 4
		x0 = x1
	}
	parts[2*prec] = fnc(upper, f) * dx0
	return sum(parts) / 6
}

func fnc(val float64, f string) (result float64) {
	num := strconv.FormatFloat(val, 'G', -1, 64)
	result, _ = compute.Evaluate(strings.Replace(f, "x", num, -1))
	return result
}

func sum(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	var parts []float64
	for _, x := range v {
		var i int
		for _, p := range parts {
			sum := p + x
			var err float64
			if math.Abs(x) < math.Abs(p) {
				err = x - (sum - p)
			} else {
				err = p - (sum - x)
			}
			if err != 0 {
				parts[i] = err
				i++
			}
			x = sum
		}
		parts = append(parts[:i], x)
	}
	var sum float64
	for _, x := range parts {
		sum += x
	}
	return sum
}
