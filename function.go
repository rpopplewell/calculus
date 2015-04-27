package calculus

import (
	"github.com/TheDemx27/calculus/compute"
	"strconv"
	"strings"
	"log"
)

type Function struct {
	f string
}

func (fnc *Function) SetFunc(name string) {
    fnc.f = name
}

func (fnc Function) Name() string {
    return fnc.f
}

func (fnc Function) Eval(val float64) float64 {
	num := strconv.FormatFloat(val, 'G', -1, 64)
	result, err := compute.Evaluate(strings.Replace(fnc.f, "x", num, -1))
	if err != nil {
		log.Fatal(err)
	}
	return result
}
