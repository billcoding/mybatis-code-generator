package main

import (
	"flag"
	"fmt"
	"runtime"
)

func printUsage() {
	fmt.Printf(`Usage of mybatis-code-generator:

mybatis-code-generator -OPTIONS

Examples:

mybatis-code-generator  -au "bigboss" -dsn "root:123@tcp(127.0.0.1:3306)/test"
mybatis-code-generator  -au "bigboss" -ep "x.y.z.entity" -mp "x.y.z.mapper" -xd "xml" -dsn "root:123@tcp(127.0.0.1:3306)/test"
mybatis-code-generator  -au "bigboss" -ep "x.y.z.entity" -mp "x.y.z.mapper" -xd "xml" -o "/to/path" -dsn "root:123@tcp(127.0.0.1:3306)/test"

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
