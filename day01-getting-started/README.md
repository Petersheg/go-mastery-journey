# Day 01 – Getting Started with Go

## 📚 Source
Official Go tutorial:
https://go.dev/doc/tutorial/getting-started

## 🎯 Goal
- Install Go
- Create a Go module
- Write and run a simple Go program
- Understand basic project structure

## 🛠 What I Did

1. Installed Go on my machine.
2. Created a new directory for the project.
3. Initialized a Go module using:

   go mod init example/hello

4. Created a `hello.go` file with:

   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }

5. Ran the program using:

   go run .

## 🧠 Key Things I Learned

- Every Go program starts with a `package` declaration.
- files with the same package name are group together
- `main` is the entry point of a Go executable.
- The `import` keyword is used to include packages.
- `go mod init` creates a `go.mod` file to manage dependencies.
- `go run .` compiles and runs the program.

## 🚀 Outcome

Successfully ran my first Go program and understood how modules work.
