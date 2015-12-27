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

////////////////////////////////////Get Toks////////////////////////////////////

func (fnc Function) GetToksStringLit() string {
	return fnc.GetToksString(1)
}

func (fnc Function) GetToksStringAbstract() string {
	return fnc.GetToksString(0)
}

func (fnc Function) GetToksArrayLit() []string {
	return fnc.GetToksArray(1)
}

func (fnc Function) GetToksArrayAbstract() []string {
	return fnc.GetToksArray(0)
}

func (fnc Function) GetToksArray(j int) []string {
	var tokArray []string
	fnc.Toks = fnc.ParseChars()
	for i := 0; i < len(fnc.Toks); i++ {
		tokArray = append(tokArray, fnc.Toks[i][j])
	}
	return tokArray
}

func (fnc Function) GetToksString(j int) string {
	var tokString string
	var tokArray = fnc.GetToksArray(j)
	for i := 0; i < len(tokArray); i++ {
		tokString = tokString + tokArray[i]
	}
	return tokString
}

////////////////////////////////////////////////////////////////////////////////

/*
* Evaluates a function string e.g. 2*x at a given 'val' for x.
*/
func (fnc Function) Eval(val float64) float64 {
	num := strconv.FormatFloat(val, 'G', -1, 64)
	result, err := compute.Evaluate(strings.Replace(fnc.F, "x", num, -1))
	if err != nil {
		log.Fatal(err)
	}
	return result
}
