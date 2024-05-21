 Here is a simple Golang program that will print "Hello, World!" to the console:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
```

This program defines a function `main` which is the entry point of your application. Inside this function, we first print "Hello, World!" to the console using the `fmt.Println()` function. This is a short form of `fmt.Fprintf(os.Stdout, "%s\n", "Hello, World!")`.

The program then exits with a status code of 0 (zero) because it's a simple console application that will always print something to the console regardless of any errors or exceptions that might occur during its execution.