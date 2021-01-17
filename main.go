package main

import (
	"github.com/sadp2021/ghf/app"
)

func main() {
	pipe := app.Pipe{}
	pipe.Exec("./example/example.go")
}
