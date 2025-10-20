# 📘 Square Root Approximation using Newton's Method (Go)

## 🧮 Overview

This exercise demonstrates how to compute the square root of a number **without using built-in math functions**, by implementing **Newton’s Method** (also known as the **Newton–Raphson method**).
The goal is to iteratively improve an estimate `z` such that `z²` is as close as possible to a given value `x`.

---

## 🚀 Problem Statement

Given a number `x`, we want to find a number `z` such that:

[
z^2 \approx x
]

Computers often calculate square roots using a loop that repeatedly improves a guess `z` using the Newton iteration formula:

[
z = z - \frac{z^2 - x}{2z}
]

Repeating this update quickly converges to the true square root of `x`.

---

## 💡 Implementation Details

### Function: `Sqrt(x, z0, tol, maxIter, verbose)`

| Parameter | Description                                                                           |
| --------- | ------------------------------------------------------------------------------------- |
| `x`       | The number whose square root is to be computed.                                       |
| `z0`      | Initial guess for the root (e.g., `1.0`, `x`, or `x/2`).                              |
| `tol`     | Tolerance value — loop stops when the change between iterations is smaller than this. |
| `maxIter` | Maximum number of iterations (to avoid infinite loops).                               |
| `verbose` | If `true`, prints each iteration to show convergence progress.                        |

### Return Values

* `z`: The computed approximation of √x.
* `it`: The number of iterations performed.

---

## ⚙️ Algorithm

1. Start with an initial guess `z0`.
2. Repeatedly update `z` using:

   ```go
   z -= (z*z - x) / (2*z)
   ```
3. Stop if:

   * The change between iterations is less than `tol`, or
   * The number of iterations reaches `maxIter`.

---

## 🧭 Example Output

```
Tracing iterations for x = 2 with z0 = 1.0:
iter  1: z = 1.5
iter  2: z = 1.41666666667
iter  3: z = 1.41421568627
iter  4: z = 1.41421356237
...
Summary comparison:
x = 2          | z0 = 1         | iters = 4  | mine = 1.41421356237 | math.Sqrt = 1.41421356237 | absErr = 2.22e-16 | relErr = 1.57e-16
```

---

## 🧪 Try Different Values

You can experiment with different initial guesses and tolerance values:

```go
Sqrt(2, 1.0, 1e-9, 100, true)
Sqrt(10, 10.0, 1e-12, 100, false)
Sqrt(1e-6, 1.0, 1e-15, 100, false)
```

---

## 🧠 Notes

* The method is derived from **Newton’s Method for solving equations**:
  ( f(z) = z^2 - x = 0 ).
* The term `(z*z - x)` measures the distance from the target,
  while dividing by `2*z` scales the correction based on the derivative ( f'(z) = 2z ).
* Works very efficiently for square roots and converges rapidly.

---

## 🧰 Requirements

* Go 1.18+
* Standard library only (`fmt`, `math`)

---

## 🏁 Running the Program

```bash
go run main.go
```

