# calculus
<a href="https://godoc.org/github.com/TheDemx27/calculus"><img src="https://godoc.org/github.com/TheDemx27/calculus?status.svg" alt="GoDoc"></a>

A golang library supporting single variable definite integration and differentiation. The `AntiDiff()` function uses Simpson's Rule, meaning integration of quadratics and other "curved" functions are much more accurate for a given sample size when compared to the Riemann sum method.
<h3>Usage</h3>
```go
func main() {
  fmt.Println(clc.AntiDiff("1/(x^2+1)", 0, 20, 50))
  fmt.Println(clc.Diff("cos(x)", math.Pi/2))
}
```
Integrating 1/x<sup>2</sup>+1 with respect to x from 0 to 20 with 50 samples,
differentiating cos(x) at the point pi/2. Sample size is optional for both functions

---
`go get github.com/TheDemx27/calculus`
