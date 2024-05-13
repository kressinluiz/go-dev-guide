- How do we run the code in our project?
  Write go run <name-of-the-file> inside the terminal
  The command 'go run' compiles and executes one or more files
  or you can write 'go build <name-of-the-file>' e executar depois.
- What does 'package main' mean?
  It's like being in a project or workspace. A collection of common source code files.
  There's two types of packages:
    - Executable: generates a file that we can run
    - Reusable: code used as 'helpers'. good place to put reusable logic.
  
  The word main is used specifically to produce a executable package.
- What does 'import "fmt"' mean?
  fmt is a standard library of the language. short of format
  docs of the std lib: golang.org/pkg
- What's that func thing?
  it's how we declare a function
- How is the main.go file organized?
  1. package declaration
  2. import other packages that we need
  3. declare functions, add actions