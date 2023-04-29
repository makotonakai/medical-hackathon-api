package main

import (
	"github.com/makotonakai/medical-hackathon/app/router"
)

func main() {
	e := router.Initialize()
	e.Logger.Fatal(e.Start(":1323"))
}
