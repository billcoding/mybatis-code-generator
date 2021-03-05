package main

import (
	"flag"
	"fmt"
	"runtime"
)

func printUsage() {
	fmt.Printf(`Usage of mybatis-code-generator:

mybatis-code-generator -dsn DSN -db DATABASE -OPTIONS

Examples:

mybatis-code-generator -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database"
mybatis-code-generator -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database" -o "/to/path" 
mybatis-code-generator -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database" -au "bigboss" -o "/to/path" 

Supports options:
`)
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf(`
github.com/billcoding/mybatis-code-generator
%s
`, runtime.Version())
}
