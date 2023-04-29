package main

import (
	"github.com/makotonakai/medical-hackathon/router"
)

func main() {
	e := router.Initialize()
	e.Logger.Fatal(e.Start(":1323"))
}
