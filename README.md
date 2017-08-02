# Go Calculus

A golang library supporting definite integration and differentiation of real valued, single variable, elementary and non-elementary functions.
<h3>Usage</h3>

`~$ go get github.com/TheDemx27/calculus`

```go
package main

import (
  "fmt"
  "math"
  clc "github.com/TheDemx27/calculus"
)

func main() {
  // Create a new function.
  f := clc.NewFunc("1/(x^2+1)")

  // Integrate from 0 to 20.
  fmt.Println(f.AntiDiff(0, 20))

  // Differentiate at point Pi/2.
  fmt.Println(f.Diff(math.Pi/2))
}
```
Alternatively, you can define any function that takes and returns a `float64`, (i.e. a function of the form `type usrFunction func(float64) float64`), and then perform calculus on it. This allows you to do calculus on practically any mapping.
```go
func fn(x float64) float64 {
  return 2*math.Floor(x)
}

func main() {
  clc.AntiDiff(fn, 0, 10)
  clc.Diff(fn, 0.5)
}
```
In all cases, adding an additional `int` argument will specify how many samples you want to use.
You can also evaluate the function at a certain point using `f.Eval()`, and return the expression defining the function using `f.Name()`, e.g.
```go
  fmt.Printf("The value of %s when x = 20 is %g\n", f.Name(), f.Eval(20))
```
<h3>TODO</h3>
* Gaussian Quadrature ✓
* Implement Risch Algorithm
* Implement Symbolic Differentiation Patterns
