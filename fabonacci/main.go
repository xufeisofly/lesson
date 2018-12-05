package main

import (
	"bufio"
	"fmt"
	"lession/fabonacci/fib"
	"os"
)

func WriteFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else if {
			fmt.Println("Error:", err.Error())
		}
	}
	defer file.Close()
	f := fib.Fibonacci()

	writer := bufio.NewWriter(file)
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
	writer.Flush()
}

func main() {
	WriteFile("fib.txt")
}
