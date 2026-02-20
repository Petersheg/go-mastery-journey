# 📘 Go Learning Journey – Day 5
## Writing Tests & Compiling/Installing Go Applications

---

## 📅 Overview

On Day 5 of my Go learning journey, I learned about:

- Writing unit tests in Go
- Using the `testing` package
- Running tests with the `go test` command
- Compiling a Go application
- Installing a Go application using Go tooling

References:

- https://go.dev/doc/tutorial/add-a-test  
- https://go.dev/doc/tutorial/compile-install  

---

# 🧪 1. Writing Tests in Go

Testing is an essential part of building reliable software.  
Today, I learned how to write automated tests for my Go functions.

Go provides a built-in `testing` package for writing tests.

---

## 🏗 Test File Naming Convention

- Test files must end with `_test.go`
- Tests must be in the same package as the code being tested
- Tests should verify expected output

## 🏗 Runing Test
- To run tests in the module: `go test`
- To see verbose output: `go test -v`

## Compile the Application
- `go build` : This creates an executable file in the current directory.

## Installing a Go Application
- `go install`: This Compiles the application,
Places the executable in the Go workspace's bin directory
The installed application can then be executed from anywhere in the terminal (if the bin directory is in your PATH).

