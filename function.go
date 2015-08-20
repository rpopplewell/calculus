package calculus

import (
	"github.com/TheDemx27/calculus/compute"
	"log"
	"strconv"
	"strings"
)

type Function struct {
	F    string
	Toks [][]string
}

func NewFunc(F string) Function {
	var emptySlice [][]string
	fnc := Function{F, emptySlice}
	return fnc
}

func (fnc *Function) SetFunc(name string) {
	fnc.F = name
}

func (fnc Function) GetFunc() string {
	return fnc.F
}

func (fnc Function) GetToks() []string {
	fnc.Parse()
	var RegToks []string
	for i := 0; i < len(fnc.Toks); i++ {
		RegToks = append(RegToks, fnc.Toks[i][1])
	}
	return RegToks
}

func (fnc Function) GetToksAbstract() []string {
	fnc.Parse()
	var ToksAbs []string
	for i := 0; i < len(fnc.Toks); i++ {
		ToksAbs = append(ToksAbs, fnc.Toks[i][0])
	}
	return ToksAbs
}

func (fnc Function) Eval(val float64) float64 {
	num := strconv.FormatFloat(val, 'G', -1, 64)
	result, err := compute.Evaluate(strings.Replace(fnc.F, "x", num, -1))
	if err != nil {
		log.Fatal(err)
	}
	return result
}
