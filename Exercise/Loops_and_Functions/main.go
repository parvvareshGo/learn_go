package main

import (
	"fmt"
	"math"
)

func Sqrt(x, z0, tol float64, maxIter int, verbose bool) (z float64, it int) {
	if x < 0 {
		return math.NaN(), 0
	}
	if x == 0 {
		return 0, 0
	}
	if z0 == 0 || math.IsNaN(z0) || math.IsInf(z0, 0) {
		z0 = 1.0
	}

	z = z0
	for i := 0; i < maxIter; i++ {
		prev := z
		z -= (z*z - x) / (2 * z)

		if verbose {
			fmt.Printf("iter %2d: z = %.12g\n", i+1, z)
		}

		if math.Abs(z-prev) <= tol {
			return z, i + 1
		}
	}
	return z, maxIter
}

func report(x, z0 float64, tol float64, maxIter int, verbose bool) {
	z, it := Sqrt(x, z0, tol, maxIter, verbose)
	ref := math.Sqrt(x)
	absErr := math.Abs(z - ref)
	var relErr float64
	if ref != 0 {
		relErr = absErr / ref
	}
	fmt.Printf("x = %-10g | z0 = %-10g | iters = %-2d | mine = %.12g | math.Sqrt = %.12g | absErr = %.3e | relErr = %.3e\n",
		x, z0, it, z, ref, absErr, relErr)
}

func main() {
	tol := 1e-12
	maxIter := 100

	xs := []float64{0, 1, 2, 3, 10, 1e-6, 1e6}

	for _, x := range xs {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("Tracing iterations for x = %g with z0 = 1.0:\n", x)
		_, _ = Sqrt(x, 1.0, tol, 10, true)

		fmt.Println("\nSummary comparison:")
		report(x, 1.0, tol, maxIter, false)
		report(x, x, tol, maxIter, false)
		report(x, x/2, tol, maxIter, false)
		fmt.Println()
	}
}
