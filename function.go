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

func (fnc Function) GetToksLit() string {
	return fnc.GetToks(1)
}

func (fnc Function) GetToksAbstract() string {
	return fnc.GetToks(0)
}

func (fnc Function) GetToks(j int) string {
	var tokArray []string
	var tokString string
	
	fnc.Toks = fnc.Parse()	
	for i := 0; i < len(fnc.Toks); i++ {
		tokArray = append(tokArray, fnc.Toks[i][j])
	}
	for i := 0; i < len(tokArray); i++ {
		tokString = tokString + tokArray[i]
	}
	return tokString
}

func (fnc Function) Eval(val float64) float64 {
	num := strconv.FormatFloat(val, 'G', -1, 64)
	result, err := compute.Evaluate(strings.Replace(fnc.F, "x", num, -1))
	if err != nil {
		log.Fatal(err)
	}
	return result
}
