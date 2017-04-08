package main

import (
	"fmt"
	"os"
)

func main() {
	args, space := "", " "
	for i := 1; i < len(os.Args); i++ {
		args += os.Args[i]
		args += space
	}
	fmt.Println(args)
}
