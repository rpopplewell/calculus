package calculus

import (
	"github.com/TheDemx27/calculus/compute"
	"strconv"
	"strings"
	"log"
)

type Function struct {
	F string
}

func (fnc *Function) SetFunc(name string) {
  fnc.F = name
}

func (fnc Function) Name() string {
  return fnc.F
}

func (fnc Function) Eval(val float64) float64 {
	num := strconv.FormatFloat(val, 'G', -1, 64)
	result, err := compute.Evaluate(strings.Replace(fnc.F, "x", num, -1))
	if err != nil {
		log.Fatal(err)
	}
	return result
}
