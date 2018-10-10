package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shamkrd/phoneBookTest/startup"
)

func main() {

	startup := startup.Programm{}
	startup.Run()
}
