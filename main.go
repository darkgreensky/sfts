package main

import (
	"sfts/initial"
	"sfts/router"
)

func main() {
	if err := initial.InitMySQL(); err != nil {
		panic(err)
	}
	if err := initial.InitMinio(); err != nil {
		panic(err)
	}

	router.InitRouter()
}
