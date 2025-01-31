# Uninitialized Map Access in Go

This repository demonstrates a common error in Go: accessing a map before it's initialized.  Attempting to access a key in an uninitialized map will result in a runtime panic.

## Bug

The `bug.go` file shows the problem. The `main` function attempts to assign a value to a map key before declaring and initializing the map. This leads to a panic when the program runs.

## Solution

The `bugSolution.go` file demonstrates the correct approach. The map `m` is explicitly initialized before use, preventing the runtime panic.