package calculus

import (
	"strings"
)

/////////////////////////////////HANDLING TOKS/////////////////////////////////

/*
* Returns toks.
* i => tok index in function
* [i][0] => constant/variable
* [i][1] => value
*/
func (fnc Function) ParseChars() [][]string {
	chars := strings.Split(fnc.F, "")
	chars = append(chars, " ")
	var tok string

	const ops = "+-/*^()"
	fncs := [...]string{"sin", "cos", "tan",
		"cot", "sec", "csc",
		"arcsin", "arccos", "arctan",
		"arccot", "arcsec", "arccsc",
		"ln", "log"}
	tokNumber := CountElements(ops, chars)
	toks := make([][]string, tokNumber)

	tokIndex := 0
	for i := 0; i < len(chars); i++ {
		tok += chars[i]

		if IsElement(ops, chars, i) {
			if strings.Contains(ops, tok) {
				toks[tokIndex] = append(toks[tokIndex], tok)
			} else if !strings.Contains(tok, "x") {
				var appended bool = false
				for k := 0; k < len(fncs)-1; k++ {
					if fncs[k] == tok {
						toks[tokIndex] = append(toks[tokIndex], fncs[k])
						appended = true
					}
				}
				if appended == false {
					toks[tokIndex] = append(toks[tokIndex], "C")
				}
			} else if strings.Contains(tok, "x") {
				toks[tokIndex] = append(toks[tokIndex], "V")
			}
			toks[tokIndex] = append(toks[tokIndex], tok)
			tok = ""
			if i == len(chars)-2 {
				break
			}
			tokIndex++
		}
	}
	return toks
}

/*
* Counts the number of elements, e.g. 2*x+345 -> [2,*,x,+,345] -> 5
*/
func CountElements(ops string, elements []string) int {
	elementNum := 0
	element := ""
	for i := 0; i < len(elements); i++ {
		element += elements[i]
		if IsElement(ops, elements, i) {
			elementNum++
			element = ""
			if i == len(elements)-2 {
				break
			}
		}
	}
	return elementNum
}

/*
* Determines if a given character is the last character in an element.
*/
func IsElement(ops string, elements []string, i int) bool {
	if strings.Contains(ops, elements[i+1]) || strings.Contains(ops, elements[i]) || i == len(elements)-2 {
		return true
	} else {
		return false
	}
}

/////////////////////////////////HANDLING TERMS/////////////////////////////////

func (fnc Function) ParseToks() []string {
	const sepOps = "+-"
	tokArrayAbs := fnc.GetToksArrayAbstract()
	termNum := CountElements(sepOps, tokArrayAbs)
	terms := make([]string, termNum)
	var term string

	termIndex := 0
	for i := 0; i < len(tokArrayAbs); i++ {
		if IsElement(sepOps, tokArrayAbs, termNum) {
			terms[termIndex] = term
		} else {
			term += tokArrayAbs[i]
		}
	}
	return terms
}
