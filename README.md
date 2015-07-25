# Go Calculus
<a href="https://godoc.org/github.com/TheDemx27/calculus"><img src="https://godoc.org/github.com/TheDemx27/calculus?status.svg" alt="GoDoc"></a>
[![Build Status](https://drone.io/github.com/TheDemx27/calculus/status.png)](https://drone.io/github.com/TheDemx27/calculus/latest)

A golang library supporting single variable definite integration and differentiation. The `AntiDiff()` function uses Simpson's Rule, meaning integration of quadratics and other "curved" functions are much more accurate for a given sample size when compared to the Riemann sum method.
<h3>Usage</h3>
`~$ go get github.com/TheDemx27/calculus`
```go
package main

import (
  "fmt"
  clc "github.com/TheDemx27/calculus"
)

func main() {
  // Create a new function.
  f := clc.Function{"1/(x^2+1)"}

  // Integrate from 0 to 20.
  fmt.Println(f.AntiDiff(0, 20))

  // Differentiate at point Pi/2.
  fmt.Println(f.Diff(math.Pi/2))
}
```
You can also set the sample size by entering another argument at the end of either function call.
In addition, you can also evaluate the function at a certain point using `f.Eval()`, and return the expression defining the function using `f.Name()`, e.g.
```go
  fmt.Printf("The value of %s when x = 20 is %g\n", f.Name(), f.Eval(20))
```
<h3>TODO</h3>
* Implement Risch Algorithm
* Implement Symbolic Differentiation Patterns
