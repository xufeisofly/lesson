package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err.Error())
		} else {
			panic(fmt.Sprintf("don't know %v", err))
		}
	}()
}

func main() {
	tryRecover()
}
