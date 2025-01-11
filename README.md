# Go Defer Statement Gotcha

This example demonstrates a common pitfall when using Go's `defer` statement.  The deferred function captures the *value* of the variable `x` at the time the `defer` statement is executed, not at the time the deferred function is called.

This can lead to unexpected behavior if the variable's value changes before the deferred function runs, as shown in the code example below. The solution involves using a closure or a copy of the value to preserve the value at the time of the defer call.

## How to reproduce

Simply run `bug.go`.

## Solution

See the `bugSolution.go` file for one possible correction.
