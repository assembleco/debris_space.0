package main

import (
	_ "github.com/assembleco/debris_space/engine/demos/animation"
	_ "github.com/assembleco/debris_space/engine/demos/audio"
	_ "github.com/assembleco/debris_space/engine/demos/experimental/physics"
	_ "github.com/assembleco/debris_space/engine/demos/geometry"
	_ "github.com/assembleco/debris_space/engine/demos/gui"
	_ "github.com/assembleco/debris_space/engine/demos/helper"
	_ "github.com/assembleco/debris_space/engine/demos/light"
	_ "github.com/assembleco/debris_space/engine/demos/loader"
	_ "github.com/assembleco/debris_space/engine/demos/material"
	_ "github.com/assembleco/debris_space/engine/demos/other"
	_ "github.com/assembleco/debris_space/engine/demos/shader"
	_ "github.com/assembleco/debris_space/engine/demos/tests"
	_ "github.com/assembleco/debris_space/engine/demos/texture"

	"github.com/assembleco/debris_space/engine/app"
)

func main() {

	// Create and run application
	app.Create().Run()
}
