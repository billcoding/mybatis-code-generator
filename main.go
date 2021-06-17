package main

import (
	"fmt"
	"github.com/billcoding/mybatis-code-generator/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
