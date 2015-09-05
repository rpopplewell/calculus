package calculus

import (
	"strings"
	// "regexp"
)

// i => term index in function
// [i][0] => constant/variable
// [i][1] => value

func (fnc Function) Parse() [][]string {
	chars := strings.Split(fnc.F, "")
	chars = append(chars, " ")
	var tok string

	const ops = "+-/*^()"
	fncs := [...]string{"sin", "cos", "tan",
		"cot", "sec", "csc",
		"arcsin", "arccos", "arctan",
		"arccot", "arcsec", "arccsc",
		"ln", "log"}
	termNumber := CountTerms(ops, fncs, chars)
	toks := make([][]string, termNumber)

  	termIndex := 0
	for i := 0; i < len(chars); i++ {
		tok += chars[i]

		if IsTerm(ops, chars, i) {
			if strings.Contains(ops, tok) {
				toks[termIndex] = append(toks[termIndex], tok)
			} else if !strings.Contains(tok, "x") {
				var appended bool = false
				for k := 0; k < len(fncs)-1; k++ {
					if fncs[k] == tok {
						toks[termIndex] = append(toks[termIndex], fncs[k])
						appended = true
					}
				}
				if appended == false {
					toks[termIndex] = append(toks[termIndex], "C")
				}
			} else if strings.Contains(tok, "x") {
				toks[termIndex] = append(toks[termIndex], "V")
			}
			toks[termIndex] = append(toks[termIndex], tok)
			tok = ""
			if i == len(chars)-2 {
				break
			}
			termIndex++
		}
	}
	return toks
}

// func GroupTerms(exp string) string {
// 	for i := 0; i < len(toks); i++ {
// 		if IsTerm(ops, chars, i) {
// 		}
// 	}
// }

func CountTerms(ops string, fncs [14]string, chars []string) int {
	var termNum int = 0
	var tok string = ""
	for i := 0; i < len(chars); i++ {
		tok += chars[i]
		if IsTerm(ops, chars, i) {
			termNum++
			tok = ""
			if i == len(chars)-2 {
				break
			}
		}
	}
	return termNum
}

func IsTerm(ops string, chars []string, i int) bool {
	if strings.Contains(ops, chars[i+1]) || strings.Contains(ops, chars[i]) || i == len(chars)-2 {
		return true
	}
	return false
}
