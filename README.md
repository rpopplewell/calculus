# calculus
A golang library supporting single variable definite integration and differentiation. 
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
