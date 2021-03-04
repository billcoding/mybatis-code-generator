package config

import (
	. "github.com/billcoding/gobatis"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var Ba *Batis
var envDSN = "DSN"
var DSN = "wangmiao:WHnsDEbkPU6gH0HS@tcp(192.168.1.250:3306)/nterp_wms"

func initEnv() {
	if os.Getenv(envDSN) != "" {
		DSN = os.Getenv(envDSN)
	}
}

func init() {
	initEnv()
	Ba = Default().DSN(DSN)
}
